package v0

import (
	"fmt"
	"github.com/metamatex/metamate/metamate/pkg/v0/boot"
	"github.com/metamatex/metamate/metamate/pkg/v0/config"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		go func() {
			err := func() (err error) {
				c:= types.Config{}

				switch hasConfig(gArgs) {
				case true:
					c, err = readProjectConfig(gArgs)
					if err != nil {
						return
					}
				case false:
					c = config.DefaultConfig
				}

				d, err := boot.NewDependencies(c, version)
				if err != nil {
					return
				}

				fmt.Printf("version: %v\nvcommit: %v\ndate: %v\n\n", version.Version, version.Commit, version.Date)

				for _, r := range d.Routes {
					for _, m := range r.Methods {
						fmt.Printf("%v: %v:%v%v\n", m, c.Host.Bind, c.Host.HttpPort, r.Path)
					}
				}

				err = http.ListenAndServe(fmt.Sprintf(":%v", c.Host.HttpPort), d.Router)
				if err != nil {
					return
				}

				return
			}()
			if err != nil {
				panic(err)
			}
		}()

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs

		return
	},
}

func addServe(parentCmd *cobra.Command) {
	parentCmd.AddCommand(serveCmd)
}
