package consumer

import (
	"github.com/emorydu/edgoms/pkg/core/messaging/consumer"
	"github.com/emorydu/edgoms/pkg/core/messaging/types"
	serializer "github.com/emorydu/edgoms/pkg/core/serializer"
	"github.com/emorydu/edgoms/pkg/logger"
	"github.com/emorydu/edgoms/pkg/rabbitmq/config"
	consumerConfigurations "github.com/emorydu/edgoms/pkg/rabbitmq/consumer/configurations"
	"github.com/emorydu/edgoms/pkg/rabbitmq/consumer/consumercontracts"
	types2 "github.com/emorydu/edgoms/pkg/rabbitmq/types"
)

type consumerFactory struct {
	connection      types2.IConnection
	eventSerializer serializer.MessageSerializer
	logger          logger.Logger
	rabbitmqOptions *config.RabbitmqOptions
}

func NewConsumerFactory(
	rabbitmqOptions *config.RabbitmqOptions,
	connection types2.IConnection,
	eventSerializer serializer.MessageSerializer,
	l logger.Logger,
) consumercontracts.ConsumerFactory {
	return &consumerFactory{
		rabbitmqOptions: rabbitmqOptions,
		logger:          l,
		eventSerializer: eventSerializer,
		connection:      connection,
	}
}

func (c *consumerFactory) CreateConsumer(
	consumerConfiguration *consumerConfigurations.RabbitMQConsumerConfiguration,
	isConsumedNotifications ...func(message types.IMessage),
) (consumer.Consumer, error) {
	return NewRabbitMQConsumer(
		c.rabbitmqOptions,
		c.connection,
		consumerConfiguration,
		c.eventSerializer,
		c.logger,
		isConsumedNotifications...)
}

func (c *consumerFactory) Connection() types2.IConnection {
	return c.connection
}
