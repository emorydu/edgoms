package consumercontracts

import (
	"github.com/emorydu/edgoms/pkg/core/messaging/consumer"
	messagingTypes "github.com/emorydu/edgoms/pkg/core/messaging/types"
	"github.com/emorydu/edgoms/pkg/rabbitmq/consumer/configurations"
	"github.com/emorydu/edgoms/pkg/rabbitmq/types"
)

type ConsumerFactory interface {
	CreateConsumer(
		consumerConfiguration *configurations.RabbitMQConsumerConfiguration,
		isConsumedNotifications ...func(message messagingTypes.IMessage),
	) (consumer.Consumer, error)

	Connection() types.IConnection
}
