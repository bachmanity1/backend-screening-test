package controller

import (
	"pandita/conf"
	"pandita/service"

	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	pandita       *conf.ViperConfig
	cardService   service.CardService
	columnService service.ColumnService
}

func newHTTPHandler(eg *echo.Group,
	pandita *conf.ViperConfig,
	cardService service.CardService,
	columnService service.ColumnService) {
	handler := &HTTPHandler{pandita, cardService, columnService}

	columnGroup := eg.Group("/column")
	newHTTPColumnHandler(columnGroup, handler)
	cardGroup := columnGroup.Group("/:columnid/card")
	newHTTPCardHandler(cardGroup, handler)
}
