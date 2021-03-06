package commands_test

import (
	"fmt"
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
	require.Nil(t, tr.LeftNode)
	require.Nil(t, tr.RightNode)
	require.Nil(t, tr.Value)
}

func TestSimpleTableTree_2(t *testing.T) {
	dir, _ := os.Getwd()
	path := strings.Split(dir, "/commands")[0]
	fmt.Println(dir)
	path = filepath.Join(path, "definitions/testdata/simple_account.go")
	fmt.Println(path)
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
	require.Nil(t, tr.LeftNode)
	require.Nil(t, tr.RightNode)
	require.Nil(t, tr.Value)

	tr.AddNodes(obj)

	nodes := tr.Traverse()
	require.Equal(t, nodes.Len(), 11)
}
