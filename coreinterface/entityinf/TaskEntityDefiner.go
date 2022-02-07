package entityinf

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/internal/internalinterface"
)

type TaskEntityDefiner interface {
	internalinterface.UsernameGetter
	internalinterface.AnyValueGetter
	internalinterface.ErrorGetter
	corejson.Jsoner
}
