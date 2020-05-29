package v0

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/business/get"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils"
	"github.com/spf13/cobra"
	"strings"
)

var getArgs = types.GetArgs{}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get entities",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var o types.Output
		errs := func() (errs []error) {
			getArgs.TypePlural = strings.ToLower(args[0])
			o, err = get.Get(d.MessageReport, gArgs.VerbosityLevel, d.RootNode, d.Factory, getArgs)
			if err != nil {
				return
			}

			return
		}()
		if len(errs) != 0 {
			d.MessageReport.AddError(errs)
		}

		utils.HandleReport(gArgs, *d.MessageReport, o, gArgs.VerbosityLevel)

		return
	},
}

func addGet(parentCmd *cobra.Command) {
	getCmd.PersistentFlags().StringVarP(&getArgs.Id, "id", "", "", "serviceId id in the format of serviceName/value")
	getCmd.PersistentFlags().StringVarP(&getArgs.Name, "name", "", "", "name id")
	getCmd.PersistentFlags().StringVarP(&getArgs.Username, "username", "", "", "username id")
	getCmd.PersistentFlags().StringVarP(&getArgs.Services, "services", "", "", "serviceFilter in the format of serviceName,serviceName")
	getCmd.PersistentFlags().StringVarP(&getArgs.Host, "host", "", "", "host of MetaMate instance")
	getCmd.PersistentFlags().StringVarP(&getArgs.Password, "password", "", "", "basic auth password")
	getCmd.PersistentFlags().StringVarP(&getArgs.User, "user", "", "", "basic auth user")
	getCmd.PersistentFlags().StringVarP(&getArgs.Path, "path", "", "", "relations path")
	getCmd.PersistentFlags().StringVarP(&getArgs.Search, "search", "", "", "search term")

	parentCmd.AddCommand(getCmd)
}
