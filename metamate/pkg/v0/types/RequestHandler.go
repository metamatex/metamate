package types

import (
	"context"
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"
)

type RequestHandler func(context.Context, string, generic.Generic) (generic.Generic, error)
