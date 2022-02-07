package entityinf

import "gitlab.com/evatix-go/core/internal/internalinterface"

type RecordEntityDefiner interface {
	BaseRecordEntityDefiner
	internalinterface.DefaultsInjector
	internalinterface.RawPayloadsGetter
}
