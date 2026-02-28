package isany

import "gitlab.com/auk-go/core/internal/reflectinternal"

func NotDeepEqual(
	left, right any,
) (isNotEqual bool) {
	return !reflectinternal.Is.AnyEqual(left, right)
}
