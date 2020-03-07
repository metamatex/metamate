package gen

import (
	"errors"
	"fmt"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/utils"
	"github.com/spf13/afero"
	"os"
	"strings"
	"sync"

	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"

	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
)

func Gen(report *types.MessageReport, fs afero.Fs, version string, rn *graph.RootNode, tasks []types.RenderTask) (errs []error) {
	wg := sync.WaitGroup{}

	err := resetDirectories(fs, tasks)
	if err != nil {
		errs = append(errs, err)

		return
	}

	for _, task := range tasks {
		wg.Add(1)

		go func(task types.RenderTask) {
			defer wg.Done()

			err := func() (err error) {
				err = validateTaskBase(fs, task)
				if err != nil {
					return
				}

				if task.Iterate != nil {
					errs0 := Iterate(report, fs, task, version, rn)
					errs = append(errs, errs0...)
				} else {
					err = IterateFalse(report, fs, task, version, rn)
					if err != nil {
						errs = append(errs, err)
					}
				}

				return
			}()
			if err != nil {
				errs = append(errs, err)
			}
		}(task)
	}

	wg.Wait()

	return
}

func validateTaskIterate(task types.RenderTask) (err error) {
	if task.Filter == nil {
		return
	}

	kind, err := getIterateKind(*task.Iterate)
	if err != nil {
		return
	}

	switch kind {
	case graph.BASIC_TYPE:
		if task.Filter.BasicTypes == nil {
			err = errors.New("only filter.basicTypes can and must be set")
		}

		break
	case graph.ENDPOINT:
		if task.Filter.Endpoints == nil {
			err = errors.New("only filter.endpoints can and must be set")
		}

		break
	case graph.ENUM:
		if task.Filter.Enums == nil {
			err = errors.New("only filter.enums can and must be set")
		}

		break
	case graph.FIELD:
		if task.Filter.Fields == nil {
			err = errors.New("only filter.fields can and must be set")
		}

		break
	case graph.RELATION:
		if task.Filter.Relations == nil {
			err = errors.New("only filter.relations can and must be set")
		}

		break
	case graph.TYPE:
		if task.Filter.Types == nil {
			err = errors.New("only filter.types can and must be set")
		}

		break
	}

	return
}

func getIterateKind(s string) (kind string, err error) {
	switch strings.ToLower(s) {
	case graph.BASIC_TYPE, utils.Plural(graph.BASIC_TYPE):
		kind = graph.BASIC_TYPE

		break
	case graph.ENDPOINT, utils.Plural(graph.ENDPOINT):
		kind = graph.ENDPOINT

		break
	case graph.ENUM, utils.Plural(graph.ENUM):
		kind = graph.ENUM

		break
	case graph.FIELD, utils.Plural(graph.FIELD):
		kind = graph.FIELD

		break
	case graph.RELATION, utils.Plural(graph.RELATION):
		kind = graph.RELATION

		break
	case graph.TYPE, utils.Plural(graph.TYPE):
		kind = graph.TYPE

		break
	default:
		err = errors.New(fmt.Sprintf("unknown kind %v", s))

		break
	}

	return
}

func validateTaskIterateFalse(t types.RenderTask) (err error) {
	return
}

func validateTaskBase(fs afero.Fs, t types.RenderTask) (err error) {
	if t.Template == nil &&
		t.TemplateData == nil {
		err = errors.New("invalid task, template and templateData is missing")

		return
	}

	if t.Template != nil &&
		t.TemplateData != nil {
		err = errors.New("invalid task, template and templateData set")

		return
	}

	if t.Template != nil {
		_, err = fs.Stat(*t.Template)
		if err != nil {
			err = errors.New(fmt.Sprintf("invalid task, template file %v not found", *t.Template))

			return
		}
	}

	if t.Out == nil {
		err = errors.New("invalid task, out is missing")

		return
	}

	if t.Filter != nil && t.Dependencies != nil {
		err = errors.New("invalid task, filter and dependencies set")
	}

	if t.Dependencies != nil && (t.Dependencies.Types != nil && t.Dependencies.Endpoints != nil) {
		err = errors.New("invalid task, dependencies.types and dependencies.endpoints set")
	}

	return
}

