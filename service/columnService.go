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
	if ccolumn, err = c.repo.GetColumnByID(ctx, column.ID); err != nil {
		return nil, err
	}
	ccolumn.Update(column)
	if err = c.repo.UpdateColumn(ctx, ccolumn); err != nil {
		return nil, err
	}
	return c.repo.GetColumnByID(ctx, ccolumn.ID)
}

// GetColumnByID ...
func (c *columnUsecase) GetColumnByID(ctx context.Context, id uint64) (column *model.Column, err error) {
	return c.repo.GetColumnByID(ctx, id)
}

// GetColumnList ...
func (c *columnUsecase) GetColumnList(ctx context.Context) (columns model.ColumnList, err error) {
	return c.repo.GetColumnList(ctx)
}

// DeleteColumn ...
func (c *columnUsecase) DeleteColumn(ctx context.Context, id uint64) (err error) {
	return c.repo.DeleteColumn(ctx, id)
}

// PutAfter ...
func (c *columnUsecase) PutAfter(ctx context.Context, id uint64, prev string) (column *model.Column, err error) {
	if column, err = c.repo.GetColumnByID(ctx, id); err != nil {
		return nil, err
	}
	next, err := c.repo.GetNextOrder(ctx, prev)
	if err != nil {
		return nil, err
	}
	column.UpdateOrder(prev, next)
	if err = c.repo.UpdateColumn(ctx, column); err != nil {
		return nil, err
	}
	return c.repo.GetColumnByID(ctx, id)
}
