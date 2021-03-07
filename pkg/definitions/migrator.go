package definitions

// Migrator ..
type Migrator interface {
	RunIntialIntent() error
	RunAfterIntialIntent() error
}
