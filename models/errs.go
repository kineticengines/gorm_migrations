package models

import "errors"

var (
	// ErrUnableToGetWorkingDirectory ...
	ErrUnableToGetWorkingDirectory = errors.New("unable to get working directory")

	// ErrGormgxYamlExists ...
	ErrGormgxYamlExists = errors.New("gormgx.yaml already exists")
)
