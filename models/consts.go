package models

const (
	GormgxYamlFileName = "gormgx.yaml"
	AppName            = "Gorm Migrations [gormgx]"
	AppDescription     = `Gorm Migrations CLI utility manages SQL migrations for gorm. It implements the same API as Django migrations, so you should be at home if coming from Django.`
	AppUsage           = "Making database changes ease, managable and maintainable"
)

// commands
const (
	IntializeCmd      = "init"
	MakemigrationsCmd = "make-migrations"
	ApplyCmd          = "apply"
	RevertCmd         = "revert"
	RevertToCmd       = "revert-to"
	ShowMigrationsCmd = "show-migrations"
)

// commands usage descriptions
const (
	IntializeCmdUsage      = "intializes the default gormgx configuration. It create gormgx.yaml in the current working directory"
	MakemigrationsCmdUsage = "analyzes models and create migrations"
	ApplyCmdUsage          = "commits migrations to the database"
	RevertCmdUsage         = "undoes the previously performed migration"
	RevertToCmdUsage       = "reverts migrations to a specific migration"
	ShowMigrationsCmdUsage = "shows all migrations"
)

// flags
const (
	FileFlag = "file"
)

// flags usage descriptions
const (
	FileFlagUsage = "loads configuration from file path"
)
