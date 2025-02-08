package producercontracts

import (
	"github.com/emorydu/edgoms/pkg/core/messaging/producer"
	types2 "github.com/emorydu/edgoms/pkg/core/messaging/types"
	"github.com/emorydu/edgoms/pkg/rabbitmq/producer/configurations"
)

type ProducerFactory interface {
	CreateProducer(
		rabbitmqProducersConfiguration map[string]*configurations.RabbitMQProducerConfiguration,
		isProducedNotifications ...func(message types2.IMessage),
	) (producer.Producer, error)
}
