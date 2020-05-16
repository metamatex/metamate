package generic

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactory(t *testing.T) {
	var err error
	rn, err = asg.New()
	if err != nil {
		panic(err)
	}

	f = NewFactory(rn)

	FTestFactoryNew(t, rn, f)
	FTestFactoryFromStruct(t, rn, f)
	FTestFactoryFromStringInterfaceMap(t, rn, f)
	FTestFactoryUnflatten(t, rn, f)
}

func FTestFactoryNew(t *testing.T, rn *graph.RootNode, f Factory) {
	t.Run("TestFactoryNew", func(t *testing.T) {
		t.Parallel()
		tn := rn.Types.MustByName(typenames.Dummy)

		g := f.New(rn.Types.MustByName(typenames.Dummy))

		assert.Equal(t, tn.Name(), g.Type().Name())
	})
}

func FTestFactoryFromStruct(t *testing.T, rn *graph.RootNode, f Factory) {
	t.Run("TestFactoryFromStruct", func(t *testing.T) {
		t.Parallel()
		err := func() (err error) {
			g := f.MustFromStruct(w)

			w0 := mql.Dummy{}
			g.MustToStruct(&w0)

			assert.Equal(t, w, w0)

			return
		}()
		if err != nil {
			t.Error(err)
		}
	})
}

func FTestFactoryFromStringInterfaceMap(t *testing.T, rn *graph.RootNode, f Factory) {
	t.Run("FTestFactoryFromStringInterfaceMap", func(t *testing.T) {
		t.Parallel()
		err := func() (err error) {
			g, err := f.FromStruct(w)
			if err != nil {
				return
			}

			m := g.ToStringInterfaceMap()

			g0, err := f.FromStringInterfaceMap(rn.Types.MustByName(typenames.Dummy), m)
			if err != nil {
				return
			}

			w0 := mql.Dummy{}
			err = g0.ToStruct(&w0)
			if err != nil {
				return
			}

			assert.Equal(t, w, w0)

			return
		}()
		if err != nil {
			t.Error(err)
		}

		return
	})
}

func FTestFactoryUnflatten(t *testing.T, rn *graph.RootNode, f Factory) {
	t.Run("FTestFactoryUnflatten", func(t *testing.T) {
		t.Parallel()
		err := func() (err error) {
			g, err := f.FromStruct(w)
			if err != nil {
				return
			}

			d := "."
			m, err := g.Flatten(d)
			if err != nil {
				return
			}

			g0, err := f.Unflatten(rn.Types.MustByName(typenames.Dummy), d, m)
			if err != nil {
				return
			}

			w0 := mql.Dummy{}
			err = g0.ToStruct(&w0)
			if err != nil {
				return
			}

			assert.Equal(t, w, w0)

			return
		}()
		if err != nil {
			t.Error(err)
		}

		return
	})
}
