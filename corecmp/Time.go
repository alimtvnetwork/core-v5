package corecmp

import (
	"time"

	"github.com/alimtvnetwork/core/corecomparator"
)

func Time(left, right time.Time) corecomparator.Compare {
	if left.Before(right) {
		return corecomparator.LeftLess
	} else if left.After(right) {
		return corecomparator.LeftGreater
	} else if left.Equal(right) {
		return corecomparator.Equal
	}

	return corecomparator.NotEqual
}
