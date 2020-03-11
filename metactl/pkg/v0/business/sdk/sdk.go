package sdk

import (
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/metactl/pkg/v0/business/gen"
	_go "github.com/metamatex/metamate/metactl/pkg/v0/business/sdk/go"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils/ptr"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"sort"
)

func GetSdks() (sdks []types.Sdk) {
	sdks = append(sdks, _go.GetSdks()...)

	sort.Slice(sdks, func(i, j int) bool {
		return sdks[i].Name < sdks[j].Name
	})

	return sdks
}

func Format(sdks []types.Sdk) (s string) {
	for _, sdk := range sdks {
		s += sdk.Name + " : " + sdk.Description + "\n"
	}

	return
}

func Reset(names []string) (errs []error) {
	for _, name := range names {
		sdk, err := getSdk(name)
		if err != nil {
			errs = append(errs, err)

			return
		}

		err = sdk.Reset()
		if err != nil {
			errs = append(errs, err)

			return
		}
	}

	return
}

func Gen(report *types.MessageReport, fs afero.Fs, version string, rn *graph.RootNode, name string, data map[string]interface{}, endpointFilter *graph.Filter, typeFilter *graph.Filter) (errs []error) {
	sdk, err := getSdk(name)
	if err != nil {
		errs = append(errs, err)

		return
	}

	err = sdk.Init(&sdk, data)
	if err != nil {
		errs = append(errs, err)

		return
	}

	if endpointFilter != nil {
		for i, _ := range sdk.Tasks {
			if sdk.Tasks[i].Dependencies == nil {
				sdk.Tasks[i].Dependencies = &types.RenderTaskDependencies{}
			}

			sdk.Tasks[i].Dependencies.Endpoints = endpointFilter
		}
	}

	if typeFilter != nil {
		for i, _ := range sdk.Tasks {
			if sdk.Tasks[i].Filter == nil {
				sdk.Tasks[i].Dependencies = &types.RenderTaskDependencies{}
			}

			sdk.Tasks[i].Dependencies.Types = typeFilter
		}
	}

	for i, _ := range sdk.Tasks {
		sdk.Tasks[i].Reset = ptr.Bool(false)
	}

	for i, _ := range sdk.Dependencies {
		sdk.Dependencies[i].Reset = ptr.Bool(false)
	}

	errs = gen.Gen(report, fs, version, rn, sdk.Dependencies)
	if len(errs) != 0 {
		return
	}

	errs = gen.Gen(report, fs, version, rn, sdk.Tasks)
	if len(errs) != 0 {
		return
	}

	return
}

func getSdk(name string) (sdk types.Sdk, err error) {
	for _, sdk = range GetSdks() {
		if sdk.Name == name {
			return
		}
	}

	err = errors.New(fmt.Sprintf("sdk %v not found", name))

	return
}
