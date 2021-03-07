package migrator

import (
	"sync"

	"github.com/kineticengines/gorm-migrations/pkg/definitions"
)

// Worker is responsible for migrating and mainting migrations state
type Worker struct {
	Name       string
	tree       *definitions.TableTree
	stateCache *sync.Map
	verbose    bool
}

// NewMigratorWorker instantiates MigrationWorker
func NewMigratorWorker(name string, tree *definitions.TableTree, verbose bool) *Worker {
	return &Worker{
		Name:       name,
		tree:       tree,
		verbose:    verbose,
		stateCache: &sync.Map{},
	}
}

// RunIntialIntent kickstarts the migration process.
func (w *Worker) RunIntialIntent() error {
	// todo: implement
	return nil
}

// RunAfterIntialIntent kickstarts the migration process.
func (w *Worker) RunAfterIntialIntent() error {
	// todo: implement
	_ = w.validate()
	return nil
}

// Validate checks for the validality of tags from the tree
func (w *Worker) validate() error {
	return nil
}
