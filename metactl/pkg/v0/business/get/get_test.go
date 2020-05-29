package get

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/boot"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"testing"
)

func TestGet(t *testing.T) {
	err := func() (err error) {
		args := types.GetArgs{
			Host: "https://metamate.one",
			//Search: "books",
			//Services: "hackernews",
			TypePlural: "socialaccounts",
			Id:         "reddit/TheMrZZ0",
			Path:       "authorsposts",
		}

		d := boot.GetDependencies(0, types.Version{})

		o, err := Get(d.MessageReport, 0, d.RootNode, d.Factory, args)
		if err != nil {
			return
		}

		handleReport(*d.MessageReport, o)

		return
	}()
	if err != nil {
		t.Error(err)
	}
}
