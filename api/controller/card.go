package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func newHTTPCardHandler(eg *echo.Group, handler *HTTPHandler) {
	// Prefix : /api/v1/card
	eg.GET("/:cardid", handler.GetCardByID)
}

// GetCardByID ...
func (h *HTTPHandler) GetCardByID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	uid, err := strconv.ParseUint(c.Param("uid"), 10, 64)
	if err != nil {
		mlog.Errorw("GetCardByID", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}

	card, err := h.cardService.GetCardByID(ctx, uid)
	if err != nil {
		mlog.Errorw("GetCardByID", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "GetCardByID OK", card)
}
