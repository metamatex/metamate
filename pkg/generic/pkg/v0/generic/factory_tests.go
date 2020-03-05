package generic

import (
	"github.com/metamatex/asg/pkg/v0/asg/typenames"
	"testing"

	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/stretchr/testify/assert"

	"github.com/metamatex/metamatemono/gen/v0/sdk"
)

func TestFactory(t *testing.T, rn *graph.RootNode, f Factory) {
	t.Parallel()

	t.Run("TestFactoryNew", func(t *testing.T) {
		t.Parallel()

		TestFactoryNew(t, rn, f)
	})

	t.Run("TestFactoryFromStruct", func(t *testing.T) {
		t.Parallel()

		TestFactoryFromStruct(t, rn, f)
	})

	t.Run("TestFactoryFromStringInterfaceMap", func(t *testing.T) {
		t.Parallel()

		TestFactoryFromStringInterfaceMap(t, rn, f)
	})
}

func TestFactoryNew(t *testing.T, rn *graph.RootNode, f Factory) {
	tn := rn.Types.MustByName(typenames.Whatever)

	g := f.New(rn.Types.MustByName(typenames.Whatever))

	assert.Equal(t, tn.Name(), g.Type().Name())
}

func TestFactoryFromStruct(t *testing.T, rn *graph.RootNode, f Factory) {
	err := func() (err error) {
		g := f.MustFromStruct(w)

		w0 := sdk.Whatever{}
		g.MustToStruct(&w0)

		assert.Equal(t, w, w0)

		return
	}()
	if err != nil {
	    t.Error(err)
	}
}

func TestFactoryFromStringInterfaceMap(t *testing.T, rn *graph.RootNode, f Factory) {
	err := func() (err error) {
		g, err := f.FromStruct(w)
		if err != nil {
			return
		}

		m := g.ToStringInterfaceMap()

		g0, err := f.FromStringInterfaceMap(rn.Types.MustByName(typenames.Whatever), m)
		if err != nil {
		    return
		}

		w0 := sdk.Whatever{}
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
}

func TestFactoryUnflatten(t *testing.T, rn *graph.RootNode, f Factory) {
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

		g0, err := f.Unflatten(rn.Types.MustByName(typenames.Whatever), d, m)
		if err != nil {
			return
		}

		w0 := sdk.Whatever{}
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
}

//FromStructs(interface{}) (Generic, error)
