package explorer

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type TemplateData struct {
	GraphqlPath  string
	ExplorerPath string
	DefaultQuery string
}

func GetStaticHandler(explorerPath string) http.Handler {
	fs := FS(false)

	return http.StripPrefix(explorerPath, http.FileServer(fs))
}

func MustGetIndexHandlerFunc(graphqlPath, explorerPath, explorerDefaultQuery string) (h http.HandlerFunc) {
	err := func() (err error) {
		fs := FS(false)

		f, err := fs.Open("/index.html")
		if err != nil {
			return
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			return
		}

		tmpl := template.Must(template.New("index.html").Parse(string(b)))
		data := TemplateData{
			GraphqlPath:  fmt.Sprintf(graphqlPath),
			ExplorerPath: strings.TrimSuffix(explorerPath, "/"),
			DefaultQuery: explorerDefaultQuery,
		}

		h = func(w http.ResponseWriter, r *http.Request) {
			err := tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}
		}

		return
	}()
	if err != nil {
		panic(err)
	}

	return
}
