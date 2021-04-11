package service

import (
	"context"
	"pandita/model"
	repo "pandita/repository"
	"time"
)

type columnUsecase struct {
	repo       repo.ColumnRepository
	ctxTimeout time.Duration
}

// NewColumnService ...
func NewColumnService(repo repo.ColumnRepository, timeout time.Duration) ColumnService {
	return &columnUsecase{
		repo:       repo,
		ctxTimeout: timeout,
	}
}

// NewColumn ...
func (c *columnUsecase) NewColumn(ctx context.Context, column *model.Column) (ccolumn *model.Column, err error) {
	return c.repo.NewColumn(ctx, column)
}

// UpdateColumn ...
func (c *columnUsecase) UpdateColumn(ctx context.Context, column *model.Column) (ccolumn *model.Column, err error) {
	return c.repo.UpdateColumn(ctx, column)
}

// GetColumnByID ...
func (c *columnUsecase) GetColumnByID(ctx context.Context, id uint64) (column *model.Column, err error) {
	return c.repo.GetColumnByID(ctx, id)
}

// DeleteColumn ...
func (c *columnUsecase) DeleteColumn(ctx context.Context, id uint64) (err error) {
	return c.repo.DeleteColumn(ctx, id)
}
