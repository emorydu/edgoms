package serializer

import "github.com/emorydu/edgoms/pkg/core/metadata"

type MetadataSerializer interface {
	Serialize(meta metadata.Metadata) ([]byte, error)
	Deserialize(bytes []byte) (metadata.Metadata, error)
}
