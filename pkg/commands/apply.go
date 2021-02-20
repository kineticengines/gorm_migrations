package commands

import (
	models "github.com/kineticengines/gorm-migrations/pkg/definitions"
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
