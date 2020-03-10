package v0

import (
	"errors"
	"fmt"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/boot"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var d = boot.GetDependencies(0)
var gArgs = types.GlobalArgs{}
var version = types.Version{}

var v0Cmd = &cobra.Command{
	Use:   "v0",
	Short: "",
	Long:  "",
}

func AddV0(cmd *cobra.Command, prefix bool, v types.Version) {
	parentCmd := cmd
	if prefix {
		parentCmd.AddCommand(v0Cmd)

		parentCmd = v0Cmd
	}

	version = v

	addGen(parentCmd)
	addInit(parentCmd)
	addAsg(parentCmd)
	addSdk(parentCmd)
	addGet(parentCmd)
	addVersion(parentCmd)
	addUpdate(parentCmd)

	parentCmd.PersistentFlags().StringVarP(&gArgs.GlobalConfigPath, "config", "c", "$HOME/.metactl/config", "config file")
	parentCmd.PersistentFlags().StringVar(&gArgs.Addr, "addr", "", "")
	parentCmd.PersistentFlags().StringVar(&gArgs.Token, "token", "", "")
	parentCmd.PersistentFlags().CountVarP(&gArgs.VerbosityLevel, "verbose", "v", "")
	parentCmd.PersistentFlags().BoolVar(&gArgs.NoColor, "no-color", false, "")
	parentCmd.PersistentFlags().StringVarP(&gArgs.OutputFormat, "output", "o", "default", "Output format. One of: default|json|yaml")
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

func handleReport(r types.MessageReport, o types.Output, verbosityLevel int) {
	printR := types.MessageReport{
		Debug:   r.Debug,
		Info:    r.Info,
		Warning: r.Warning,
		Error:   r.Error,
	}

	switch verbosityLevel {
	case 0:
		printR.Debug = []string{}
	default:
	}

	err := utils.PrintReport(gArgs.NoColor, gArgs.OutputFormat, gArgs.ReturnData(), printR, o)
	if err != nil {
		log.Fatal(err)
	}

	if len(r.Error) != 0 {
		os.Exit(1)
	}
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
