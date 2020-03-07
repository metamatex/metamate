package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var root = NewRoot()

func Test(t *testing.T) {
	root.Types.ById("querymode").Edges.Enums.Dependencies().Each(func(tn *EnumNode) {
		println(tn.Name())
	})

	//getEdges(root.Types.ByName(Person).Edges)

	//errs := root.Validate()
	//spew.Dump(errs)

	//spew.Dump(getFields(root.Types.ByName(Person)))
	//
	//spew.Dump(getFields(root.Types.ByName(Status)))

	//root.Types.ByName(Status).ComposeFields().BroadcastPrint()

	//root.Types.ByName(Person).ComposeFields().BroadcastPrint()
}

func TestGetEdgeMaps(t *testing.T) {
	t.Parallel()

	root := NewRoot()

	_, manyEdges := GetEdgeMaps(root.Types.ByName(Whatever).GetEdges())

	if assert.Contains(t, manyEdges, "Types") {
		assert.Contains(t, manyEdges["Types"], "Dependencies")
	}

	_, manyEdges = GetEdgeMaps(root.Types.ByName(Whatever).GetEdges(), TypeDependenciesTypes)

	assert.NotContains(t, manyEdges, "Types")
}

func TestFilter(t *testing.T) {
	name := root.Types.Filter(Filter{
		Names: &NamesSubset{
			Or: []string{Whatever},
		},
	}).ByName(Whatever).Name()

	assert.Equal(t, Whatever, name)
}
