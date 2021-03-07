package definitions

import (
	"fmt"
)

// SchemaDataType ...
type SchemaDataType interface {
	DataTypeOf() string
}

// PostgresSchemaDataType ....
type PostgresSchemaDataType struct{}

// DataTypeOf ...
func (p *PostgresSchemaDataType) DataTypeOf(field OpColumn) string {
	switch field.FieldType {
	case Bool:
		return "boolean"
	case Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64:
		size := field.Size
		if field.FieldType == Uint {
			size++
		}
		if field.IsAutoIncrement {
			switch {
			case size <= 16:
				return "smallserial"
			case size <= 32:
				return "serial"
			default:
				return "bigserial"
			}
		} else {
			switch {
			case size <= 16:
				return "smallint"
			case size <= 32:
				return "integer"
			default:
				return "bigint"
			}
		}
	case Float32, Float64:
		if field.Precision > 0 {
			if field.Scale > 0 {
				return fmt.Sprintf("numeric(%d, %d)", field.Precision, field.Scale)
			}
			return fmt.Sprintf("numeric(%d)", field.Precision)
		}
		return "decimal"
	case String:
		if field.Size > 0 {
			return fmt.Sprintf("varchar(%d)", field.Size)
		}
		return "text"
	case Time:
		return "timestamptz"
	case Bytes:
		return "bytea"
	}

	return string(field.FieldType)
}
