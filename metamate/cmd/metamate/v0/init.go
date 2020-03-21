package v0

import (
	"errors"
	"fmt"
	"github.com/metamatex/metamate/metamate/pkg/v0/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: fmt.Sprintf("initialize %v", config.ConfigFile),
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		_, err = fs.Stat(config.ConfigFile)
		if err == nil {
			err = errors.New(fmt.Sprintf("%v already present", config.ConfigFile))

			return
		}

		f, err := fs.Create(config.ConfigFile)
		if err != nil {
			return
		}

		err = yaml.NewEncoder(f).Encode(config.DefaultConfig)
		if err != nil {
			return
		}

		return
	},
}

func addInit(parentCmd *cobra.Command) {
	parentCmd.AddCommand(initCmd)
}
