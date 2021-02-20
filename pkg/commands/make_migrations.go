package commands

import (
	"go/types"
	"os"
	"path/filepath"
	"sync"

	"github.com/kineticengines/gorm-migrations/pkg/definitions"
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
		return NewMgxMaker(configPath, verbose).Migrate()
	},
}

// MakeMigration ...
type MakeMigration interface {
	Migrate() error

	hasError() error
	loadYaml() MakeMigration
	setIntent() (MakeMigration, error)
	buildIntialIntent() (MakeMigration, error)
	buildCreateIntent() (MakeMigration, error)
}

// MgxMaker ...
type MgxMaker struct {
	modelsPath     string
	verbose        bool
	gormgxFilePath string
	errorsCache    *sync.Map
	config         *definitions.Config
	intent         definitions.Intent
	modelsPkgs     *[]*types.Package
}

// NewMgxMaker ...
func NewMgxMaker(path string, verbose bool) MakeMigration {
	return &MgxMaker{modelsPath: path, verbose: verbose, errorsCache: &sync.Map{}, modelsPkgs: &[]*types.Package{}}
}

// Migrate ...
func (m *MgxMaker) Migrate() error {
	// read gormgx.yaml file
	if err := m.hasError(); err != nil {
		return err
	}
	_, err := m.loadYaml().setIntent()
	_, err = m.buildIntialIntent()
	_, err = m.buildCreateIntent()
	return err
}

func (m *MgxMaker) hasError() error {
	v, ok := m.errorsCache.Load(errorKey)
	if ok {
		return v.(error)
	}
	return nil
}

func (m *MgxMaker) loadYaml() MakeMigration {
	printVerbose(m.verbose, log.InfoLevel, "Reading gormgx.yaml configuration file")
	cfg, err := readYamlToconfig()
	if err != nil {
		m.errorsCache.Store(errorKey, err)
		return m
	}
	m.config = cfg
	return m
}

func (m *MgxMaker) setIntent() (MakeMigration, error) {
	printVerbose(m.verbose, log.InfoLevel, "Setting intent")
	if err := m.hasError(); err != nil {
		return nil, err
	}
	if !checkIntialMIgrationExists() {
		m.intent = definitions.IntialIntent
		return m, nil
	}

	m.intent = definitions.CreateIntent
	return m, nil
}

func (m *MgxMaker) buildIntialIntent() (MakeMigration, error) {
	if err := m.hasError(); err != nil {
		return nil, err
	}
	if m.intent == definitions.IntialIntent {
		if err := readIntentModels(m.modelsPkgs, m.config.Models, m.verbose); err != nil {
			m.errorsCache.Store(errorKey, err)
			return nil, err
		}
		// analyze package
		// todo: use sync.WaitGroup to loop
		pkgs := *m.modelsPkgs
		pkg := pkgs[0]
		analyzePkg(pkg, m.verbose)
		return m, nil
	}
	return m, nil
}

func (m *MgxMaker) buildCreateIntent() (MakeMigration, error) {
	if err := m.hasError(); err != nil {
		return nil, err
	}
	if m.intent == definitions.CreateIntent {
		return m, nil
	}
	return m, nil
}
