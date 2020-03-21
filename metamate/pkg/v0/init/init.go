package init

import (
	"github.com/metamatex/metamate/metamate/pkg/v0/config"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

func Init(fs afero.Fs) (err error) {
	_, err = fs.Stat("metamate.yaml")
	if err == nil {
		err = errors.New("metamate.yaml already present")

		return
	}

	f, err := fs.Create("metamate.yaml")
	if err != nil {
		return
	}

	err = yaml.NewEncoder(f).Encode(config.DefaultConfig)
	if err != nil {
		return
	}

	return
}
