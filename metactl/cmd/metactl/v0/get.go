package v0

import (
	"github.com/metamatex/metamatemono/metactl/pkg/v0/business/get"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		o, err := get.Get(
			d.RootNode,
			d.Factory,
			gArgs.Addr,
			gArgs.Token,
			args[0],
		)
		if err != nil {
			d.MessageReport.AddError(err)
		}

		handleReport(*d.MessageReport, o, gArgs.VerbosityLevel)

		return
	},
}

func addGet(parentCmd *cobra.Command) {
	parentCmd.AddCommand(getCmd)
}
