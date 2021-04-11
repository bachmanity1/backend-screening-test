// Package service ...
//
// This layer will act as the business process handler.
// Any process will handled here. This layer will decide, which repository layer will use.
// And have responsibility to provide data to serve into delivery.
// Process the data doing calculation or anything will done here.
//
// Service layer will accept any input from Delivery layer,
// that already sanitized, then process the input could be storing into DB ,
// or Fetching from DB ,etc.
//
// This Service layer will depends to Repository Layer
package service

import (
	"context"
	"pandita/model"
	"pandita/util"
)

var mlog *util.MLogger

func InitService(env string) {
	mlog, _ = util.InitLog("service", env)
}

// ColumnService ...
type ColumnService interface {
	NewColumn(ctx context.Context, column *model.Column) (*model.Column, error)
	UpdateColumn(ctx context.Context, column *model.Column) (*model.Column, error)
	GetColumnByID(ctx context.Context, id uint64) (*model.Column, error)
	DeleteColumn(ctx context.Context, id uint64) error
}

// CardService ...
type CardService interface {
	NewCard(ctx context.Context, card *model.Card) (*model.Card, error)
	UpdateCard(ctx context.Context, card *model.Card) (*model.Card, error)
	GetCardByID(ctx context.Context, id uint64) (*model.Card, error)
	DeleteCard(ctx context.Context, id uint64) error
}
