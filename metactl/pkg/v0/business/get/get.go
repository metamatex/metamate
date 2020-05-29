package get

import (
	"errors"
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils"
	"net/http"
	"sort"
	"strings"
)

type PrintConfig struct {
	Keys            []string
	GenerateColumns func(gEntities generic.Slice) (m map[string][]string)
}

var printConfigs = map[string]PrintConfig{
	mql.TypeNames.Service: {
		Keys: []string{"id.serviceName", "id.value", "name", "isVirtual", "url.value", "port", "sdkVersion", "endpoints"},
		GenerateColumns: func(gEntities generic.Slice) (m map[string][]string) {
			m = map[string][]string{}
			es := []string{}
			for _, g := range gEntities.Get() {
				gEndpoints, ok := g.Generic(mql.FieldNames.Endpoints)
				if !ok {
					es = append(es, "")

					continue
				}

				es0 := []string{}
				for _, n := range gEndpoints.FieldNames() {
					es0 = append(es0, n)
				}

				sort.Strings(es0)

				es = append(es, strings.Join(es0, ", "))
			}

			m["endpoints"] = es

			return
		},
	},
	mql.TypeNames.PostFeed: {
		Keys: []string{"id.serviceName", "id.value", "info.name.value"},
	},
	mql.TypeNames.Post: {
		Keys: []string{"id.serviceName", "id.value", "kind", "title.value", "relations.favoredBySocialAccounts.count", "totalWasRepliedToByPostsCount"},
	},
	mql.TypeNames.SocialAccount: {
		Keys: []string{"id.serviceName", "id.value", "username", "points", "relations.authorsPosts.count"},
	},
}

func Get(report *types.MessageReport, verbosity int, rn *graph.RootNode, f generic.Factory, args types.GetArgs) (o types.Output, err error) {
	err = args.Validate()
	if err != nil {
		return
	}

	en, err := rn.Endpoints.ById(graph.NodeId("get" + args.TypePlural))
	if err != nil {
		return
	}

	forTn := en.Edges.Type.For()
	var relationsFn *graph.FieldNode
	var printTn *graph.TypeNode
	gReq := f.New(forTn.Edges.Type.GetRequest())
	gReq.MustSetGeneric([]string{mql.FieldNames.Mode}, f.MustFromStruct(args.GetMode()))
	gReq.MustSetGeneric([]string{mql.FieldNames.ServiceFilter}, f.MustFromStruct(args.GetSvcFilter()))

	if args.Path != "" {
		getRelationsTn := forTn.Edges.Type.GetRelations()
		if getRelationsTn == nil {
			err = errors.New(fmt.Sprintf("--path set but %v doesn't have any relations", forTn.Name()))

			return
		}

		relationsFn, err = getRelationsTn.Edges.Fields.Holds().ByName(args.Path)
		if err != nil {
			return
		}

		printTn = relationsFn.Edges.Type.Holds().Edges.Type.For()

		gReq.MustSetGeneric([]string{mql.FieldNames.Relations, relationsFn.Name()}, f.New(relationsFn.Edges.Type.Holds()))
	} else {
		printTn = forTn
	}

	if verbosity > 0 {
		gReq.Print()
	}

	gRsp, err := generic.Send(f, &http.Client{}, args.Host+"/httpjson", args.User, args.Password, gReq)
	if err != nil {
		return
	}

	gErrors, ok := gRsp.GenericSlice(fieldnames.Errors)
	if ok {
		addGErrorsToReport(report, gErrors)
	}

	gEntities, ok := gRsp.GenericSlice(forTn.PluralFieldName())
	if !ok {
		err = errors.New("no data")

		return
	}

	var printGEntities generic.Slice
	if args.Path != "" {
		printGEntities = f.NewSlice(printTn)
		for _, g := range gEntities.Get() {
			gCollection, ok := g.Generic(mql.FieldNames.Relations, relationsFn.Name())
			gErrors, ok := gCollection.GenericSlice(fieldnames.Errors)
			if ok {
				addGErrorsToReport(report, gErrors)
			}

			gSlice, ok := gCollection.GenericSlice(printTn.PluralFieldName())
			if ok {
				printGEntities.Append(gSlice.Get()...)
			}
		}
	} else {
		printGEntities = gEntities
	}

	m, err := printGEntities.Flatten(".")
	if err != nil {
		return
	}

	c, ok := printConfigs[printTn.Name()]
	if !ok {
		err = errors.New(fmt.Sprintf("no print config keys for %v", printTn.Name()))

		return
	}

	if c.GenerateColumns != nil {
		m0 := c.GenerateColumns(gEntities)
		for k, v := range m0 {
			m[k] = v
		}
	}

	keysInfo := []string{}
	for k, _ := range m {
		keysInfo = append(keysInfo, k)
	}

	columns := [][]string{}
	header := []string{}

	for _, v := range c.Keys {
		columns = append(columns, m[v])
		header = append(header, strings.ToUpper(v))
	}

	rows := utils.ToRows(columns)

	rows = append([][]string{header}, rows...)

	o.Data = gRsp.ToStringInterfaceMap()
	o.Text = utils.GetTableString([]string{}, rows, []int{})

	return
}

func addGErrorsToReport(r *types.MessageReport, gErrs generic.Slice) {
	var errs []mql.Error
	gErrs.MustToStructs(&errs)

	for _, e := range errs {
		var name string
		var value string
		if e.Service != nil && e.Service.Id != nil {
			if e.Service.Id.ServiceName != nil {
				name = *e.Service.Id.ServiceName
			}

			if e.Service.Id.Value != nil {
				value = *e.Service.Id.Value
			}
		}

		var message string
		if e.Message != nil {
			message = *e.Message
		}

		r.AddError(fmt.Sprintf("%v/%v: %v", name, value, message))
	}
}
