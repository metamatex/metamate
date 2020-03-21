package main

import (
	"fmt"
	"github.com/metamatex/metamate/metamate/cmd/metamate/v0"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/spf13/cobra"
	"os"
)

var (
	version = "dev-0.0.0"
	commit  = "dev"
	date    = "dev"
)

var rootCmd = &cobra.Command{
	Use:   "metamate",
	Short: "",
	Long:  "",
}

func main() {
	v := types.Version{
		Version: version,
		Commit:  commit,
		Date:    date,
	}

	v0.AddV0(rootCmd, false, v)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
