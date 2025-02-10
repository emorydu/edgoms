package test

import (
	"github.com/emorydu/edgoms/pkg/config/environment"
	"github.com/emorydu/edgoms/pkg/fxapp"
	"github.com/emorydu/edgoms/pkg/fxapp/contracts"

	"go.uber.org/fx/fxtest"
)

type TestApplicationBuilder struct {
	contracts.ApplicationBuilder
	TB fxtest.TB
}

func NewTestApplicationBuilder(tb fxtest.TB) *TestApplicationBuilder {
	return &TestApplicationBuilder{
		TB:                 tb,
		ApplicationBuilder: fxapp.NewApplicationBuilder(environment.Test),
	}
}

func (a *TestApplicationBuilder) Build() contracts.Application {
	app := NewTestApplication(
		a.TB,
		a.GetProvides(),
		a.GetDecorates(),
		a.Options(),
		a.Logger(),
		environment.Test,
	)

	return app
}
