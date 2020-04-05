package v0

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
	"os"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update metamate",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = func() (err error) {
			latest, found, err := selfupdate.DetectLatest("metamatex/metamate")
			if err != nil {
				err = errors.New(fmt.Sprintf("error occurred while detecting version: %v", err.Error()))

				return
			}

			vString := version.Version
			if vString == "dev-0.0.0" {
				vString = "0.0.0"
			}

			v := semver.MustParse(vString)
			if !found || latest.Version.LTE(v) {
				fmt.Printf("current version %v is the latest", version.Version)

				return
			}


			fmt.Printf("do you want to update to v%v ? (y/n): ", latest.Version)
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

			fmt.Printf("successfully updated to version: v%v", latest.Version)

			return
		}()
		if err != nil {
		    return
		}

		return
	},
}

func addUpdate(parentCmd *cobra.Command) {
	parentCmd.AddCommand(updateCmd)
}
