package spec

import (
	"context"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/gen/v0/mql"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostClientAccounts(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string) {
	t.Run("TestPostClientAccounts", func(t *testing.T) {
		t.Parallel()

		requirePostClientAccount(t, ctx, f, h, svcName, "metamate@metamate.io", "secret")
	})
}

func requirePostWhatevers(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string, whatevers []sdk.Whatever) (postRsp sdk.PostWhateversResponse) {
	postReq := sdk.PostWhateversRequest{
		ServiceFilter: &sdk.ServiceFilter{
			Id: &sdk.ServiceIdFilter{
				Value: &sdk.StringFilter{
					Is: sdk.String(svcName),
				},
			},
		},
		Mode: &sdk.PostMode{
			Kind: &sdk.PostModeKind.Collection,
			Collection: &sdk.CollectionPostMode{},
		},
		Select: &sdk.PostWhateversResponseSelect{
			Meta: GetResponseMetaSelect(),
			Whatevers: &sdk.WhateverSelect{
				Id: &sdk.ServiceIdSelect{
					Value: sdk.Bool(true),
				},
				AlternativeIds: &sdk.IdSelect{
					Kind: sdk.Bool(true),
					Name: sdk.Bool(true),
					Email: &sdk.EmailSelect{
						Value: sdk.Bool(true),
					},
				},
			},
		},
		Whatevers: whatevers,
	}

	gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
	if err != nil {
		t.Error(err)

		return
	}

	requirePostRsp(t, gPostRsp)

	gPostRsp.MustToStruct(&postRsp)

	return
}

func requirePostBlueWhatevers(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string, blueWhatevers []sdk.BlueWhatever) (postRsp sdk.PostBlueWhateversResponse) {
	postReq := sdk.PostBlueWhateversRequest{
		ServiceFilter: &sdk.ServiceFilter{
			Id: &sdk.ServiceIdFilter{
				Value: &sdk.StringFilter{
					Is: sdk.String(svcName),
				},
			},
		},
		Mode: &sdk.PostMode{
			Kind: &sdk.PostModeKind.Collection,
			Collection: &sdk.CollectionPostMode{},
		},
		Select: &sdk.PostBlueWhateversResponseSelect{
			Meta: GetResponseMetaSelect(),
			BlueWhatevers: &sdk.BlueWhateverSelect{
				Id: &sdk.ServiceIdSelect{
					Value: sdk.Bool(true),
				},
				AlternativeIds: &sdk.IdSelect{
					Kind: sdk.Bool(true),
					Name: sdk.Bool(true),
					Email: &sdk.EmailSelect{
						Value: sdk.Bool(true),
					},
				},
			},
		},
		BlueWhatevers: blueWhatevers,
	}

	gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
	if err != nil {
		t.Error(err)

		return
	}

	requirePostRsp(t, gPostRsp)
	gPostRsp.MustToStruct(&postRsp)

	return
}

func requirePostClientAccount(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName, email, password string) (postRsp sdk.PostClientAccountsRequest) {
	postReq := sdk.PostClientAccountsRequest{
		ServiceFilter: &sdk.ServiceFilter{
			Id: &sdk.ServiceIdFilter{
				Value: &sdk.StringFilter{
					Is: sdk.String(svcName),
				},
			},
		},
		Mode: &sdk.PostMode{
			Kind: &sdk.PostModeKind.Collection,
			Collection: &sdk.CollectionPostMode{},
		},
		Select: &sdk.PostClientAccountsResponseSelect{
			Meta: GetResponseMetaSelect(),
			ClientAccounts: &sdk.ClientAccountSelect{
				Id: &sdk.ServiceIdSelect{
					Value: sdk.Bool(true),
				},
				AlternativeIds: &sdk.IdSelect{
					Kind: sdk.Bool(true),
					Name: sdk.Bool(true),
					Email: &sdk.EmailSelect{
						Value: sdk.Bool(true),
					},
				},
				Password: &sdk.PasswordSelect{
					IsHashed:     sdk.Bool(true),
					HashFunction: sdk.Bool(true),
					Value:        sdk.Bool(true),
				},
			},
		},
		ClientAccounts: []sdk.ClientAccount{
			{
				AlternativeIds: []sdk.Id{
					{
						Kind: &sdk.IdKind.Email,
						Email: &sdk.Email{
							Value: sdk.String(email),
						},
					},
				},
				Password: &sdk.Password{
					Value: sdk.String(password),
				},
			},
		},
	}

	gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
	if err != nil {
		t.Error(err)

		return
	}

	requirePostRsp(t, gPostRsp)

	gPostRsp.MustToStruct(&postRsp)

	return
}

