package server

import (
	"fmt"
	"net/http"
	"strings"
)

// TODO test the headers are correct
func (s *Server) staticFileHandler(h http.Handler) http.Handler {
	// https://create-react-app.dev/docs/production-build
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		paths := strings.Split(r.URL.Path, "/")
		if paths[0] == "static" {
			// Cache-Control: max-age=31536000 for your build/static assets
			w.Header().Set("Cache-Control", "max-age=31536000")
		} else {
			//Cache-Control: no-cache for everything else
			w.Header().Set("Cache-Control", "no-cache")
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
