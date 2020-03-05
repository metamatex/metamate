package v0

import (
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/business/asg"
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/utils"
	"github.com/spf13/cobra"
)

var asgFlagsCmd = &cobra.Command{
	Use:   "flags",
	Short: "print available flags",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		o := asg.Flags(c.ReturnData())

		err = utils.PrintReport(c.NoColor, c.OutputFormat, c.ReturnData(), *d.MessageReport, o)
		if err != nil {
			return
		}

		return
	},
}

func addAsgFlags(parentCmd *cobra.Command) {
	parentCmd.AddCommand(asgFlagsCmd)
}
