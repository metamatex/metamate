package types

import (
	"context"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
)

type MockService interface {
	LookupService(context.Context, sdk.LookupServiceRequest) sdk.LookupServiceResponse
	PostClientAccounts(context.Context, sdk.PostClientAccountsRequest) sdk.PostClientAccountsResponse
	GetClientAccounts(context.Context, sdk.GetClientAccountsRequest) sdk.GetClientAccountsResponse
	AuthenticateClientAccount(context.Context, sdk.AuthenticateClientAccountRequest) sdk.AuthenticateClientAccountResponse
	VerifyToken(context.Context, sdk.VerifyTokenRequest) sdk.VerifyTokenResponse
	GetWhatevers(context.Context, sdk.GetWhateversRequest) sdk.GetWhateversResponse
	PostWhatevers(context.Context, sdk.PostWhateversRequest) sdk.PostWhateversResponse
	PipeWhatevers(context.Context, sdk.PipeWhateversRequest) sdk.PipeWhateversResponse
	PutWhatevers(context.Context, sdk.PutWhateversRequest) sdk.PutWhateversResponse
	DeleteWhatevers(context.Context, sdk.DeleteWhateversRequest) sdk.DeleteWhateversResponse
	GetBlueWhatevers(context.Context, sdk.GetBlueWhateversRequest) sdk.GetBlueWhateversResponse
	PostBlueWhatevers(context.Context, sdk.PostBlueWhateversRequest) sdk.PostBlueWhateversResponse
	PutBlueWhatevers(context.Context, sdk.PutBlueWhateversRequest) sdk.PutBlueWhateversResponse
	DeleteBlueWhatevers(context.Context, sdk.DeleteBlueWhateversRequest) sdk.DeleteBlueWhateversResponse
	PostServiceAccounts(context.Context, sdk.PostServiceAccountsRequest) sdk.PostServiceAccountsResponse
	PutServiceAccounts(context.Context, sdk.PutServiceAccountsRequest) sdk.PutServiceAccountsResponse
	GetServiceAccounts(context.Context, sdk.GetServiceAccountsRequest) sdk.GetServiceAccountsResponse
}
