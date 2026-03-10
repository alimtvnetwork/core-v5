package coredynamic

import (
	"reflect"

	"github.com/alimtvnetwork/core/errcore"
)

func NotAcceptedTypesErr(
	input any,
	acceptedTypes ...reflect.Type,
) error {
	currentRv := reflect.TypeOf(input)
	isNotMatchingAcceptedType := !IsAnyTypesOf(
		currentRv,
		acceptedTypes...)

	if isNotMatchingAcceptedType {
		return errcore.
			ExpectingSimpleNoTypeError(
				"type doesn't match, accepting types are listed.",
				acceptedTypes,
				currentRv)
	}

	return nil
}
