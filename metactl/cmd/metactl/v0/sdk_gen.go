package v0

import (
	"errors"
	"fmt"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/business/sdk"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
	"github.com/spf13/cobra"
	"strings"
)

var packageArg string
var endpointsArg string
var typesArgs string

func init() {
	sdkGenCmd.PersistentFlags().StringVarP(&packageArg, "package", "p", "", "")
	sdkGenCmd.PersistentFlags().StringVarP(&endpointsArg, "endpoints", "e", "", "")
	sdkGenCmd.PersistentFlags().StringVarP(&typesArgs, "types", "t", "", "")
}

func addSdkGen(parentCmd *cobra.Command) {
	parentCmd.AddCommand(sdkGenCmd)
}

var sdkGenCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate sdk",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		errs := func() (errs []error) {
			data := map[string]interface{}{}

			if packageArg != "" {
				data["package"] = packageArg
			}

			typeFilter, errs := getTypeFilter(typesArgs)
			if len(errs) != 0 {
				return
			}

			endpointFilter, errs := getEndpointFilter(endpointsArg)
			if len(errs) != 0 {
				return
			}

			sdkNames := strings.Split(args[0], ",")

			errs = sdk.Reset(sdkNames)
			if len(errs) != 0 {
				return
			}

			for _, sdkName := range strings.Split(args[0], ",") {
				errs = sdk.Gen(d.MessageReport, d.FileSystem, d.Version, d.RootNode, sdkName, data, endpointFilter, typeFilter)
				if len(errs) != 0 {
					return
				}
			}

			return
		}()
		if len(errs) != 0 {
			d.MessageReport.AddError(errs)
		}

		handleReport(*d.MessageReport, types.Output{}, gArgs.VerbosityLevel)

		return
	},
}

func getEndpointFilter(endpointsArg string) (f *graph.Filter, errs []error) {
	endpointNames := strings.Split(endpointsArg, ",")
	if len(endpointNames) == 1 && endpointNames[0] == "" {
		return
	}

	for _, endpointName := range endpointNames {
		if !d.RootNode.Endpoints.HasName(endpointName) {
			errs = append(errs, errors.New(fmt.Sprintf("endpoint %v unknown", endpointName)))
		}
	}
	if len(errs) != 0 {
		return
	}

	f = &graph.Filter{
		Names: &graph.NamesSubset{
			Or: endpointNames,
		},
	}

	return
}

func getTypeFilter(typesArg string) (f *graph.Filter, errs []error) {
	typeNames := strings.Split(typesArg, ",")
	if len(typeNames) == 1 && typeNames[0] == "" {
		return
	}

	for _, typeName := range typeNames {
		if !d.RootNode.Types.HasName(typeName) {
			errs = append(errs, errors.New(fmt.Sprintf("type %v unknown", typeName)))
		}
	}
	if len(errs) != 0 {
		return
	}

	f = &graph.Filter{
		Names: &graph.NamesSubset{
			Or: typeNames,
		},
	}

	return
}
