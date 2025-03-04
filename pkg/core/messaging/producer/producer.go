package producer

import (
	"context"

	"github.com/emorydu/edgoms/pkg/core/messaging/types"
	"github.com/emorydu/edgoms/pkg/core/metadata"
)

type Producer interface {
	PublishMessage(ctx context.Context, message types.IMessage, meta metadata.Metadata) error
	PublishMessageWithTopicName(
		ctx context.Context,
		message types.IMessage,
		meta metadata.Metadata,
		topicOrExchangeName string,
	) error
	IsProduced(func(message types.IMessage))
}
