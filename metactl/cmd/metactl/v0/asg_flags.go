package v0

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/business/asg"
	"github.com/spf13/cobra"
)

var asgFlagsCmd = &cobra.Command{
	Use:   "flags",
	Short: "print available flags",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		o := asg.Flags(gArgs.ReturnData())

		handleReport(*d.MessageReport, o, gArgs.VerbosityLevel)

		return
	},
}

func addAsgFlags(parentCmd *cobra.Command) {
	parentCmd.AddCommand(asgFlagsCmd)
}
