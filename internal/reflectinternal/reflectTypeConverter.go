package reflectinternal

import (
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type reflectTypeConverter struct{}

func (it reflectTypeConverter) SafeName(anyItem any) string {
	rt := reflect.TypeOf(anyItem)

	if Is.Null(rt) {
		return ""
	}

	return rt.String()
}

func (it reflectTypeConverter) SafeTypeNameOfSliceOrSingle(
	isSingle bool,
	anyItem any,
) string {
	if isSingle {
		return it.SafeName(anyItem)
	}

	return it.SliceFirstItemTypeName(anyItem)
}

// SliceFirstItemTypeName
//
// Gets slice element type name, reduce ptr slice as well.
func (it reflectTypeConverter) SliceFirstItemTypeName(slice any) string {
	rt := reflect.TypeOf(slice)

	if Is.Null(rt) {
		return ""
	}

	if rt.Kind() == reflect.Ptr || rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	return rt.Elem().String()
}

func (it reflectTypeConverter) NamesStringUsingReflectType(
	isFullName bool,
	reflectTypes ...reflect.Type,
) string {
	if len(reflectTypes) == 0 {
		return ""
	}

	return strings.Join(
		it.NamesUsingReflectType(isFullName, reflectTypes...),
		constants.CommaSpace,
	)
}

func (it reflectTypeConverter) TypeNamesString(
	isFullName bool,
	anyItems ...any,
) string {
	if len(anyItems) == 0 {
		return ""
	}

	return strings.Join(
		TypeNames(isFullName, anyItems...),
		constants.CommaSpace,
	)
}

func (it reflectTypeConverter) NamesUsingReflectType(
	isFullName bool,
	reflectTypes ...reflect.Type,
) []string {
	if len(reflectTypes) == 0 {
		return []string{}
	}

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

func (it reflectTypeConverter) NamesReferenceString(
	isFullName bool,
	anyItems ...any,
) string {
	if len(anyItems) == 0 {
		return ""
	}

	return "Reference (Types): " + strings.Join(
		TypeNames(isFullName, anyItems...),
		constants.CommaSpace,
	)
}

func (it reflectTypeConverter) Names(
	isFullName bool,
	anyItems ...any,
) []string {
	if len(anyItems) == 0 {
		return []string{}
	}

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

func (it reflectTypeConverter) Name(anyItem any) string {
	rf := reflect.TypeOf(anyItem)

	if rf == nil {
		return ""
	}

	return rf.String()
}

func (it reflectTypeConverter) NameUsingFmt(anyItem any) string {
	return fmt.Sprintf(constants.SprintTypeFormat, anyItem)
}
