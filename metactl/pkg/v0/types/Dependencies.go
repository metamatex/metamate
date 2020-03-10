package types

import (
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/generic/pkg/v0/generic"

	"github.com/spf13/afero"
)

type Dependencies struct {
	VerbosityLevel int
	Version        string
	FileSystem     afero.Fs
	ProjectConfig  *ProjectConfig
	GlobalConfig   GlobalConfig
	MessageReport  *MessageReport
	RootNode       *graph.RootNode
	Factory        generic.Factory
	GlobalArgs     GlobalArgs
	//Index          bleve.Index
}
