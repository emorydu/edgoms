package redis

import (
	"context"
	"testing"

	"github.com/emorydu/edgoms/pkg/config"
	"github.com/emorydu/edgoms/pkg/config/environment"
	"github.com/emorydu/edgoms/pkg/core"
	"github.com/emorydu/edgoms/pkg/logger/external/fxlog"
	"github.com/emorydu/edgoms/pkg/logger/zap"
	redis2 "github.com/emorydu/edgoms/pkg/redis"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func Test_Custom_Redis_Container(t *testing.T) {
	ctx := context.Background()
	var redisClient redis.UniversalClient

	fxtest.New(t,
		config.ModuleFunc(environment.Test),
		zap.Module,
		fxlog.FxLogger,
		core.Module,
		redis2.Module,
		fx.Decorate(RedisContainerOptionsDecorator(t, ctx)),
		fx.Populate(&redisClient),
	).RequireStart()

	assert.NotNil(t, redisClient)
}
