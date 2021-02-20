package commands

import (
	"github.com/kineticengines/gorm-migrations/models"
	"github.com/urfave/cli/v2"
)

// ShowMigrationsCmd ...
var ShowMigrationsCmd = &cli.Command{
	Name:  models.ShowMigrationsCmd,
	Usage: models.ShowMigrationsCmdUsage,
	Action: func(c *cli.Context) error {
		return nil
	},
}
