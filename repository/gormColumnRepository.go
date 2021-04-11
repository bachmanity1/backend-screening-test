package repository

import (
	"context"
	"pandita/model"

	"github.com/juju/errors"
	"gorm.io/gorm"
)

type gormColumnRepository struct {
	Conn *gorm.DB
}

// NewGormColumnRepository ...
func NewGormColumnRepository(conn *gorm.DB) ColumnRepository {
	migrations := []interface{}{
		&model.Column{},
	}
	if err := conn.Migrator().AutoMigrate(migrations...); err != nil {
		mlog.Panicw("Unable to AutoMigrate ColumnRepository", "error", err)
	}
	return &gormColumnRepository{Conn: conn}
}

// NewColumn ...
func (g *gormColumnRepository) NewColumn(ctx context.Context, column *model.Column) (ccolumn *model.Column, err error) {
	mlog.With(ctx).Debugw("gormColumn NewColumn", "column", column)
	scope := g.Conn.WithContext(ctx)
	if err = scope.Create(&column).Error; err != nil {
		mlog.With(ctx).Errorw("gormColumn NewColumn", "error", err)
		return nil, err
	}
	return column, nil
}

// UpdateColumn ...
func (g *gormColumnRepository) UpdateColumn(ctx context.Context, column *model.Column) (ccolumn *model.Column, err error) {
	scope := g.Conn.WithContext(ctx)
	if err = scope.Updates(column).Error; err != nil {
		mlog.With(ctx).Errorw("gormColumn NewColumn", "error", err)
		return nil, err
	}
	scope = scope.Where("id = ?", column.ID).Find(&column)
	if scope.Error != nil || scope.RowsAffected == 0 {
		return nil, errors.NotFoundf("columnID[%d]", column.ID)
	}
	return column, nil
}

// GetColumnByID ...
func (g *gormColumnRepository) GetColumnByID(ctx context.Context, id uint64) (column *model.Column, err error) {
	scope := g.Conn.WithContext(ctx)
	scope = scope.Where("id = ?", id).Find(&column)
	if scope.Error != nil || scope.RowsAffected == 0 {
		return nil, errors.NotFoundf("columnID[%d]", id)
	}
	return column, nil
}

// DeleteColumn ...
func (g *gormColumnRepository) DeleteColumn(ctx context.Context, id uint64) error {
	scope := g.Conn.WithContext(ctx)
	if err := scope.Delete(&model.Column{}, id).Error; err != nil {
		return errors.Annotatef(err, "Internal Server Error")
	}
	return nil
}
