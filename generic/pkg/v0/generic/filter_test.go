package generic

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
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
	dummies := []mql.Dummy{
		{},
	}

	table := []struct {
		name     string
		filter   mql.DummyFilter
		expected int
	}{
		{
			name: "set:true",
			filter: mql.DummyFilter{
				Set: mql.Bool(true),
			},
			expected: 1,
		},
		{
			name: "set:false",
			filter: mql.DummyFilter{
				Set: mql.Bool(false),
			},
			expected: 0,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(dummies)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestNestedTypeFilter(t *testing.T) {
	dummies := []mql.Dummy{
		{
			UnionField: &mql.DummyUnion{},
		},
		{},
	}

	table := []struct {
		name     string
		filter   mql.DummyFilter
		expected int
	}{
		{
			name: "set:true",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					Set: mql.Bool(true),
				},
			},
			expected: 1,
		},
		{
			name: "set:false",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					Set: mql.Bool(false),
				},
			},
			expected: 1,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(dummies)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestStringFilter(t *testing.T) {
	dummies := []mql.Dummy{
		{},
		{
			StringField: mql.String("a"),
		},
		{
			StringField: mql.String("A"),
		},
		{
			StringField: mql.String("b"),
		},
		{
			StringField: mql.String("B"),
		},
		{
			StringField: mql.String("c"),
		},
		{
			StringField: mql.String("C"),
		},
		{
			StringField: mql.String("yxy"),
		},
		{
			StringField: mql.String("yXy"),
		},
	}

	table := []struct {
		name     string
		filter   mql.DummyFilter
		expected int
	}{
		{
			name: "set:true",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					Set: mql.Bool(true),
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "set:false",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					Set: mql.Bool(false),
				},
				UnionField: &mql.DummyUnionFilter{},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,is:a",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(false),
					Is:            mql.String("a"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,is: a",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(true),
					Is:            mql.String("a"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,not:a",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(false),
					Not:           mql.String("a"),
				},
			},
			expected: len(dummies) - 2,
		},
		{
			name: "caseSensitive:true,not:a",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(true),
					Not:           mql.String("a"),
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "caseSensitive:false,contains:x",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(false),
					Contains:      mql.String("x"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,contains:x",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(true),
					Contains:      mql.String("x"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notContains:x",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(false),
					NotContains:   mql.String("x"),
				},
			},
			expected: len(dummies) - 2,
		},
		{
			name: "caseSensitive:true,notContains:x",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(true),
					NotContains:   mql.String("x"),
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "caseSensitive:false,startsWith:yx",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(false),
					StartsWith:    mql.String("yx"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,startsWith:yx",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(true),
					StartsWith:    mql.String("yx"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notStartsWith:yx",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(false),
					NotStartsWith: mql.String("yx"),
				},
			},
			expected: len(dummies) - 2,
		},
		{
			name: "caseSensitive:true,notStartsWith:yx",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(true),
					NotStartsWith: mql.String("yx"),
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "caseSensitive:false,endsWith:xy",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(false),
					EndsWith:      mql.String("xy"),
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,endsWith:xy",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(true),
					EndsWith:      mql.String("xy"),
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notEndsWith:xy",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(false),
					NotEndsWith:   mql.String("xy"),
				},
			},
			expected: len(dummies) - 2,
		},
		{
			name: "caseSensitive:true,notEndsWith:xy",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(true),
					NotEndsWith:   mql.String("xy"),
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "caseSensitive:false,in:[a]",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(false),
					In:            []string{"a"},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,in:[a]",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(true),
					In:            []string{"a"},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notIn:[a]",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(false),
					NotIn:         []string{"a"},
				},
			},
			expected: len(dummies) - 2,
		},
		{
			name: "caseSensitive:true,notIn:[a]",
			filter: mql.DummyFilter{
				StringField: &mql.StringFilter{
					CaseSensitive: mql.Bool(true),
					NotIn:         []string{"a"},
				},
			},
			expected: len(dummies) - 1,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(dummies)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestNestedStringFilter(t *testing.T) {
	dummies := []mql.Dummy{
		{
			UnionField: &mql.DummyUnion{},
		},
		{
			UnionField: &mql.DummyUnion{
				StringField: mql.String("a"),
			},
		},
		{
			UnionField: &mql.DummyUnion{
				StringField: mql.String("A"),
			},
		},
		{
			UnionField: &mql.DummyUnion{
				StringField: mql.String("b"),
			},
		},
		{
			UnionField: &mql.DummyUnion{
				StringField: mql.String("B"),
			},
		},
		{
			UnionField: &mql.DummyUnion{
				StringField: mql.String("c"),
			},
		},
		{
			UnionField: &mql.DummyUnion{
				StringField: mql.String("C"),
			},
		},
		{
			UnionField: &mql.DummyUnion{
				StringField: mql.String("yxy"),
			},
		},
		{
			UnionField: &mql.DummyUnion{
				StringField: mql.String("yXy"),
			},
		},
	}

	table := []struct {
		name     string
		filter   mql.DummyFilter
		expected int
	}{
		{
			name: "set:true",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						Set: mql.Bool(true),
					},
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "set:false",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						Set: mql.Bool(false),
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,is:a",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(false),
						Is:            mql.String("a"),
					},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,is: a",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(true),
						Is:            mql.String("a"),
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,not:a",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(false),
						Not:           mql.String("a"),
					},
				},
			},
			expected: len(dummies) - 2,
		},
		{
			name: "caseSensitive:true,not:a",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(true),
						Not:           mql.String("a"),
					},
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "caseSensitive:false,contains:x",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(false),
						Contains:      mql.String("x"),
					},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,contains:x",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(true),
						Contains:      mql.String("x"),
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notContains:x",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(false),
						NotContains:   mql.String("x"),
					},
				},
			},
			expected: len(dummies) - 2,
		},
		{
			name: "caseSensitive:true,notContains:x",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(true),
						NotContains:   mql.String("x"),
					},
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "caseSensitive:false,startsWith:yx",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(false),
						StartsWith:    mql.String("yx"),
					},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,startsWith:yx",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(true),
						StartsWith:    mql.String("yx"),
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notStartsWith:yx",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(false),
						NotStartsWith: mql.String("yx"),
					},
				},
			},
			expected: len(dummies) - 2,
		},
		{
			name: "caseSensitive:true,notStartsWith:yx",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(true),
						NotStartsWith: mql.String("yx"),
					},
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "caseSensitive:false,endsWith:xy",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(false),
						EndsWith:      mql.String("xy"),
					},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,endsWith:xy",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(true),
						EndsWith:      mql.String("xy"),
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notEndsWith:xy",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(false),
						NotEndsWith:   mql.String("xy"),
					},
				},
			},
			expected: len(dummies) - 2,
		},
		{
			name: "caseSensitive:true,notEndsWith:xy",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(true),
						NotEndsWith:   mql.String("xy"),
					},
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "caseSensitive:false,in:[a]",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(false),
						In:            []string{"a"},
					},
				},
			},
			expected: 2,
		},
		{
			name: "caseSensitive:true,in:[a]",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(true),
						In:            []string{"a"},
					},
				},
			},
			expected: 1,
		},
		{
			name: "caseSensitive:false,notIn:[a]",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(false),
						NotIn:         []string{"a"},
					},
				},
			},
			expected: len(dummies) - 2,
		},
		{
			name: "caseSensitive:true,notIn:[a]",
			filter: mql.DummyFilter{
				UnionField: &mql.DummyUnionFilter{
					StringField: &mql.StringFilter{
						CaseSensitive: mql.Bool(true),
						NotIn:         []string{"a"},
					},
				},
			},
			expected: len(dummies) - 1,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(dummies)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestEnumFilter(t *testing.T) {
	dummies := []mql.Dummy{
		{},
		{
			EnumField: mql.String(mql.DummyKind.Red),
		},
		{
			EnumField: mql.String(mql.DummyKind.Blue),
		},
		{
			EnumField: mql.String(mql.DummyKind.Green),
		},
	}

	table := []struct {
		name     string
		filter   mql.DummyFilter
		expected int
	}{
		{
			name: "set:true",
			filter: mql.DummyFilter{
				EnumField: &mql.EnumFilter{
					Set: mql.Bool(true),
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "set:false",
			filter: mql.DummyFilter{
				EnumField: &mql.EnumFilter{
					Set: mql.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "is:red",
			filter: mql.DummyFilter{
				EnumField: &mql.EnumFilter{
					Is: mql.String(mql.DummyKind.Red),
				},
			},
			expected: 1,
		},
		{
			name: "not:red",
			filter: mql.DummyFilter{
				EnumField: &mql.EnumFilter{
					Not: mql.String(mql.DummyKind.Red),
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "in:[red,blue]",
			filter: mql.DummyFilter{
				EnumField: &mql.EnumFilter{
					In: []string{mql.DummyKind.Red, mql.DummyKind.Blue},
				},
			},
			expected: 2,
		},
		{
			name: "notIn:[red,blue]",
			filter: mql.DummyFilter{
				EnumField: &mql.EnumFilter{
					NotIn: []string{mql.DummyKind.Red, mql.DummyKind.Blue},
				},
			},
			expected: len(dummies) - 2,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(dummies)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestInt32Filter(t *testing.T) {
	dummies := []mql.Dummy{
		{},
		{
			Int32Field: mql.Int32(0),
		},
		{
			Int32Field: mql.Int32(1),
		},
		{
			Int32Field: mql.Int32(2),
		},
	}

	table := []struct {
		name     string
		filter   mql.DummyFilter
		expected int
	}{
		{
			name: "set:false",
			filter: mql.DummyFilter{
				Int32Field: &mql.Int32Filter{
					Set: mql.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "set:true",
			filter: mql.DummyFilter{
				Int32Field: &mql.Int32Filter{
					Set: mql.Bool(true),
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "is:1",
			filter: mql.DummyFilter{
				Int32Field: &mql.Int32Filter{
					Is: mql.Int32(1),
				},
			},
			expected: 1,
		},
		{
			name: "not:1",
			filter: mql.DummyFilter{
				Int32Field: &mql.Int32Filter{
					Not: mql.Int32(1),
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "lt:1",
			filter: mql.DummyFilter{
				Int32Field: &mql.Int32Filter{
					Lt: mql.Int32(1),
				},
			},
			expected: 1,
		},
		{
			name: "lte:1",
			filter: mql.DummyFilter{
				Int32Field: &mql.Int32Filter{
					Lte: mql.Int32(1),
				},
			},
			expected: 2,
		},
		{
			name: "gt:1",
			filter: mql.DummyFilter{
				Int32Field: &mql.Int32Filter{
					Gt: mql.Int32(1),
				},
			},
			expected: 1,
		},
		{
			name: "gte:1",
			filter: mql.DummyFilter{
				Int32Field: &mql.Int32Filter{
					Gte: mql.Int32(1),
				},
			},
			expected: 2,
		},
		{
			name: "in:[1,2]",
			filter: mql.DummyFilter{
				Int32Field: &mql.Int32Filter{
					In: []int32{1, 2},
				},
			},
			expected: 2,
		},
		{
			name: "notIn:[1,2]",
			filter: mql.DummyFilter{
				Int32Field: &mql.Int32Filter{
					In: []int32{1, 2},
				},
			},
			expected: len(dummies) - 2,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(dummies)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestFloat64Filter(t *testing.T) {
	dummies := []mql.Dummy{
		{},
		{
			Float64Field: mql.Float64(0),
		},
		{
			Float64Field: mql.Float64(1),
		},
		{
			Float64Field: mql.Float64(2),
		},
	}

	table := []struct {
		name     string
		filter   mql.DummyFilter
		expected int
	}{
		{
			name: "set:false",
			filter: mql.DummyFilter{
				Float64Field: &mql.Float64Filter{
					Set: mql.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "set:true",
			filter: mql.DummyFilter{
				Float64Field: &mql.Float64Filter{
					Set: mql.Bool(true),
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "is:1",
			filter: mql.DummyFilter{
				Float64Field: &mql.Float64Filter{
					Is: mql.Float64(1),
				},
			},
			expected: 1,
		},
		{
			name: "not:1",
			filter: mql.DummyFilter{
				Float64Field: &mql.Float64Filter{
					Not: mql.Float64(1),
				},
			},
			expected: len(dummies) - 1,
		},
		{
			name: "lt:1",
			filter: mql.DummyFilter{
				Float64Field: &mql.Float64Filter{
					Lt: mql.Float64(1),
				},
			},
			expected: 1,
		},
		{
			name: "lte:1",
			filter: mql.DummyFilter{
				Float64Field: &mql.Float64Filter{
					Lte: mql.Float64(1),
				},
			},
			expected: 2,
		},
		{
			name: "gt:1",
			filter: mql.DummyFilter{
				Float64Field: &mql.Float64Filter{
					Gt: mql.Float64(1),
				},
			},
			expected: 1,
		},
		{
			name: "gte:1",
			filter: mql.DummyFilter{
				Float64Field: &mql.Float64Filter{
					Gte: mql.Float64(1),
				},
			},
			expected: 2,
		},
		{
			name: "in:[1,2]",
			filter: mql.DummyFilter{
				Float64Field: &mql.Float64Filter{
					In: []float64{1, 2},
				},
			},
			expected: 2,
		},
		{
			name: "notIn:[1,2]",
			filter: mql.DummyFilter{
				Float64Field: &mql.Float64Filter{
					In: []float64{1, 2},
				},
			},
			expected: len(dummies) - 2,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(dummies)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

func TestBoolFilter(t *testing.T) {
	dummies := []mql.Dummy{
		{},
		{
			BoolField: mql.Bool(true),
		},
		{
			BoolField: mql.Bool(false),
		},
	}

	table := []struct {
		name     string
		filter   mql.DummyFilter
		expected int
	}{
		{
			name: "set:false",
			filter: mql.DummyFilter{
				BoolField: &mql.BoolFilter{
					Set: mql.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "set:true",
			filter: mql.DummyFilter{
				BoolField: &mql.BoolFilter{
					Set: mql.Bool(true),
				},
			},
			expected: 2,
		},
		{
			name: "is:true",
			filter: mql.DummyFilter{
				BoolField: &mql.BoolFilter{
					Is: mql.Bool(true),
				},
			},
			expected: 1,
		},
		{
			name: "is:false",
			filter: mql.DummyFilter{
				BoolField: &mql.BoolFilter{
					Is: mql.Bool(false),
				},
			},
			expected: 1,
		},
		{
			name: "not:true",
			filter: mql.DummyFilter{
				BoolField: &mql.BoolFilter{
					Not: mql.Bool(true),
				},
			},
			expected: 2,
		},
		{
			name: "not:false",
			filter: mql.DummyFilter{
				BoolField: &mql.BoolFilter{
					Not: mql.Bool(false),
				},
			},
			expected: 2,
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			gSlice := f.MustFromStructs(dummies)

			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))

			assert.Equal(t, c.expected, len(gSlice.Get()))
		})
	}

	return
}

//func TestAndFilter(t *testing.T) {
//	dummies := []mql.Dummy{
//		{
//			BoolField: mql.Bool(true),
//			Int32Field: mql.Int32(0),
//		},
//		{
//			BoolField: mql.Bool(true),
//			Int32Field: mql.Int32(1),
//		},
//	}
//
//	table := []struct {
//		name     string
//		filter   mql.DummyFilter
//		expected int
//	}{
//		{
//			name: "boolField.is:true && int32Field.is:0",
//			filter: mql.DummyFilter{
//				And: []mql.DummyFilter{
//					{
//						BoolField: &mql.BoolFilter{
//							Is: mql.Bool(true),
//						},
//					},
//					{
//						Int32Field: &mql.Int32Filter{
//							Is: mql.Int32(0),
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
//			gSlice := f.MustFromStructs(dummies)
//
//			gSlice = gSlice.Filter(false, f.MustFromStruct(c.filter))
//
//			assert.Equal(t, c.expected, len(gSlice.Get()))
//		})
//	}
//
//	return
//}
