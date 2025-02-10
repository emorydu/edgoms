package consumer

import (
	"context"

	"github.com/emorydu/edgoms/pkg/core/messaging/types"
)

type Consumer interface {
	Start(ctx context.Context) error
	Stop() error
	ConnectHandler(handler ConsumerHandler)
	IsConsumed(func(message types.IMessage))
	GetName() string
}
