package micro

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"github.com/wizk3y/go-micro/logger"
	"github.com/wizk3y/go-micro/middleware"
	"github.com/wizk3y/go-micro/transhttp"
	"github.com/wizk3y/go-micro/util"
)

// NewAPIApp returns new app
func NewAPIApp(name string) (*App, *http.Server) {
	app := App{
		Name: name,
		Type: AppTypeAPI,
	}
	// init default flags, then parse
	app.InitDefaultFlags()
	app.ParseFlags()
	// init for micro app
	app.InitCommon()

	logger.Infof("%v. Starting...", app.Info())

	addr := fmt.Sprintf("%s:%d", app.HTTPHost, app.HTTPPort)
	hs := http.Server{
		Addr: addr,
		// Note: when using ABL on AWS https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#connection-idle-timeout
		// IdleTimeout must be greater than ALB Idle Timeout
		IdleTimeout:  70 * time.Second,
		ReadTimeout:  40 * time.Second,
		WriteTimeout: 70 * time.Second,
	}

	// handle sigterm
	util.HandleSigterm(func() {
		logger.Infof("%s. Stopping...", app.Info())
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		err := hs.Shutdown(ctx)
		if err != nil {
			logger.Errorw("An error occured when http server is being shutdown", "err", err)
		}
		for _, cb := range app.OnCloseCallback {
			cb()
		}
	})

	return &app, &hs
}

func (app *App) AddServer(s transhttp.WebServer) {
	app.AddRoutes(s.InitRoutes(), s.InitMiddlewares()...)
	app.OnClose(s.OnClose)
}

// AddRoutes -- add HTTP routes
func (app *App) AddRoutes(routes transhttp.Routes, mrs ...negroni.Handler) {
	app.Routes = append(app.Routes, routes)
	app.Middlewares = append(app.Middlewares, mrs)
}

// RunAPI -- run API server
func (app *App) RunAPI(hs *http.Server, mids ...negroni.Handler) (err error) {
	// init router + middlewares
	mr := transhttp.NewRouter(app.Routes, app.Middlewares)

	n := app.InitGlobalAPIMiddlewares(mr, mids...)
	hs.Handler = n

	// ListenAndServe
	hs.SetKeepAlivesEnabled(true)
	logger.Infof("Starting HTTP service %v at %v", app.Name, hs.Addr)

	err = hs.ListenAndServe()
	return
}

// InitGlobalAPIMiddlewares --
func (app *App) InitGlobalAPIMiddlewares(handler http.Handler, mids ...negroni.Handler) *negroni.Negroni {
	// using negroni mids for all requests
	n := negroni.New()

	logDisabled := viper.GetBool("api.disable_trace_log")

	n.Use(middleware.NewCorsMiddleware())
	n.Use(middleware.NewPreFlightMiddleware())
	if logDisabled {
		logger.Info("Disabled request & response log middleware")
	} else {
		// use logger, print each incoming request and response
		n.Use(middleware.NewHttpLogMiddleware())
	}

	// use recovery, catches panics and responds with a 500 response code
	n.Use(middleware.NewRecoveryMiddleware(app.IsDevelopment))

	// more mids
	if len(mids) > 0 {
		for _, mid := range mids {
			n.Use(mid)
		}
	}

	n.UseHandler(handler)

	return n
}
