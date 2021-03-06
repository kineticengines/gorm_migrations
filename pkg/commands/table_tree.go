package commands

import (
	"container/list"
	"go/types"

	"github.com/kineticengines/gorm-migrations/pkg/definitions"
)

// TableTree tracks the columns of an SQL table
type TableTree struct {
	LeftNode  *TableTree // basic types
	RightNode *TableTree // compound types
	Value     *definitions.FieldMeta
}

// AddNodes adds a node to the table tree by checking its type
func (t *TableTree) AddNodes(u *types.Struct) {
	t.addNodesHelper(u, u.NumFields(), 0)
}

func (t *TableTree) addNodesHelper(u *types.Struct, numOfFields int, index int) {
	fieldType := t.computeBasicType(u.Field(index).Type().Underlying())
	if fieldType != definitions.Compound {
		t.LeftNode = &TableTree{Value: &definitions.FieldMeta{FieldName: u.Field(index).Name(),
			Tag: u.Tag(index), FieldType: fieldType}}
		if index+1 < numOfFields {
			t.LeftNode.addNodesHelper(u, numOfFields, index+1)
		}
	}
}

func (t *TableTree) computeBasicType(u types.Type) definitions.BasicType {
	switch x := u.(type) {
	case *types.Struct:
		return definitions.Compound
	case *types.Pointer:
		elem := x.Underlying().(*types.Pointer).Elem()
		return t.computeBasicType(elem)
	case *types.Basic:
		switch x.Kind() {
		case types.Int:
			return definitions.Int
		case types.Int8:
			return definitions.Int8
		case types.Int16:
			return definitions.Int16
		case types.Int32:
			return definitions.Int32
		case types.Int64:
			return definitions.Int64
		case types.Uint:
			return definitions.Uint
		case types.Uint8:
			return definitions.Uint8
		case types.Uint16:
			return definitions.Uint16
		case types.Uint32:
			return definitions.Uint32
		case types.Uint64:
			return definitions.Uint64
		case types.Float32:
			return definitions.Float32
		case types.Float64:
			return definitions.Float64
		case types.Complex64:
			return definitions.Complex64
		case types.Complex128:
			return definitions.Complex128
		case types.String:
			return definitions.String
		case types.Bool:
			return definitions.Bool
		}
	}
	return definitions.Nil
}

// Traverse ...
func (t *TableTree) Traverse() list.List {
	var m list.List
	return t.traverseHelper(m)
}

func (t *TableTree) traverseHelper(m list.List) list.List {
	if t.LeftNode != nil {
		m.PushBack(t.LeftNode.Value)
		return t.LeftNode.traverseHelper(m)
	}
	return m
}
