package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/kineticengines/gorm-migrations/pkg/definitions"
	"github.com/kineticengines/gorm-migrations/pkg/engine"
	"github.com/urfave/cli/v2"
)

// InitializeCmd ...
var InitializeCmd = &cli.Command{
	Name:  definitions.InitializeCmd,
	Usage: definitions.InitializeCmdUsage,
	Action: func(c *cli.Context) error {
		i := initializer{runner: engine.NewRunner()}
		return i.initialize()
	},
}

var initTemplate = template.Must(template.New("name").Parse(
	`# Where all migrations files will be located
migrations: migrations/*.gormgx

# Where your gorm models are located
models:
  - models/models.go  

# Which database type to use. Defaults to "postgres". Choices : [postgres,mysql,sqlite]
dialect: postgres

# Which driver to use. Defaults to "default". Choices : [default,cloudsqlpostgres]
driver_name: default

# Optional : set the time zone for time.Time fields. Defaults to "Africa/Nairobi"
time_zone: Africa/Nairobi

`))

type initializer struct {
	runner definitions.Worker
}

func (i *initializer) initialize() error {
	exists, err := i.checkIfInitialized()
	if err != nil {
		return err
	}
	if !exists {
		return i.createGormgxYamlFile()
	}
	return definitions.ErrGormgxYamlExists
}

// checkIfInitialized checks for the presence of gormgx.yaml file
// returns an error if it absent
func (i *initializer) checkIfInitialized() (bool, error) {
	file, err := i.runner.GormgxFilePath()
	if err != nil {
		return false, err
	}
	if _, err := os.Stat(*file); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func (i *initializer) createGormgxYamlFile() error {
	path, err := i.runner.GormgxFilePath()
	if err != nil {
		return err
	}
	_, err = os.OpenFile(*path, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := initTemplate.Execute(&buf, ""); err != nil {
		return err
	}

	if err := ioutil.WriteFile(*path, buf.Bytes(), 0600); err != nil {
		return fmt.Errorf("unable to write gormgx file: " + err.Error())
	}

	fmt.Println(definitions.AfterInitializeMessage)
	return nil
}