func TestAuthenticateClientAccount(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), storageSvc string, authSvc string) {
	t.Run("TestAuthenticateClientAccount", func(t *testing.T) {
		t.Parallel()
		requirePostClientAccount(t, ctx, f, h, storageSvc, "metamate@metamate.io", "secret")

		requireAuthenticateClientAccount(t, ctx, f, h, authSvc, "metamate@metamate.io", "secret")
	})
}

func requireAuthenticateClientAccount(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName, email, password string) (rsp sdk.AuthenticateClientAccountResponse) {
	authRequest := sdk.AuthenticateClientAccountRequest{
		ServiceFilter: &sdk.ServiceFilter{
			Id: &sdk.ServiceIdFilter{
				Value: &sdk.StringFilter{
					Is: sdk.String(svcName),
				},
			},
		},
		Select: &sdk.AuthenticateClientAccountResponseSelect{
			Meta: GetResponseMetaSelect(),
			Output: &sdk.AuthenticateClientAccountOutputSelect{
				Token: &sdk.TokenSelect{
					Value: sdk.Bool(true),
				},
			},
		},
		Input: &sdk.AuthenticateClientAccountInput{
			Id: &sdk.Id{
				Kind: &sdk.IdKind.Email,
				Email: &sdk.Email{
					Value: sdk.String(email),
				},
			},
			Password: sdk.String(password),
		},
	}

	gAuthRsp, err := h(ctx, f.MustFromStruct(authRequest))
	if err != nil {
		t.Error(err)

		return
	}

	authRsp := sdk.AuthenticateClientAccountResponse{}
	gAuthRsp.MustToStruct(&authRsp)

	require.NotNil(t, authRsp)
	require.NotNil(t, authRsp.Output)
	require.NotNil(t, authRsp.Output.Token)
	require.NotNil(t, authRsp.Output.Token.Value)
	assert.NotEqual(t, "", *authRsp.Output.Token.Value)

	rsp = authRsp

	return
}

