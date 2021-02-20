package commands

import (
	"os"
	"path/filepath"

	"github.com/kineticengines/gorm-migrations/models"
	"github.com/urfave/cli/v2"
)

// MakeMigrationCmd ...
var MakeMigrationCmd = &cli.Command{
	Name:  models.MakemigrationsCmd,
	Usage: models.MakemigrationsCmdUsage,
	Action: func(c *cli.Context) error {
		var configPath string
		if c.String(models.FileFlag) != "" {
			configPath = c.String(models.FileFlag)
		} else {
			dir, _ := os.Getwd()
			configPath = filepath.Join(dir, models.GormgxYamlFileName)
		}
		NewMgxMaker(configPath).Migrate()
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  models.FileFlag,
			Usage: models.FileFlagUsage,
		},
	},
}

// MakeMigration ...
type MakeMigration interface {
	Migrate()

	loadYaml() MakeMigration
	checkIntial() MakeMigration
}

// MgxMaker ...
type MgxMaker struct {
	workPath string
}

// NewMgxMaker ...
func NewMgxMaker(path string) MakeMigration {
	return &MgxMaker{workPath: path}
}

// Migrate ...
func (m *MgxMaker) Migrate() {

}

func (m *MgxMaker) loadYaml() MakeMigration {
	return m
}

func (m *MgxMaker) checkIntial() MakeMigration {
	return m
}
