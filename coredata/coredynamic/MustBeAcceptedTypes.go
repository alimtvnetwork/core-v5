package coredynamic

import (
	"reflect"

	"github.com/alimtvnetwork/core/errcore"
)

func MustBeAcceptedTypes(
	input any,
	acceptedTypes ...reflect.Type,
) {
	err := NotAcceptedTypesErr(
		input,
		acceptedTypes...)
	errcore.HandleErr(err)
}
