package reflectinternal

import (
	"errors"
	"fmt"
	"reflect"
)

type reflectGetter struct{}

// PublicValuesMapStruct
//
//	returns structs fields map[string]any
//	map[string:fieldName]any:PublicValue
//
//	Only public values will be collected into map values
func (it reflectGetter) PublicValuesMapStruct(anyItem any) (
	map[string]any, error,
) {
	if Is.Null(anyItem) {
		return map[string]any{},
			errors.New("null given to expand map[name]value, failed")
	}

	return ReflectGetterUsingReflectValue.PublicValuesMapStruct(
		reflect.ValueOf(anyItem),
	)
}

// FieldNameWithValuesMap
//
//	returns structs fields map[string]any
//	map[string:fieldName]reflect.Type:fieldType
//
//	unlike PublicValuesMapStruct to map it collects
//	all fields with values including the private ones.
//
// However, this one will be slower in performance than PublicValuesMapStruct.
func (it reflectGetter) FieldNameWithValuesMap(anyItem any) (
	r map[string]any, error error,
) {
	if Is.Null(anyItem) {
		return map[string]any{},
			it.nullError(r)
	}

	return ReflectGetterUsingReflectValue.FieldNameWithValuesMap(
		reflect.ValueOf(anyItem),
	)
}

func (it reflectGetter) nullError(i any) error {
	return fmt.Errorf("null given to expand %T, failed", i)
}

// FieldNamesMap
//
//	returns structs fields map[string]bool names
//	map[string:fieldName]bool:exists
func (it reflectGetter) FieldNamesMap(
	anyItem any,
) (
	r map[string]bool, err error,
) {
	if Is.Null(anyItem) {
		return map[string]bool{},
			it.nullError(r)
	}

	return ReflectGetterUsingReflectValue.FieldNamesMap(
		reflect.ValueOf(anyItem),
	)
}

// StructFieldsMap
//
//	returns structs all fields (public, private) map[string]reflect.StructField
//	map[string:fieldName]reflect.StructField:StructField
func (it reflectGetter) StructFieldsMap(
	anyItem any,
) map[string]reflect.StructField {
	if Is.Null(anyItem) {
		return map[string]reflect.StructField{}
	}

	return ReflectGetterUsingReflectValue.StructFieldsMap(
		reflect.ValueOf(anyItem),
	)
}

// NullFieldsMap
//
//	returns structs all fields (public, private) map[string]bool
//	null fields map only
func (it reflectGetter) NullFieldsMap(
	anyItem any,
) map[string]bool {
	if Is.Null(anyItem) {
		return map[string]bool{}
	}

	return ReflectGetterUsingReflectValue.NullFieldsMap(
		defaultMaxLevelOfReflection,
		reflect.ValueOf(anyItem),
	)
}

// NullOrZeroFieldsMap
//
//	returns structs all fields (public, private) map[string]bool
//	null or zero fields map only
func (it reflectGetter) NullOrZeroFieldsMap(
	anyItem any,
) map[string]bool {
	if Is.Null(anyItem) {
		return map[string]bool{}
	}

	return ReflectGetterUsingReflectValue.NullOrZeroFieldsMap(
		defaultMaxLevelOfReflection,
		reflect.ValueOf(anyItem),
	)
}
