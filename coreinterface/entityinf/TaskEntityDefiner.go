package entityinf

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coreinterface"
	"github.com/alimtvnetwork/core/internal/internalinterface"
)

type TaskEntityDefiner interface {
	internalinterface.UsernameGetter
	internalinterface.AnyValueGetter
	internalinterface.ErrorGetter

	coreinterface.ReflectSetter
	Deserialize(
		anyPointer any,
	) error

	corejson.Jsoner
}
