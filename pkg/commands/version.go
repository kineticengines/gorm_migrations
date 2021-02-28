package commands

import (
	"fmt"

	models "github.com/kineticengines/gorm-migrations/pkg/definitions"
	"github.com/urfave/cli/v2"
)

// Version ...
var (
	version   = "dev"
	buildtime = "dev"
)

// VersionCmd ...
var VersionCmd = &cli.Command{
	Name:  models.VersionCmd,
	Usage: models.VersionCmdUsage,
	Action: func(c *cli.Context) error {
		fmt.Printf(`Gormgx version  : %v
Gormgx buildtime : %v
`, version, buildtime)
		return nil
	},
}
