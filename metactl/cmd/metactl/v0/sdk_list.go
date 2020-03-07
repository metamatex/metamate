package v0

import (
	"github.com/metamatex/metamatemono/metactl/pkg/v0/business/sdk"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/utils"
	"github.com/spf13/cobra"
)

var sdkListCmd = &cobra.Command{
	Use:   "list",
	Short: "list available sdks",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		sdks := sdk.GetSdks()

		o := types.Output{}
		o.Data = sdks
		o.Text = sdk.Format(sdks)

		err = utils.PrintReport(c.NoColor, c.OutputFormat, c.ReturnData(), *d.MessageReport, o)
		if err != nil {
			return
		}

		return
	},
}

func addSdkList(parentCmd *cobra.Command) {
	parentCmd.AddCommand(sdkListCmd)
}
