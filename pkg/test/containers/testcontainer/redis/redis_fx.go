package redis

import (
	"context"
	"testing"

	"github.com/emorydu/edgoms/pkg/logger"
	"github.com/emorydu/edgoms/pkg/redis"
)

var RedisContainerOptionsDecorator = func(t *testing.T, ctx context.Context) interface{} {
	return func(c *redis.RedisOptions, logger logger.Logger) (*redis.RedisOptions, error) {
		return NewRedisTestContainers(logger).PopulateContainerOptions(ctx, t)
	}
}
