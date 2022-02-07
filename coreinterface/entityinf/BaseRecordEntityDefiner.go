package entityinf

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/internal/internalinterface"
	"gitlab.com/evatix-go/core/internal/internalinterface/internalserializer"
)

type BaseRecordEntityDefiner interface {
	internalinterface.IdentifierWithEqualer
	internalinterface.TypeNameWithEqualer
	internalinterface.EntityTypeNameWithEqualer
	internalinterface.CategoryNameWithEqualer
	internalinterface.TableNamer
	corejson.Jsoner
	internalserializer.FieldBytesToPointerDeserializer
}
