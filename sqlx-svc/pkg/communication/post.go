package communication

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/sqlx-svc/pkg/persistence"
)

func handlePost(ctx context.Context, db sqlx.Ext, f generic.Factory, gRequest generic.Generic) (gResponse generic.Generic, errs []error) {
	pluralFieldName := gRequest.Type().Edges.Type.For().PluralFieldName()

	gSlice := gRequest.MustGenericSlice(pluralFieldName)

	gs := gSlice.Get()

	for i, _ := range gs {
		id, err := persistence.Insert(db, gs[i])
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = gs[i].SetString([]string{fieldnames.Id, fieldnames.Value}, fmt.Sprintf("%v", id))
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}

	gSlice.Set(gs)

	gResponse = f.New(gRequest.Type().Edges.Type.Response())
	err := gResponse.SetGenericSlice([]string{pluralFieldName}, gSlice)
	if err != nil {
		errs = append(errs, err)

		return
	}

	return
}
