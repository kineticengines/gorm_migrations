package commands

import (
	"github.com/kineticengines/gorm-migrations/models"
	"github.com/urfave/cli/v2"
)

// IntializeCmd ...
var IntializeCmd = &cli.Command{
	Name:  models.IntializeCmd,
	Usage: models.IntializeCmdUsage,
	Action: func(c *cli.Context) error {
		return nil
	},
}
