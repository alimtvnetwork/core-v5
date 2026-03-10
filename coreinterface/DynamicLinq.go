package coreinterface

import "github.com/alimtvnetwork/core/internal/internalinterface"

type DynamicLinq interface {
	CountGetter
	LengthGetter
	FirstDynamic() any
	LastDynamic() any
	FirstOrDefaultDynamic() any
	LastOrDefaultDynamic() any
	SkipDynamic(skippingItemsCount int) any
	TakeDynamic(takeDynamicItems int) any
	// LimitDynamic alias for TakeDynamic
	LimitDynamic(limit int) any
}

type DyanmicLinqer interface {
	internalinterface.DyanmicLinqer
}
