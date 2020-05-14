package pkg

import (
	"context"
	"github.com/mattn/go-mastodon"
	"github.com/metamatex/metamate/gen/v0/mql"
	
)

func getSocialAccountId(ctx context.Context, c *mastodon.Client, req mql.GetSocialAccountsRequest) (rsp mql.GetSocialAccountsResponse) {
	err := func() (err error) {
		var account *mastodon.Account
		switch *req.Mode.Id.Kind {
		case mql.IdKind.ServiceId:
			account, err = c.GetAccount(ctx, mastodon.ID(*req.Mode.Id.ServiceId.Value))
			if err != nil {
				return
			}

			break
		case mql.IdKind.Email:
			break
		case mql.IdKind.Name:
			break
		case mql.IdKind.Username:
			break
		case mql.IdKind.Ean:
			break
		case mql.IdKind.Url:
			break
		case mql.IdKind.Me:
			account, err = c.GetAccountCurrentUser(ctx)
			if err != nil {
				return
			}

			break
		}

		rsp.SocialAccounts = []mql.SocialAccount{MapSocialAccountFromMastodonAccount(*account)}

		return
	}()
	if err != nil {
		rsp.Errors = append(rsp.Errors, mql.Error{
			Message: mql.String(err.Error()),
		})
	}

	return
}

func getSocialAccountsSearch(ctx context.Context, c *mastodon.Client, req mql.GetSocialAccountsRequest) (rsp mql.GetSocialAccountsResponse) {
	err := func() (err error) {
		accounts, err := c.AccountsSearch(ctx, *req.Mode.Search.Term, 100)
		if err != nil {
			return
		}

		rsp.SocialAccounts = MapSocialAccountsFromMastodonAccounts(accounts)

		return
	}()
	if err != nil {
		rsp.Errors = append(rsp.Errors, mql.Error{
			Message: mql.String(err.Error()),
		})
	}

	return
}

func getSocialAccountsRelation(ctx context.Context, c *mastodon.Client, req mql.GetSocialAccountsRequest) (rsp mql.GetSocialAccountsResponse) {
	var accounts []*mastodon.Account

	//var pagination *mql.Pagination
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
		case mql.SocialAccountRelationName.SocialAccountBlocksSocialAccounts:
			accounts, err = c.GetBlocks(ctx, pg)
			if err != nil {
				return
			}
		case mql.SocialAccountRelationName.SocialAccountFollowedBySocialAccounts:
			accounts, err = c.GetAccountFollowers(ctx, mastodon.ID(*req.Mode.Relation.Id.Value), pg)
			if err != nil {
				return
			}
		case mql.SocialAccountRelationName.SocialAccountFollowsSocialAccounts:
			accounts, err = c.GetAccountFollowing(ctx, mastodon.ID(*req.Mode.Relation.Id.Value), pg)
			if err != nil {
				return
			}
		case mql.SocialAccountRelationName.SocialAccountMutesSocialAccounts:
			accounts, err = c.GetMutes(ctx, pg)
			if err != nil {
				return
			}
		case mql.SocialAccountRelationName.SocialAccountRequestedToBeFollowedBySocialAccounts:
			accounts, err = c.GetFollowRequests(ctx, pg)
			if err != nil {
				return
			}
		case mql.PostRelationName.PostFavoredBySocialAccounts:
			accounts, err = c.GetFavouritedBy(ctx, mastodon.ID(*req.Mode.Relation.Id.Value), pg)
			if err != nil {
				return
			}
		}

		return
	}()
	if err != nil {
		rsp.Errors = append(rsp.Errors, mql.Error{
			Message: mql.String(err.Error()),
		})
	}

	//if pg != nil {
	//	pagination := &mql.Pagination{
	//		Previous: &mql.Page{
	//			Kind: &mql.PageKind.CursorPage,
	//			CursorPage: &mql.CursorPage{
	//				Value: mql.String(string(pg.SinceID)),
	//			},
	//		},
	//		Next: &mql.Page{
	//			Kind: &mql.PageKind.CursorPage,
	//			CursorPage: &mql.CursorPage{
	//				Value: mql.String(string(pg.MaxID)),
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