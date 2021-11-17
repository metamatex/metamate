package embedded

import (
	"bytes"
	"net/http"
)

type responseWriter struct {
	header     http.Header
	b          bytes.Buffer
	statusCode int
}

func newResponseWriter() *responseWriter {
	return &responseWriter{
		header: map[string][]string{},
	}
}

func (w *responseWriter) Header() http.Header {
	return w.header
}

func (w *responseWriter) Write(b []byte) (int, error) {
	return w.b.Write(b)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}
