package service

import (
	"context"
	"terra/model"
	repo "terra/repository"
	"time"
)

type cardUsecase struct {
	repo       repo.CardRepository
	colRepo    repo.ColumnRepository
	ctxTimeout time.Duration
}

// NewCardService ...
func NewCardService(repo repo.CardRepository, colRepo repo.ColumnRepository, timeout time.Duration) CardService {
	return &cardUsecase{
		repo:       repo,
		colRepo:    colRepo,
		ctxTimeout: timeout,
	}
}

// NewCard ...
func (c *cardUsecase) NewCard(ctx context.Context, card *model.Card) (ccard *model.Card, err error) {
	if _, err = c.colRepo.GetColumnByID(ctx, card.ColumnID); err != nil {
		return nil, err
	}
	return c.repo.NewCard(ctx, card)
}

// UpdateCard ...
func (c *cardUsecase) UpdateCard(ctx context.Context, card *model.Card) (ccard *model.Card, err error) {
	if ccard, err = c.repo.GetCardByID(ctx, card.ColumnID, card.ID); err != nil {
		return nil, err
	}
	ccard.Update(card)
	if err = c.repo.UpdateCard(ctx, ccard); err != nil {
		return nil, err
	}
	return c.repo.GetCardByID(ctx, ccard.ColumnID, ccard.ID)
}

// GetCardByID ...
func (c *cardUsecase) GetCardByID(ctx context.Context, columnID, cardID uint64) (card *model.Card, err error) {
	return c.repo.GetCardByID(ctx, columnID, cardID)
}

// DeleteCard ...
func (c *cardUsecase) DeleteCard(ctx context.Context, columnID, cardID uint64) (err error) {
	return c.repo.DeleteCard(ctx, columnID, cardID)
}

// PutAfter ...
func (c *cardUsecase) PutAfter(ctx context.Context, columnID, cardID uint64, prev string) (card *model.Card, err error) {
	if card, err = c.repo.GetCardByID(ctx, columnID, cardID); err != nil {
		return nil, err
	}
	next, err := c.repo.GetNextOrder(ctx, columnID, prev)
	if err != nil {
		return nil, err
	}
	card.SetOrder(prev, next)
	if err = c.repo.UpdateCard(ctx, card); err != nil {
		return nil, err
	}
	return c.repo.GetCardByID(ctx, columnID, cardID)
}