func IterateFalse(report *types.MessageReport, fs afero.Fs, task types.RenderTask, version string, rn *graph.RootNode) (err error) {
	err = validateTaskIterateFalse(task)
	if err != nil {
		return
	}

	renderContext := getRenderContext(rn, version, task)

	templateData, err := getTemplateData(fs, task)
	if err != nil {
		return
	}

	err = render(report, fs, task, renderContext, templateData)
	if err != nil {
		return
	}

	return
}

func Iterate(report *types.MessageReport, fs afero.Fs, task types.RenderTask, version string, rn *graph.RootNode) (errs []error) {
	err := validateTaskIterate(task)
	if err != nil {
		errs = append(errs, err)

		return
	}

	wg := sync.WaitGroup{}

	renderContexts := getIterateRenderContexts(version, rn, task)

	for _, renderContext := range renderContexts {
		go func(renderContext types.IterateRenderContext) {
			defer wg.Done()

			err := func(renderContext types.IterateRenderContext) (err error) {
				manualPath, err := getRenderPath(task, renderContext, false)
				if err != nil {
					return
				}

				_, err = fs.Stat(manualPath)
				if err == nil || !os.IsNotExist(err) {
					return
				}
				if err != nil && os.IsNotExist(err) {
					var templateData string
					templateData, err = getTemplateData(fs, task)
					if err != nil {
						return
					}

					err = render(report, fs, task, renderContext, templateData)
					if err != nil {
						return
					}
				}

				return
			}(renderContext)
			if err != nil {
				errs = append(errs, err)
			}

		}(renderContext)

		wg.Add(1)
	}

	wg.Wait()

	return
}

func getIterateRenderContexts(version string, rn *graph.RootNode, task types.RenderTask) (renderContexts []types.IterateRenderContext) {
	if task.Filter != nil {
		if task.Filter.BasicTypes != nil {
			nm := rn.BasicTypes.Filter(*task.Filter.BasicTypes)

			for _, n := range nm {
				renderContexts = append(renderContexts, types.IterateRenderContext{
					Version:   &version,
					Data:      task.Data,
					BasicType: n,
				})
			}

			return
		}

		if task.Filter.Endpoints != nil {
			nm := rn.Endpoints.Filter(*task.Filter.Endpoints)

			for _, n := range nm {
				renderContexts = append(renderContexts, types.IterateRenderContext{
					Version:  &version,
					Data:     task.Data,
					Endpoint: n,
				})
			}

			return
		}

		if task.Filter.Enums != nil {
			nm := rn.Enums.Filter(*task.Filter.Enums)

			for _, n := range nm {
				renderContexts = append(renderContexts, types.IterateRenderContext{
					Version: &version,
					Data:    task.Data,
					Enum:    n,
				})
			}

			return
		}

		if task.Filter.Fields != nil {
			nm := rn.Fields.Filter(*task.Filter.Fields)

			for _, n := range nm {
				renderContexts = append(renderContexts, types.IterateRenderContext{
					Version: &version,
					Data:    task.Data,
					Field:   n,
				})
			}

			return
		}

		if task.Filter.Relations != nil {
			nm := rn.Relations.Filter(*task.Filter.Relations)

			for _, n := range nm {
				renderContexts = append(renderContexts, types.IterateRenderContext{
					Version:  &version,
					Data:     task.Data,
					Relation: n,
				})
			}

			return
		}

		if task.Filter.Types != nil {
			nm := rn.Types.Filter(*task.Filter.Types)

			for _, n := range nm {
				renderContexts = append(renderContexts, types.IterateRenderContext{
					Version: &version,
					Data:    task.Data,
					Type:    n,
				})
			}

			return
		}

		if task.Filter.Types != nil {
			nm := rn.Types.Filter(*task.Filter.Types)

			for _, n := range nm {
				renderContexts = append(renderContexts, types.IterateRenderContext{
					Version: &version,
					Data:    task.Data,
					Type:    n,
				})
			}

			return
		}
	}

	iterateKind, err := getIterateKind(*task.Iterate)
	if err != nil {
		return
	}

	if task.Dependencies != nil {
		switch iterateKind {
		case graph.ENUM:
			if task.Dependencies.Endpoints != nil {
				graph.GetEnumDependenciesFromEndpointIds(rn, rn.Endpoints.Filter(*task.Dependencies.Endpoints)).Each(func(n *graph.EnumNode) {
					renderContexts = append(renderContexts, types.IterateRenderContext{
						Version: &version,
						Data:    task.Data,
						Enum:    n,
					})
				})
			}

			if task.Dependencies.Types != nil {
				graph.GetEnumDependenciesFromTypeIds(rn, rn.Types.Filter(*task.Dependencies.Types)).Each(func(n *graph.EnumNode) {
					renderContexts = append(renderContexts, types.IterateRenderContext{
						Version: &version,
						Data:    task.Data,
						Enum:    n,
					})
				})
			}

			break
		case graph.TYPE:
			if task.Dependencies.Endpoints != nil {
				graph.GetTypeDependenciesFromEndpointIds(rn, rn.Endpoints.Filter(*task.Dependencies.Endpoints)).Each(func(n *graph.TypeNode) {
					renderContexts = append(renderContexts, types.IterateRenderContext{
						Version: &version,
						Data:    task.Data,
						Type:    n,
					})
				})
			}

			if task.Dependencies.Types != nil {
				graph.GetTypeDependenciesFromTypeIds(rn, rn.Types.Filter(*task.Dependencies.Types)).Each(func(n *graph.TypeNode) {
					renderContexts = append(renderContexts, types.IterateRenderContext{
						Version: &version,
						Data:    task.Data,
						Type:    n,
					})
				})
			}

			break
		case graph.ENDPOINT:
			rn.Endpoints.Filter(*task.Dependencies.Endpoints).Each(func(n *graph.EndpointNode) {
				renderContexts = append(renderContexts, types.IterateRenderContext{
					Version:  &version,
					Data:     task.Data,
					Endpoint: n,
				})
			})

			break
		}
	}

	return
}

