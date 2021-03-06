package definitions

// BasicType ...
type BasicType string

// Basic types
const (
	Bool       BasicType = "bool"
	Int        BasicType = "int"
	Int8       BasicType = "int8"
	Int16      BasicType = "int16"
	Int32      BasicType = "int32"
	Int64      BasicType = "int64"
	Uint       BasicType = "uint"
	Uint8      BasicType = "uint8"
	Uint16     BasicType = "uint16"
	Uint32     BasicType = "uint32"
	Uint64     BasicType = "uint64"
	Float32    BasicType = "float32"
	Float64    BasicType = "float64"
	Complex64  BasicType = "complex64"
	Complex128 BasicType = "complex128"
	String     BasicType = "string"
	Compound   BasicType = "compound"
	Time       BasicType = "time"
	Nil        BasicType = "nil"
)
