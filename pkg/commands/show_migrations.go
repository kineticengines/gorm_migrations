package commands

import (
	"github.com/kineticengines/gorm-migrations/pkg/definitions"
	"github.com/urfave/cli/v2"
)

// ShowMigrationsCmd ...
var ShowMigrationsCmd = &cli.Command{
	Name:  definitions.ShowMigrationsCmd,
	Usage: definitions.ShowMigrationsCmdUsage,
	Action: func(c *cli.Context) error {
		return nil
	},
}
