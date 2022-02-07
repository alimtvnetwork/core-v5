package enuminf

import "gitlab.com/evatix-go/core/coredata/corejson"

type BasicEnumer interface {
	enumNameStinger
	nameValuer
	IsNameEqualer
	IsAnyNameOfChecker
	ToNumberStringer
	IsValidInvalidChecker
	IsBothEnumEqualer
	BasicEnumValuer
	EnumFormatter
	corejson.JsonMarshaller
}
