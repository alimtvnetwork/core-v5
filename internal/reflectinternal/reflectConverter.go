package reflectinternal

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/refeflectcore/reflectmodel"
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
	if Is.Null(rv) {
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

func (it reflectConverter) InterfacesToTypesNamesWithValues(items []interface{}) []string {
	if len(items) == 0 {
		return []string{}
	}

	var output []string

	for i, item := range items {
		toType := reflect.TypeOf(item)
		compiledString := fmt.Sprintf(
			"%d. %s [value: %s]",
			i,
			toType.Name(),
			convertinteranl.AnyTo.SmartString(item),
		)

		output = append(output, compiledString)

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

func ReflectValToInterfaces(
	isSkipOnNil bool,
	reflectVal reflect.Value,
) []interface{} {
	if reflectVal.Kind() == reflect.Ptr {
		return ReflectValToInterfaces(
			isSkipOnNil,
			reflect.Indirect(reflect.ValueOf(reflectVal)),
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []interface{}{}
	}

	length := reflectVal.Len()
	slice := make([]interface{}, 0, length)

	if length == 0 {
		return slice
	}

	for i := 0; i < length; i++ {
		value := reflectVal.Index(i)

		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		valueInf := value.Interface()

		if isSkipOnNil && Is.Null(value) {
			continue
		}

		slice = append(slice, valueInf)
	}

	return slice
}

func ReflectValToInterfacesAsync(
	reflectVal reflect.Value,
) []interface{} {
	if reflectVal.Kind() == reflect.Ptr {
		return ReflectValToInterfacesAsync(
			reflect.Indirect(reflect.ValueOf(reflectVal)),
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []interface{}{}
	}

	length := reflectVal.Len()
	slice := make([]interface{}, length)

	if length == 0 {
		return slice
	}

	wg := sync.WaitGroup{}
	setterIndexFunc := func(index int) {
		value := reflectVal.Index(index)

		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		valueInf := value.Interface()
		slice[index] = valueInf

		wg.Done()
	}

	wg.Add(length)
	for i := 0; i < length; i++ {
		go setterIndexFunc(i)
	}

	wg.Wait()

	return slice
}

func ReflectValToInterfacesUsingProcessor(
	isSkipOnNil bool,
	processorFunc func(item interface{}) (result interface{}, isTake, isBreak bool),
	reflectVal reflect.Value,
) []interface{} {
	if reflectVal.Kind() == reflect.Ptr {
		return ReflectValToInterfaces(
			isSkipOnNil,
			reflect.Indirect(reflect.ValueOf(reflectVal)),
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	if !isSliceOrArray {
		return []interface{}{}
	}

	length := reflectVal.Len()
	slice := make([]interface{}, 0, length)

	if length == 0 {
		return slice
	}

	for i := 0; i < length; i++ {
		value := reflectVal.Index(i)

		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		valueInf := value.Interface()

		if isSkipOnNil && Is.Null(valueInf) {
			continue
		}

		rs, isTake, isBreak :=
			processorFunc(valueInf)

		if isTake {
			slice = append(slice, rs)
		}

		if isBreak {
			return slice
		}
	}

	return slice
}

func ReflectInterfaceVal(any interface{}) interface{} {
	rVal := reflect.ValueOf(any)

	if rVal.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
	}

	return rVal.Interface()
}

func SafeTypeName(any interface{}) string {
	rt := reflect.TypeOf(any)

	if Is.Null(rt) {
		return ""
	}

	return rt.String()
}

func SafeTypeNameOfSliceOrSingle(
	isSingle bool,
	any interface{},
) string {
	if isSingle {
		return SafeTypeName(any)
	}

	return SafeSliceToTypeName(any)
}

// SafeSliceToTypeName
//
// Gets slice element type name, reduce ptr slice as well.
func SafeSliceToTypeName(slice interface{}) string {
	rt := reflect.TypeOf(slice)

	if Is.Null(rt) {
		return ""
	}

	if rt.Kind() == reflect.Ptr || rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	return rt.Elem().String()
}
