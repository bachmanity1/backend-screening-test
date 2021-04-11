package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func newHTTPColumnHandler(eg *echo.Group, handler *HTTPHandler) {
	// Prefix : /api/v1/column
	eg.GET("/:columnid", handler.GetColumnByID)
}

// GetColumnByID ...
func (h *HTTPHandler) GetColumnByID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	uid, err := strconv.ParseUint(c.Param("uid"), 10, 64)
	if err != nil {
		mlog.Errorw("GetColumnByID", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}

	column, err := h.columnService.GetColumnByID(ctx, uid)
	if err != nil {
		mlog.Errorw("GetColumnByID", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "GetColumnByID OK", column)
}
