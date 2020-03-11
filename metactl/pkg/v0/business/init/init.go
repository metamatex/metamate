package init

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/metactl/pkg/v0/business/sdk"
	_go "github.com/metamatex/metamate/metactl/pkg/v0/business/sdk/go"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
	"html/template"
)

type templateData struct {
	Sdks []types.Sdk
	Config string
}

var metactlyaml = `# this file is used by metactl to generate client and service sdks
#
# available sdks:
#
{{- range $i, $sdk := .Sdks }}
# {{ $sdk.Name }}: {{ $sdk.Description }} 
{{- end }}

{{ .Config }} 
`

var initialConfig = types.V0Project{
	Gen: types.Gen{
		Sdks: []types.ProjectSdk{
			{
				Names: []string{_go.SdkHttpJsonService},
				Data: map[string]interface{}{
					"name": "socialservice",
					"package": "github.com/somebody/socialservice",
				},
				Endpoints: &graph.Filter{
					Names: &graph.NamesSubset{
						Or: []string{"GetFeeds", "GetPeople", "PutPeople", "GetStatuses", "PostStatuses", "PutStatuses", "DeleteStatuses"},
					},
				},
			},
			{
				Names: []string{_go.SdkHttpJsonClient},
				Data: map[string]interface{}{
					"package": "github.com/somebody/socialclient",
				},
				Endpoints: &graph.Filter{
					Names: &graph.NamesSubset{
						Or: []string{"GetFeeds", "GetPeople", "PutPeople", "GetStatuses", "PostStatuses", "PutStatuses", "DeleteStatuses"},
					},
				},
			},
		},
	},
}

func Init(fs afero.Fs, report *types.MessageReport) (err error) {
	_, err = fs.Stat("metactl.yaml")
	if err == nil {
		err = errors.New("metactl.yaml already present")

		return
	}

	f, err := fs.Create("metactl.yaml")
	if err != nil {
		return
	}

	b, err := yaml.Marshal(initialConfig)
	if err != nil {
		return
	}

	err = template.Must(template.New("").Parse(metactlyaml)).Execute(f, templateData{Config:string(b), Sdks: sdk.GetSdks()})
	if err != nil {
		return
	}

	report.AddInfo("created metactl.yaml")

	return
}
