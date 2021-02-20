package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/kineticengines/gorm-migrations/pkg/definitions"
	"github.com/urfave/cli/v2"
)

// IntializeCmd ...
var IntializeCmd = &cli.Command{
	Name:  definitions.IntializeCmd,
	Usage: definitions.IntializeCmdUsage,
	Action: func(c *cli.Context) error {
		return initialize()
	},
}

var initTemplate = template.Must(template.New("name").Parse(
	`# Where all migrations files will be located
migrations: migrations/*.gormgx

# Where your gorm models are located
models:
  - models/models.go  

# Optional: set to add "gorm.Model" to your models
add_gorm_model: true

`))

func initialize() error {
	exists, err := checkIfInitialized()
	if err != nil {
		return err
	}
	if !exists {
		return createGormgxYamlFile()
	}
	return definitions.ErrGormgxYamlExists
}

// checkIfInitialized checks for the presence of gormgx.yaml file
// returns an error if it absent
func checkIfInitialized() (bool, error) {
	file, err := gormgxFilePath()
	if err != nil {
		return false, err
	}
	if _, err := os.Stat(*file); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func createGormgxYamlFile() error {
	path, err := gormgxFilePath()
	if err != nil {
		return err
	}
	_, err = os.OpenFile(*path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := initTemplate.Execute(&buf, ""); err != nil {
		return err
	}

	if err := ioutil.WriteFile(*path, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("unable to write gormgx file: " + err.Error())
	}

	fmt.Println(definitions.AfterIntializeMessage)
	return nil
}
