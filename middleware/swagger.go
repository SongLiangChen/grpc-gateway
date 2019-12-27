package middleware

import (
	"net/http"
	"path"
	"strings"
)

func SwaggerServer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			http.NotFound(w, r)
			return
		}

		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("proto/", p)
		http.ServeFile(w, r, p)
	}
}
