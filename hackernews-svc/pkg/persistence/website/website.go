package website

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/sdk"
	"net/http"
)

func GetSocialAccountBookmarksPosts(c *http.Client, username string, sp *sdk.ServicePage) (ss []sdk.Post, pagination *sdk.Pagination, errs []sdk.Error) {
	err := func() (err error) {
		if sp == nil {
			sp = &sdk.ServicePage{
				Page: &sdk.Page{
					Kind: &sdk.PageKind.IndexPage,
					IndexPage: &sdk.IndexPage{
						Value: sdk.Int32(0),
					},
				},
			}
		}

		u := fmt.Sprintf("https://news.ycombinator.com/favorites?id=%v&p=%v", username, *sp.Page.IndexPage.Value+1)

		rsp, err := http.Get(u)
		if err != nil {
			return
		}
		defer rsp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(rsp.Body)
		if err != nil {
			return
		}

		doc.Find("table.itemlist tr.athing").Each(func(i int, s *goquery.Selection) {
			v, ok := s.Attr("id")
			if !ok {
				return
			}

			ss = append(ss, sdk.Post{
				Id: &sdk.ServiceId{
					Value: &v,
				},
			})
		})

		pagination = &sdk.Pagination{}
		pagination.Current = []sdk.ServicePage{*sp}

		if doc.Find("table.itemlist a.morelink").Length() == 1 {
			pagination.Next = []sdk.ServicePage{
				{
					Page: &sdk.Page{
						Kind: &sdk.PageKind.IndexPage,
						IndexPage: &sdk.IndexPage{
							Value: sdk.Int32(*sp.Page.IndexPage.Value + 1),
						},
					},
				},
			}
		}

		if *sp.Page.IndexPage.Value != 0 {
			pagination.Previous = []sdk.ServicePage{
				{
					Page: &sdk.Page{
						Kind: &sdk.PageKind.IndexPage,
						IndexPage: &sdk.IndexPage{
							Value: sdk.Int32(*sp.Page.IndexPage.Value - 1),
						},
					},
				},
			}
		}

		return
	}()
	if err != nil {
		errs = append(errs, sdk.Error{
			Message: sdk.String(err.Error()),
		})
	}

	return
}
