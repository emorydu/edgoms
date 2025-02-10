package eventstoredb

import (
	"context"
	"testing"

	"github.com/emorydu/edgoms/pkg/config"
	"github.com/emorydu/edgoms/pkg/config/environment"
	"github.com/emorydu/edgoms/pkg/core"
	"github.com/emorydu/edgoms/pkg/eventstroredb"
	"github.com/emorydu/edgoms/pkg/logger/external/fxlog"
	"github.com/emorydu/edgoms/pkg/logger/zap"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func Test_Custom_EventStoreDB_Container(t *testing.T) {
	var esdbClient *esdb.Client
	ctx := context.Background()

	fxtest.New(t,
		config.ModuleFunc(environment.Test),
		zap.Module,
		fxlog.FxLogger,
		core.Module,
		eventstroredb.ModuleFunc(func() {
		}),
		fx.Decorate(EventstoreDBContainerOptionsDecorator(t, ctx)),
		fx.Populate(&esdbClient),
	).RequireStart()
}
