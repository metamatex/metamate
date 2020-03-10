package v0

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
	"os"

	//"log"
)

var versiona = "0.0.1"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = func() (err error) {
			latest, found, err := selfupdate.DetectLatest("metamatex/metamate")
			if err != nil {
				err = errors.New(fmt.Sprintf("error occurred while detecting version: %v", err.Error()))

				return
			}

			v := semver.MustParse(versiona)
			if !found || latest.Version.LTE(v) {
				d.MessageReport.AddInfo(fmt.Sprintf("current version %v is the latest", versiona))

				return
			}

			prefix := "\u001B[36mask\u001b[0m"
			if gArgs.NoColor {
				prefix = "ask"
			}

			fmt.Printf("%v do you want to update to v%v? (y/n): ", prefix, latest.Version)
			input, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil || (input != "y\n" && input != "n\n") {
				err = errors.New("invalid input")

				return
			}
			if input == "n\n" {
				return
			}

			exe, err := os.Executable()
			if err != nil {
				err = errors.New("could not locate executable path")

				return
			}
			err = selfupdate.UpdateTo(latest.AssetURL, exe)
			if err != nil {
				err = errors.New(fmt.Sprintf("error occurred while updating binary: %v", err.Error()))

			    return
			}

			d.MessageReport.AddInfo(fmt.Sprintf("successfully updated to version: v%v", latest.Version))

			return
		}()
		if err != nil {
			d.MessageReport.AddError(err)

		    return
		}

		handleReport(*d.MessageReport, types.Output{}, gArgs.VerbosityLevel)

		return
	},
}

func addUpdate(parentCmd *cobra.Command) {
	parentCmd.AddCommand(updateCmd)
}
