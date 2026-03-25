package enumimpl

import (
	"github.com/alimtvnetwork/core/constants"
)

func JoinPrependUsingDot(
	prepend any,
	anyItems ...any,
) string {
	return PrependJoin(
		constants.Dot,
		prepend,
		anyItems...)
}
