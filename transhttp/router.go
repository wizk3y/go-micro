package transhttp

import (
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"github.com/wizk3y/go-micro/middleware"
)

const (
	DefaultTimeout = 10 * time.Second
)

// Route -- Defines a single route, e.g. a human readable name, HTTP method,
// pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	BasePath    string
	Pattern     string
	Handler     http.Handler
	Middlewares []negroni.Handler
	Timeout     time.Duration
}

// Routes -- Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// NewRouter -- load all routers
func NewRouter(routes []Routes, middlewares [][]negroni.Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for i, r := range routes {
		sr := router.NewRoute().Subrouter()

		mw := middlewares[i]

		addSubRouter(sr, r, mw)
	}

	return router
}

func addSubRouter(router *mux.Router, routes Routes, middlewares []negroni.Handler) {
	for _, r := range routes {
		handlers := make([]negroni.Handler, 0)

		timeout := r.Timeout
		if timeout >= 0 {
			if timeout == 0 {
				timeout = time.Duration(viper.GetInt64("api.timeout")) * time.Millisecond
			}
			if timeout == 0 {
				timeout = DefaultTimeout
			}

			handlers = append(handlers, middleware.NewTimeoutMiddleware(timeout))
		}

		// subrouter middleware
		for _, m := range middlewares {
			handlers = append(handlers, m)
		}

		// handler middleware
		for _, mw := range r.Middlewares {
			handlers = append(handlers, mw)
		}

		handlers = append(handlers, negroni.Wrap(r.Handler))

		router.
			Methods(r.Method).
			Path(path.Join(r.BasePath, r.Pattern)).
			Handler(negroni.New(handlers...))
	}
}
