package communication

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"
)

func handleDelete(ctx context.Context, db sqlx.Ext, m generic.Factory, gRequest generic.Generic) (gResponse generic.Generic, errs []error) {
	return
}
