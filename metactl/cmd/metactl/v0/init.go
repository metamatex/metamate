package v0

import (
	init0 "github.com/metamatex/metamate/metactl/pkg/v0/business/init"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize metactl.yaml",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = init0.Init(d.FileSystem, d.MessageReport)
		if err != nil {
			d.MessageReport.AddError(err)
		}

		utils.HandleReport(gArgs, *d.MessageReport, types.Output{}, gArgs.VerbosityLevel)

		return
	},
}

func addInit(parentCmd *cobra.Command) {
	parentCmd.AddCommand(initCmd)
}
