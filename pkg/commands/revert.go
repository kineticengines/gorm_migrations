package commands

import (
	"github.com/kineticengines/gorm-migrations/pkg/definitions"
	"github.com/urfave/cli/v2"
)

// RevertCmd ...
var RevertCmd = &cli.Command{
	Name:  definitions.RevertCmd,
	Usage: definitions.RevertCmdUsage,
	Action: func(c *cli.Context) error {
		return nil
	},
}
