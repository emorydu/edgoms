package rabbitmq

import (
	"context"
	"testing"

	"github.com/emorydu/edgoms/pkg/logger"
	"github.com/emorydu/edgoms/pkg/rabbitmq/config"
)

var RabbitmqContainerOptionsDecorator = func(t *testing.T, ctx context.Context) interface{} {
	return func(c *config.RabbitmqOptions, logger logger.Logger) (*config.RabbitmqOptions, error) {
		rabbitmqHostOptions, err := NewRabbitMQTestContainers(logger).PopulateContainerOptions(ctx, t)
		c.RabbitmqHostOptions = rabbitmqHostOptions

		return c, err
	}
}
