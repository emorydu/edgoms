package eventstroredb

import (
	"github.com/emorydu/edgoms/pkg/es/contracts/projection"
)

type ProjectionsConfigurations struct {
	Projections []projection.IProjection
}
