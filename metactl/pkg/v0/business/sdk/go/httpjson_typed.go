package _go

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils/ptr"
)

const (
	TaskTypedHttpJsonClient  = "TaskTypedHttpJsonClient"
	TaskTypedHttpJsonService = "TaskTypedHttpJsonService"
)

func init() {
	tasks[TaskTypedHttpJsonClient] = types.RenderTask{
		TemplateData: &goTypedHttpJsonClientTpl,
		Out:          ptr.String("httpjson_client_.go"),
	}

	tasks[TaskTypedHttpJsonService] = types.RenderTask{
		Name:         ptr.String(TaskTypedHttpJsonService),
		TemplateData: &goTypedHttpJsonServiceTpl,
		Out:          ptr.String("{{ index .Data \"name\" }}_httpjson_server_.go"),
	}
}

var goTypedHttpJsonClientTpl = `package mql

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
)

type HttpJsonClient struct {
	opts HttpJsonClientOpts
}

type HttpJsonClientOpts struct {
	HttpClient	*http.Client
	Addr	string
}

func NewHttpJsonClient(opts HttpJsonClientOpts) (Client) {
	if opts.HttpClient == nil {
		opts.HttpClient = &http.Client{}
	}
	
	return HttpJsonClient{opts: opts}
}

func (c HttpJsonClient) send(req interface{}, rsp interface{}) (err error) {
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(req)
	if err != nil {
		return
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.opts.Addr, b)
	if err != nil {
		return
	}

	httpReq.Header.Set(ContentTypeHeader, ContentTypeJson)
	httpReq.Header.Set(MetamateTypeHeader, reflect.TypeOf(req).Name())

	res, err := c.opts.HttpClient.Do(httpReq)
	if err != nil {
		return
	}

	err = json.NewDecoder(res.Body).Decode(rsp)
	if err != nil {
		return
	}

	return
}

{{- range $i, $endpoint := .Endpoints.Slice.Sort }}
func (c HttpJsonClient) {{ $endpoint.Name }}(ctx context.Context, req {{ $endpoint.Edges.Type.Request.Name }}) (rsp *{{ $endpoint.Edges.Type.Response.Name }}, err error) {
	err = c.send(req, &rsp)

	return
}
{{- end }}`

var goTypedHttpJsonServiceTpl = `package mql

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type {{ index .Data "name" | title }}HttpJsonServer struct {
	opts {{ index .Data "name" | title }}HttpJsonServerOpts
}

type {{ index .Data "name" | title }}HttpJsonServerOpts struct {
	Service {{ index .Data "name" | title }}Service
}

func New{{ index .Data "name" | title }}HttpJsonServer(opts {{ index .Data "name" | title }}HttpJsonServerOpts) (http.Handler) {
	return {{ index .Data "name" | title }}HttpJsonServer{opts: opts}
}

func (s {{ index .Data "name" | title }}HttpJsonServer) send(w http.ResponseWriter, rsp interface{}) (err error) {
	w.Header().Set(ContentTypeHeader, ContentTypeJson)
	w.Header().Set(MetamateTypeHeader, reflect.TypeOf(rsp).Name())

	err = json.NewEncoder(w).Encode(rsp)
	if err != nil {
	    return
	}

	return
}

func (s {{ index .Data "name" | title }}HttpJsonServer) getService() (Service) {
{{- range $ei, $endpoint := .Endpoints.Slice.Sort }}
{{- if ne $endpoint.Name "LookupService" }}
	{{ $endpoint.FieldName }}Endpoint := s.opts.Service.Get{{ $endpoint.Name }}Endpoint()
{{- end }}
{{- end }}

	return Service{
		Name: String(s.opts.Service.Name()),
		SdkVersion: String(Version),
		Endpoints: &Endpoints{
			LookupService: &LookupServiceEndpoint{},
{{- range $ei, $endpoint := .Endpoints.Slice.Sort }}
{{- if ne $endpoint.Name "LookupService" }}
			{{ $endpoint.Name }}: &{{ $endpoint.FieldName }}Endpoint,
{{- end }}
{{- end }}
		},
	}
}

func (s {{ index .Data "name" | title }}HttpJsonServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get(MetamateTypeHeader) {
	case LookupServiceRequestName:
			var req LookupServiceRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				return
			}
	
			svc := s.getService()
			rsp := LookupServiceResponse{
				Output: &LookupServiceOutput{
					Service: &svc,
				},
			}
	
			err = s.send(w, rsp)
			if err != nil {
				return
			}
{{- range $ei, $endpoint := .Endpoints.Slice.Sort }}
{{- if ne $endpoint.Name "LookupService" }}
    case {{ $endpoint.Edges.Type.Request.Name }}Name:
        var req {{ $endpoint.Edges.Type.Request.Name }}
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            return
        }

        rsp := s.opts.Service.{{ $endpoint.Name }}(r.Context(), req)

        err = s.send(w, rsp)
        if err != nil {
            return
        }
{{- end }}
{{- end }}
	}
}`
