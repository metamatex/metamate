package v0

import (
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/business/get"
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/utils"
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
			c.Addr,
			c.Token,
			args[0],
		)
		if err != nil {
			d.MessageReport.AddError(err)
		}

		err = utils.PrintReport(c.NoColor, c.OutputFormat, c.ReturnData(), *d.MessageReport, o)
		if err != nil {
			return
		}

		return
	},
}

func addGet(parentCmd *cobra.Command) {
	parentCmd.AddCommand(getCmd)
}
