package commands_test

import (
	"go/types"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kineticengines/gorm-migrations/pkg/commands"
	"github.com/stretchr/testify/require"
)

func TestSimpleTableTree_1(t *testing.T) {
	tr := new(commands.TableTree)
	require.Nil(t, tr.ChildNode)
	require.Nil(t, tr.Value)
}

func TestSimpleTableTree_2(t *testing.T) {
	dir, _ := os.Getwd()
	path := strings.Split(dir, "/commands")[0]
	path = filepath.Join(path, "definitions/testdata/simple_account.go")
	pkg, err := commands.ReadModelsFromPath(path)
	require.Nil(t, err)
	require.NotNil(t, pkg)

	// Find all named types on file level.
	var allNamed []*types.Named
	for _, name := range pkg.Scope().Names() {
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {
			allNamed = append(allNamed, obj.Type().(*types.Named))
		}
	}

	require.Equal(t, 1, len(allNamed))

	namedType := allNamed[0]
	obj := namedType.Underlying().(*types.Struct)

	tr := new(commands.TableTree)
	require.Nil(t, tr.ChildNode)
	require.Nil(t, tr.Value)

	tr.AddNodes(obj)

	nodes := tr.Traverse()
	require.Equal(t, nodes.Len(), 11)
}

func TestComplexTableTree(t *testing.T) {
	dir, _ := os.Getwd()
	path := strings.Split(dir, "/commands")[0]
	path = filepath.Join(path, "definitions/testdata/complex_account1.go")

	pkg, err := commands.ReadModelsFromPath(path)
	require.Nil(t, err)
	require.NotNil(t, pkg)

	// Find all named types on file level.
	var allNamed []*types.Named
	for _, name := range pkg.Scope().Names() {
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {
			allNamed = append(allNamed, obj.Type().(*types.Named))
		}
	}

	require.Equal(t, 2, len(allNamed))

	namedType := allNamed[0]
	obj := namedType.Underlying().(*types.Struct)

	tr := new(commands.TableTree)
	require.Nil(t, tr.ChildNode)
	require.Nil(t, tr.Value)

	tr.AddNodes(obj)

	nodes := tr.Traverse()
	require.Equal(t, 2, nodes.Len())
}

func TestComplexTableTree_2(t *testing.T) {
	dir, _ := os.Getwd()
	path := strings.Split(dir, "/commands")[0]
	path = filepath.Join(path, "definitions/testdata/complex_account2.go")

	pkg, err := commands.ReadModelsFromPath(path)
	require.Nil(t, err)
	require.NotNil(t, pkg)

	// Find all named types on file level.
	var allNamed []*types.Named
	for _, name := range pkg.Scope().Names() {
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {
			allNamed = append(allNamed, obj.Type().(*types.Named))
		}
	}

	require.Equal(t, 2, len(allNamed))

	namedType := allNamed[0]
	obj := namedType.Underlying().(*types.Struct)

	tr := new(commands.TableTree)
	require.Nil(t, tr.ChildNode)
	require.Nil(t, tr.Value)

	tr.AddNodes(obj)

	nodes := tr.Traverse()

	require.Equal(t, 5, nodes.Len())
}

func TestComplexTableTree_3(t *testing.T) {
	dir, _ := os.Getwd()
	path := strings.Split(dir, "/commands")[0]
	path = filepath.Join(path, "definitions/testdata/complex_account3.go")

	pkg, err := commands.ReadModelsFromPath(path)
	require.Nil(t, err)
	require.NotNil(t, pkg)

	// Find all named types on file level.
	var allNamed []*types.Named
	for _, name := range pkg.Scope().Names() {
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {
			allNamed = append(allNamed, obj.Type().(*types.Named))
		}
	}

	require.Equal(t, 1, len(allNamed))

	namedType := allNamed[0]
	obj := namedType.Underlying().(*types.Struct)

	tr := new(commands.TableTree)
	require.Nil(t, tr.ChildNode)
	require.Nil(t, tr.Value)

	tr.AddNodes(obj)

	nodes := tr.Traverse()

	require.Equal(t, 2, nodes.Len())
}

func TestComplexTableTree_4(t *testing.T) {
	dir, _ := os.Getwd()
	path := strings.Split(dir, "/commands")[0]
	path = filepath.Join(path, "definitions/testdata/complex_account4.go")

	pkg, err := commands.ReadModelsFromPath(path)
	require.Nil(t, err)
	require.NotNil(t, pkg)

	// Find all named types on file level.
	var allNamed []*types.Named
	for _, name := range pkg.Scope().Names() {
		if obj, ok := pkg.Scope().Lookup(name).(*types.TypeName); ok {
			allNamed = append(allNamed, obj.Type().(*types.Named))
		}
	}

	require.Equal(t, 1, len(allNamed))

	namedType := allNamed[0]
	obj := namedType.Underlying().(*types.Struct)

	tr := new(commands.TableTree)
	require.Nil(t, tr.ChildNode)
	require.Nil(t, tr.Value)

	tr.AddNodes(obj)

	nodes := tr.Traverse()

	require.Equal(t, 9, nodes.Len())
}
