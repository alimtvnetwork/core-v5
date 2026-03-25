package coredynamic

import (
	"reflect"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
)

func TypeNotEqualErr(
	left, right any,
) error {
	leftRt := reflect.TypeOf(left)
	rightRt := reflect.TypeOf(right)

	if leftRt == rightRt {
		return nil
	}

	return errcore.
		TypeMismatchType.
		SrcDestinationErr(
			"left, right type doesn't match!",
			constants.LeftLower, leftRt.String(),
			constants.Right, rightRt.String(),
		)
}
