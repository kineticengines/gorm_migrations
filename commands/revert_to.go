package commands

import (
	"github.com/kineticengines/gorm-migrations/models"
	"github.com/urfave/cli/v2"
)

// RevertToCmd ...
var RevertToCmd = &cli.Command{
	Name:  models.RevertToCmd,
	Usage: models.RevertToCmdUsage,
	Action: func(c *cli.Context) error {
		return nil
	},
}
