package v0

import (
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/spf13/cobra"
)

var asgTypeCmd = &cobra.Command{
	Use:   "type",
	Short: "print type info",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		tn, err := d.RootNode.Types.ById(graph.ToNodeId(args[0]))
		if err != nil {
			return
		}

		tn.Print()

		return
	},
}

func addAsgType(parentCmd *cobra.Command) {
	parentCmd.AddCommand(asgTypeCmd)
}
