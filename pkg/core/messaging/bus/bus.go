package bus

import (
	consumer2 "github.com/emorydu/edgoms/pkg/core/messaging/consumer"
	"github.com/emorydu/edgoms/pkg/core/messaging/producer"
)

type Bus interface {
	producer.Producer
	consumer2.BusControl
	consumer2.ConsumerConnector
}
