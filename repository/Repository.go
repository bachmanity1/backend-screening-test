// Package repository ...
//
// Repository will store any Database handler.
// Querying, or Creating/ Inserting into any database will stored here.
// This layer will act for CRUD to database only.
// No business process happen here. Only plain function to Database.
//
// This layer also have responsibility to choose what DB will used in Application.
// Could be Mysql, MongoDB, MariaDB, Postgresql whatever, will decided here.
//
// If using ORM, this layer will control the input, and give it directly to ORM services.
//
// If calling microservices, will handled here. Create HTTP Request to other services, and sanitize the data.
// This layer, must fully act as a repository. Handle all data input - output no specific logic happen.
//
// This Repository layer will depends to Connected DB , or other microservices if exists.
package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"pandita/conf"
	"pandita/model"
	"pandita/util"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mlog *util.MLogger

func InitRepository(env string) {
	mlog, _ = util.InitLog("repository", env)
}

// InitDB ...
func InitDB(pandita *conf.ViperConfig) *gorm.DB {
	mlog.Debugw("InitDB ",
		"host", pandita.GetString("db_host"),
		"user", pandita.GetString("db_user"),
		"name", pandita.GetString("db_name"),
	)

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      getLogLevel(pandita.GetString("loglevel")),
			Colorful:      true,
		},
	)

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC",
		pandita.GetString("db_user"),
		pandita.GetString("db_pass"),
		pandita.GetString("db_host"),
		pandita.GetInt("db_port"),
		pandita.GetString("db_name"),
	)

	dbConn, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		mlog.Errorw("InitDB Open", "err", err)
		retryCount := 0
		for pandita.GetBool("db_retry") {
			time.Sleep(3 * time.Second)
			dbConn, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{
				Logger: dbLogger,
			})
			if err == nil {
				break
			}
			mlog.Errorw("InitDB Open", "err", err, "retry", retryCount, "dsn", dbURI)
			if retryCount > 3 {
				os.Exit(1)
			}
			retryCount++
		}
		if dbConn == nil {
			os.Exit(1)
		}
	}
	maxopen := pandita.GetInt("db_maxopen")
	db, _ := dbConn.DB()
	db.SetMaxIdleConns(int(float32(maxopen) * 0.9))
	db.SetMaxOpenConns(maxopen)
	db.SetConnMaxLifetime(time.Duration(pandita.GetInt("db_maxlife")) * time.Second)
	return dbConn
}

func getLogLevel(logLevel string) logger.LogLevel {
	l := strings.ToLower(logLevel)
	if strings.Contains(l, "sql_info") {
		return logger.Info
	}

	return logger.Silent
}

// ColumnRepository ...
type ColumnRepository interface {
	NewColumn(ctx context.Context, column *model.Column) (*model.Column, error)
	UpdateColumn(ctx context.Context, column *model.Column) (*model.Column, error)
	GetColumnByID(ctx context.Context, id uint64) (*model.Column, error)
	GetColumnList(ctx context.Context) (model.ColumnList, error)
	DeleteColumn(ctx context.Context, id uint64) error
}

// CardRepository ...
type CardRepository interface {
	NewCard(ctx context.Context, card *model.Card) (*model.Card, error)
	UpdateCard(ctx context.Context, card *model.Card) (*model.Card, error)
	GetCardByID(ctx context.Context, columnID, cardID uint64) (*model.Card, error)
	DeleteCard(ctx context.Context, columnID, cardID uint64) error
}
