package pkg

import (
	"context"
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamate/gen/v0/sdk"
	
)

func getSocialAccountId(ctx context.Context, c *mastodon.Client, req sdk.GetSocialAccountsRequest) (rsp sdk.GetSocialAccountsResponse) {
	err := func() (err error) {
		var account *mastodon.Account
		switch *req.Mode.Id.Kind {
		case sdk.IdKind.ServiceId:
			account, err = c.GetAccount(ctx, mastodon.ID(*req.Mode.Id.ServiceId.Value))
			if err != nil {
				return
			}

			break
		case sdk.IdKind.Email:
			break
		case sdk.IdKind.Name:
			break
		case sdk.IdKind.Username:
			break
		case sdk.IdKind.Ean:
			break
		case sdk.IdKind.Url:
			break
		case sdk.IdKind.Me:
			account, err = c.GetAccountCurrentUser(ctx)
			if err != nil {
				return
			}

			break
		}

		rsp.SocialAccounts = []sdk.SocialAccount{MapSocialAccountFromMastodonAccount(*account)}

		return
	}()
	if err != nil {
		rsp.Errors = append(rsp.Errors, sdk.Error{
			Message: &sdk.Text{
				Value: sdk.String(err.Error()),
			},
		})
	}

	return
}

func getSocialAccountsSearch(ctx context.Context, c *mastodon.Client, req sdk.GetSocialAccountsRequest) (rsp sdk.GetSocialAccountsResponse) {
	err := func() (err error) {
		accounts, err := c.AccountsSearch(ctx, *req.Mode.Search.Term, 100)
		if err != nil {
			return
		}

		rsp.SocialAccounts = MapSocialAccountsFromMastodonAccounts(accounts)

		return
	}()
	if err != nil {
		rsp.Errors = append(rsp.Errors, sdk.Error{
			Message: &sdk.Text{
				Value: sdk.String(err.Error()),
			},
		})
	}

	return
}

func getSocialAccountsRelation(ctx context.Context, c *mastodon.Client, req sdk.GetSocialAccountsRequest) (rsp sdk.GetSocialAccountsResponse) {
	var accounts []*mastodon.Account

	//var pagination *sdk.Pagination
	pg := &mastodon.Pagination{}

	//if len(req.Pages) > 0 {
	//	pagination = req.Pages[0].Page
	//}
	//
	//if pagination != nil &&
	//	pagination.Next != nil &&
	//	pagination.Next.CursorPage != nil {
	//	pg.MaxID = mastodon.ID(*pagination.Next.CursorPage.Value)
	//}
	//
	//if pagination != nil &&
	//	pagination.Previous != nil &&
	//	pagination.Previous.CursorPage != nil {
	//	pg.SinceID = mastodon.ID(*pagination.Previous.CursorPage.Value)
	//}

	err := func() (err error) {
		switch *req.Mode.Relation.Relation {
		case sdk.SocialAccountRelationName.SocialAccountBlocksSocialAccounts:
			accounts, err = c.GetBlocks(ctx, pg)
			if err != nil {
				return
			}
		case sdk.SocialAccountRelationName.SocialAccountFollowedBySocialAccounts:
			accounts, err = c.GetAccountFollowers(ctx, mastodon.ID(*req.Mode.Relation.Id.Value), pg)
			if err != nil {
				return
			}
		case sdk.SocialAccountRelationName.SocialAccountFollowsSocialAccounts:
			accounts, err = c.GetAccountFollowing(ctx, mastodon.ID(*req.Mode.Relation.Id.Value), pg)
			if err != nil {
				return
			}
		case sdk.SocialAccountRelationName.SocialAccountMutesSocialAccounts:
			accounts, err = c.GetMutes(ctx, pg)
			if err != nil {
				return
			}
		case sdk.SocialAccountRelationName.SocialAccountRequestedToBeFollowedBySocialAccounts:
			accounts, err = c.GetFollowRequests(ctx, pg)
			if err != nil {
				return
			}
		case sdk.PostRelationName.PostFavoredBySocialAccounts:
			accounts, err = c.GetFavouritedBy(ctx, mastodon.ID(*req.Mode.Relation.Id.Value), pg)
			if err != nil {
				return
			}
		}

		return
	}()
	if err != nil {
		rsp.Errors = append(rsp.Errors, sdk.Error{
			Message: &sdk.Text{
				Value: sdk.String(err.Error()),
			},
		})
	}

	//if pg != nil {
	//	pagination := &sdk.Pagination{
	//		Previous: &sdk.Page{
	//			Kind: &sdk.PageKind.CursorPage,
	//			CursorPage: &sdk.CursorPage{
	//				Value: sdk.String(string(pg.SinceID)),
	//			},
	//		},
	//		Next: &sdk.Page{
	//			Kind: &sdk.PageKind.CursorPage,
	//			CursorPage: &sdk.CursorPage{
	//				Value: sdk.String(string(pg.MaxID)),
	//			},
	//		},
	//	}
	//
	//	if pagination != nil {
	//		pagination.Current = pagination
	//	}
	//
	//	rsp.Pagination = pagination
	//}

	rsp.SocialAccounts = MapSocialAccountsFromMastodonAccounts(accounts)

	return
}