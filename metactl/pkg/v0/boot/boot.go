package boot

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/expansion"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/spf13/afero"
)

func GetTestDependencies() (d types.Dependencies) {
	d = getDefaultDependencies()

	d.FileSystem = afero.NewMemMapFs()

	return d
}

func GetDependencies(verbosity int, version types.Version) (d types.Dependencies) {
	d = getDefaultDependencies()

	d.FileSystem = afero.NewOsFs()
	d.Version = version

	root := graph.NewRoot()
	root.Wire()

	err := expansion.Expand(verbosity, root)
	if err != nil {
		d.MessageReport.AddError(err)
	}

	errs := root.Validate()
	if len(errs) != 0 {
		d.MessageReport.AddError(errs)
	}

	d.RootNode = root

	d.Factory = generic.NewFactory(root)

	return d
}

func getDefaultDependencies() (d types.Dependencies) {
	d.MessageReport = &types.MessageReport{}

	return
}

func GetGlobalConfig() (c types.GlobalConfig) {
	c = types.GlobalConfig{
		V0: types.V0Global{
			Index: types.Index{
				NGramMinimum: 5,
				NGramMaximum: 15,
				Path:         "/Users/philippwordehoff/.metactl/v0/index",
			},
		},
	}

	return
}

func GetProjectConfig() (c types.ProjectConfig) {
	return
}
