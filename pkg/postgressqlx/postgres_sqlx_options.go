package postgressqlx

import (
	"github.com/emorydu/edgoms/pkg/config"
	"github.com/emorydu/edgoms/pkg/config/environment"
	typeMapper "github.com/emorydu/edgoms/pkg/reflection/typemapper"

	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[PostgresSqlxOptions]())

type PostgresSqlxOptions struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"dbName"`
	SSLMode  bool   `mapstructure:"sslMode"`
	Password string `mapstructure:"password"`
}

func provideConfig(environment environment.Environment) (*PostgresSqlxOptions, error) {
	return config.BindConfigKey[*PostgresSqlxOptions](optionName, environment)
}
