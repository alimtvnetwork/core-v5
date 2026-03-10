package enuminf

import "github.com/alimtvnetwork/core/coredata/corejson"

type BaseEnumer interface {
	enumNameStinger
	SimpleEnumer
	NameValuer
	IsNameEqualer
	IsAnyNameOfChecker
	ToNumberStringer
	IsValidInvalidChecker
	BasicEnumValuer
	RangeNamesCsvGetter
	corejson.JsonMarshaller
}
