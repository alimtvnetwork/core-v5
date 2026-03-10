package entityinf

import "github.com/alimtvnetwork/core/internal/internalinterface"

type RecordEntityDefiner interface {
	BaseRecordEntityDefiner
	internalinterface.DefaultsInjector
	internalinterface.RawPayloadsGetter
}
