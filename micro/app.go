package micro

import (
	"flag"
	"fmt"
	"log"

	"github.com/urfave/negroni"
	"github.com/wizk3y/go-micro/config"
	"github.com/wizk3y/go-micro/logger"
	"github.com/wizk3y/go-micro/transhttp"
)

type AppType string

type OnCloseCallback func()

const (
	AppTypeAPI AppType = "api"
)

type App struct {
	// Common
	Name          string
	Version       string
	Type          AppType
	IsDevelopment bool
	// Configuration
	ConfigLoaded bool
	ConfigFile   string
	// HTTP
	HTTPHost    string
	HTTPPort    int64
	Middlewares [][]negroni.Handler
	Routes      []transhttp.Routes
	// Timeout
	DefaultAPITimeout int64
	// Callback
	OnCloseCallback []OnCloseCallback
}

// InitDefaultFlags -- init Default flags
func (app *App) InitDefaultFlags() {
	flag.StringVar(&app.Version, "version", "1.0.0", "App version")
	flag.BoolVar(&app.IsDevelopment, "dev", false, "Development mode")

	// flags for config
	flag.StringVar(&app.ConfigFile, "config-file", "", "Configuration file")

	// flags for http
	flag.StringVar(&app.HTTPHost, "http-host", "", "HTTP listen host")
	flag.Int64Var(&app.HTTPPort, "http-port", 8888, "HTTP listen port")

	// timeout
	flag.Int64Var(&app.DefaultAPITimeout, "api-timeout", 0, "Timeout of API in ms")
}

// ParseFlags -- parse flags
func (app *App) ParseFlags() {
	flag.Parse()
}

// Info -- show info of app
func (app *App) Info() string {
	return fmt.Sprintf("App type: %s. App name: %s. App version %s", app.Type, app.Name, app.Version)
}

// InitCommon -- initializes common things: logger, config, service discovery, tracing
func (app *App) InitCommon() {
	app.PreCheck()
	app.InitConfig()
	app.InitLogger()
}

// PreCheck -- pre check some conditional
func (app *App) PreCheck() {
	if app.Name == "" {
		log.Panic("must set micro app name")
	}
	if string(app.Type) == "" {
		log.Panic("must set micro app type")
	}
}

// InitConfig -- initializes config
func (app *App) InitConfig() {
	result := false

	if app.ConfigFile == "" {
		// read by default
		result = config.ReadConfig("config", "./conf", ".")
	} else {
		// read by input file
		result = config.ReadConfigByFile(app.ConfigFile)
	}

	if !result {
		logger.Panic("Could not load config")
	}

	app.ConfigLoaded = true
	logger.Info("Config loaded")
}

// InitLogger -- initializes logger
func (app *App) InitLogger() {
	if app.IsDevelopment {
		logger.InitLoggerDefaultDev()
	} else {
		logger.InitLoggerDefault()
	}
}

// OnClose -- register callback
func (app *App) OnClose(f OnCloseCallback) {
	app.OnCloseCallback = append(app.OnCloseCallback, f)
}
