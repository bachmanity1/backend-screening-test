package repository

import (
	"context"
	"terra/model"

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

	lastColumn := &model.Column{}
	if err = scope.Order("columns.order desc").Find(&lastColumn).Error; err != nil {
		mlog.With(ctx).Errorw("gormColumn NewColumn", "error", err)
		return nil, err
	}
	column.Order = lastColumn.Order + 1

	if err = scope.Create(&column).Error; err != nil {
		mlog.With(ctx).Errorw("gormColumn NewColumn", "error", err)
		return nil, err
	}
	column.Cards = model.CardList{}
	return column, nil
}

// UpdateColumn ...
func (g *gormColumnRepository) UpdateColumn(ctx context.Context, column *model.Column) (err error) {
	scope := g.Conn.WithContext(ctx)
	if err = scope.Updates(column).Error; err != nil {
		mlog.With(ctx).Errorw("gormColumn NewColumn", "error", err)
		return err
	}
	return nil
}

// GetColumnByID ...
func (g *gormColumnRepository) GetColumnByID(ctx context.Context, id uint64) (column *model.Column, err error) {
	scope := g.Conn.WithContext(ctx)
	column = &model.Column{}
	scope = scope.Preload("Cards", func(db *gorm.DB) *gorm.DB {
		return db.Order("cards.order").Where("cards.status != ?", model.Archived)
	}).Where("id = ?", id).Find(&column)
	if scope.Error != nil || scope.RowsAffected == 0 {
		return nil, errors.NotFoundf("columnID[%d]", id)
	}
	return column, nil
}

// GetColumnList ...
func (g *gormColumnRepository) GetColumnList(ctx context.Context) (columns model.ColumnList, err error) {
	scope := g.Conn.WithContext(ctx)
	columns = model.ColumnList{}
	scope = scope.Preload("Cards", func(db *gorm.DB) *gorm.DB {
		return db.Order("cards.order").Where("cards.status != ?", model.Archived)
	}).Order("columns.order").Find(&columns)
	return columns, scope.Error
}

// DeleteColumn ...
func (g *gormColumnRepository) DeleteColumn(ctx context.Context, id uint64) error {
	scope := g.Conn.WithContext(ctx)
	if err := scope.Delete(&model.Column{}, id).Error; err != nil {
		return errors.Annotatef(err, "Internal Server Error")
	}
	return nil
}

// UpdateColumnOrder ...
func (g *gormColumnRepository) UpdateColumnOrder(ctx context.Context, id, prev uint64) error {
	scope := g.Conn.WithContext(ctx)
	scope = scope.Begin()
	if err := scope.Model(&model.Column{}).
		Where("columns.order > ?", prev).
		UpdateColumn("columns.order", gorm.Expr("columns.order + ?", 1)).Error; err != nil {
		scope.Rollback()
		return errors.Annotatef(err, "Internal Server Error")
	}
	if err := scope.Model(&model.Column{}).
		Where("id = ?", id).
		Update("columns.order", prev+1).Error; err != nil {
		scope.Rollback()
		return errors.Annotatef(err, "Internal Server Error")
	}
	scope.Commit()
	return nil
}

// // GetNextOrder ...
// func (g *gormColumnRepository) GetNextOrder(ctx context.Context, prev string) (order string, err error) {
// 	scope := g.Conn.WithContext(ctx)
// 	column := &model.Column{}
// 	if err = scope.Where("columns.order > ?", prev).
// 		Order("columns.order").Limit(1).Find(&column).Error; err != nil {
// 		mlog.With(ctx).Errorw("GetNextOrder", "error", err)
// 		return "", err
// 	}
// 	return column.Order, nil
// }
