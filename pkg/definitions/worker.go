package definitions

import (
	"go/types"

	log "github.com/sirupsen/logrus"
)

// Worker ...
type Worker interface {
	PrintVerbose(verbose bool, logLevel log.Level, message interface{})

	GormgxFilePath() (*string, error)

	ReadYamlToconfig() (*Config, error)

	CheckIntialMigrationExists() bool

	ReadIntentModels(modelsPkgs *[]*types.Package, paths []string, verbose bool) error

	ReadModelsFromPath(path string) (*types.Package, error)

	ReadFileSet(path string) (*types.Package, error)

	ReadInterfaceFile() []*types.Named

	AnalyzePkg(pkg *types.Package, verbose bool) map[string]*TableTree

	NameTypeFieldsMeta(v *types.Named) *TableTree

	SplitTypedNameToObjectName(t *types.Named) string
}
