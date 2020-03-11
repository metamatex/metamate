package gen

import (
	"path/filepath"

	"github.com/spf13/afero"

	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"

	"github.com/metamatex/metamate/metactl/pkg/v0/types"
)

func resetDirectories(fs afero.Fs, tasks []types.RenderTask) (err error) {
	resettedDirs := map[string]bool{}

	for _, task := range tasks {
		if task.Reset != nil {
			if !*task.Reset {
				continue
			}
		}

		var wildcardPath string
		wildcardPath, err = getWildcardPath(task, nil)
		if err != nil {
			return
		}

		dir := filepath.Dir(wildcardPath)

		resetted := resettedDirs[dir]
		if !resetted {
			err = reset(fs, task)
			if err != nil {
				return
			}

			resettedDirs[dir] = true
		}
	}

	return
}

func reset(fs afero.Fs, renderTask types.RenderTask) (err error) {
	p, err := getWildcardPath(renderTask, nil)
	if err != nil {
		return
	}

	files, err := afero.Glob(fs, p)
	if err != nil {
		return
	}

	for _, file := range files {
		err = fs.RemoveAll(file)
		if err != nil {
			return
		}
	}

	return
}

func getWildcardPath(renderTask types.RenderTask, renderContext interface{}) (path string, err error) {
	if renderContext == nil {
		bn := graph.NewBasicTypeNode()
		bn.SetName("*")

		en := graph.NewEndpointNode()
		en.SetName("*")

		enumN := graph.NewEnumNode()
		enumN.SetName("*")

		fn := graph.NewFieldNode()
		fn.SetName("*")

		rn := graph.NewRelationNode()
		rn.SetName("*")

		tn := graph.NewTypeNode()
		tn.SetName("*")

		renderContext = &types.IterateRenderContext{
			BasicType: bn,
			Endpoint:  en,
			Enum:      enumN,
			Field:     fn,
			Relation:  rn,
			Type:      tn,
		}
	}

	return getRenderPath(renderTask, renderContext, true)
}
