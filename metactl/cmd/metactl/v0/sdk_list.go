package v0

import (
	"github.com/metamatex/metamatemono/metactl/pkg/v0/business/sdk"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
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

		handleReport(*d.MessageReport, o, c.VerbosityLevel)

		return
	},
}

func addSdkList(parentCmd *cobra.Command) {
	parentCmd.AddCommand(sdkListCmd)
}
