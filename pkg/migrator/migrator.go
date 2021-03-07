package migrator

import (
	"fmt"
	"sync"

	"github.com/kineticengines/gorm-migrations/pkg/definitions"
	log "github.com/sirupsen/logrus"
)

// Worker is responsible for migrating and mainting migrations state
type Worker struct {
	Name       string
	tree       *definitions.TableTree
	stateCache *sync.Map
	verbose    bool
	runner     definitions.Worker
}

// NewMigratorWorker instantiates MigrationWorker
func NewMigratorWorker(name string, tree *definitions.TableTree, verbose bool,
	runner definitions.Worker) *Worker {
	return &Worker{
		Name:       name,
		tree:       tree,
		verbose:    verbose,
		stateCache: &sync.Map{},
		runner:     runner,
	}
}

// RunIntialIntent kickstarts the migration process.
func (w *Worker) RunIntialIntent() error {
	w.runner.PrintVerbose(w.verbose, log.InfoLevel, fmt.Sprintf("Running initial migration for %s", w.Name))
	columns := w.tree.Traverse()
	for e := columns.Front(); e != nil; e = e.Next() {
		meta, ok := e.Value.(*definitions.FieldMeta)
		if ok && meta != nil {
			fmt.Println(meta)
		}
	}
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
