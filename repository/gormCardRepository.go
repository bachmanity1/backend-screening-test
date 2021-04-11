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
	return
}

// UpdateCard ...
func (g *gormCardRepository) UpdateCard(ctx context.Context, card *model.Card) (ccard *model.Card, err error) {
	return
}

// GetCardByID ...
func (g *gormCardRepository) GetCardByID(ctx context.Context, id uint64) (card *model.Card, err error) {
	scope := g.Conn.WithContext(ctx)
	scope = scope.Where("id = ?", id).Find(&card)
	if scope.RowsAffected == 0 {
		return nil, errors.NotFoundf("cardID [%d]", id)
	}
	return card, nil
}

// DeleteCard ...
func (g *gormCardRepository) DeleteCard(ctx context.Context, id uint64) error {
	scope := g.Conn.WithContext(ctx)
	if err := scope.Delete(&model.Card{}, id).Error; err != nil {
		return errors.Annotatef(err, "Internal Server Error")
	}
	return nil
}