package definitions

import (
	"errors"

	"github.com/kyokomi/emoji"
)

var (
	// ErrUnableToGetWorkingDirectory ...
	ErrUnableToGetWorkingDirectory = errors.New(emoji.Sprint(`:grimace: Unable to get working directory`))

	// ErrGormgxYamlExists ...
	ErrGormgxYamlExists = errors.New(emoji.Sprint(`:clown_face: Seems like gormgx.yaml already exists`))

	// ErrFailedToFetchGormgxPath ...
	ErrFailedToFetchGormgxPath = errors.New(`:clown_face: Failed to fetch gormgx.yaml path`)

	// ErrFailedToReadGormgxFile ...
	ErrFailedToReadGormgxFile = errors.New(`:clown_face: Failed to read gormgx.yaml. Try running  "gormgx init" first`)

	// ErrFailedToUnmarshalGormgxFile ...
	ErrFailedToUnmarshalGormgxFile = errors.New(`:clown_face: Failed to unmarshal gormgx.yaml`)

	// ErrNilType ...
	ErrNilType = errors.New("nil type found. This cause by embedding structs into models")
)
