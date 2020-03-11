package v0

import (
	"fmt"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version info",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var t string
		switch gArgs.VerbosityLevel {
		case 0:
			t = version.Version
		default:
			t = fmt.Sprintf("version: %v\ncommit: %v\ndate: %v\n", version.Version, version.Commit, version.Date)
		}

		o := types.Output{
			Data: version,
			Text: t,
		}

		handleReport(*d.MessageReport, o, gArgs.VerbosityLevel)

		return
	},
}

func addVersion(parentCmd *cobra.Command) {
	parentCmd.AddCommand(versionCmd)
}
