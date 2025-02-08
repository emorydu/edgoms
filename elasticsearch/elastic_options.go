package elasticsearch

import (
	"github.com/emorydu/edgoms/pkg/config"
	"github.com/emorydu/edgoms/pkg/config/environment"
	typeMapper "github.com/emorydu/edgoms/pkg/reflection/typemapper"

	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[ElasticOptions]())

type ElasticOptions struct {
	URL string `mapstructure:"url"`
}

func provideConfig(environment environment.Environment) (*ElasticOptions, error) {
	return config.BindConfigKey[*ElasticOptions](optionName, environment)
}
