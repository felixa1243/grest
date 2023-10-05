package libs

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type Router struct {
	routes []Routes
}
type Routes struct {
	Path    *regexp.Regexp
	Method  string
	Handler http.HandlerFunc
}

func (r *Router) Route(method, path string, handler http.HandlerFunc) {
	exactPath := regexp.MustCompile("^" + path + "$")
	r.routes = append(r.routes, Routes{
		Path:    exactPath,
		Method:  method,
		Handler: handler,
	})
}
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, e := range r.routes {
		if match := e.Match(req); !match {
			continue
		}
		handler := e.Handler
		handler.ServeHTTP(w, req)
		return
	}
	http.NotFound(w, req)
}
func (r *Routes) Match(req *http.Request) bool {
	if r.Method != req.Method {
		return false
	}
	match := r.Path.FindStringSubmatch(req.URL.Path)
	return match != nil
}
func (r *Router) Json(data interface{}, status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("status code", http.StatusText(status))
		res, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		w.Write(res)
	}
}
