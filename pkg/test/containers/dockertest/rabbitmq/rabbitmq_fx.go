package rabbitmq

import (
	"context"
	"testing"

	"github.com/emorydu/edgoms/pkg/logger"
	"github.com/emorydu/edgoms/pkg/rabbitmq/config"
)

var RabbitmqDockerTestContainerOptionsDecorator = func(t *testing.T, ctx context.Context) interface{} {
	return func(c *config.RabbitmqOptions, logger logger.Logger) (*config.RabbitmqOptions, error) {
		rabbitmqHostOptions, err := NewRabbitMQDockerTest(logger).PopulateContainerOptions(ctx, t)
		c.RabbitmqHostOptions = rabbitmqHostOptions

		return c, err
	}
}
