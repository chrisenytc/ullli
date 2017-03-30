package router

import (
	"context"
	"time"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/gorillamux"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"gopkg.in/kataras/iris.v6/middleware/recover"

	"github.com/chrisenytc/ullli/config"
	"github.com/chrisenytc/ullli/middlewares"
)

func Load() {
	log.Info("Loading server configs.")

	// Create proxy
	app := iris.New()

	log.Info("Loading HTTP request multiplexer.")

	// Enable devlogs
	if config.IsDevelopment() {
		app.Adapt(iris.DevLogger())
	}

	// Enable http handler
	app.Adapt(gorillamux.New())

	// Enable recovery
	app.Use(recover.New())

	// Configure views
	tmpl := view.HTML("./views", ".html")
	tmpl.Layout("layout.html")

	// Enable view engine
	app.Adapt(tmpl)

	app.StaticWeb("/assets", "./assets")
	app.StaticWeb("/images", "./images")

	log.Info("Loading request logger.")

	// Logger
	app.UseFunc(middlewares.LoggerMiddleware)

	log.Info("Loading routes.")

	log.Info("Loading error handlers.")

	// Load not found handler
	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		middlewares.LoggerMiddleware(ctx)

		err := ctx.RenderWithStatus(iris.StatusNotFound, "404.html", nil)

		if err != nil {
			log.Panicf("Render error on not found: %s", err)
		}
	})

	// Load bad request handler
	app.OnError(iris.StatusBadRequest, func(ctx *iris.Context) {
		middlewares.LoggerMiddleware(ctx)

		err := ctx.RenderWithStatus(iris.StatusBadRequest, "400.html", nil)

		if err != nil {
			log.Panicf("Render error on bad request: %s", err)
		}
	})

	// Load error handler
	app.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		middlewares.LoggerMiddleware(ctx)

		err := ctx.RenderWithStatus(iris.StatusInternalServerError, "500.html", nil)

		if err != nil {
			log.Panicf("Render error on internal server error: %s", err)
		}
	})

	// Define routes
	LoadRoutes(app)

	// Enable graceful shutdown
	app.Adapt(iris.EventPolicy{
		Interrupted: func(*iris.Framework) {
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			app.Shutdown(ctx)
		},
	})

	log.Info("Loading server.")

	log.Infof("Running on environment: %s.", config.Get().Environment)

	log.Infof("Listening on port %s.", config.Get().Port)

	// Start proxy
	app.Listen(":" + config.Get().Port)
}
