package v0

import (
	"github.com/blang/semver"
	"github.com/davecgh/go-spew/spew"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
	//"log"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		v := semver.MustParse("0.0.1")

		spew.Dump(v.String())

		r, b, err := selfupdate.DetectLatest("metamatex/metamate")
		if err != nil {
		    return
		}

		spew.Dump(b)
		spew.Dump(r)

		//latest, err := selfupdate.UpdateSelf(v, "metamatex/metamate")
		//if err != nil {
		//	log.Println("Binary update failed:", err)
		//	return
		//}
		//if latest.Version.Equals(v) {
		//	log.Println("Current binary is the latest version", version)
		//} else {
		//	log.Println("Successfully updated to version", latest.Version)
		//	log.Println("Release note:\n", latest.ReleaseNotes)
		//}

		return
	},
}

func addUpdate(parentCmd *cobra.Command) {
	parentCmd.AddCommand(updateCmd)
}
