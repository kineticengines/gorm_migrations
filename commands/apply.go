package commands

import (
	"github.com/kineticengines/gorm-migrations/models"
	"github.com/urfave/cli/v2"
)

// ApplyCmd ...
var ApplyCmd = &cli.Command{
	Name:  models.ApplyCmd,
	Usage: models.ApplyCmdUsage,
	Action: func(c *cli.Context) error {
		return nil
	},
}
