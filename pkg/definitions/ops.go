package definitions

// Op identifier of operations that can be invoked
type Op int

const (

	// CreateModelOp creates a new model and the corresponding database table
	CreateModelOp Op = iota

	// DeleteModelOp deletes a model and drops its database table
	DeleteModelOp

	// RenameModelOp renames a model and renames its database table
	RenameModelOp

	// AlterModelTableOp renames the database table for a model
	AlterModelTableOp

	// AlterUniqueTogetherOp changes the unique constraints of a model
	AlterUniqueTogetherOp

	// AlterIndexTogetherOp changes the indexes of a model
	AlterIndexTogetherOp

	// AlterOrderWithRespectToOp creates or deletes the _order column for a model
	AlterOrderWithRespectToOp

	// AlterModelOptionsOp changes various model options without affecting the database
	AlterModelOptionsOp

	// AddFieldOp adds a field to a model and the corresponding column in the database
	AddFieldOp

	// RemoveFieldOp removes a field from a model and drops the corresponding column from the database
	RemoveFieldOp

	// AlterFieldOp changes a field’s definition and alters its database column if necessary
	AlterFieldOp

	// RenameFieldOp renames a field and, if necessary, also its database column
	RenameFieldOp

	// AddIndexOp creates an index in the database table for the model
	AddIndexOp

	// RemoveIndexOp removes an index in the database table for the model
	RemoveIndexOp
)

// FieldConstraint ...
type FieldConstraint int

const (

	// OnDeleteCascadeConstraint ...
	OnDeleteCascadeConstraint FieldConstraint = 1

	// OnUpdateCascaseConstraint ...
	OnUpdateCascaseConstraint
)

// OpColumn ...
type OpColumn struct {

	// this is the column name. If name is provided by the tags, gormgx will use that.
	// Otherwise, lowercase, underscore separated name from struct will be used
	CanonicalName string

	FieldType BasicType

	// the type of the field that corresponds to what the database expects
	DatabaseType string

	// whether this column can take nulls
	IsNull bool

	// the maximum number of characters. this is for `VARCHAR`, INT and Floats types
	Size int

	// the default of the the column, It's an interface since we don't know the `DatabaseType` for hand
	Default string

	// the timezone to use for date fields. All date/time fields in gormgx are timezone by default
	TimeZone string

	IsNullable bool

	IsPrimaryKey bool

	IsAutoIncrement bool

	// whether this column should considered as an index
	IsIndex bool

	IsUnique bool

	IsForeignKey bool

	// the table that this column is foreignKey to. the value will be like `Model.Field`
	ForeignKeyTo string

	ForeignKeyConstraints FieldConstraint

	// applied for floats data types
	Precision int

	Scale int
}

// Operation describes a single operation with all it's relevant metadata
type Operation struct {
	Op        Op
	TableName string
	Columns   []*OpColumn
}

// MigrationDump ...
type MigrationDump struct {
	IsInitial    bool        `yaml:"is_initial"`
	Dependencies []string    `yaml:"dependencies"`
	Operations   []Operation `yaml:"operations"`
}
