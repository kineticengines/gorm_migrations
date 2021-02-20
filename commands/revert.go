package commands

import (
	"github.com/kineticengines/gorm-migrations/models"
	"github.com/urfave/cli/v2"
)

// RevertCmd ...
var RevertCmd = &cli.Command{
	Name:  models.RevertCmd,
	Usage: models.RevertCmdUsage,
	Action: func(c *cli.Context) error {
		return nil
	},
}
