package migrator

import (
	"fmt"
	"strings"

	"github.com/kineticengines/gorm-migrations/pkg/definitions"
	log "github.com/sirupsen/logrus"
)

// Worker is responsible for migrating and mainting migrations state
type Worker struct {
	Name    string
	tree    *definitions.TableTree
	verbose bool
	runner  definitions.Worker
	opDump  *definitions.MigrationDump
}

// NewMigratorWorker instantiates MigrationWorker
func NewMigratorWorker(name string, tree *definitions.TableTree, verbose bool,
	runner definitions.Worker) *Worker {
	return &Worker{
		Name:    name,
		tree:    tree,
		verbose: verbose,
		runner:  runner,
		opDump:  &definitions.MigrationDump{},
	}
}

// RunInitialIntent kickstarts the migration process.
func (w *Worker) RunInitialIntent() error {
	w.runner.PrintVerbose(w.verbose, log.InfoLevel, fmt.Sprintf("Running initial migration for %s", w.Name))
	w.opDump.IsInitial = true

	operations := definitions.Operation{}
	operations.TableName = strings.ToLower(w.Name)

	columns := w.tree.Traverse()
	for e := columns.Front(); e != nil; e = e.Next() {
		meta, ok := e.Value.(*definitions.FieldMeta)
		if ok && meta != nil { // remove nils from the previous step
			fmt.Println(meta)
			col, err := w.decomposeField(meta)
			if err != nil {
				return err
			}
			operations.Columns = append(operations.Columns, col)
		}
	}
	return nil
}

// RunAfterIntialIntent kickstarts the migration process.
func (w *Worker) RunAfterIntialIntent() error {
	// todo: implement

	return nil
}

func (w *Worker) decomposeField(f *definitions.FieldMeta) (*definitions.OpColumn, error) {
	column := definitions.OpColumn{}
	colName, err := w.extractColumnNameFromTag(f)
	if err != nil {
		return nil, err
	}

	column.CanonicalName = *colName
	column.FieldType = f.FieldType

	hasNull, err := w.extractIsNullFromTag(f)
	if err != nil {
		return nil, err
	}

	column.IsNull = !hasNull

	// todo: continue from here

	// column.DatabaseType

	// column.Size int
	// column.Default
	// column.TimeZone
	// column.IsNullable
	// column.IsPrimaryKey
	// column.IsAutoIncrement
	// column.IsIndex
	// column.IsUnique
	// column.IsForeignKey
	// column.ForeignKeyTo
	// column.ForeignKeyConstraints
	// column.Precision
	// column.Scale

	return &column, nil
}

// determines the name of the column
func (w *Worker) extractColumnNameFromTag(f *definitions.FieldMeta) (*string, error) {
	tag := f.Tag
	if len(tag) != 0 {
		if !w.validate(tag) {
			return nil, fmt.Errorf("invalid tag for column: %v", f.FieldName)
		}
		parts := w.splitIntoParts(tag)
		colName := w.findColumnPart(parts)
		if colName == nil {
			// return column name as the field name
			col := strings.ToLower(f.FieldName)
			return &col, nil
		}
		return colName, nil
	}
	col := strings.ToLower(f.FieldName)
	return &col, nil
}

// determines whether a column is nullable
func (w *Worker) extractIsNullFromTag(f *definitions.FieldMeta) (bool, error) {
	tag := f.Tag
	if len(tag) != 0 {
		if !w.validate(tag) {
			return false, fmt.Errorf("invalid tag for column: %v", f.FieldName)
		}
		parts := w.splitIntoParts(tag)
		hasNotNull := w.findNotNullPart(parts)
		return hasNotNull, nil
	}
	return false, nil
}

// validate checks for the validality of tags from the tree
func (w *Worker) validate(tag string) bool {
	return strings.Contains(tag, "gorm")
}

func (w *Worker) splitIntoParts(tag string) []string {
	return strings.Split(strings.Split(strings.Trim(strings.Split(tag, "gorm")[1], `"`), `"`)[1], `;`)
}

func (w *Worker) findColumnPart(parts []string) *string {
	for _, v := range parts {
		if strings.Contains(v, "column") {
			colName := strings.Split(v, ":")[1]
			return &colName
		}
	}
	return nil
}

func (w *Worker) findNotNullPart(parts []string) bool {
	for _, v := range parts {
		if strings.Contains(v, "not null") {
			return true
		}
	}
	return false
}
