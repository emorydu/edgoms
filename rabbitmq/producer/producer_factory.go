package producer

import (
	"github.com/emorydu/edgoms/pkg/core/messaging/producer"
	"github.com/emorydu/edgoms/pkg/core/messaging/types"
	serializer "github.com/emorydu/edgoms/pkg/core/serializer"
	"github.com/emorydu/edgoms/pkg/logger"
	"github.com/emorydu/edgoms/pkg/rabbitmq/config"
	producerConfigurations "github.com/emorydu/edgoms/pkg/rabbitmq/producer/configurations"
	"github.com/emorydu/edgoms/pkg/rabbitmq/producer/producercontracts"
	types2 "github.com/emorydu/edgoms/pkg/rabbitmq/types"
)

type producerFactory struct {
	connection      types2.IConnection
	logger          logger.Logger
	eventSerializer serializer.MessageSerializer
	rabbitmqOptions *config.RabbitmqOptions
}

func NewProducerFactory(
	rabbitmqOptions *config.RabbitmqOptions,
	connection types2.IConnection,
	eventSerializer serializer.MessageSerializer,
	l logger.Logger,
) producercontracts.ProducerFactory {
	return &producerFactory{
		rabbitmqOptions: rabbitmqOptions,
		logger:          l,
		connection:      connection,
		eventSerializer: eventSerializer,
	}
}

func (p *producerFactory) CreateProducer(
	rabbitmqProducersConfiguration map[string]*producerConfigurations.RabbitMQProducerConfiguration,
	isProducedNotifications ...func(message types.IMessage),
) (producer.Producer, error) {
	return NewRabbitMQProducer(
		p.rabbitmqOptions,
		p.connection,
		rabbitmqProducersConfiguration,
		p.logger,
		p.eventSerializer,
		isProducedNotifications...)
}
