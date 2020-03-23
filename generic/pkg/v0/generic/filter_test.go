package generic

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/sdk"

	"github.com/stretchr/testify/assert"
	"testing"
)

var rn *graph.RootNode
var f Factory

func init() {
	var err error
	rn, err = asg.New()
	if err != nil {
		panic(err)
	}

	f = NewFactory(rn)
}

func TestTypeFilter(t *testing.T) {
	whatevers := []sdk.Whatever{
		{
		},
	}

	table := []struct {
		name     string
		filter   sdk.WhateverFilter
		expected int
	}{
		{
			name: "set:true",
			filter: sdk.WhateverFilter{
				Set: sdk.Bool(true),
			},
			expected: 1,
		},
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				Set: sdk.Bool(false),
			},
			expected: 0,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestNestedTypeFilter(t *testing.T) {
	whatevers := []sdk.Whatever{
		{
			UnionField: &sdk.WhateverUnion{},
		},
		{
		},
	}

	table := []struct {
		name     string
		filter   sdk.WhateverFilter
		expected int
	}{
		{
			name: "set:true",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					Set: sdk.Bool(true),
				},
			},
			expected: 1,
		},
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					Set: sdk.Bool(false),
				},
			},
			expected: 1,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestStringFilter(t *testing.T) {
	whatevers := []sdk.Whatever{
		{
		},
		{
			StringField: sdk.String("a"),
		},
		{
			StringField: sdk.String("A"),
		},
		{
			StringField: sdk.String("b"),
		},
		{
			StringField: sdk.String("B"),
		},
		{
			StringField: sdk.String("c"),
		},
		{
			StringField: sdk.String("C"),
		},
		{
			StringField: sdk.String("yxy"),
		},
		{
			StringField: sdk.String("yXy"),
		},
	}

	table := []struct {
		name     string
		filter   sdk.WhateverFilter
		expected int
	}{
		{
			name: "set:true",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					Set: sdk.Bool(true),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					Set: sdk.Bool(false),
				},
				UnionField: &sdk.WhateverUnionFilter{},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,is:a",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(false),
					Is:            sdk.String("a"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,is: a",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(true),
					Is:            sdk.String("a"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,not:a",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(false),
					Not:           sdk.String("a"),
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,not:a",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(true),
					Not:           sdk.String("a"),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,contains:x",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(false),
					Contains:      sdk.String("x"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,contains:x",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(true),
					Contains:      sdk.String("x"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notContains:x",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(false),
					NotContains:   sdk.String("x"),
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notContains:x",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(true),
					NotContains:   sdk.String("x"),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,startsWith:yx",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(false),
					StartsWith:    sdk.String("yx"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,startsWith:yx",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(true),
					StartsWith:    sdk.String("yx"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notStartsWith:yx",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(false),
					NotStartsWith: sdk.String("yx"),
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notStartsWith:yx",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(true),
					NotStartsWith: sdk.String("yx"),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,endsWith:xy",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(false),
					EndsWith:      sdk.String("xy"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,endsWith:xy",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(true),
					EndsWith:      sdk.String("xy"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notEndsWith:xy",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(false),
					NotEndsWith:   sdk.String("xy"),
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notEndsWith:xy",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(true),
					NotEndsWith:   sdk.String("xy"),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,in:[a]",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(false),
					In:            []string{"a"},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,in:[a]",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(true),
					In:            []string{"a"},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notIn:[a]",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(false),
					NotIn:         []string{"a"},
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notIn:[a]",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: sdk.Bool(true),
					NotIn:         []string{"a"},
				},
			},
			expected: len(whatevers) - 1,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestNestedStringFilter(t *testing.T) {
	whatevers := []sdk.Whatever{
		{
			UnionField: &sdk.WhateverUnion{
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: sdk.String("a"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: sdk.String("A"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: sdk.String("b"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: sdk.String("B"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: sdk.String("c"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: sdk.String("C"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: sdk.String("yxy"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: sdk.String("yXy"),
			},
		},
	}

	table := []struct {
		name     string
		filter   sdk.WhateverFilter
		expected int
	}{
		{
			name: "set:true",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						Set: sdk.Bool(true),
					},
				},
			},
			expected: len(whatevers) -1,
		},
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						Set: sdk.Bool(false),
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,is:a",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(false),
						Is:            sdk.String("a"),
					},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,is: a",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(true),
						Is:            sdk.String("a"),
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,not:a",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(false),
						Not:           sdk.String("a"),
					},
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,not:a",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(true),
						Not:           sdk.String("a"),
					},
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,contains:x",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(false),
						Contains:      sdk.String("x"),
					},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,contains:x",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(true),
						Contains:      sdk.String("x"),
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notContains:x",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(false),
						NotContains:   sdk.String("x"),
					},
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notContains:x",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(true),
						NotContains:   sdk.String("x"),
					},
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,startsWith:yx",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(false),
						StartsWith:    sdk.String("yx"),
					},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,startsWith:yx",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(true),
						StartsWith:    sdk.String("yx"),
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notStartsWith:yx",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(false),
						NotStartsWith: sdk.String("yx"),
					},
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notStartsWith:yx",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(true),
						NotStartsWith: sdk.String("yx"),
					},
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,endsWith:xy",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(false),
						EndsWith:      sdk.String("xy"),
					},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,endsWith:xy",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(true),
						EndsWith:      sdk.String("xy"),
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notEndsWith:xy",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(false),
						NotEndsWith:   sdk.String("xy"),
					},
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notEndsWith:xy",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(true),
						NotEndsWith:   sdk.String("xy"),
					},
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,in:[a]",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(false),
						In:            []string{"a"},
					},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,in:[a]",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(true),
						In:            []string{"a"},
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notIn:[a]",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(false),
						NotIn:         []string{"a"},
					},
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notIn:[a]",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					StringField: &sdk.StringFilter{
						CaseSensitive: sdk.Bool(true),
						NotIn:         []string{"a"},
					},
				},
			},
			expected: len(whatevers) - 1,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestEnumFilter(t *testing.T) {
	whatevers := []sdk.Whatever{
		{
		},
		{
			EnumField: sdk.String(sdk.WhateverKind.Red),
		},
		{
			EnumField: sdk.String(sdk.WhateverKind.Blue),
		},
		{
			EnumField: sdk.String(sdk.WhateverKind.Green),
		},
	}

	table := []struct {
		name     string
		filter   sdk.WhateverFilter
		expected int
	}{
		{
			name: "set:true",
			filter: sdk.WhateverFilter{
				EnumField: &sdk.EnumFilter{
					Set: sdk.Bool(true),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				EnumField: &sdk.EnumFilter{
					Set: sdk.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "is:red",
			filter: sdk.WhateverFilter{
				EnumField: &sdk.EnumFilter{
					Is: sdk.String(sdk.WhateverKind.Red),
				},
			},
			expected: 1,
		},
		{
			name: "not:red",
			filter: sdk.WhateverFilter{
				EnumField: &sdk.EnumFilter{
					Not: sdk.String(sdk.WhateverKind.Red),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "in:[red,blue]",
			filter: sdk.WhateverFilter{
				EnumField: &sdk.EnumFilter{
					In: []string{sdk.WhateverKind.Red, sdk.WhateverKind.Blue},
				},
			},
			expected: 2,
		},
		{
			name: "notIn:[red,blue]",
			filter: sdk.WhateverFilter{
				EnumField: &sdk.EnumFilter{
					NotIn: []string{sdk.WhateverKind.Red, sdk.WhateverKind.Blue},
				},
			},
			expected: len(whatevers) - 2,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestInt32Filter(t *testing.T) {
	whatevers := []sdk.Whatever{
		{
		},
		{
			Int32Field: sdk.Int32(0),
		},
		{
			Int32Field: sdk.Int32(1),
		},
		{
			Int32Field: sdk.Int32(2),
		},
	}

	table := []struct {
		name     string
		filter   sdk.WhateverFilter
		expected int
	}{
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Set: sdk.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "set:true",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Set: sdk.Bool(true),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "is:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Is: sdk.Int32(1),
				},
			},
			expected: 1,
		},
		{
			name: "not:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Not: sdk.Int32(1),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "lt:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Lt: sdk.Int32(1),
				},
			},
			expected: 1,
		},
		{
			name: "lte:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Lte: sdk.Int32(1),
				},
			},
			expected: 2,
		},
		{
			name: "gt:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Gt: sdk.Int32(1),
				},
			},
			expected: 1,
		},
		{
			name: "gte:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Gte: sdk.Int32(1),
				},
			},
			expected: 2,
		},
		{
			name: "in:[1,2]",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					In: []int32{1, 2},
				},
			},
			expected: 2,
		},
		{
			name: "notIn:[1,2]",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					In: []int32{1, 2},
				},
			},
			expected: len(whatevers) - 2,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestFloat64Filter(t *testing.T) {
	whatevers := []sdk.Whatever{
		{
		},
		{
			Float64Field: sdk.Float64(0),
		},
		{
			Float64Field: sdk.Float64(1),
		},
		{
			Float64Field: sdk.Float64(2),
		},
	}

	table := []struct {
		name     string
		filter   sdk.WhateverFilter
		expected int
	}{
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Set: sdk.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "set:true",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Set: sdk.Bool(true),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "is:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Is: sdk.Float64(1),
				},
			},
			expected: 1,
		},
		{
			name: "not:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Not: sdk.Float64(1),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "lt:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Lt: sdk.Float64(1),
				},
			},
			expected: 1,
		},
		{
			name: "lte:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Lte: sdk.Float64(1),
				},
			},
			expected: 2,
		},
		{
			name: "gt:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Gt: sdk.Float64(1),
				},
			},
			expected: 1,
		},
		{
			name: "gte:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Gte: sdk.Float64(1),
				},
			},
			expected: 2,
		},
		{
			name: "in:[1,2]",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					In: []float64{1, 2},
				},
			},
			expected: 2,
		},
		{
			name: "notIn:[1,2]",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					In: []float64{1, 2},
				},
			},
			expected: len(whatevers) - 2,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestBoolFilter(t *testing.T) {
	whatevers := []sdk.Whatever{
		{
		},
		{
			BoolField: sdk.Bool(true),
		},
		{
			BoolField: sdk.Bool(false),
		},
	}

	table := []struct {
		name     string
		filter   sdk.WhateverFilter
		expected int
	}{
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				BoolField: &sdk.BoolFilter{
					Set: sdk.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "set:true",
			filter: sdk.WhateverFilter{
				BoolField: &sdk.BoolFilter{
					Set: sdk.Bool(true),
				},
			},
			expected: 2,
		},
		{
			name: "is:true",
			filter: sdk.WhateverFilter{
				BoolField: &sdk.BoolFilter{
					Is: sdk.Bool(true),
				},
			},
			expected: 1,
		},
		{
			name: "is:false",
			filter: sdk.WhateverFilter{
				BoolField: &sdk.BoolFilter{
					Is: sdk.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "not:true",
			filter: sdk.WhateverFilter{
				BoolField: &sdk.BoolFilter{
					Not: sdk.Bool(true),
				},
			},
			expected: 2,
		},
		{
			name: "not:false",
			filter: sdk.WhateverFilter{
				BoolField: &sdk.BoolFilter{
					Not: sdk.Bool(false),
				},
			},
			expected: 2,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

//func TestAndFilter(t *testing.T) {
//	whatevers := []sdk.Whatever{
//		{
//			BoolField: sdk.Bool(true),
//			Int32Field: sdk.Int32(0),
//		},
//		{
//			BoolField: sdk.Bool(true),
//			Int32Field: sdk.Int32(1),
//		},
//	}
//
//	table := []struct {
//		name     string
//		filter   sdk.WhateverFilter
//		expected int
//	}{
//		{
//			name: "boolField.is:true && int32Field.is:0",
//			filter: sdk.WhateverFilter{
//				And: []sdk.WhateverFilter{
//					{
//						BoolField: &sdk.BoolFilter{
//							Is: sdk.Bool(true),
//						},
//					},
//					{
//						Int32Field: &sdk.Int32Filter{
//							Is: sdk.Int32(0),
//						},
//					},
//				},
//			},
//			expected: 1,
//		},
//	}
//
//	for _, c := range table {
//		t.Run(c.name, func(t *testing.T) {
//			gSlice := f.MustFromStructs(whatevers)
//
//			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))
//
//			assert.Equal(t, c.expected, len(gSlice.Get()))
//		})
//	}
//
//	return
//}