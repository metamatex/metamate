//go:generate yarn build
//go:generate esc -pkg admin -o static.go -prefix build build
package admin

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type TemplateData struct {
	Path string
}

func GetStaticHandler(path string) (http.Handler) {
	fs := FS(false)

	return http.StripPrefix(path, http.FileServer(fs))
}

func MustGetIndexHandlerFunc(path string) (h http.HandlerFunc) {
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
			Path: fmt.Sprintf(path),
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
