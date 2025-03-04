package configurations

import (
	"github.com/emorydu/edgoms/pkg/core/messaging/consumer"
	"github.com/emorydu/edgoms/pkg/core/messaging/types"
)

type RabbitMQConsumerConnector interface {
	consumer.ConsumerConnector
	// ConnectRabbitMQConsumer Add a new consumer to existing message type consumers. if there is no consumer, will create a new consumer for the message type
	ConnectRabbitMQConsumer(
		messageType types.IMessage,
		consumerBuilderFunc RabbitMQConsumerConfigurationBuilderFuc,
	) error
}
