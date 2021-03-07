package definitions

import (
	"container/list"
	"go/types"
	"strings"
)

// TableTree tracks the columns of an SQL table
type TableTree struct {
	childNode *TableTree

	value *FieldMeta
}

// AddNodes adds a node to the table tree by checking its type
func (t *TableTree) AddNodes(u *types.Struct) {
	t.addNodesHelper(u, u.NumFields(), 0)
}

func (t *TableTree) addNodesHelper(u *types.Struct, numOfFields int, index int) {
	fieldType := t.computeBasicType(u.Field(index).Type().Underlying())
	switch fieldType {
	case Compound:
		compoundField := u.Field(index).Type().Underlying().(*types.Struct)
		compoundFieldTree := new(TableTree)
		compoundFieldTree.AddNodes(compoundField)
		compoundFieldNodes := compoundFieldTree.Traverse()

		refNode := t
		for e := compoundFieldNodes.Front(); e != nil; e = e.Next() {
			if refNode.childNode == nil {
				meta, ok := e.Value.(*FieldMeta)
				if ok {
					refNode.childNode = &TableTree{value: meta}
					refNode = refNode.childNode
				}
			}
		}

		if index+1 < numOfFields {
			// this adds a child whose value is nil. when traversing the data structure, expect to have child nodes with no values
			// this is not a major pain point since the nil are remove at a later stage.
			refNode.childNode = new(TableTree)
			refNode.childNode.addNodesHelper(u, numOfFields, index+1)
		}

	default:
		t.childNode = &TableTree{value: &FieldMeta{FieldName: u.Field(index).Name(),
			Tag: u.Tag(index), FieldType: fieldType}}
		if index+1 < numOfFields {
			t.childNode.addNodesHelper(u, numOfFields, index+1)
		}

	}

}

func (t *TableTree) computeBasicType(u types.Type) BasicType {
	switch x := u.(type) {
	case *types.Struct:
		if t.isOfTimeType(x) || t.isOfNullableTimeType(x) {
			return Time
		}
		return Compound
	case *types.Pointer:
		elem := x.Underlying().(*types.Pointer).Elem()
		return t.computeBasicType(elem)
	case *types.Basic:
		return t.whichTypesBasic(x)
	}
	return Nil
}

func (t *TableTree) whichTypesBasic(x *types.Basic) BasicType {
	switch x.Kind() {
	case types.Int:
		return Int
	case types.Int8:
		return Int8
	case types.Int16:
		return Int16
	case types.Int32:
		return Int32
	case types.Int64:
		return Int64
	case types.Uint:
		return Uint
	case types.Uint8:
		return Uint8
	case types.Uint16:
		return Uint16
	case types.Uint32:
		return Uint32
	case types.Uint64:
		return Uint64
	case types.Float32:
		return Float32
	case types.Float64:
		return Float64
	case types.Complex64:
		return Complex64
	case types.Complex128:
		return Complex128
	case types.String:
		return String
	case types.Bool:
		return Bool
	default:
		return Nil
	}
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
	if t.childNode != nil {
		m.PushBack(t.childNode.value)
		return t.childNode.traverseHelper(m)
	}
	return m
}
