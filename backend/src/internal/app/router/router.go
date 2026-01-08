package router

import "net/http"

type Route func(*http.ServeMux)

func New(routes ...Route) *http.ServeMux {
	mux := http.NewServeMux()

	for _, route := range routes {
		route(mux)
	}

	return mux
}

func Handle(pattern string, h http.Handler) Route {
	return func(mux *http.ServeMux) {
		mux.Handle(pattern, h)
	}
}

func HandleFunc(pattern string, h http.HandlerFunc) Route {
	return func(mux *http.ServeMux) {
		mux.HandleFunc(pattern, h)
	}
}
