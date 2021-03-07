package commands

import (
	"go/types"
	"os"
	"path/filepath"
	"sync"

	"github.com/kineticengines/gorm-migrations/pkg/definitions"
	"github.com/kineticengines/gorm-migrations/pkg/engine"
	"github.com/kineticengines/gorm-migrations/pkg/migrator"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const errorKey = "error"

// MakeMigrationCmd ...
var MakeMigrationCmd = &cli.Command{
	Name:  definitions.MakemigrationsCmd,
	Usage: definitions.MakemigrationsCmdUsage,
	Action: func(c *cli.Context) error {
		var configPath string
		dir, _ := os.Getwd()
		configPath = filepath.Join(dir, definitions.GormgxYamlFileName)
		verbose := c.Bool("v")
		return NewMgxMaker(configPath, engine.NewRunner(), verbose).Migrate()
	},
}

// MakeMigration ...
type MakeMigration interface {
	Migrate() error

	hasError() error
	loadYaml() MakeMigration
	setIntent() (MakeMigration, error)
	buildIntialIntent() (MakeMigration, error)
	buildAfterIntialIntent() (MakeMigration, error)
}

// MgxMaker ...
type MgxMaker struct {
	modelsPath  string
	verbose     bool
	errorsCache *sync.Map
	config      *definitions.Config
	intent      definitions.Intent
	modelsPkgs  *[]*types.Package
	tables      map[string]*definitions.TableTree
	runner      definitions.Worker
}

// NewMgxMaker ...
func NewMgxMaker(path string, runner definitions.Worker, verbose bool) MakeMigration {
	return &MgxMaker{modelsPath: path, verbose: verbose, errorsCache: &sync.Map{}, runner: runner,
		modelsPkgs: &[]*types.Package{}, tables: map[string]*definitions.TableTree{}}
}

// Migrate ...
func (m *MgxMaker) Migrate() error {
	// read gormgx.yaml file
	if err := m.hasError(); err != nil {
		return err
	}
	if _, err := m.loadYaml().setIntent(); err != nil {
		return err
	}
	if _, err := m.buildIntialIntent(); err != nil {
		return err
	}
	if _, err := m.buildAfterIntialIntent(); err != nil {
		return err
	}
	// called to catch any errors in the final steps
	if err := m.hasError(); err != nil {
		return err
	}
	return nil
}

func (m *MgxMaker) hasError() error {
	v, ok := m.errorsCache.Load(errorKey)
	if ok {
		return v.(error)
	}
	return nil
}

func (m *MgxMaker) loadYaml() MakeMigration {
	m.runner.PrintVerbose(m.verbose, log.InfoLevel, "Reading gormgx.yaml configuration file")
	cfg, err := m.runner.ReadYamlToconfig()
	if err != nil {
		m.errorsCache.Store(errorKey, err)
		return m
	}
	m.config = cfg
	return m
}

func (m *MgxMaker) setIntent() (MakeMigration, error) {
	if err := m.hasError(); err != nil {
		return nil, err
	}

	m.runner.PrintVerbose(m.verbose, log.InfoLevel, "Setting intent")
	if !m.runner.CheckIntialMigrationExists() {
		m.intent = definitions.IntialIntent
		return m, nil
	}

	m.intent = definitions.AfterIntialIntent
	return m, nil
}

func (m *MgxMaker) buildIntialIntent() (MakeMigration, error) {
	if err := m.hasError(); err != nil {
		return nil, err
	}
	if m.intent == definitions.IntialIntent {
		if err := m.runner.ReadIntentModels(m.modelsPkgs, m.config.Models, m.verbose); err != nil {
			m.errorsCache.Store(errorKey, err)
			return nil, err
		}
		// analyze package
		pkgs := *m.modelsPkgs
		var wg sync.WaitGroup
		var mutex = &sync.Mutex{}

		for _, pkg := range pkgs {
			wg.Add(1)
			go func(w *sync.WaitGroup, mtx *sync.Mutex, pkg *types.Package, mgx *MgxMaker) {
				defer w.Done()
				tbl := m.runner.AnalyzePkg(pkg, mgx.verbose)
				for k, v := range tbl {
					mtx.Lock()
					mgx.tables[k] = v
					mtx.Unlock()
				}
			}(&wg, mutex, pkg, m)
		}
		wg.Wait()
	}

	return m.createMigrationFiles()
}

func (m *MgxMaker) buildAfterIntialIntent() (MakeMigration, error) {
	if err := m.hasError(); err != nil {
		return nil, err
	}
	if m.intent == definitions.AfterIntialIntent {
		return m, nil
	}
	return m, nil
}

func (m *MgxMaker) createMigrationFiles() (MakeMigration, error) {
	if m.intent == definitions.IntialIntent {
		var wg sync.WaitGroup
		for tableName, tableTree := range m.tables {
			wg.Add(1)
			go func(w *sync.WaitGroup, tn string, tt *definitions.TableTree, mgx *MgxMaker) {
				defer wg.Done()
				if err := migrator.NewMigratorWorker(tn, tt, m.verbose).RunIntialIntent(); err != nil {
					m.errorsCache.Store(errorKey, err)
				}

			}(&wg, tableName, tableTree, m)

		}
		wg.Wait()
	}

	return m, nil
}
