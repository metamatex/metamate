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
		Out:          ptr.String("transport/httpjson_client_.go"),
	}

	tasks[TaskTypedHttpJsonService] = types.RenderTask{
		Name:         ptr.String(TaskTypedHttpJsonService),
		TemplateData: &goTypedHttpJsonServiceTpl,
		Out:          ptr.String("transport/services/{{ index .Data \"name\" }}/httpjson_server_.go"),
	}
}

var goTypedHttpJsonClientTpl = `package transport
{{ $package := index .Data "package" }}
import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"{{ $package }}/gen/v0/sdk"
)

type HttpJsonClient struct {
	opts HttpJsonClientOpts
}

type HttpJsonClientOpts struct {
	HttpClient	*http.Client
	Token	string
	Addr	string
}

func NewHttpJsonClient(opts HttpJsonClientOpts) (Client) {
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

	httpReq.Header.Set(CONTENT_TYPE_HEADER, CONTENT_TYPE_JSON)
	httpReq.Header.Set(METAMATE_TYPE_HEADER, reflect.TypeOf(req).Name())
	httpReq.Header.Set(AUTHORIZATION_HEADER, "Bearer " + c.opts.Token)

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
func (c HttpJsonClient) {{ $endpoint.Name }}(ctx context.Context, req sdk.{{ $endpoint.Edges.Type.Request.Name }}) (rsp *sdk.{{ $endpoint.Edges.Type.Response.Name }}, err error) {
	err = c.send(req, &rsp)

	return
}
{{- end }}`

var goTypedHttpJsonServiceTpl = `package {{ index .Data "name" }}
{{ $package := index .Data "package" }}
import (
	"encoding/json"
	"net/http"
	"reflect"
	"{{ $package }}/gen/v0/sdk"
	"{{ $package }}/gen/v0/sdk/utils/ptr"
	"{{ $package }}/gen/v0/sdk/transport"
)

type HttpJsonServer struct {
	opts HttpJsonServerOpts
}

type HttpJsonServerOpts struct {
	Service Service
}

func NewHttpJsonServer(opts HttpJsonServerOpts) (http.Handler) {
	return HttpJsonServer{opts: opts}
}

func (s HttpJsonServer) send(w http.ResponseWriter, rsp interface{}) (err error) {
	w.Header().Set(transport.CONTENT_TYPE_HEADER, transport.CONTENT_TYPE_JSON)
	w.Header().Set(transport.METAMATE_TYPE_HEADER, reflect.TypeOf(rsp).Name())

	err = json.NewEncoder(w).Encode(rsp)
	if err != nil {
	    return
	}

	return
}

func (s HttpJsonServer) getService() (sdk.Service) {
{{- range $ei, $endpoint := .Endpoints.Slice.Sort }}
{{- if ne $endpoint.Name "LookupService" }}
	{{ $endpoint.FieldName }}Endpoint := s.opts.Service.Get{{ $endpoint.Name }}Endpoint()
{{- end }}
{{- end }}

	return sdk.Service{
		Name: ptr.String(s.opts.Service.Name()),
		Endpoints: &sdk.Endpoints{
			LookupService: &sdk.LookupServiceEndpoint{},
{{- range $ei, $endpoint := .Endpoints.Slice.Sort }}
{{- if ne $endpoint.Name "LookupService" }}
			{{ $endpoint.Name }}: &{{ $endpoint.FieldName }}Endpoint,
{{- end }}
{{- end }}
		},
	}
}

func (s HttpJsonServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get(transport.METAMATE_TYPE_HEADER) {
	case sdk.LookupServiceRequestName:
			var req sdk.LookupServiceRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				return
			}
	
			svc := s.getService()
			rsp := sdk.LookupServiceResponse{
				Output: &sdk.LookupServiceOutput{
					Service: &svc,
				},
			}
	
			err = s.send(w, rsp)
			if err != nil {
				return
			}
{{- range $ei, $endpoint := .Endpoints.Slice.Sort }}
{{- if ne $endpoint.Name "LookupService" }}
    case sdk.{{ $endpoint.Edges.Type.Request.Name }}Name:
        var req sdk.{{ $endpoint.Edges.Type.Request.Name }}
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
