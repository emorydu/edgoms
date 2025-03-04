//go:build unit
// +build unit

package messagepersistence

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/emorydu/edgoms/pkg/config"
	"github.com/emorydu/edgoms/pkg/config/environment"
	"github.com/emorydu/edgoms/pkg/core"
	"github.com/emorydu/edgoms/pkg/core/messaging/persistmessage"
	"github.com/emorydu/edgoms/pkg/logger"
	defaultLogger "github.com/emorydu/edgoms/pkg/logger/defaultlogger"
	"github.com/emorydu/edgoms/pkg/logger/external/fxlog"
	"github.com/emorydu/edgoms/pkg/logger/zap"
	"github.com/emorydu/edgoms/pkg/mapper"
	"github.com/emorydu/edgoms/pkg/postgresgorm"
	"github.com/emorydu/edgoms/pkg/postgresgorm/helpers/gormextensions"

	"emperror.dev/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"gorm.io/gorm"
)

type postgresMessageServiceTest struct {
	suite.Suite
	DB                  *gorm.DB
	logger              logger.Logger
	messagingRepository persistmessage.MessagePersistenceService
	dbContext           *PostgresMessagePersistenceDBContext
	storeMessages       []*persistmessage.StoreMessage
	ctx                 context.Context
	dbFilePath          string
	app                 *fxtest.App
}

func TestPostgresMessageService(t *testing.T) {
	suite.Run(
		t,
		&postgresMessageServiceTest{logger: defaultLogger.GetLogger()},
	)
}

//func (c *postgresMessageServiceTest) SetupSuite() {
//	opts, err := gorm2.NewGormTestContainers(defaultLogger.GetLogger()).
//		PopulateContainerOptions(context.Background(), c.T())
//	c.Require().NoError(err)
//
//	gormDB, err := postgresgorm.NewGorm(opts)
//	c.Require().NoError(err)
//	c.DB = gormDB
//
//	err = migrationDatabase(gormDB)
//	c.Require().NoError(err)
//
//	c.dbContext = NewPostgresMessagePersistenceDBContext(gormDB)
//	c.messagingRepository = NewPostgresMessageService(
//		c.dbContext,
//		defaultLogger.GetLogger(),
//	)
//}

func (c *postgresMessageServiceTest) SetupTest() {
	var gormDBContext *PostgresMessagePersistenceDBContext
	var gormOptions *postgresgorm.GormOptions

	app := fxtest.New(
		c.T(),
		config.ModuleFunc(environment.Test),
		zap.Module,
		fxlog.FxLogger,
		core.Module,
		postgresgorm.Module,
		fx.Decorate(
			func(cfg *postgresgorm.GormOptions) (*postgresgorm.GormOptions, error) {
				// using sql-lite with a database file
				cfg.UseSQLLite = true

				return cfg, nil
			},
		),
		fx.Provide(NewPostgresMessagePersistenceDBContext),
		fx.Populate(&gormDBContext),
		fx.Populate(&gormOptions),
	).RequireStart()

	c.dbContext = gormDBContext
	c.dbFilePath = gormOptions.Dns()
	c.app = app

	c.initDB()
}

func (c *postgresMessageServiceTest) TearDownTest() {
	err := c.cleanupDB()
	c.Require().NoError(err)

	mapper.ClearMappings()

	c.app.RequireStop()
}

//func (c *postgresMessageServiceTest) SetupTest() {
//	ctx := context.Background()
//	c.ctx = ctx
//	p, err := seedData(context.Background(), c.DB)
//	c.Require().NoError(err)
//	c.storeMessages = p
//}
//
//func (c *postgresMessageServiceTest) TearDownTest() {
//	err := c.cleanupPostgresData()
//	c.Require().NoError(err)
//}

func (c *postgresMessageServiceTest) BeginTx() {
	c.logger.Info("starting transaction")
	tx := c.dbContext.DB().Begin()
	gormContext := gormextensions.SetTxToContext(c.ctx, tx)
	c.ctx = gormContext
}

func (c *postgresMessageServiceTest) CommitTx() {
	tx := gormextensions.GetTxFromContextIfExists(c.ctx)
	if tx != nil {
		c.logger.Info("committing transaction")
		tx.Commit()
	}
}

func (c *postgresMessageServiceTest) Test_Add() {
	message := &persistmessage.StoreMessage{
		ID:            uuid.NewV4(),
		MessageStatus: persistmessage.Processed,
		Data:          "test data 3",
		DataType:      "string",
		CreatedAt:     time.Now(),
		DeliveryType:  persistmessage.Outbox,
	}

	c.BeginTx()
	err := c.messagingRepository.Add(c.ctx, message)
	c.CommitTx()

	c.Require().NoError(err)

	m, err := c.messagingRepository.GetById(c.ctx, message.ID)
	if err != nil {
		return
	}

	c.Assert().NotNil(m)
	c.Assert().Equal(message.ID, m.ID)
}

func (c *postgresMessageServiceTest) initDB() {
	err := migrateGorm(c.dbContext.DB())
	c.Require().NoError(err)

	storeMessages, err := seedData(c.dbContext.DB())
	c.Require().NoError(err)

	c.storeMessages = storeMessages
}

func (c *postgresMessageServiceTest) cleanupDB() error {
	sqldb, _ := c.dbContext.DB().DB()
	e := sqldb.Close()
	c.Require().NoError(e)

	// removing sql-lite file
	err := os.Remove(c.dbFilePath)

	return err
}

func migrateGorm(db *gorm.DB) error {
	err := db.AutoMigrate(&persistmessage.StoreMessage{})
	if err != nil {
		return err
	}

	return nil
}

func seedData(
	db *gorm.DB,
) ([]*persistmessage.StoreMessage, error) {
	messages := []*persistmessage.StoreMessage{
		{
			ID:            uuid.NewV4(),
			MessageStatus: persistmessage.Processed,
			Data:          "test data",
			DataType:      "string",
			CreatedAt:     time.Now(),
			DeliveryType:  persistmessage.Outbox,
		},
		{
			ID:            uuid.NewV4(),
			MessageStatus: persistmessage.Processed,
			Data:          "test data 2",
			DataType:      "string",
			CreatedAt:     time.Now(),
			DeliveryType:  persistmessage.Outbox,
		},
	}

	// seed data
	err := db.CreateInBatches(messages, len(messages)).Error
	if err != nil {
		return nil, errors.Wrap(err, "error in seed database")
	}

	return messages, nil
}
