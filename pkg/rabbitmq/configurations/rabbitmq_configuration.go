package configurations

import (
	consumerConfigurations "github.com/emorydu/edgoms/pkg/rabbitmq/consumer/configurations"
	producerConfigurations "github.com/emorydu/edgoms/pkg/rabbitmq/producer/configurations"
)

type RabbitMQConfiguration struct {
	ProducersConfigurations []*producerConfigurations.RabbitMQProducerConfiguration
	ConsumersConfigurations []*consumerConfigurations.RabbitMQConsumerConfiguration
}
