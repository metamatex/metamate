package v0

import (
	"github.com/spf13/cobra"
)

var asgCmd = &cobra.Command{
	Use:   "asg",
	Short: "inspect the abstract schema graph",
	Long:  "",
}

func addAsg(parentCmd *cobra.Command) {
	addAsgInfo(asgCmd)
	addAsgFlags(asgCmd)
	addAsgType(asgCmd)

	parentCmd.AddCommand(asgCmd)
}
