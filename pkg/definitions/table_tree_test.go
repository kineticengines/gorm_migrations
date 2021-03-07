package definitions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleTableTree_1(t *testing.T) {
	tr := new(TableTree)
	require.Nil(t, tr.childNode)
	require.Nil(t, tr.value)
}
