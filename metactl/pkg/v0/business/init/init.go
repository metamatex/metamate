package init

import (
	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

var metactlyaml = `v0:
  gen:
	sdks:
	  - name: go_httpjson_typed
		args:
		  package: github.com/metamatex/metamate
          endpoints: [Whatever, Person]
    tasks:
      - template: path/to/template
        # when iterating, an out template needs to be provided
        out: "path/to/render/to{{ .Fields.Name }}_.go"
        # create one file per node
        iterate: true
        # specify what nodes to provide to the template
        # when iterating, only one kind of nodes can be used
        filter:
          basicTypes:
            flags:
              or: []
              and: []
              nor: []
            names:
              or: []
              nor: []
          endpoints:
          enums:
          fields:
          interfaces:
          relations:
          types:
      - template: path/to/template
        # when not iterating, no template needs to be provided
        out: "path/to/render/to_.go"
        iterate: false
        # when not iterating, all kind of nodes can be provided
        filter: 
          basicTypes:
            flags:
              or: []
              and: []
              nor: []
            names:
              or: []
              nor: []
          endpoints:
          enums:
          fields:
          interfaces:
          relations:
          types:
`

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

	_, err = f.WriteString(metactlyaml)
	if err != nil {
		return
	}

	report.AddInfo("created metactl.yaml")

	return
}
