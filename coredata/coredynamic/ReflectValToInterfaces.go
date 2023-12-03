package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// ReflectValToInterfaces
//
// Assuming passing reflect val is an array or slice
// loop using reflection and returns the interfaces slice
func ReflectValToInterfaces(
	isSkipOnNil bool,
	reflectVal reflect.Value,
) []interface{} {
	return reflectinternal.Converter.ReflectValToInterfaces(
		isSkipOnNil,
		reflectVal,
	)
}
