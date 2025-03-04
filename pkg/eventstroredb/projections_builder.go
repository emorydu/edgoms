package eventstroredb

import (
	"github.com/emorydu/edgoms/pkg/es/contracts/projection"
)

type ProjectionsBuilder interface {
	AddProjection(projection projection.IProjection) ProjectionsBuilder
	AddProjections(projections []projection.IProjection) ProjectionsBuilder
	Build() *ProjectionsConfigurations
}
