package isany

import (
	"reflect"
)

func TypeSame(
	left, right any,
) bool {
	leftRt := reflect.TypeOf(left)
	rightRt := reflect.TypeOf(right)

	return leftRt == rightRt
}
