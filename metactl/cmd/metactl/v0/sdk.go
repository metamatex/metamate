package v0

import (
	"github.com/spf13/cobra"
)

var sdkCmd = &cobra.Command{
	Use:   "sdk",
	Short: "generate sdk",
	Long:  "",
}

func addSdk(parentCmd *cobra.Command) {
	addSdkList(sdkCmd)
	addSdkGen(sdkCmd)

	parentCmd.AddCommand(sdkCmd)
}
