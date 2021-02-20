package main

import (
	"os"
	"sort"

	"github.com/kineticengines/gorm-migrations/commands"
	"github.com/kineticengines/gorm-migrations/models"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        models.AppName,
		Description: models.AppDescription,
		Usage:       models.AppUsage,
		Commands: []*cli.Command{
			commands.IntializeCmd,
			commands.MakeMigrationCmd,
			commands.ApplyCmd,
			commands.RevertCmd,
			commands.RevertToCmd,
			commands.ShowMigrationsCmd,
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
