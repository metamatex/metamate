package generic

import (
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/typenames"
	"testing"

	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/stretchr/testify/assert"

	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
)

var w = sdk.Whatever{
	StringField:  ptr.String("a"),
	Int32Field:   ptr.Int32(1),
	Float64Field: ptr.Float64(1),
	BoolField:    ptr.Bool(true),
	EnumField:    ptr.String(sdk.WhateverKind.Red),
	UnionField: &sdk.WhateverUnion{
		StringField:  ptr.String("a"),
		Int32Field:   ptr.Int32(1),
		Float64Field: ptr.Float64(1),
		BoolField:    ptr.Bool(true),
	},
	StringList:  []string{"a"},
	Int32List:   []int32{1},
	Float64List: []float64{1},
	BoolList:    []bool{true},
	EnumList:    []string{sdk.WhateverKind.Red},
	UnionList: []sdk.WhateverUnion{
		{
			StringField:  ptr.String("a"),
			Int32Field:   ptr.Int32(1),
			Float64Field: ptr.Float64(1),
			BoolField:    ptr.Bool(true),
		},
		{
			StringField:  ptr.String("b"),
			Int32Field:   ptr.Int32(2),
			Float64Field: ptr.Float64(2),
			BoolField:    ptr.Bool(true),
		},
	},
}

func TestGeneric(t *testing.T, rn *graph.RootNode, f Factory) {
	t.Parallel()

	t.Run("TestGenericSetGet", func(t *testing.T) {
		t.Parallel()

		TestGenericSetGet(t, rn, f)
	})

	t.Run("TestGenericNestedSetGet", func(t *testing.T) {
		t.Parallel()

		TestGenericNestedSetGet(t, rn, f)
	})
}

func TestGenericSetGet(t *testing.T, rn *graph.RootNode, factory Factory) {
	g := factory.New(rn.Types.MustByName(typenames.Whatever))

	s := "a"
	err := g.SetString([]string{fieldnames.StringField}, s)
	if err != nil {
		t.Error(err)
	}
	s0, ok := g.String(fieldnames.StringField)
	assert.Equal(t, true, ok)
	assert.Equal(t, s, s0)

	i := int32(1)
	err = g.SetInt32([]string{fieldnames.Int32Field}, i)
	if err != nil {
		t.Error(err)
	}
	i0, ok := g.Int32(fieldnames.Int32Field)
	assert.Equal(t, true, ok)
	assert.Equal(t, i, i0)

	f := float64(1)
	err = g.SetFloat64([]string{fieldnames.Float64Field}, f)
	if err != nil {
		t.Error(err)
	}
	f0, ok := g.Float64(fieldnames.Float64Field)
	assert.Equal(t, true, ok)
	assert.Equal(t, f, f0)

	b := true
	err = g.SetBool([]string{fieldnames.BoolField}, b)
	if err != nil {
		t.Error(err)
	}
	b0, ok := g.Bool(fieldnames.BoolField)
	assert.Equal(t, true, ok)
	assert.Equal(t, b, b0)
}

func TestGenericNestedSetGet(t *testing.T, rn *graph.RootNode, factory Factory) {
	g := factory.New(rn.Types.MustByName(typenames.Whatever))

	s := "a"
	err := g.SetString([]string{fieldnames.UnionField, fieldnames.StringField}, s)
	if err != nil {
		t.Error(err)
	}

	g0, ok := g.Generic(fieldnames.UnionField)
	if assert.Equal(t, true, ok) {
		assert.Equal(t, typenames.WhateverUnion, g0.Type().Name())
	}

	s0, ok := g.String(fieldnames.UnionField, fieldnames.StringField)
	assert.Equal(t, true, ok)
	assert.Equal(t, s, s0)

	i := int32(1)
	err = g.SetInt32([]string{fieldnames.UnionField, fieldnames.Int32Field}, i)
	if err != nil {
		t.Error(err)
	}

	i0, ok := g.Int32(fieldnames.UnionField, fieldnames.Int32Field)
	assert.Equal(t, true, ok)
	assert.Equal(t, i, i0)
}

func TestGenericToStringInterfaceMap(t *testing.T, rn *graph.RootNode, factory Factory) {
	g, err := factory.FromStruct(w)
	if err != nil {
		t.Error(err)

		return
	}

	m := g.ToStringInterfaceMap()

	println(m)
}

func TestGenericFlatten(t *testing.T, rn *graph.RootNode, factory Factory) {
	g, err := factory.FromStruct(w)
	if err != nil {
		t.Error(err)

		return
	}

	m, err := g.Flatten("_")
	if err != nil {
		t.Error(err)

		return
	}

	println(m)
}
