package reflectinternal

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"github.com/alimtvnetwork/core/internal/convertinternal"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

type reflectConverter struct{}

func (it reflectConverter) ArgsToReflectValues(args []any) []reflect.Value {
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
) []any {
	if len(reflectValues) == 0 {
		return []any{}
	}

	list := make(
		[]any,
		len(reflectValues),
	)

	for i, rv := range reflectValues {
		list[i] = it.ReflectValueToAnyValue(rv)
	}

	return list
}

func (it reflectConverter) ReflectValueToAnyValue(rv reflect.Value) any {
	if Is.Null(rv) {
		return nil
	}

	k := rv.Kind()

	switch k {
	case reflect.Ptr, reflect.Interface:
		return rv.Elem().Interface()
	case reflect.String:
		return rv.String()
	case reflect.Int:
		return rv.Int()
	default:
		return rv.Interface()
	}
}

func (it reflectConverter) InterfacesToTypes(items []any) []reflect.Type {
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

func (it reflectConverter) InterfacesToTypesNames(items []any) []string {
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

func (it reflectConverter) InterfacesToTypesNamesWithValues(items []any) []string {
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
			convertinternal.AnyTo.SmartString(item),
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

func (it reflectConverter) ToPtrRvIfNotAlready(
	rv reflect.Value,
) reflect.Value {
	if rv.Kind() == reflect.Ptr {
		return rv
	}

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
	anyItem any,
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
	anyItem any,
) *reflectmodel.ReflectValueKind {
	return it.ReducePointer(anyItem, 3)
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
	anyItem any,
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

// ReflectValToInterfaces
//
// Assuming passing reflect val is an array or slice
// loop using reflection and returns the interfaces slice
func (it reflectConverter) ReflectValToInterfaces(
	isSkipOnNil bool,
	reflectVal reflect.Value,
) []any {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ReflectValToInterfaces(
			isSkipOnNil,
			reflect.Indirect(reflectVal),
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	isNotSliceOrArray := !isSliceOrArray

	if isNotSliceOrArray {
		return []any{}
	}

	length := reflectVal.Len()
	slice := make([]any, 0, length)

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

func (it reflectConverter) ReflectValToInterfacesAsync(
	reflectVal reflect.Value,
) []any {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ReflectValToInterfacesAsync(
			reflect.Indirect(reflectVal),
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	isNotSliceOrArray := !isSliceOrArray

	if isNotSliceOrArray {
		return []any{}
	}

	length := reflectVal.Len()
	slice := make([]any, length)

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

func (it reflectConverter) ReflectValToInterfacesUsingProcessor(
	isSkipOnNil bool,
	processorFunc func(item any) (result any, isTake, isBreak bool),
	reflectVal reflect.Value,
) []any {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ReflectValToInterfaces(
			isSkipOnNil,
			reflect.Indirect(reflectVal),
		)
	}

	k := reflectVal.Kind()
	isSliceOrArray := k == reflect.Slice ||
		k == reflect.Array

	isNotSliceOrArray := !isSliceOrArray

	if isNotSliceOrArray {
		return []any{}
	}

	length := reflectVal.Len()
	slice := make([]any, 0, length)

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

func (it reflectConverter) ReflectInterfaceVal(anyItem any) any {
	rVal := reflect.ValueOf(anyItem)

	if rVal.Kind() == reflect.Ptr {
		rVal = rVal.Elem()
	}

	return rVal.Interface()
}

func (it reflectConverter) ToPointerRv(
	anyItem any,
) *reflect.Value {
	if anyItem == nil {
		return nil
	}

	rv := reflect.ValueOf(anyItem)
	newRv := it.ReflectValueToPointerReflectValue(rv)

	return &newRv
}

func (it reflectConverter) ToPointer(
	anyItem any,
) any {
	if anyItem == nil {
		return anyItem
	}

	rv := reflect.ValueOf(anyItem)
	newRv := it.ReflectValueToPointerReflectValue(rv)

	return newRv.Interface()
}
