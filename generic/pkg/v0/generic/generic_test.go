package generic

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/stretchr/testify/assert"
	"testing"
)

var w = mql.Dummy{
	StringField:  mql.String("a"),
	Int32Field:   mql.Int32(1),
	Float64Field: mql.Float64(1),
	BoolField:    mql.Bool(true),
	EnumField:    mql.String(mql.DummyKind.Red),
	UnionField: &mql.DummyUnion{
		StringField:  mql.String("a"),
		Int32Field:   mql.Int32(1),
		Float64Field: mql.Float64(1),
		BoolField:    mql.Bool(true),
	},
	StringList:  []string{"a"},
	Int32List:   []int32{1},
	Float64List: []float64{1},
	BoolList:    []bool{true},
	EnumList:    []string{mql.DummyKind.Red},
	UnionList: []mql.DummyUnion{
		{
			StringField:  mql.String("a"),
			Int32Field:   mql.Int32(1),
			Float64Field: mql.Float64(1),
			BoolField:    mql.Bool(true),
		},
		{
			StringField:  mql.String("b"),
			Int32Field:   mql.Int32(2),
			Float64Field: mql.Float64(2),
			BoolField:    mql.Bool(true),
		},
	},
}

func TestGeneric(t *testing.T) {
	var err error
	rn, err = asg.New()
	if err != nil {
		panic(err)
	}

	f = NewFactory(rn)

	SetGet(t, rn, f)
	NestedSetGet(t, rn, f)
	GetHash(t, rn, f)
}

func SetGet(t *testing.T, rn *graph.RootNode, factory Factory) {
	name := "SetGet"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		g := factory.New(rn.Types.MustByName(typenames.Dummy))

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
	})

}

func NestedSetGet(t *testing.T, rn *graph.RootNode, factory Factory) {
	name := "NestedSetGet"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		g := factory.New(rn.Types.MustByName(typenames.Dummy))

		s := "a"
		err := g.SetString([]string{fieldnames.UnionField, fieldnames.StringField}, s)
		if err != nil {
			t.Error(err)
		}

		g0, ok := g.Generic(fieldnames.UnionField)
		if assert.Equal(t, true, ok) {
			assert.Equal(t, typenames.DummyUnion, g0.Type().Name())
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
	})
}

func ToStringInterfaceMap(t *testing.T, rn *graph.RootNode, factory Factory) {
	name := "ToStringInterfaceMap"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		g, err := factory.FromStruct(w)
		if err != nil {
			t.Error(err)

			return
		}

		m := g.ToStringInterfaceMap()

		println(m)
	})
}

func Flatten(t *testing.T, rn *graph.RootNode, factory Factory) {
	name := "Flatten"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

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
	})
}

func GetHash(t *testing.T, rn *graph.RootNode, factory Factory) {
	name := "GetHash"
	t.Run(name, func(t *testing.T) {
		err := func() (err error) {
			g, err := factory.FromStruct(w)
			if err != nil {
				return
			}

			assert.True(t, g.GetHash() == g.GetHash())

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}
