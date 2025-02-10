package postgresmessaging

import (
	"github.com/emorydu/edgoms/pkg/core/messaging/persistmessage"
	"github.com/emorydu/edgoms/pkg/postgresmessaging/messagepersistence"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module(
	"postgresmessagingfx",
	fx.Provide(
		messagepersistence.NewPostgresMessagePersistenceDBContext,
		messagepersistence.NewPostgresMessageService,
	),
	fx.Invoke(migrateMessaging),
)

func migrateMessaging(db *gorm.DB) error {
	err := db.Migrator().AutoMigrate(&persistmessage.StoreMessage{})

	return err
}
