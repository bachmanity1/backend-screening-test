package controller

import (
	"fmt"
	"net/http"
	mw "terra/api/middleware"
	"terra/conf"
	repo "terra/repository"
	"terra/service"
	"terra/util"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	// terraStatus for common response status
	terraStatus struct {
		TRID       string      `json:"trID" example:"20200213052007345858"`
		ResultCode string      `json:"resultCode" example:"0000"`
		ResultMsg  string      `json:"resultMsg" example:"Request OK"`
		ResultData interface{} `json:"resultData,omitempty"`
	}
)

var mlog *util.MLogger
var timeout = 10 * time.Second

func InitControler(env string) {
	mlog, _ = util.InitLog("controller", env)
}

// InitHandler ...
func InitHandler(terra *conf.ViperConfig, e *echo.Echo, db *gorm.DB) (err error) {
	api := e.Group("/api")
	ver := api.Group("/v1")
	ver.Use(mw.TransID())

	cardRepo := repo.NewGormCardRepository(db)
	columnRepo := repo.NewGormColumnRepository(db)
	cardService := service.NewCardService(cardRepo, columnRepo, timeout)
	columnService := service.NewColumnService(columnRepo, timeout)
	newHTTPHandler(ver, terra, cardService, columnService)
	return nil
}

func response(c echo.Context, code int, resMsg string, result ...interface{}) error {
	resCode := "0000"
	if code != http.StatusOK {
		resCode = fmt.Sprintf("1%d", code)
	}

	id, ok := c.Request().Context().Value(util.TransIDKey).(string)
	if !ok {
		id = util.NewID()
	}

	res := terraStatus{
		TRID:       id,
		ResultCode: resCode,
		ResultMsg:  resMsg,
	}

	if result != nil {
		res.ResultData = result[0]
	}
	return c.JSON(code, res)
}
