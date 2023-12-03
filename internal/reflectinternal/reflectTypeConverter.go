package reflectinternal

import (
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type reflectTypeConverter struct{}

func (it reflectConverter) SafeTypeName(any interface{}) string {
	rt := reflect.TypeOf(any)

	if Is.Null(rt) {
		return ""
	}

	return rt.String()
}

func (it reflectConverter) SafeTypeNameOfSliceOrSingle(
	isSingle bool,
	any interface{},
) string {
	if isSingle {
		return it.SafeTypeName(any)
	}

	return it.SliceFirstItemTypeName(any)
}

// SliceFirstItemTypeName
//
// Gets slice element type name, reduce ptr slice as well.
func (it reflectConverter) SliceFirstItemTypeName(slice interface{}) string {
	rt := reflect.TypeOf(slice)

	if Is.Null(rt) {
		return ""
	}

	if rt.Kind() == reflect.Ptr || rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	return rt.Elem().String()
}

func (it reflectConverter) TypeNamesStringUsingReflectType(
	isFullName bool,
	reflectTypes ...reflect.Type,
) string {
	return strings.Join(
		it.TypeNamesUsingReflectType(isFullName, reflectTypes...),
		constants.CommaSpace,
	)
}

func (it reflectConverter) TypeNamesString(
	isFullName bool,
	anyItems ...interface{},
) string {
	return strings.Join(
		TypeNames(isFullName, anyItems...),
		constants.CommaSpace,
	)
}

func (it reflectConverter) TypeNamesUsingReflectType(
	isFullName bool,
	reflectTypes ...reflect.Type,
) []string {
	slice := make([]string, len(reflectTypes))

	if isFullName {
		for i, item := range reflectTypes {
			slice[i] = item.String()
		}

		return slice
	}

	for i, item := range reflectTypes {
		slice[i] = item.Name()
	}

	return slice
}

func (it reflectConverter) TypeNamesReferenceString(
	isFullName bool,
	anyItems ...interface{},
) string {
	return "Reference (Types): " + strings.Join(
		TypeNames(isFullName, anyItems...),
		constants.CommaSpace,
	)
}

func (it reflectConverter) Names(
	isFullName bool,
	anyItems ...interface{},
) []string {
	slice := make([]string, len(anyItems))

	if isFullName {
		for i, item := range anyItems {
			slice[i] = reflect.TypeOf(item).String()
		}

		return slice
	}

	for i, item := range anyItems {
		slice[i] = reflect.TypeOf(item).Name()
	}

	return slice
}

func (it reflectConverter) Name(any interface{}) string {
	rf := reflect.TypeOf(any)

	if rf == nil {
		return ""
	}

	return rf.String()
}
