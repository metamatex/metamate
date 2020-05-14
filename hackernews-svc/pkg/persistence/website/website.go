package website

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/metamatex/metamate/hackernews-svc/gen/v0/mql"
	"net/http"
)

func GetSocialAccountBookmarksPosts(c *http.Client, username string, sp *mql.ServicePage) (ss []mql.Post, pagination *mql.Pagination, errs []mql.Error) {
	err := func() (err error) {
		if sp == nil {
			sp = &mql.ServicePage{
				Page: &mql.Page{
					Kind: &mql.PageKind.IndexPage,
					IndexPage: &mql.IndexPage{
						Value: mql.Int32(0),
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

			ss = append(ss, mql.Post{
				Id: &mql.ServiceId{
					Value: &v,
				},
			})
		})

		pagination = &mql.Pagination{}
		pagination.Current = []mql.ServicePage{*sp}

		if doc.Find("table.itemlist a.morelink").Length() == 1 {
			pagination.Next = []mql.ServicePage{
				{
					Page: &mql.Page{
						Kind: &mql.PageKind.IndexPage,
						IndexPage: &mql.IndexPage{
							Value: mql.Int32(*sp.Page.IndexPage.Value + 1),
						},
					},
				},
			}
		}

		if *sp.Page.IndexPage.Value != 0 {
			pagination.Previous = []mql.ServicePage{
				{
					Page: &mql.Page{
						Kind: &mql.PageKind.IndexPage,
						IndexPage: &mql.IndexPage{
							Value: mql.Int32(*sp.Page.IndexPage.Value - 1),
						},
					},
				},
			}
		}

		return
	}()
	if err != nil {
		errs = append(errs, mql.Error{
			Message: mql.String(err.Error()),
		})
	}

	return
}
