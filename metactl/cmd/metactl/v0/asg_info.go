package v0

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/business/asg"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils"
	"github.com/spf13/cobra"
)

var asgInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "print info",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		r, err := asg.Info(d.RootNode)
		if err != nil {
			d.MessageReport.AddError(err)
		}

		o := types.Output{}
		o.Data = r

		utils.HandleReport(gArgs, *d.MessageReport, o, gArgs.VerbosityLevel)

		return
	},
}

func addAsgInfo(parentCmd *cobra.Command) {
	parentCmd.AddCommand(asgInfoCmd)
}
