package graphql

import (
	"context"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg"
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"
	"testing"
)

func TestGraphql(t *testing.T) {
	err := func() (err error) {
		rn, err := asg.New()
		if err != nil {
			return
		}

		f := generic.NewFactory(rn)

		schema, err := GetSchema(f, func(ctx context.Context, gCliReq generic.Generic) (gCliRsp generic.Generic) {
			gCliReq.Print()

			return f.New(gCliReq.Type().Edges.Type.Response())
		}, rn)
		if err != nil {
			return
		}

		q := `query MyQuery {
  getFeeds {
    feeds {
      id {
        serviceName
        value
      }
      relations {
        containsStatuses(filter: {id: {value: {is: "home"}}}) {
          statuses {
            id {
              serviceName
              value
            }
          }
        }
      }
    }
  }
}`

		_, err = ExecuteQuery(schema, q)
		if err != nil {
			return
		}

		return
	}()
	if err != nil {
		t.Error(err)
	}
}
