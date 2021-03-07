package definitions

import (
	"github.com/kyokomi/emoji"
)

// context constants
const (
	AppName                 = "Gorm Migrations [gormgx]"
	AppDescription          = `Gorm Migrations CLI utility manages SQL migrations for gorm. It implements the same API as Django migrations, so you should be at home if coming from Django.`
	AppUsage                = "Making database changes ease, manageable and maintainable"
	GormgxYamlFileName      = "gormgx.yaml"
	IntialMigrationFileName = "00001_init.gormgx"
	DefaultMIgrationsPath   = "migrations"
)

// commands
const (
	IntializeCmd      = "init"
	MakemigrationsCmd = "make-migrations"
	ApplyCmd          = "apply"
	RevertCmd         = "revert"
	RevertToCmd       = "revert-to"
	ShowMigrationsCmd = "show-migrations"
	VersionCmd        = "version"
)

// commands usage descriptions
const (
	IntializeCmdUsage      = "intializes the default gormgx configuration. It create gormgx.yaml in the current working directory"
	MakemigrationsCmdUsage = "analyzes models and create migrations"
	ApplyCmdUsage          = "commits migrations to the database"
	RevertCmdUsage         = "undoes the previously performed migration"
	RevertToCmdUsage       = "reverts migrations to a specific migration"
	ShowMigrationsCmdUsage = "shows all migrations"
	VersionCmdUsage        = "prints the version number"
)

// flags
const (
	VerboseFlag = "v"
)

// flags usage descriptions
const (
	VerboseFlagUsage = "allows verbose mode"
)

// Intent ...
type Intent int

const (
	// IntialIntent ...
	IntialIntent Intent = iota

	// CreateIntent ...
	CreateIntent
)

// response message
var (
	AfterIntializeMessage = emoji.Sprint(`:beer: Hurray!!! Gormgx has been intialized. Check "gormgx.yaml" file and amend it to your needs. Remember not to remove it`)
)
