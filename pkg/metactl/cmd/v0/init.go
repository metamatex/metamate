package v0

import (
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/types"
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/utils"
	"github.com/spf13/cobra"

	init0 "github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/business/init"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create project c",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = init0.Init(d.FileSystem, d.MessageReport)
		if err != nil {
			d.MessageReport.AddError(err)
		}

		err = utils.PrintReport(c.NoColor, c.OutputFormat, c.ReturnData(), *d.MessageReport, types.Output{})
		if err != nil {
			return
		}

		return
	},
}

func addInit(parentCmd *cobra.Command) {
	parentCmd.AddCommand(initCmd)
}
