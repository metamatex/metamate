package v0

import (
	"errors"
	"fmt"
	"github.com/metamatex/metamate/metactl/pkg/v0/boot"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var d types.Dependencies
var gArgs = types.GlobalArgs{}
var version = types.Version{}

var v0Cmd = &cobra.Command{
	Use:   "v0",
	Short: "",
	Long:  "",
}

func AddV0(cmd *cobra.Command, prefix bool, v types.Version) {
	d = boot.GetDependencies(0, v)

	parentCmd := cmd
	if prefix {
		parentCmd.AddCommand(v0Cmd)

		parentCmd = v0Cmd
	}

	version = v

	addInit(parentCmd)
	addGen(parentCmd)
	addGet(parentCmd)
	addAsg(parentCmd)
	addVersion(parentCmd)
	addUpdate(parentCmd)

	parentCmd.PersistentFlags().StringVarP(&gArgs.GlobalConfigPath, "config", "c", "$HOME/.metactl/config", "config file")
	parentCmd.PersistentFlags().CountVarP(&gArgs.VerbosityLevel, "verbose", "v", "")
	parentCmd.PersistentFlags().BoolVar(&gArgs.NoColor, "no-color", false, "")
	parentCmd.PersistentFlags().StringVarP(&gArgs.OutputFormat, "output", "o", "default", "one of default|json|yaml")
}

func readProjectConfig(args types.GlobalArgs) (c types.ProjectConfig, err error) {
	v := viper.New()

	if args.ProjectConfigPath != "" {
		v.SetConfigFile(args.ProjectConfigPath)
	} else {
		v.AddConfigPath(".")
		v.SetConfigName("metactl")
	}

	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	d.MessageReport.AddDebug(fmt.Sprintf("using %v", v.ConfigFileUsed()))

	err = v.Unmarshal(&c)
	if err != nil {
		return
	}

	return
}

func hasProjectConfig(args types.GlobalArgs) (err error) {
	v := viper.New()

	if args.ProjectConfigPath != "" {
		v.SetConfigFile(args.ProjectConfigPath)
	} else {
		v.AddConfigPath(".")
		v.SetConfigName("metactl")
	}

	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	return
}

func requireProjectConfig(gArgs types.GlobalArgs) (c types.ProjectConfig, err error) {
	err = hasProjectConfig(gArgs)
	if err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			err = errors.New("metactl.yaml not found, run `metactl init` to create it")
		default:
		}

		return
	}

	c, err = readProjectConfig(gArgs)
	if err != nil {
		return
	}

	return
}
