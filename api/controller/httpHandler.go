package controller

import (
	"terra/conf"
	"terra/service"

	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	terra         *conf.ViperConfig
	cardService   service.CardService
	columnService service.ColumnService
}

func newHTTPHandler(eg *echo.Group,
	terra *conf.ViperConfig,
	cardService service.CardService,
	columnService service.ColumnService) {
	handler := &HTTPHandler{terra, cardService, columnService}

	columnGroup := eg.Group("/column")
	newHTTPColumnHandler(columnGroup, handler)
	cardGroup := columnGroup.Group("/:columnid/card")
	newHTTPCardHandler(cardGroup, handler)
}
