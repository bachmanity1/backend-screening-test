package controller

import (
	"net/http"
	"pandita/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func newHTTPCardHandler(eg *echo.Group, handler *HTTPHandler) {
	// Prefix : /api/v1/card
	eg.POST("", handler.NewCard)
	eg.GET("/:id", handler.GetCardByID)
	eg.PUT("/:id", handler.UpdateCard)
	eg.DELETE("/:id", handler.DeleteCard)
}

// NewCard ...
func (h *HTTPHandler) NewCard(c echo.Context) (err error) {
	ctx := c.Request().Context()

	card := &model.Card{}
	if err := c.Bind(card); err != nil {
		mlog.With(ctx).Infow("NewCard", "error", err)
		return response(c, http.StatusBadRequest, err.Error())
	}

	card, err = h.cardService.NewCard(ctx, card)
	if err != nil {
		mlog.With(ctx).Errorw("NewCard", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "NewCard OK", card)
}

// GetCardByID ...
func (h *HTTPHandler) GetCardByID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("GetCardByID", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}

	card, err := h.cardService.GetCardByID(ctx, id)
	if err != nil {
		mlog.With(ctx).Errorw("GetCardByID", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "GetCardByID OK", card)
}

// UpdateCard ...
func (h *HTTPHandler) UpdateCard(c echo.Context) (err error) {
	ctx := c.Request().Context()

	card := &model.Card{}
	if err := c.Bind(card); err != nil {
		mlog.With(ctx).Infow("UpdateCard", "error", err)
		return response(c, http.StatusBadRequest, err.Error())
	}

	card, err = h.cardService.UpdateCard(ctx, card)
	if err != nil {
		mlog.With(ctx).Errorw("UpdateCard", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "UpdateCard OK", card)
}

// DeleteCard ...
func (h *HTTPHandler) DeleteCard(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("DeleteCard", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}

	if err := h.cardService.DeleteCard(ctx, id); err != nil {
		mlog.With(ctx).Errorw("DeleteCard", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "DeleteCard OK")
}
