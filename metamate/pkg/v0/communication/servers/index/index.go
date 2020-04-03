package index

import (
	"fmt"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.New("index.html").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>MetaMate - Index</title>
	<link rel="icon" href="/static/logo/blue/favicon.ico" />
</head>
<body>
<p>version: {{ .Version.Version }}</p>
<p>commit: {{ .Version.Commit }}</p>
<p>date: {{ .Version.Date }}</p>
<ul>
{{ range $i, $link := .Links }}
	<li>
		<a href="{{ $link.Href }}">{{ $link.Label }}</a>
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

type templateData struct {
	Links []link
	Version types.Version
}

func GetIndexHandlerFunc(rs []types.Route, v types.Version) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		var ls []link

		for _, r := range rs {
			ls = append(ls, link{Label: r.Path, Href: fmt.Sprintf("http://%v%v", req.Host, r.Path)})
		}

		err := tpl.Execute(writer, templateData{Version: v, Links:ls})
		if err != nil {
			panic(err)
		}
	}
}

func GetStaticHandler() http.Handler {
	fs := FS(false)

	return http.FileServer(fs)
}
