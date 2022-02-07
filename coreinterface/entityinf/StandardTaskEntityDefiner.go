package entityinf

import (
	"gitlab.com/evatix-go/core/internal/internalinterface"
	"gitlab.com/evatix-go/core/internal/internalinterface/internalserializer"
)

type StandardTaskEntityDefiner interface {
	TaskEntityDefiner
	internalinterface.AnyAttributesGetter
	internalinterface.AnyAttributesReflectSetter
	internalinterface.IdAsStringer
	internalinterface.IdIntegerGetter
	internalinterface.HasErrorChecker
	IsStandardTaskEntityEqual(entity StandardTaskEntityDefiner) bool
	internalinterface.ValueReflectSetter
	internalserializer.SelfBytesSerializerDeserializer
}
