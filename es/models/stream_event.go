package models

import (
	"github.com/emorydu/edgoms/pkg/core/domain"
	"github.com/emorydu/edgoms/pkg/core/metadata"

	uuid "github.com/satori/go.uuid"
)

type StreamEvent struct {
	EventID  uuid.UUID
	Version  int64
	Position int64
	Event    domain.IDomainEvent
	Metadata metadata.Metadata
}
