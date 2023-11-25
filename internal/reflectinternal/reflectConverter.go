package reflectinternal

import (
	"reflect"
	"unsafe"

	"gitlab.com/auk-go/core/internal/reflectmodel"
)

type reflectConverter struct{}

func (it reflectConverter) ArgsToReflectValues(args []interface{}) []reflect.Value {
	if len(args) == 0 {
		return []reflect.Value{}
	}

	list := make(
		[]reflect.Value,
		len(args),
	)

	for i, arg := range args {
		list[i] = reflect.ValueOf(arg)
	}

	return list
}

func (it reflectConverter) ReflectValuesToInterfaces(
	reflectValues []reflect.Value,
) []interface{} {
	if len(reflectValues) == 0 {
		return []interface{}{}
	}

	list := make(
		[]interface{},
		len(reflectValues),
	)

	for i, rv := range reflectValues {
		list[i] = it.ReflectValueToAnyValue(rv)
	}

	return list
}

func (it reflectConverter) ReflectValueToAnyValue(rv reflect.Value) interface{} {
	if IsNull(rv) {
		return nil
	}

	k := rv.Kind()

	switch k {
	case reflect.Ptr, reflect.Interface:
		return rv.Elem().Interface()
	default:
		return rv.Interface()
	}
}

func (it reflectConverter) InterfacesToTypes(items []interface{}) []reflect.Type {
	if len(items) == 0 {
		return []reflect.Type{}
	}

	var output []reflect.Type

	for _, item := range items {
		toType := reflect.TypeOf(item)
		output = append(output, toType)

	}

	return output
}

func (it reflectConverter) InterfacesToTypesNames(items []interface{}) []string {
	if len(items) == 0 {
		return []string{}
	}

	var output []string

	for _, item := range items {
		toType := reflect.TypeOf(item)
		output = append(output, toType.Name())

	}

	return output
}

func (it reflectConverter) ReflectValueToPointerReflectValue(
	rv reflect.Value,
) reflect.Value {
	toInterface := rv.Interface()
	toPointer := &toInterface
	unsafePtr := unsafe.Pointer(&toPointer)

	return reflect.NewAt(rv.Type(), unsafePtr)
}

// ReducePointer
//
// anyItem must be a struct or pointer to struct
//
// level means how many ****Struct to reduce to Struct
func (it reflectConverter) ReducePointer(
	anyItem interface{},
	level int,
) *reflectmodel.ReflectValueKind {
	rv := reflect.ValueOf(anyItem) // can be a pointer or non pointer

	return it.ReducePointerRv(rv, level)
}

// ReducePointerRv
//
// # Rv must be a struct or pointer to struct
//
// level means how many ****Struct to reduce to Struct
func (it reflectConverter) ReducePointerRv(
	rv reflect.Value,
	level int,
) *reflectmodel.ReflectValueKind {
	return Looper.ReducePointerRv(rv, level)
}

// ReducePointerDefault
//
// anyItem must be a struct or pointer to struct
//
// Default means level 3 at max
func (it reflectConverter) ReducePointerDefault(
	anyItem interface{},
) *reflectmodel.ReflectValueKind {
	return it.ReducePointerDefault(anyItem)
}

// ReducePointerRvDefault
//
// # Rv must be a struct or pointer to struct
//
// level means how many ****Struct to reduce to Struct
//
// Default means level 3
func (it reflectConverter) ReducePointerRvDefault(
	rv reflect.Value,
) *reflectmodel.ReflectValueKind {
	return Looper.ReducePointerRvDefault(rv)
}

func (it reflectConverter) ReducePointerDefaultToType(
	anyItem interface{},
) *reflect.Type {
	rv := reflect.ValueOf(anyItem)

	return it.ReducePointerRvDefaultToType(rv)
}

// ReducePointerRvDefaultToType
//
// # Rv must be a struct or pointer to struct
//
// level means how many ****Struct to reduce to Struct
//
// Default means level 3
func (it reflectConverter) ReducePointerRvDefaultToType(
	rv reflect.Value,
) *reflect.Type {
	result := Looper.ReducePointerRvDefault(rv)

	if result != nil {
		toType := result.FinalReflectVal.Type()

		return &toType
	}

	return nil
}
