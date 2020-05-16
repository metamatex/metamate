package typescript

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils/ptr"
)

const (
	TaskClient = "TaskClient"
)

func init() {
	tasks[TaskClient] = types.RenderTask{
		TemplateData: &ClientTpl,
		Out:          ptr.String("mql_.ts"),
	}
}

var ClientTpl = `import * as axios from 'axios';

export interface ClientOpts {
    client: axios.AxiosInstance;
    addr: string;
}

export class Client {
    opts: ClientOpts;

    constructor(opts: ClientOpts) {
        this.opts = opts;
    }
    {{ range $i, $endpoint := .Endpoints }}
    async {{ $endpoint.Name }}(req: {{ $endpoint.Edges.Type.Request.Name }}): Promise<{{ $endpoint.Edges.Type.Response.Name }}> {
        let rsp = await this.opts.client.request<{{ $endpoint.Edges.Type.Response.Name }}>({
            url: this.opts.addr,
            method: "post",
            data: req,
            headers: {
                "X-Asg-type": "{{ $endpoint.Edges.Type.Request.Name }}",
                "Content-Type": "application/json; charset=utf-8",
            },
        });

        return rsp.data
    }
    {{ end }}
}
{{ range $i, $enum := .Enums }}
export const {{ $enum.Name }} =  Object.freeze({
{{- range $vi, $value := sortAlpha $enum.Data.Values }}
    {{ camel $value }}: "{{ $value }}",
{{- end }}
});
{{ end }}

{{- range $i, $type := .Types }}
export interface {{ $type.Name }} {
{{- range $i, $fn := $type.Edges.Fields.Holds.Slice.Sort }}
{{- if $fn.IsBool }}
    {{ $fn.Name }}?: boolean;
{{- end }}
{{- if $fn.IsFloat64 }}
    {{ $fn.Name }}?: number;
{{- end }}
{{- if $fn.IsString }}
    {{ $fn.Name }}?: string;
{{- end }}
{{- if $fn.IsInt32 }}
    {{ $fn.Name }}?: number;
{{- end }}
{{- if $fn.IsType }}
    {{ $fn.Name }}?: {{ camel $fn.Edges.Type.Holds.Name }};
{{- end }}
{{- if $fn.IsEnum }}
    {{ $fn.Name }}?: string;
{{- end }}
{{- if $fn.IsStringList }}
    {{ $fn.Name }}?: string[];
{{- end }}
{{- if $fn.IsInt32List }}
    {{ $fn.Name }}?: number[];
{{- end }}
{{- if $fn.IsFloat64List }}
    {{ $fn.Name }}?: number[];
{{- end }}
{{- if $fn.IsBoolList }}
    {{ $fn.Name }}?: boolean[];
{{- end }}
{{- if $fn.IsTypeList }}
    {{ $fn.Name }}?: {{ camel $fn.Edges.Type.Holds.Name }}[];
{{- end }}
{{- if $fn.IsEnumList }}
    {{ $fn.Name }}?: string[];
{{- end }}
{{- end }}
}
{{ end }}`
