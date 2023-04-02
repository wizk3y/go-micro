package transhttp

import "github.com/urfave/negroni"

// WebServer --
type WebServer interface {
	InitRoutes() Routes
	InitMiddlewares() []negroni.Handler
	OnClose()
}
