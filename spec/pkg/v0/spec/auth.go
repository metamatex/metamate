package spec

import (
	"context"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"

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

func requirePostDummies(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string, whatevers []mql.Whatever) (postRsp mql.PostDummiesResponse) {
	postReq := mql.PostDummiesRequest{
		ServiceFilter: &mql.ServiceFilter{
			Id: &mql.ServiceIdFilter{
				Value: &mql.StringFilter{
					Is: mql.String(svcName),
				},
			},
		},
		Mode: &mql.PostMode{
			Kind:       &mql.PostModeKind.Collection,
			Collection: &mql.CollectionPostMode{},
		},
		Select: &mql.PostDummiesResponseSelect{
			Meta: GetResponseMetaSelect(),
			Dummies: &mql.WhateverSelect{
				Id: &mql.ServiceIdSelect{
					Value: mql.Bool(true),
				},
				AlternativeIds: &mql.IdSelect{
					Kind: mql.Bool(true),
					Name: mql.Bool(true),
					Email: &mql.EmailSelect{
						Value: mql.Bool(true),
					},
				},
			},
		},
		Dummies: whatevers,
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

func requirePostBlueDummies(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName string, blueDummies []mql.BlueWhatever) (postRsp mql.PostBlueDummiesResponse) {
	postReq := mql.PostBlueDummiesRequest{
		ServiceFilter: &mql.ServiceFilter{
			Id: &mql.ServiceIdFilter{
				Value: &mql.StringFilter{
					Is: mql.String(svcName),
				},
			},
		},
		Mode: &mql.PostMode{
			Kind:       &mql.PostModeKind.Collection,
			Collection: &mql.CollectionPostMode{},
		},
		Select: &mql.PostBlueDummiesResponseSelect{
			Meta: GetResponseMetaSelect(),
			BlueDummies: &mql.BlueWhateverSelect{
				Id: &mql.ServiceIdSelect{
					Value: mql.Bool(true),
				},
				AlternativeIds: &mql.IdSelect{
					Kind: mql.Bool(true),
					Name: mql.Bool(true),
					Email: &mql.EmailSelect{
						Value: mql.Bool(true),
					},
				},
			},
		},
		BlueDummies: blueDummies,
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

func requirePostClientAccount(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName, email, password string) (postRsp mql.PostClientAccountsRequest) {
	postReq := mql.PostClientAccountsRequest{
		ServiceFilter: &mql.ServiceFilter{
			Id: &mql.ServiceIdFilter{
				Value: &mql.StringFilter{
					Is: mql.String(svcName),
				},
			},
		},
		Mode: &mql.PostMode{
			Kind:       &mql.PostModeKind.Collection,
			Collection: &mql.CollectionPostMode{},
		},
		Select: &mql.PostClientAccountsResponseSelect{
			Meta: GetResponseMetaSelect(),
			ClientAccounts: &mql.ClientAccountSelect{
				Id: &mql.ServiceIdSelect{
					Value: mql.Bool(true),
				},
				AlternativeIds: &mql.IdSelect{
					Kind: mql.Bool(true),
					Name: mql.Bool(true),
					Email: &mql.EmailSelect{
						Value: mql.Bool(true),
					},
				},
				Password: &mql.PasswordSelect{
					IsHashed:     mql.Bool(true),
					HashFunction: mql.Bool(true),
					Value:        mql.Bool(true),
				},
			},
		},
		ClientAccounts: []mql.ClientAccount{
			{
				AlternativeIds: []mql.Id{
					{
						Kind: &mql.IdKind.Email,
						Email: &mql.Email{
							Value: mql.String(email),
						},
					},
				},
				Password: &mql.Password{
					Value: mql.String(password),
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

func requireAuthenticateClientAccount(t *testing.T, ctx context.Context, f generic.Factory, h func(ctx context.Context, gReq generic.Generic) (gRsp generic.Generic, err error), svcName, email, password string) (rsp mql.AuthenticateClientAccountResponse) {
	authRequest := mql.AuthenticateClientAccountRequest{
		ServiceFilter: &mql.ServiceFilter{
			Id: &mql.ServiceIdFilter{
				Value: &mql.StringFilter{
					Is: mql.String(svcName),
				},
			},
		},
		Select: &mql.AuthenticateClientAccountResponseSelect{
			Meta: GetResponseMetaSelect(),
			Output: &mql.AuthenticateClientAccountOutputSelect{
				Token: &mql.TokenSelect{
					Value: mql.Bool(true),
				},
			},
		},
		Input: &mql.AuthenticateClientAccountInput{
			Id: &mql.Id{
				Kind: &mql.IdKind.Email,
				Email: &mql.Email{
					Value: mql.String(email),
				},
			},
			Password: mql.String(password),
		},
	}

	gAuthRsp, err := h(ctx, f.MustFromStruct(authRequest))
	if err != nil {
		t.Error(err)

		return
	}

	authRsp := mql.AuthenticateClientAccountResponse{}
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

		postReq := mql.PostServiceAccountsRequest{
			ServiceFilter: &mql.ServiceFilter{
				Id: &mql.ServiceIdFilter{
					Value: &mql.StringFilter{
						Is: mql.String(storageSvcName),
					},
				},
			},
			Mode: &mql.PostMode{
				Kind:       &mql.PostModeKind.Collection,
				Collection: &mql.CollectionPostMode{},
			},
			Select: &mql.PostServiceAccountsResponseSelect{
				Meta: GetResponseMetaSelect(),
				ServiceAccounts: &mql.ServiceAccountSelect{
					Id: &mql.ServiceIdSelect{
						Value: mql.Bool(true),
					},
					AlternativeIds: &mql.IdSelect{
						Kind: mql.Bool(true),
						Name: mql.Bool(true),
						Email: &mql.EmailSelect{
							Value: mql.Bool(true),
						},
					},
					Url: &mql.UrlSelect{
						Value: mql.Bool(true),
					},
					Handle: mql.Bool(true),
					Password: &mql.PasswordSelect{
						IsHashed: mql.Bool(true),
						Value:    mql.Bool(true),
					},
				},
			},
			ServiceAccounts: []mql.ServiceAccount{
				{
					Url: &mql.Url{
						Value: mql.String("example.com"),
					},
					Password: &mql.Password{
						Value: mql.String("example123"),
					},
				},
			},
		}

		gPostRsp, err := h(ctx, f.MustFromStruct(postReq))
		if err != nil {
			t.Error(err)

			return
		}

		postSvcRsp := mql.PostServiceAccountsResponse{}
		gPostRsp.MustToStruct(&postSvcRsp)

		putReq := mql.PutServiceAccountsRequest{
			ServiceFilter: &mql.ServiceFilter{
				Id: &mql.ServiceIdFilter{
					Value: &mql.StringFilter{
						Is: mql.String(storageSvcName),
					},
				},
			},
			Mode: &mql.PutMode{
				Kind: &mql.PutModeKind.Relation,
				Relation: &mql.RelationPutMode{
					Operation: &mql.RelationOperation.Add,
					Id:        postSysRsp.ClientAccounts[0].Id,
					Ids:       []mql.ServiceId{*postSvcRsp.ServiceAccounts[0].Id},
					Relation:  &mql.ClientAccountRelationName.ClientAccountOwnsServiceAccounts,
				},
			},
			Select: &mql.PutServiceAccountsResponseSelect{
				Meta: GetResponseMetaSelect(),
			},
		}

		_, err = h(ctx, f.MustFromStruct(putReq))
		if err != nil {
			t.Error(err)

			return
		}

		getReq := mql.GetDummiesRequest{
			Mode: &mql.GetMode{
				Kind:       &mql.GetModeKind.Collection,
				Collection: &mql.CollectionGetMode{},
			},
			Auth: &mql.Auth{
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

		verifyReq := mql.VerifyTokenRequest{
			ServiceFilter: &mql.ServiceFilter{
				Id: &mql.ServiceIdFilter{
					Value: &mql.StringFilter{
						Is: mql.String(svcName),
					},
				},
			},
			Select: &mql.VerifyTokenResponseSelect{
				Meta: GetResponseMetaSelect(),
				Output: &mql.VerifyTokenOutputSelect{
					IsValid: mql.Bool(true),
					ClientAccountId: &mql.ServiceIdSelect{
						ServiceName: mql.Bool(true),
						Value:       mql.Bool(true),
					},
				},
			},
			Input: &mql.VerifyTokenInput{
				Token: authRsp.Output.Token,
			},
		}

		gVerifyRsp, err := h(ctx, f.MustFromStruct(verifyReq))
		if err != nil {
			t.Error(err)

			return
		}

		verfiyRsp := mql.VerifyTokenResponse{}
		gVerifyRsp.MustToStruct(&verfiyRsp)

		require.NotNil(t, verfiyRsp.Output)

		require.NotNil(t, verfiyRsp.Output.IsValid)
		assert.True(t, *verfiyRsp.Output.IsValid)

		require.NotNil(t, verfiyRsp.Output.ClientAccountId)
		assert.NotNil(t, verfiyRsp.Output.ClientAccountId.Value)
		//assert.NotNil(t, verfiyRsp.Output.ClientAccountId.ServiceName)
	})
}
