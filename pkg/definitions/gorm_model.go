package definitions

//GormModel defines a valid model
type GormModel interface {
	IsModel() bool
}
