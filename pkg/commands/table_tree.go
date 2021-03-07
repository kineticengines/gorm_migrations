package commands

import (
	"container/list"
	"go/types"
	"strings"

	"github.com/kineticengines/gorm-migrations/pkg/definitions"
)

// TableTree tracks the columns of an SQL table
type TableTree struct {
	ChildNode *TableTree

	Value *definitions.FieldMeta
}

// AddNodes adds a node to the table tree by checking its type
func (t *TableTree) AddNodes(u *types.Struct) {
	t.addNodesHelper(u, u.NumFields(), 0)
}

func (t *TableTree) addNodesHelper(u *types.Struct, numOfFields int, index int) {
	fieldType := t.computeBasicType(u.Field(index).Type().Underlying())
	switch fieldType {
	case definitions.Compound:
		compoundField := u.Field(index).Type().Underlying().(*types.Struct)
		compoundFieldTree := new(TableTree)
		compoundFieldTree.AddNodes(compoundField)
		compoundFieldNodes := compoundFieldTree.Traverse()

		refNode := t
		for e := compoundFieldNodes.Front(); e != nil; e = e.Next() {
			if refNode.ChildNode == nil {
				meta, ok := e.Value.(*definitions.FieldMeta)
				if ok {
					refNode.ChildNode = &TableTree{Value: meta}
					refNode = refNode.ChildNode
				}
			}
		}

		if index+1 < numOfFields {
			// this adds a child whose value is nil. when traversing the data structure, expect to have child nodes with no values
			// this is not a major pain point since the nil are remove at a later stage.
			refNode.ChildNode = new(TableTree)
			refNode.ChildNode.addNodesHelper(u, numOfFields, index+1)
		}

	default:
		t.ChildNode = &TableTree{Value: &definitions.FieldMeta{FieldName: u.Field(index).Name(),
			Tag: u.Tag(index), FieldType: fieldType}}
		if index+1 < numOfFields {
			t.ChildNode.addNodesHelper(u, numOfFields, index+1)
		}

	}

}

func (t *TableTree) computeBasicType(u types.Type) definitions.BasicType {
	switch x := u.(type) {
	case *types.Struct:
		if t.isOfTimeType(x) || t.isOfNullableTimeType(x) {
			return definitions.Time
		}
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

func (t *TableTree) isOfTimeType(x types.Type) bool {
	field := x.Underlying().(*types.Struct)
	if field.NumFields() == 3 && strings.Contains(field.Field(0).String(), "field wall uint64") &&
		strings.Contains(field.Field(1).String(), "field ext int64") && strings.Contains(field.Field(2).String(), "field loc *time.Location") {
		return true
	}
	return false
}

func (t *TableTree) isOfNullableTimeType(x types.Type) bool {
	field := x.Underlying().(*types.Struct)
	if field.NumFields() == 2 && strings.Contains(field.Field(0).String(), "field Time time.Time") &&
		strings.Contains(field.Field(1).String(), "field Valid bool") {
		return true
	}
	return false
}

// Traverse ...
func (t *TableTree) Traverse() list.List {
	var m list.List
	return t.traverseHelper(m)
}

func (t *TableTree) traverseHelper(m list.List) list.List {
	if t.ChildNode != nil {
		m.PushBack(t.ChildNode.Value)
		return t.ChildNode.traverseHelper(m)
	}
	return m
}
