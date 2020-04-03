package types

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/sdk"
)

type MockService interface {
	LookupService(context.Context, sdk.LookupServiceRequest) sdk.LookupServiceResponse
	GetClientAccounts(context.Context, sdk.GetClientAccountsRequest) sdk.GetClientAccountsResponse
	GetWhatevers(context.Context, sdk.GetWhateversRequest) sdk.GetWhateversResponse
	GetBlueWhatevers(context.Context, sdk.GetBlueWhateversRequest) sdk.GetBlueWhateversResponse
	GetServiceAccounts(context.Context, sdk.GetServiceAccountsRequest) sdk.GetServiceAccountsResponse
}
