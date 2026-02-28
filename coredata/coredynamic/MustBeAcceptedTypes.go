package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/errcore"
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
