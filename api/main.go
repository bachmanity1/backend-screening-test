package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	ct "terra/api/controller"
	mw "terra/api/middleware"
	conf "terra/conf"
	"terra/model"
	repo "terra/repository"
	"terra/service"
	"terra/util"

	"github.com/juju/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var mlog *util.MLogger

func init() {
	// use all cpu
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	terra := conf.Terra
	mlog, _ = util.InitLog("main", terra.GetString("loglevel"))

	e := echoInit(terra)

	// Prepare Server
	env := terra.GetString("loglevel")
	ct.InitControler(env)
	model.InitModel(env)
	repo.InitRepository(env)
	service.InitService(env)
	db := repo.InitDB(terra)
	if err := ct.InitHandler(terra, e, db); err != nil {
		mlog.Errorw("InitHandler", "err", errors.Details(err))
		os.Exit(1)
	}

	// Start Server
	apiServer := fmt.Sprintf("0.0.0.0:%d", terra.GetInt("port"))
	mlog.Infow("Starting server", "listen", apiServer)
	if err := e.Start(apiServer); err != nil {
		mlog.Errorw("End server", "err", err)
	}
}

func echoInit(terra *conf.ViperConfig) (e *echo.Echo) {
	// Echo instance
	e = echo.New()

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.GET, echo.PUT, echo.DELETE},
	}))
	// Ping Check
	e.GET("/healthCheck", func(c echo.Context) error { return c.String(http.StatusAlreadyReported, "terra API Alive!\n") })
	e.POST("/healthCheck", func(c echo.Context) error { return c.String(http.StatusAlreadyReported, "terra API Alive!\n") })

	e.Use(mw.ZapLogger(mlog))
	e.HideBanner = true

	return e
}
