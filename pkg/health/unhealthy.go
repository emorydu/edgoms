package health

import (
	"context"

	"github.com/emorydu/edgoms/pkg/health/contracts"
)

type UnhealthyHealthService struct{}

func NewUnhealthyHealthService() UnhealthyHealthService {
	return UnhealthyHealthService{}
}

func (service UnhealthyHealthService) CheckHealth(
	context.Context,
) contracts.Check {
	return contracts.Check{
		"postgres": contracts.Status{Status: contracts.StatusDown},
		"redis":    contracts.Status{Status: contracts.StatusDown},
	}
}
