# go-micro

go-micro is a Go toolkit help you build an monolith application or a service in micro-services, also change from monolith to micro-service with less effort

## Install
```shell
go get github.com/wizk3y/go-micro
```

**Note:** go-micro uses [Go Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies.

## Usage
- Create file/directory with follow structure (as best practice, you can organize project by your own)
```
your-project/
├── go.mod
├── go.sum
├── cmd/
│   └── main.go
└── pkg/
    ├── app/
    │   └── webserver.go
    └── handler/
        ├── get-handler.go
        ├── post-handler.go
        └── put-handler.go
```
- Using `NewAPIApp` to create API app base, `AddServer` to add your app server, and `RunAPI` as final step to start your API server
```go
// cmd/main.go
package main

import (
	"net/http"

	"module-name/pkg/app"
	"github.com/wizk3y/go-micro/logger"
	"github.com/wizk3y/go-micro/micro"
)

func main() {
	microApp, hs := micro.NewAPIApp("your-api-server-name")

    appServer := app.NewServer()
	microApp.AddServer(appServer)

    err := microApp.RunAPI(hs)
	if err != nil {
		if err.Error() == http.ErrServerClosed.Error() {
			logger.Info(http.ErrServerClosed.Error())
		} else {
			logger.Errorf("HTTP server closed with error: %v", err)
		}
	}
}
```
- File contains an struct have `InitRoutes`, `InitMiddlewares`, `OnClose` as an implement of `transhttp.WebServer`
```go
// pkg/app/webserver.go
package app

import (
	"net/http"

	"module-name/pkg/handler"
	"github.com/urfave/negroni"
	"github.com/wizk3y/go-micro/transhttp"
)

type server struct {
	basePath string
}

func NewServer() transhttp.WebServer {
	s := server{
		basePath: "/v1/",
	}

	return &s
}

// InitRoutes -- Initialize our routes
func (s *server) InitRoutes() transhttp.Routes {
	return transhttp.Routes{
        transhttp.Route{
			Name:     "Get something handler",
			Method:   http.MethodGet,
			BasePath: s.basePath,
			Pattern:  "/path",
			Handler: &handler.GetHandler{},
		},
    }
}

// InitMiddlewares -- Initialize our middlewares
func (s *server) InitMiddlewares() []negroni.Handler {
	return []negroni.Handler{}
}

// OnClose --
func (s *server) OnClose() {

}
```
- File implement an endpoint handler with `ServerHTTP`, which is implement of `http.Handler`. Also using `RespondError` or `RespondJSON` for return response data to client, note that always `return` right after to prevent issue write response multiple times.
```go
// pkg/handler/get-handler.go
package handler

import (
	"net/http"

	"github.com/wizk3y/go-micro/transhttp"
)

type GetHandler struct {
}

// ServeHTTP --
func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // handling some logic
    if err != nil {
		transhttp.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

    transhttp.RespondJSON(w, http.StatusOK, map[string]interface{}{})
}
```

## Advance usage
- [Organize project as micro-services (single repository - best practice)](https://github.com/wizk3y/go-micro-doc/tree/master/structure/single-repository)
- [Organize project as micro-services (multiple repositories - best practice)](https://github.com/wizk3y/go-micro-doc/tree/master/structure/multiple-repositories)

## Features
- [Configuration](https://github.com/wizk3y/go-micro-doc/tree/master/feature/configuration)
- [HTTP Middleware](https://github.com/wizk3y/go-micro-doc/tree/master/feature/http-middleware)
- [PostgreSQL connection](https://github.com/wizk3y/go-micro-doc/tree/master/feature/postgresql-connection)
- [MySQL connection](https://github.com/wizk3y/go-micro-doc/tree/master/feature/mysql-connection)
- [Redis connection](https://github.com/wizk3y/go-micro-doc/tree/master/feature/redis-connection)


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)