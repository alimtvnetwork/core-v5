package isany

import "github.com/alimtvnetwork/core/internal/reflectinternal"

func NotDeepEqual(
	left, right any,
) (isNotEqual bool) {
	return !reflectinternal.Is.AnyEqual(left, right)
}
