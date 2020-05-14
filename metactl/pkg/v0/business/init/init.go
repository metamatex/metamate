package init

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/business/sdk"
	_go "github.com/metamatex/metamate/metactl/pkg/v0/business/sdk/go"
	"github.com/metamatex/metamate/metactl/pkg/v0/business/sdk/typescript"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
	"html/template"
)

type templateData struct {
	Sdks   []types.SdkGenerator
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

var initialConfig = types.ProjectConfig{
	V0: types.V0Project{
		Gen: types.Gen{
			Sdks: []types.SdkConfig{
				{
					Name: _go.SdkHttpJsonService,
					Args: map[string]interface{}{
						"name":    "socialservice",
						"package": "github.com/somebody/socialservice",
					},
					Endpoints: []string{"GetPostFeeds", "GetSocialAccounts", "GetPosts"},
				},
				{
					Name: _go.SdkHttpJsonClient,
					Args: map[string]interface{}{
						"package": "github.com/somebody/socialclient",
					},
					Endpoints: []string{"GetPostFeeds", "GetSocialAccounts", "GetPosts"},
				},
				{
					Name: typescript.SdkHttpJsonClient,
					Args: map[string]interface{}{
						"path": "src",
					},
					Endpoints: []string{"GetPostFeeds", "GetSocialAccounts", "GetPosts"},
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

	err = template.Must(template.New("").Parse(metactlyaml)).Execute(f, templateData{Config: string(b), Sdks: sdk.GetSdks()})
	if err != nil {
		return
	}

	report.AddInfo("created metactl.yaml")

	return
}
