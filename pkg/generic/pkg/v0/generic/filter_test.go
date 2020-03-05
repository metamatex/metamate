package generic

import (
	"github.com/metamatex/asg/pkg/v0/asg"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/gen/v0/sdk/utils/ptr"
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
				Set: ptr.Bool(true),
			},
			expected: 1,
		},
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				Set: ptr.Bool(false),
			},
			expected: 0,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice.Filter(false, f.MustFromStruct(c.filter))

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
					Set: ptr.Bool(true),
				},
			},
			expected: 1,
		},
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				UnionField: &sdk.WhateverUnionFilter{
					Set: ptr.Bool(false),
				},
			},
			expected: 1,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice.Filter(false, f.MustFromStruct(c.filter))

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
			StringField: ptr.String("a"),
		},
		{
			StringField: ptr.String("A"),
		},
		{
			StringField: ptr.String("b"),
		},
		{
			StringField: ptr.String("B"),
		},
		{
			StringField: ptr.String("c"),
		},
		{
			StringField: ptr.String("C"),
		},
		{
			StringField: ptr.String("yxy"),
		},
		{
			StringField: ptr.String("yXy"),
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
					Set: ptr.Bool(true),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					Set: ptr.Bool(false),
				},
				UnionField: &sdk.WhateverUnionFilter{},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,is:a",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(false),
					Is:            ptr.String("a"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,is: a",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(true),
					Is:            ptr.String("a"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,not:a",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(false),
					Not:           ptr.String("a"),
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,not:a",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(true),
					Not:           ptr.String("a"),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,contains:x",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(false),
					Contains:      ptr.String("x"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,contains:x",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(true),
					Contains:      ptr.String("x"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notContains:x",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(false),
					NotContains:   ptr.String("x"),
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notContains:x",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(true),
					NotContains:   ptr.String("x"),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,startsWith:yx",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(false),
					StartsWith:    ptr.String("yx"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,startsWith:yx",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(true),
					StartsWith:    ptr.String("yx"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notStartsWith:yx",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(false),
					NotStartsWith: ptr.String("yx"),
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notStartsWith:yx",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(true),
					NotStartsWith: ptr.String("yx"),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,endsWith:xy",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(false),
					EndsWith:      ptr.String("xy"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,endsWith:xy",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(true),
					EndsWith:      ptr.String("xy"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notEndsWith:xy",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(false),
					NotEndsWith:   ptr.String("xy"),
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notEndsWith:xy",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(true),
					NotEndsWith:   ptr.String("xy"),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "caseSensitive:false,in:[a]",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(false),
					In:            []string{"a"},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,in:[a]",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(true),
					In:            []string{"a"},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notIn:[a]",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(false),
					NotIn:         []string{"a"},
				},
			},
			expected: len(whatevers) - 2,
		},
		{
			name: "caseSensitive:true,notIn:[a]",
			filter: sdk.WhateverFilter{
				StringField: &sdk.StringFilter{
					CaseSensitive: ptr.Bool(true),
					NotIn:         []string{"a"},
				},
			},
			expected: len(whatevers) - 1,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice.Filter(false, f.MustFromStruct(c.filter))

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
				StringField: ptr.String("a"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: ptr.String("A"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: ptr.String("b"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: ptr.String("B"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: ptr.String("c"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: ptr.String("C"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: ptr.String("yxy"),
			},
		},
		{
			UnionField: &sdk.WhateverUnion{
				StringField: ptr.String("yXy"),
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
						Set: ptr.Bool(true),
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
						Set: ptr.Bool(false),
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
						CaseSensitive: ptr.Bool(false),
						Is:            ptr.String("a"),
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
						CaseSensitive: ptr.Bool(true),
						Is:            ptr.String("a"),
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
						CaseSensitive: ptr.Bool(false),
						Not:           ptr.String("a"),
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
						CaseSensitive: ptr.Bool(true),
						Not:           ptr.String("a"),
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
						CaseSensitive: ptr.Bool(false),
						Contains:      ptr.String("x"),
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
						CaseSensitive: ptr.Bool(true),
						Contains:      ptr.String("x"),
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
						CaseSensitive: ptr.Bool(false),
						NotContains:   ptr.String("x"),
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
						CaseSensitive: ptr.Bool(true),
						NotContains:   ptr.String("x"),
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
						CaseSensitive: ptr.Bool(false),
						StartsWith:    ptr.String("yx"),
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
						CaseSensitive: ptr.Bool(true),
						StartsWith:    ptr.String("yx"),
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
						CaseSensitive: ptr.Bool(false),
						NotStartsWith: ptr.String("yx"),
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
						CaseSensitive: ptr.Bool(true),
						NotStartsWith: ptr.String("yx"),
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
						CaseSensitive: ptr.Bool(false),
						EndsWith:      ptr.String("xy"),
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
						CaseSensitive: ptr.Bool(true),
						EndsWith:      ptr.String("xy"),
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
						CaseSensitive: ptr.Bool(false),
						NotEndsWith:   ptr.String("xy"),
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
						CaseSensitive: ptr.Bool(true),
						NotEndsWith:   ptr.String("xy"),
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
						CaseSensitive: ptr.Bool(false),
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
						CaseSensitive: ptr.Bool(true),
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
						CaseSensitive: ptr.Bool(false),
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
						CaseSensitive: ptr.Bool(true),
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

			gSlice.Filter(false, f.MustFromStruct(c.filter))

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
			EnumField: ptr.String(sdk.WhateverKind_RED),
		},
		{
			EnumField: ptr.String(sdk.WhateverKind_BLUE),
		},
		{
			EnumField: ptr.String(sdk.WhateverKind_GREEN),
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
					Set: ptr.Bool(true),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "set:false",
			filter: sdk.WhateverFilter{
				EnumField: &sdk.EnumFilter{
					Set: ptr.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "is:RED",
			filter: sdk.WhateverFilter{
				EnumField: &sdk.EnumFilter{
					Is: ptr.String(sdk.WhateverKind_RED),
				},
			},
			expected: 1,
		},
		{
			name: "not:RED",
			filter: sdk.WhateverFilter{
				EnumField: &sdk.EnumFilter{
					Not: ptr.String(sdk.WhateverKind_RED),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "in:[RED,BLUE]",
			filter: sdk.WhateverFilter{
				EnumField: &sdk.EnumFilter{
					In: []string{sdk.WhateverKind_RED, sdk.WhateverKind_BLUE},
				},
			},
			expected: 2,
		},
		{
			name: "notIn:[RED,BLUE]",
			filter: sdk.WhateverFilter{
				EnumField: &sdk.EnumFilter{
					NotIn: []string{sdk.WhateverKind_RED, sdk.WhateverKind_BLUE},
				},
			},
			expected: len(whatevers) - 2,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice.Filter(false, f.MustFromStruct(c.filter))

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
			Int32Field: ptr.Int32(0),
		},
		{
			Int32Field: ptr.Int32(1),
		},
		{
			Int32Field: ptr.Int32(2),
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
					Set: ptr.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "set:true",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Set: ptr.Bool(true),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "is:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Is: ptr.Int32(1),
				},
			},
			expected: 1,
		},
		{
			name: "not:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Not: ptr.Int32(1),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "lt:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Lt: ptr.Int32(1),
				},
			},
			expected: 1,
		},
		{
			name: "lte:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Lte: ptr.Int32(1),
				},
			},
			expected: 2,
		},
		{
			name: "gt:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Gt: ptr.Int32(1),
				},
			},
			expected: 1,
		},
		{
			name: "gte:1",
			filter: sdk.WhateverFilter{
				Int32Field: &sdk.Int32Filter{
					Gte: ptr.Int32(1),
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

			gSlice.Filter(false, f.MustFromStruct(c.filter))

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
			Float64Field: ptr.Float64(0),
		},
		{
			Float64Field: ptr.Float64(1),
		},
		{
			Float64Field: ptr.Float64(2),
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
					Set: ptr.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "set:true",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Set: ptr.Bool(true),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "is:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Is: ptr.Float64(1),
				},
			},
			expected: 1,
		},
		{
			name: "not:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Not: ptr.Float64(1),
				},
			},
			expected: len(whatevers) - 1,
		},
		{
			name: "lt:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Lt: ptr.Float64(1),
				},
			},
			expected: 1,
		},
		{
			name: "lte:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Lte: ptr.Float64(1),
				},
			},
			expected: 2,
		},
		{
			name: "gt:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Gt: ptr.Float64(1),
				},
			},
			expected: 1,
		},
		{
			name: "gte:1",
			filter: sdk.WhateverFilter{
				Float64Field: &sdk.Float64Filter{
					Gte: ptr.Float64(1),
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

			gSlice.Filter(false, f.MustFromStruct(c.filter))

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
			BoolField: ptr.Bool(true),
		},
		{
			BoolField: ptr.Bool(false),
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
					Set: ptr.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "set:true",
			filter: sdk.WhateverFilter{
				BoolField: &sdk.BoolFilter{
					Set: ptr.Bool(true),
				},
			},
			expected: 2,
		},
		{
			name: "is:true",
			filter: sdk.WhateverFilter{
				BoolField: &sdk.BoolFilter{
					Is: ptr.Bool(true),
				},
			},
			expected: 1,
		},
		{
			name: "is:false",
			filter: sdk.WhateverFilter{
				BoolField: &sdk.BoolFilter{
					Is: ptr.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "not:true",
			filter: sdk.WhateverFilter{
				BoolField: &sdk.BoolFilter{
					Not: ptr.Bool(true),
				},
			},
			expected: 2,
		},
		{
			name: "not:false",
			filter: sdk.WhateverFilter{
				BoolField: &sdk.BoolFilter{
					Not: ptr.Bool(false),
				},
			},
			expected: 2,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestAndFilter(t *testing.T) {
	whatevers := []sdk.Whatever{
		{
			BoolField: ptr.Bool(true),
			Int32Field: ptr.Int32(0),
		},
		{
			BoolField: ptr.Bool(true),
			Int32Field: ptr.Int32(1),
		},
	}

	table := []struct {
		name     string
		filter   sdk.WhateverFilter
		expected int
	}{
		{
			name: "boolField.is:true && int32Field.is:0",
			filter: sdk.WhateverFilter{
				And: []sdk.WhateverFilter{
					{
						BoolField: &sdk.BoolFilter{
							Is: ptr.Bool(true),
						},
					},
					{
						Int32Field: &sdk.Int32Filter{
							Is: ptr.Int32(0),
						},
					},
				},
			},
			expected: 1,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(whatevers)

			gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}