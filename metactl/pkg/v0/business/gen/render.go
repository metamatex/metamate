package gen

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/spf13/afero"
)

var funcMap = template.FuncMap{
	"plural":         inflection.Plural,
	"lowerCamel":     strcase.ToLowerCamel,
	"camel":          strcase.ToCamel,
	"screamingSnake": strcase.ToScreamingSnake,
	"snake":          strcase.ToSnake,
}

func render(report *types.MessageReport, fs afero.Fs, renderTask types.RenderTask, renderContext interface{}, templateData string) (err error) {
	outPath, err := getRenderPath(renderTask, renderContext, true)
	if err != nil {
		return
	}

	outDir := filepath.Dir(outPath)

	_, err = fs.Stat(outDir)
	if err != nil && os.IsNotExist(err) {
		err = fs.MkdirAll(outDir, 0755)
		if err != nil {
			return
		}
	}
	if err != nil && !os.IsNotExist(err) {
		return
	}

	f, err := fs.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return
	}

	tmpl := template.Must(template.New("").Funcs(template.FuncMap(sprig.FuncMap())).Funcs(funcMap).Parse(templateData))

	err = tmpl.Execute(f, renderContext)
	if err != nil {
		err = errors.New(fmt.Sprintf("error while rendering: %v", err.Error()))

		fs.Remove(outPath)

		return
	}

	report.AddDebug(fmt.Sprintf("generated %v", outPath))

	return
}

func getRenderPath(renderTask types.RenderTask, renderContext interface{}, withGeneratedIndicator bool) (path string, err error) {
	tmpl := template.Must(template.New(*renderTask.Out).Funcs(template.FuncMap(sprig.FuncMap())).Funcs(funcMap).Parse(string(*renderTask.Out)))

	var b bytes.Buffer
	err = tmpl.Execute(&b, renderContext)
	if err != nil {
		return
	}

	path = b.String()

	if !strings.Contains(path, "_") && withGeneratedIndicator {
		err = errors.New(fmt.Sprintf("%s doesn't contain the generated indicator \"_\"", *renderTask.Out))
	}

	if strings.Contains(path, "_") && !withGeneratedIndicator {
		path = strings.Replace(path, "_", "", -1)
	}

	return
}
