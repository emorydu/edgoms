package domain

import (
	"github.com/emorydu/edgoms/pkg/core/metadata"
)

type EventEnvelope struct {
	EventData interface{}
	Metadata  metadata.Metadata
}
