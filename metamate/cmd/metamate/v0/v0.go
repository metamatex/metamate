package v0

import (
	"github.com/metamatex/metamate/metamate/pkg/v0/config"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fs afero.Fs
var gArgs = types.GlobalArgs{}
var version = types.Version{}

var v0Cmd = &cobra.Command{
	Use:   "v0",
	Short: "",
	Long:  "",
}

func AddV0(cmd *cobra.Command, prefix bool, v types.Version) {
	fs = afero.NewOsFs()

	parentCmd := cmd
	if prefix {
		parentCmd.AddCommand(v0Cmd)

		parentCmd = v0Cmd
	}

	version = v

	addInit(parentCmd)
	addServe(parentCmd)
	addUpdate(parentCmd)
	addVersion(parentCmd)

	parentCmd.PersistentFlags().StringVarP(&gArgs.ConfigPath, "config", "c", config.ConfigFile, "config file")
	parentCmd.PersistentFlags().CountVarP(&gArgs.VerbosityLevel, "verbose", "v", "")
}

func readProjectConfig(args types.GlobalArgs) (c types.Config, err error) {
	v := viper.New()

	if args.ConfigPath != config.ConfigFile {
		v.SetConfigFile(args.ConfigPath)
	} else {
		v.AddConfigPath(".")
		v.SetConfigName(config.ConfigFileName)
	}

	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	err = v.Unmarshal(&c)
	if err != nil {
		return
	}

	return
}

func hasConfig(args types.GlobalArgs) (bool) {
	v := viper.New()

	if args.ConfigPath != config.ConfigFile {
		v.SetConfigFile(args.ConfigPath)
	} else {
		v.AddConfigPath(".")
		v.SetConfigName(config.ConfigFileName)
	}

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return false
	}

	return true
}