package _go

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils/ptr"
)

const (
	TaskTypedClient = "TaskTypedClient"
	TaskTypedServer = "TaskTypedServer"
)

func init() {
	tasks[TaskTypedClient] = types.RenderTask{
		TemplateData: &goTypedClientTpl,
		Out:          ptr.String("client_.go"),
	}

	tasks[TaskTypedServer] = types.RenderTask{
		Name:         ptr.String(TaskTypedServer),
		TemplateData: &goTypedServiceTpl,
		Out:          ptr.String("{{ index .Data \"name\" }}_server_.go"),
	}
}

var goTypedClientTpl = `package mql

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
)

type Client struct {
	opts ClientOpts
}

type ClientOpts struct {
	HttpClient	*http.Client
	Addr	string
}

func NewClient(opts ClientOpts) (Client) {
	if opts.HttpClient == nil {
		opts.HttpClient = &http.Client{}
	}
	
	return Client{opts: opts}
}

func (c Client) send(req interface{}, rsp interface{}) (err error) {
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
	httpReq.Header.Set(AsgTypeHeader, reflect.TypeOf(req).Name())

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
func (c Client) {{ $endpoint.Name }}(ctx context.Context, req {{ $endpoint.Edges.Type.Request.Name }}) (rsp *{{ $endpoint.Edges.Type.Response.Name }}, err error) {
	err = c.send(req, &rsp)

	return
}
{{- end }}`

var goTypedServiceTpl = `package mql

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type {{ index .Data "name" | title }}Server struct {
	opts {{ index .Data "name" | title }}ServerOpts
}

type {{ index .Data "name" | title }}ServerOpts struct {
	Service {{ index .Data "name" | title }}Service
}

func New{{ index .Data "name" | title }}Server(opts {{ index .Data "name" | title }}ServerOpts) (http.Handler) {
	return {{ index .Data "name" | title }}Server{opts: opts}
}

func (s {{ index .Data "name" | title }}Server) send(w http.ResponseWriter, rsp interface{}) (err error) {
	w.Header().Set(ContentTypeHeader, ContentTypeJson)
	w.Header().Set(AsgTypeHeader, reflect.TypeOf(rsp).Name())

	err = json.NewEncoder(w).Encode(rsp)
	if err != nil {
	    return
	}

	return
}

func (s {{ index .Data "name" | title }}Server) getService() (Service) {
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

func (s {{ index .Data "name" | title }}Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get(AsgTypeHeader) {
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
