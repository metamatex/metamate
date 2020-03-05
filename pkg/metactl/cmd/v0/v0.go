package v0

import (
	"fmt"
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/boot"
	"github.com/metamatex/metamatemono/pkg/metactl/pkg/v0/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var d = boot.GetDependencies(0)
var c = types.Config{}
var projectConfig *types.ProjectConfig

func init() {
	cobra.OnInitialize(initConfig)
}

var v0Cmd = &cobra.Command{
	Use:   "v0",
	Short: "",
	Long:  "",
}

func AddV0(cmd *cobra.Command, prefix bool) {
	parentCmd := cmd
	if prefix {
		parentCmd.AddCommand(v0Cmd)

		parentCmd = v0Cmd
	}

	addGen(parentCmd)
	addInit(parentCmd)
	addAsg(parentCmd)
	addSdk(parentCmd)
	addGet(parentCmd)

	parentCmd.PersistentFlags().StringVar(&c.GlobalConfigPath, "c", "$HOME/.metactl/c", "c file")
	parentCmd.PersistentFlags().StringVar(&c.Addr, "addr", "", "")
	parentCmd.PersistentFlags().StringVar(&c.Token, "token", "", "")
	parentCmd.PersistentFlags().CountVarP(&c.VerbosityLevel, "verbose", "v", "")
	parentCmd.PersistentFlags().BoolVar(&c.NoColor, "no-color", false, "")
	parentCmd.PersistentFlags().StringVarP(&c.OutputFormat, "output", "o", "default", "Output format. One of: default|json|yaml")
}

func initConfig() {
	err := func() (err error) {
		v := viper.New()

		if c.ProjectConfigPath != "" {
			v.SetConfigFile(c.ProjectConfigPath)
		} else {
			v.AddConfigPath(".")
			v.SetConfigName("metactl")
		}

		v.AutomaticEnv()

		err = v.ReadInConfig()
		if err != nil {
			return
		}

		d.MessageReport.AddDebug(fmt.Sprintf("using c file: %v", v.ConfigFileUsed()))

		err = v.Unmarshal(&projectConfig)
		if err != nil {
			return
		}

		return
	}()
	if err != nil {
		d.MessageReport.AddError(err)
	}
}

func requireProjectConfig() (err error) {
	if projectConfig == nil {
		err = errors.New("no metactl.yaml found, to create it, run `metactl init`")

		return
	}

	return
}
