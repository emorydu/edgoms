package config

import (
	"github.com/emorydu/edgoms/pkg/config"
	"github.com/emorydu/edgoms/pkg/config/environment"
	typeMapper "github.com/emorydu/edgoms/pkg/reflection/typemapper"

	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[GrpcOptions]())

type GrpcOptions struct {
	Port        string `mapstructure:"port"        env:"TcpPort"`
	Host        string `mapstructure:"host"        env:"Host"`
	Development bool   `mapstructure:"development" env:"Development"`
	Name        string `mapstructure:"name"        env:"ShortTypeName"`
}

func ProvideConfig(environment environment.Environment) (*GrpcOptions, error) {
	return config.BindConfigKey[*GrpcOptions](optionName, environment)
}
