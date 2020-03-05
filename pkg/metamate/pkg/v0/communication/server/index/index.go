//go:generate esc -pkg index -o static.go static
package index

import (
	"fmt"
	"github.com/metamatex/metamatemono/pkg/metamate/pkg/v0/types"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.New("index.html").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>MetaMate - Index</title>
	<link rel="icon" href="/static/favicon.ico" />
</head>
<body>
<ul>
{{ range $i, $url := . }}
	<li>
		<a href="{{ $url.Href }}">{{ $url.Label }}</a>
	</li>
{{ end }}
</ul>
</body>
</html>
`))

type link struct {
	Href  string
	Label string
}

func GetIndexHandlerFunc(port int, rs []types.Route) (http.HandlerFunc) {
	return func(writer http.ResponseWriter, req *http.Request) {
		var ls []link

		switch port {
		case 80:
			for _, r := range rs {
				ls = append(ls, link{Label: r.Path, Href: fmt.Sprintf("http://%v%v", req.Host, r.Path)})
			}
		default:
			for _, r := range rs {
				ls = append(ls, link{Label: r.Path, Href: fmt.Sprintf("http://%v:%v%v", req.Host, port, r.Path)})
			}
		}

		err := tpl.Execute(writer, ls)
		if err != nil {
			panic(err)
		}
	}
}

func GetStaticHandler() (http.Handler) {
	fs := FS(false)

	return http.FileServer(fs)
}
