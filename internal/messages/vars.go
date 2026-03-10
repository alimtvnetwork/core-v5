package messages

import (
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/errcore"
)

var (
	ComparatorOutOfRangeMessage = errcore.RangeNotMeet(
		errcore.ComparatorShouldBeWithinRangeType.String(),
		corecomparator.Min(),
		corecomparator.Max(),
		corecomparator.Ranges())
)
