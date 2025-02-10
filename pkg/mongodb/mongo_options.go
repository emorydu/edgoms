package mongodb

import (
	"github.com/emorydu/edgoms/pkg/config"
	"github.com/emorydu/edgoms/pkg/config/environment"
	typeMapper "github.com/emorydu/edgoms/pkg/reflection/typemapper"

	"github.com/iancoleman/strcase"
)

type MongoDbOptions struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	User          string `mapstructure:"user"`
	Password      string `mapstructure:"password"`
	Database      string `mapstructure:"database"`
	UseAuth       bool   `mapstructure:"useAuth"`
	EnableTracing bool   `mapstructure:"enableTracing" default:"true"`
}

func provideConfig(
	environment environment.Environment,
) (*MongoDbOptions, error) {
	optionName := strcase.ToLowerCamel(
		typeMapper.GetGenericTypeNameByT[MongoDbOptions](),
	)
	return config.BindConfigKey[*MongoDbOptions](optionName, environment)
}
