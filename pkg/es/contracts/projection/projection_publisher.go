package projection

import (
	"context"

	"github.com/emorydu/edgoms/pkg/es/models"
)

type IProjectionPublisher interface {
	Publish(ctx context.Context, streamEvent *models.StreamEvent) error
}
