package serializerinf

import "gitlab.com/evatix-go/core/internal/internalinterface/internalserializer"

type SerializerDeserializerBinder interface {
	internalserializer.SerializerDeserializer
	AsSerializerDeserializerBinder() internalserializer.SerializerDeserializerBinder
}
