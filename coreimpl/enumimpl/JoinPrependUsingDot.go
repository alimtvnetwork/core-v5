package enumimpl

import (
	"gitlab.com/auk-go/core/constants"
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
