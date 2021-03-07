package main

import (
	"os"
	"sort"

	"github.com/kineticengines/gorm-migrations/pkg/commands"
	"github.com/kineticengines/gorm-migrations/pkg/definitions"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        definitions.AppName,
		Description: definitions.AppDescription,
		Usage:       definitions.AppUsage,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  definitions.VerboseFlag,
				Usage: definitions.VerboseFlagUsage,
			},
		},
		Commands: []*cli.Command{
			commands.InitializeCmd,
			commands.MakeMigrationCmd,
			commands.ApplyCmd,
			commands.RevertCmd,
			commands.RevertToCmd,
			commands.ShowMigrationsCmd,
			commands.VersionCmd,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
