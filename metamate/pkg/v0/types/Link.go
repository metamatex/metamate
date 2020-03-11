package types

import (
	sdk "github.com/metamatex/metamate/gen/v0/sdk"
)

type LinkStore interface {
	GetLinks(relationName string, active bool, id sdk.ServiceId) (ids []sdk.ServiceId, err error)
	PostLinks(relationName string, active bool, id sdk.ServiceId, ids []sdk.ServiceId) (err error)
	DeleteLinks(relationName string, active bool, id sdk.ServiceId, ids []sdk.ServiceId) (err error)
}

type Link struct {
	Active  sdk.ServiceId
	Passive sdk.ServiceId
}
