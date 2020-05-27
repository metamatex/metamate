package types

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"

	"github.com/spf13/afero"
)

type Dependencies struct {
	Version       Version
	FileSystem    afero.Fs
	ProjectConfig *ProjectConfig
	GlobalConfig  GlobalConfig
	MessageReport *MessageReport
	RootNode      *graph.RootNode
	Factory       generic.Factory
	GlobalArgs    GlobalArgs
	//Index          bleve.Index
}
