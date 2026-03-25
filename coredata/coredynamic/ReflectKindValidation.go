package coredynamic

import (
	"reflect"

	"github.com/alimtvnetwork/core/errcore"
)

func ReflectKindValidation(
	expectedKind reflect.Kind,
	anyItem any,
) error {
	actualKind := reflect.
		ValueOf(anyItem).
		Kind()

	if actualKind == expectedKind {
		return nil
	}

	return errcore.ExpectingErrorSimpleNoType(
		"ReflectKindValidation: reflect kind validation failed",
		expectedKind,
		actualKind)
}
