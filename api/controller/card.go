package controller

import (
	"net/http"
	"strconv"
	"terra/model"

	"github.com/labstack/echo/v4"
)

func newHTTPCardHandler(eg *echo.Group, handler *HTTPHandler) {
	// Prefix : /api/v1/card
	eg.POST("", handler.NewCard)
	eg.GET("/:id", handler.GetCardByID)
	eg.PUT("/:id", handler.UpdateCard)
	eg.DELETE("/:id", handler.DeleteCard)
	eg.PUT("/:id/after/:prev", handler.PutAfterCard)
}

// NewCard ...
func (h *HTTPHandler) NewCard(c echo.Context) (err error) {
	ctx := c.Request().Context()

	columnID, err := strconv.ParseUint(c.Param("columnid"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("NewCard", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}
	card := &model.Card{}
	if err := c.Bind(card); err != nil {
		mlog.With(ctx).Errorw("NewCard", "error", err)
		return response(c, http.StatusBadRequest, err.Error())
	}
	card.ColumnID = columnID

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

	columnID, err := strconv.ParseUint(c.Param("columnid"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("GetCardByID", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}
	cardID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("GetCardByID", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}

	card, err := h.cardService.GetCardByID(ctx, columnID, cardID)
	if err != nil {
		mlog.With(ctx).Errorw("GetCardByID", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "GetCardByID OK", card)
}

// UpdateCard ...
func (h *HTTPHandler) UpdateCard(c echo.Context) (err error) {
	ctx := c.Request().Context()

	columnID, err := strconv.ParseUint(c.Param("columnid"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("UpdateCard", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}
	card := &model.Card{}
	if err := c.Bind(card); err != nil {
		mlog.With(ctx).Errorw("UpdateCard", "error", err)
		return response(c, http.StatusBadRequest, err.Error())
	}
	card.ColumnID = columnID

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

	columnID, err := strconv.ParseUint(c.Param("columnid"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("DeleteCard", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}
	cardID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("DeleteCard", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}

	if err := h.cardService.DeleteCard(ctx, columnID, cardID); err != nil {
		mlog.With(ctx).Errorw("DeleteCard", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "DeleteCard OK")
}

// PutAfterCard ...
func (h *HTTPHandler) PutAfterCard(c echo.Context) (err error) {
	ctx := c.Request().Context()

	columnID, err := strconv.ParseUint(c.Param("columnid"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("PutAfterCard", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}
	cardID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("PutAfterCard", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}
	prev, err := strconv.ParseUint(c.Param("prev"), 10, 64)
	if err != nil {
		mlog.With(ctx).Errorw("PutAfterColumn", "error", err)
		return response(c, http.StatusBadRequest, "Invalid Path Param")
	}
	column, err := h.cardService.PutAfter(ctx, columnID, cardID, prev)
	if err != nil {
		mlog.With(ctx).Errorw("PutAfterCard", "error", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "PutAfterCard OK", column)
}
