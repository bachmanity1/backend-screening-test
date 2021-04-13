package repository

import (
	"context"
	"pandita/model"

	"github.com/juju/errors"
	"gorm.io/gorm"
)

type gormCardRepository struct {
	Conn *gorm.DB
}

// NewGormCardRepository ...
func NewGormCardRepository(conn *gorm.DB) CardRepository {
	migrations := []interface{}{
		&model.Card{},
	}
	if err := conn.Migrator().AutoMigrate(migrations...); err != nil {
		mlog.Panicw("Unable to AutoMigrate CardRepository", "error", err)
	}
	return &gormCardRepository{Conn: conn}
}

// NewCard ...
func (g *gormCardRepository) NewCard(ctx context.Context, card *model.Card) (ccard *model.Card, err error) {
	mlog.With(ctx).Debugw("gormCard NewCard", "card", card)
	scope := g.Conn.WithContext(ctx)

	lastCard := &model.Card{}
	if err = scope.Where("column_id = ?", card.ColumnID).Order("cards.order desc").Limit(1).Find(&lastCard).Error; err != nil {
		mlog.With(ctx).Errorw("gormCard NewCard", "error", err)
		return nil, err
	}
	card.Order = lastCard.Order + 1

	if err = scope.Create(&card).Error; err != nil {
		mlog.With(ctx).Errorw("gormCard NewCard", "error", err)
		return nil, err
	}
	return card, nil
}

// UpdateCard ...
func (g *gormCardRepository) UpdateCard(ctx context.Context, card *model.Card) (err error) {
	scope := g.Conn.WithContext(ctx)
	if err = scope.Updates(card).Error; err != nil {
		mlog.With(ctx).Errorw("gormCard NewCard", "error", err)
		return err
	}
	return nil
}

// GetCardByID ...
func (g *gormCardRepository) GetCardByID(ctx context.Context, columnID, cardID uint64) (card *model.Card, err error) {
	scope := g.Conn.WithContext(ctx)
	scope = scope.Where("column_id = ? AND id = ?", columnID, cardID).Find(&card)
	if scope.Error != nil || scope.RowsAffected == 0 {
		return nil, errors.NotFoundf("cardID[%d]")
	}
	return card, nil
}

// DeleteCard ...
func (g *gormCardRepository) DeleteCard(ctx context.Context, columnID, cardID uint64) error {
	scope := g.Conn.WithContext(ctx)
	if err := scope.Where("column_id = ? AND id = ?", columnID, cardID).Delete(&model.Card{}).Error; err != nil {
		return errors.Annotatef(err, "Internal Server Error")
	}
	return nil
}

// UpdateCardOrder ...
func (g *gormCardRepository) UpdateCardOrder(ctx context.Context, columnID, cardID, prev uint64) error {
	scope := g.Conn.WithContext(ctx)
	scope = scope.Begin()
	if err := scope.Model(&model.Card{}).
		Where("cards.order > ?", prev).
		UpdateColumn("cards.order", gorm.Expr("cards.order + ?", 1)).Error; err != nil {
		scope.Rollback()
		return errors.Annotatef(err, "Internal Server Error")
	}
	if err := scope.Model(&model.Card{}).
		Where("column_id = ? AND id = ?", columnID, cardID).
		Update("cards.order", prev+1).Error; err != nil {
		scope.Rollback()
		return errors.Annotatef(err, "Internal Server Error")
	}
	scope.Commit()
	return nil
}

// // GetNextCardOrder ...
// func (g *gormCardRepository) GetNextOrder(ctx context.Context, columnID uint64, prev string) (order string, err error) {
// 	scope := g.Conn.WithContext(ctx)
// 	card := &model.Card{}
// 	if err = scope.Where("cards.column_id = ?", columnID).
// 		Where("cards.order > ?", prev).
// 		Order("cards.order").Limit(1).Find(&card).Error; err != nil {
// 		mlog.With(ctx).Errorw("GetNextOrder", "error", err)
// 		return "", err
// 	}
// 	return card.Order, nil
// }
