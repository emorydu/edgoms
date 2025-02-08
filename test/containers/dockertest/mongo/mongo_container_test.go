package mongo

import (
	"context"
	"testing"

	"github.com/emorydu/edgoms/pkg/config"
	"github.com/emorydu/edgoms/pkg/config/environment"
	"github.com/emorydu/edgoms/pkg/core"
	"github.com/emorydu/edgoms/pkg/logger/external/fxlog"
	"github.com/emorydu/edgoms/pkg/logger/zap"
	"github.com/emorydu/edgoms/pkg/mongodb"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func Test_Mongo_Container(t *testing.T) {
	ctx := context.Background()
	var mongoClient *mongo.Client

	fxtest.New(t,
		config.ModuleFunc(environment.Test),
		zap.Module,
		fxlog.FxLogger,
		core.Module,
		mongodb.Module,
		fx.Decorate(MongoDockerTestContainerOptionsDecorator(t, ctx)),
		fx.Populate(&mongoClient),
	).RequireStart()

	assert.NotNil(t, mongoClient)
}
