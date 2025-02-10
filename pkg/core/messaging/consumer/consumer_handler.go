package consumer

import (
	"context"

	"github.com/emorydu/edgoms/pkg/core/messaging/types"
)

type ConsumerHandler interface {
	Handle(ctx context.Context, consumeContext types.MessageConsumeContext) error
}
