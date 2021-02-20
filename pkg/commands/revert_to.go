package commands

import (
	"github.com/kineticengines/gorm-migrations/pkg/definitions"
	"github.com/urfave/cli/v2"
)

// RevertToCmd ...
var RevertToCmd = &cli.Command{
	Name:  definitions.RevertToCmd,
	Usage: definitions.RevertToCmdUsage,
	Action: func(c *cli.Context) error {
		return nil
	},
}
