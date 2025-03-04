package contracts

import (
	"context"

	"github.com/emorydu/edgoms/pkg/config/environment"
	"github.com/emorydu/edgoms/pkg/logger"

	"go.uber.org/fx"
)

type Application interface {
	Container
	RegisterHook(function interface{})
	Run()
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Wait() <-chan fx.ShutdownSignal
	Logger() logger.Logger
	Environment() environment.Environment
}
