package boot

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/typenames"
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/sqlx-svc/pkg/communication"
	"github.com/metamatex/metamatemono/sqlx-svc/pkg/persistence"
	"github.com/metamatex/metamatemono/sqlx-svc/pkg/types"

	"github.com/metamatex/metamatemono/gen/v0/sdk"
)

func NewDependencies(rn *graph.RootNode, f generic.Factory, c types.Config) (d types.Dependencies, err error) {
	supportedIdKinds := map[string]bool{
		sdk.IdKind.Ean:      true,
		sdk.IdKind.Email:    true,
		sdk.IdKind.Name:     true,
		sdk.IdKind.Url:      true,
		sdk.IdKind.Username: true,
	}

	d.RootNode = rn
	d.Factory = f

	db, err := getDb(c.DriverName, c.DataSource)
	if err != nil {
		return
	}

	db.SetMaxOpenConns(c.MaxOpenConns)

	d.Db = db

	if c.Log {
		d.Db = types.LoggedDB{Ext: d.Db}
	}

	tnm := d.RootNode.Types.Flagged(typeflags.GetEndpoints, true).ByNames(c.TypeNames...)

	err = persistence.Migrate(supportedIdKinds, d.Db, tnm, d.RootNode.Relations)
	if err != nil {
		return
	}

	gGetServiceRsp := GetGGetServiceRsp(d.RootNode, d.Factory, tnm)

	d.ServeFunc = communication.GetServer(supportedIdKinds, d.Db, rn, d.Factory, gGetServiceRsp)

	return
}

func getDb(driverName, dataSource string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open(driverName, dataSource)
	if err != nil {
		return
	}

	db, err = sqlx.Connect(driverName, dataSource)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	return
}

func NewTestConfig() (c types.Config) {
	c.Log = true
	c.DriverName = "sqlite3"
	c.DataSource = ":memory:"
	c.TypeNames = []string{typenames.Whatever, typenames.ClientAccount, typenames.ServiceAccount, typenames.BlueWhatever}
	c.MaxOpenConns = 1

	return
}

func GetGGetServiceRsp(rn *graph.RootNode, f generic.Factory, tnm graph.TypeNodeMap) (gRsp generic.Generic) {
	gEndpoints := f.New(rn.Types.MustByName(typenames.Endpoints))

	tnm.Each(func(tn *graph.TypeNode) {
		gPostEndpoint := getGPostEndpoint(f, tn)
		gEndpoints.MustSetGeneric([]string{tn.Edges.Endpoint.Post().FieldName()}, gPostEndpoint)

		gGetEndpoint := getGGetEndpoint(f, tn)
		gEndpoints.MustSetGeneric([]string{tn.Edges.Endpoint.Get().FieldName()}, gGetEndpoint)

		gPutEndpoint := getGPutEndpoint(f, tn)
		gEndpoints.MustSetGeneric([]string{tn.Edges.Endpoint.Put().FieldName()}, gPutEndpoint)

		gDeleteEndpoint := getGDeleteEndpoint(f, tn)
		gEndpoints.MustSetGeneric([]string{tn.Edges.Endpoint.Delete().FieldName()}, gDeleteEndpoint)
	})

	gSvc := f.New(rn.Types.MustByName(typenames.Service))
	gSvc.MustSetString([]string{fieldnames.Name}, "sqlx-svc")
	gSvc.MustSetGeneric([]string{fieldnames.Endpoints}, gEndpoints)

	gRsp = f.New(rn.Endpoints.MustByName(endpointnames.LookupService).Edges.Type.Response())
	gRsp.MustSetGeneric([]string{fieldnames.Output, fieldnames.Service}, gSvc)

	return
}

func getGPostEndpoint(f generic.Factory, tn *graph.TypeNode) (gEndpoint generic.Generic) {
	gEndpoint = f.New(tn.Edges.Type.PostEndpoint())

	return
}

func getGGetEndpoint(f generic.Factory, tn *graph.TypeNode) (gEndpoint generic.Generic) {
	gEndpoint = f.New(tn.Edges.Type.GetEndpoint())

	getMode := sdk.GetModeFilter{
		Kind: &sdk.EnumFilter{
			In: []string{
				sdk.GetModeKind.Collection,
				sdk.GetModeKind.Id,
				sdk.GetModeKind.Relation,
			},
		},
	}

	gEndpoint.MustSetGeneric([]string{fieldnames.Filter, fieldnames.Mode}, f.MustFromStruct(getMode))

	return
}

func getGPutEndpoint(f generic.Factory, tn *graph.TypeNode) (gEndpoint generic.Generic) {
	gEndpoint = f.New(tn.Edges.Type.GetEndpoint())

	putMode := sdk.PutModeFilter{
		Kind: &sdk.EnumFilter{
			In: []string{
				sdk.PutModeKind.Relation,
			},
		},
	}

	gEndpoint.MustSetGeneric([]string{fieldnames.Filter, fieldnames.Mode}, f.MustFromStruct(putMode))

	return
}

func getGDeleteEndpoint(f generic.Factory, tn *graph.TypeNode) (gEndpoint generic.Generic) {
	gEndpoint = f.New(tn.Edges.Type.DeleteEndpoint())

	return
}
