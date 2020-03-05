package types

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
)

type Dependencies struct {
	Db       sqlx.Ext
	Factory  generic.Factory
	RootNode *graph.RootNode
	ServeFunc   func(ctx context.Context, gRequest generic.Generic) (gResponse generic.Generic)
}
