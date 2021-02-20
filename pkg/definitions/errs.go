package definitions

import "errors"

var (
	// ErrUnableToGetWorkingDirectory ...
	ErrUnableToGetWorkingDirectory = errors.New("unable to get working directory")

	// ErrGormgxYamlExists ...
	ErrGormgxYamlExists = errors.New("gormgx.yaml already exists")

	// ErrFailedToLoadGormgxFile ...
	ErrFailedToLoadGormgxFile = errors.New("failed to load gormgx.yaml")

	// ErrFailedToReadGormgxFile ...
	ErrFailedToReadGormgxFile = errors.New("failed to read gormgx")

	// ErrFailedToUnmarshalGormgxFile ...
	ErrFailedToUnmarshalGormgxFile = errors.New("failed to read gormgx.yaml")

	// ErrNilType ...
	ErrNilType = errors.New("nil type found. This cause by embedding structs into models")
)