func getRenderContext(rn *graph.RootNode, version string, task types.RenderTask) (renderCtx types.RenderContext) {
	renderCtx.Version = &version
	renderCtx.Data = task.Data

	if task.Filter != nil {
		if task.Filter.BasicTypes != nil {
			renderCtx.BasicTypes = rn.BasicTypes.Filter(*task.Filter.BasicTypes)
		}

		if task.Filter.Endpoints != nil {
			renderCtx.Endpoints = rn.Endpoints.Filter(*task.Filter.Endpoints)
		}

		if task.Filter.Enums != nil {
			renderCtx.Enums = rn.Enums.Filter(*task.Filter.Enums)
		}

		if task.Filter.Fields != nil {
			renderCtx.Fields = rn.Fields.Filter(*task.Filter.Fields)
		}

		if task.Filter.Relations != nil {
			renderCtx.Relations = rn.Relations.Filter(*task.Filter.Relations)
		}

		if task.Filter.Types != nil {
			renderCtx.Types = rn.Types.Filter(*task.Filter.Types)
		}
	}

	if task.Dependencies != nil {
		if task.Dependencies.Endpoints != nil {
			renderCtx.Endpoints = rn.Endpoints.Filter(*task.Dependencies.Endpoints)
			renderCtx.Types = graph.GetTypeDependenciesFromEndpointIds(rn, renderCtx.Endpoints)
			//renderCtx.Types = renderCtx.Types.AddTypeNodeMap(getReqRspTypesFromEndpoints(renderCtx.Endpoints))
			renderCtx.Enums = graph.GetEnumDependenciesFromEndpointIds(rn, renderCtx.Endpoints)
		}

		if task.Dependencies.Types != nil {
			renderCtx.Types = graph.GetTypeDependenciesFromTypeIds(rn, rn.Types.Filter(*task.Dependencies.Types))
			renderCtx.Enums = graph.GetEnumDependenciesFromTypeIds(rn, rn.Types.Filter(*task.Dependencies.Types))
		}
	}

	return
}

func getReqRspTypesFromEndpoints(enm graph.EndpointNodeMap) (tnm graph.TypeNodeMap) {
	tnm = graph.TypeNodeMap{}

	for _, en := range enm {
		tnm.Add(en.Edges.Type.Request(), en.Edges.Type.Response())
	}

	return
}

func getTemplateData(fs afero.Fs, task types.RenderTask) (s string, err error) {
	if task.TemplateData != nil {
		s = *task.TemplateData

		return
	}

	b, err := afero.ReadFile(fs, *task.Template)
	if err != nil {
		return
	}

	s = string(b)

	return
}
