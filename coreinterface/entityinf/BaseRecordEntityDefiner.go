package entityinf

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/internal/internalinterface"
	"github.com/alimtvnetwork/core/internal/internalinterface/internalserializer"
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
