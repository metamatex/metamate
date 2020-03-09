package v0

import (
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
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

		handleReport(*d.MessageReport, types.Output{}, c.VerbosityLevel)

		return
	},
}

func addAsgType(parentCmd *cobra.Command) {
	parentCmd.AddCommand(asgTypeCmd)
}