func TestToken(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), authSvcName, storageSvcName string) {
	name := "TestToken"
	t.Run(name, func(t *testing.T) {
		t.Parallel()

		postSysRsp := requirePostClientAccount(t, ctx, f, h, storageSvcName, "metamate@metamate.io", "secret")

		authRsp := requireAuthenticateClientAccount(t, ctx, f, h, authSvcName, "metamate@metamate.io", "secret")

		postReq := sdk.PostServiceAccountsRequest{
			ServiceFilter: &sdk.ServiceFilter{
				Id: &sdk.ServiceIdFilter{
					Value: &sdk.StringFilter{
						Is: sdk.String(storageSvcName),
					},
				},
			},
			Mode: &sdk.PostMode{
				Kind: &sdk.PostModeKind.Collection,
				Collection: &sdk.CollectionPostMode{},
			},
			Select: &sdk.PostServiceAccountsResponseSelect{
				Meta: GetResponseMetaSelect(),
				ServiceAccounts: &sdk.ServiceAccountSelect{
					Id: &sdk.ServiceIdSelect{
						Value: sdk.Bool(true),
					},
					AlternativeIds: &sdk.IdSelect{
						Kind: sdk.Bool(true),
						Name: sdk.Bool(true),
						Email: &sdk.EmailSelect{
							Value: sdk.Bool(true),
						},
					},
					Url: &sdk.UrlSelect{
						Value: sdk.Bool(true),
					},
					Handle: sdk.Bool(true),
					Password: &sdk.PasswordSelect{
						IsHashed: sdk.Bool(true),
						Value:    sdk.Bool(true),
					},
				},
			},
			ServiceAccounts: []sdk.ServiceAccount{
				{
					Url: &sdk.Url{
						Value: sdk.String("example.com"),
					},
					Password: &sdk.Password{
						Value: sdk.String("example123"),
					},
				},
			},
		}

		gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
		if err != nil {
			t.Error(err)

			return
		}

		postSvcRsp := sdk.PostServiceAccountsResponse{}
		gPostRsp.MustToStruct(&postSvcRsp)

		putReq := sdk.PutServiceAccountsRequest{
			ServiceFilter: &sdk.ServiceFilter{
				Id: &sdk.ServiceIdFilter{
					Value: &sdk.StringFilter{
						Is: sdk.String(storageSvcName),
					},
				},
			},
			Mode: &sdk.PutMode{
				Kind: &sdk.PutModeKind.Relation,
				Relation: &sdk.RelationPutMode{
					Operation: &sdk.RelationOperation.Add,
					Id:        postSysRsp.ClientAccounts[0].Id,
					Ids:       []sdk.ServiceId{*postSvcRsp.ServiceAccounts[0].Id},
					Relation:  &sdk.ClientAccountRelationName.ClientAccountOwnsServiceAccounts,
				},
			},
			Select: &sdk.PutServiceAccountsResponseSelect{
				Meta: GetResponseMetaSelect(),
			},
		}

		_, err = h(ctx, f.MustFromStruct(putReq))
		if err != nil {
			t.Error(err)

			return
		}

		getReq := sdk.GetWhateversRequest{
			Mode: &sdk.GetMode{
				Kind: &sdk.GetModeKind.Collection,
				Collection: &sdk.CollectionGetMode{},
			},
			Auth: &sdk.Auth{
				Token: authRsp.Output.Token,
			},
		}

		_, err = h(ctx, f.MustFromStruct(getReq))
		if err != nil {
			t.Error(err)

			return
		}
	})
}

func TestVerifyToken(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string) {
	t.Run("TestVerifyToken", func(t *testing.T) {
		t.Parallel()

		requirePostClientAccount(t, ctx, f, h, svcName, "metamate@metamate.io", "secret")

		authRsp := requireAuthenticateClientAccount(t, ctx, f, h, svcName, "metamate@metamate.io", "secret")

		verifyReq := sdk.VerifyTokenRequest{
			ServiceFilter: &sdk.ServiceFilter{
				Id: &sdk.ServiceIdFilter{
					Value: &sdk.StringFilter{
						Is: sdk.String(svcName),
					},
				},
			},
			Select: &sdk.VerifyTokenResponseSelect{
				Meta: GetResponseMetaSelect(),
				Output: &sdk.VerifyTokenOutputSelect{
					IsValid: sdk.Bool(true),
					ClientAccountId: &sdk.ServiceIdSelect{
						ServiceName: sdk.Bool(true),
						Value:       sdk.Bool(true),
					},
				},
			},
			Input: &sdk.VerifyTokenInput{
				Token: authRsp.Output.Token,
			},
		}

		gVerifyRsp, err := h(ctx, f.MustFromStruct(verifyReq))
		if err != nil {
			t.Error(err)

			return
		}

		verfiyRsp := sdk.VerifyTokenResponse{}
		gVerifyRsp.MustToStruct(&verfiyRsp)

		require.NotNil(t, verfiyRsp.Output)

		require.NotNil(t, verfiyRsp.Output.IsValid)
		assert.True(t, *verfiyRsp.Output.IsValid)

		require.NotNil(t, verfiyRsp.Output.ClientAccountId)
		assert.NotNil(t, verfiyRsp.Output.ClientAccountId.Value)
		//assert.NotNil(t, verfiyRsp.Output.ClientAccountId.ServiceName)
	})
}
