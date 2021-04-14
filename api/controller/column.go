package controller

import (
	"net/http"
	"strconv"
	"terra/model"
	"terra/util"

	"github.com/labstack/echo/v4"
)

func newHTTPColumnHandler(eg *echo.Group, handler *HTTPHandler) {
	// Prefix : /api/v1/column
	eg.POST("", handler.NewColumn)
	eg.GET("", handler.GetColumnList)
	eg.GET("/:id", handler.GetColumnByID)
	eg.PUT("/:id", handler.UpdateColumn)
	eg.PUT("/:id/after/:prev", handler.PutAfterColumn)
	eg.DELETE("/:id", handler.DeleteColumn)
}

// NewColumn ...
func (h *HTTPHandler) NewColumn(c echo.Context) (err error) {
	ctx := c.Request().Context()

	column := &model.Column{}
	if err := c.Bind(column); err != nil {
		mlog.With(ctx).Infow("NewColumn", "error", err)
		return response(c, http.StatusBadRequest, err.Error())
	}

	column, err = h.columnService.NewColumn(ctx, column)
	if err != nil {
		mlog.With(ctx).Errorw("NewColumn", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "NewColumn OK", column)
}

// GetColumnByID ...
func (h *HTTPHandler) GetColumnByID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("GetColumnByID", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}

	column, err := h.columnService.GetColumnByID(ctx, id)
	if err != nil {
		mlog.With(ctx).Errorw("GetColumnByID", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "GetColumnByID OK", column)
}

// UpdateColumn ...
func (h *HTTPHandler) UpdateColumn(c echo.Context) (err error) {
	ctx := c.Request().Context()

	column := &model.Column{}
	if err := c.Bind(column); err != nil {
		mlog.With(ctx).Infow("UpdateColumn", "error", err)
		return response(c, http.StatusBadRequest, err.Error())
	}

	column, err = h.columnService.UpdateColumn(ctx, column)
	if err != nil {
		mlog.With(ctx).Errorw("UpdateColumn", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "UpdateColumn OK", column)
}

// DeleteColumn ...
func (h *HTTPHandler) DeleteColumn(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("DeleteColumn", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}

	if err := h.columnService.DeleteColumn(ctx, id); err != nil {
		mlog.With(ctx).Errorw("DeleteColumn", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "DeleteColumn OK")
}

// GetColumnList ...
func (h *HTTPHandler) GetColumnList(c echo.Context) (err error) {
	ctx := c.Request().Context()

	columns, err := h.columnService.GetColumnList(ctx)
	if err != nil {
		mlog.With(ctx).Errorw("GetColumnList", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "GetColumnList OK", columns)
}

// PutAfterColumn ...
func (h *HTTPHandler) PutAfterColumn(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("PutAfterColumn", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}
	prev, ok := util.ParseRank(c.Param("prev"))
	if !ok {
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}
	column, err := h.columnService.PutAfter(ctx, id, prev)
	if err != nil {
		mlog.With(ctx).Errorw("PutAfterColumn", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "PutAfterColumn OK", column)
}
