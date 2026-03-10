package isany

import (
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

func DeepEqual(
	left, right any,
) (isEqual bool) {
	return reflectinternal.Is.AnyEqual(left, right)
}
