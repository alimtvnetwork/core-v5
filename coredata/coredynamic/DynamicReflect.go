package coredynamic

import (
	"reflect"

	"gitlab.com/auk-go/core/defaulterr"
)

// DynamicReflect.go — Reflection-based operations, loops, filters, and
// conversion methods extracted from Dynamic.go.

func (it *Dynamic) ReflectValue() *reflect.Value {
	if it.reflectVal != nil {
		return it.reflectVal
	}

	reflectValueOfAny := reflect.ValueOf(it.innerData)
	it.reflectVal = &reflectValueOfAny

	return it.reflectVal
}

func (it *Dynamic) MapToKeyVal() (*KeyValCollection, error) {
	return MapAsKeyValSlice(*it.ReflectValue())
}

func (it *Dynamic) ReflectKind() reflect.Kind {
	return it.ReflectValue().Kind()
}

func (it *Dynamic) ReflectTypeName() string {
	return it.typeName.Value()
}

func (it *Dynamic) ReflectType() reflect.Type {
	if it.reflectType != nil {
		return it.reflectType
	}

	reflectType := reflect.TypeOf(it.innerData)
	it.reflectType = reflectType

	return it.reflectType
}

func (it *Dynamic) IsReflectTypeOf(
	typeRequest reflect.Type,
) bool {
	return it.ReflectType() == typeRequest
}

func (it *Dynamic) IsReflectKind(checkingKind reflect.Kind) bool {
	return it.ReflectKind() == checkingKind
}

func (it *Dynamic) ItemReflectValueUsingIndex(index int) reflect.Value {
	return it.ReflectValue().Index(index)
}

func (it *Dynamic) ItemReflectValueUsingKey(key any) reflect.Value {
	return it.ReflectValue().MapIndex(reflect.ValueOf(key))
}

func (it *Dynamic) ItemUsingIndex(index int) any {
	return it.ReflectValue().Index(index).Interface()
}

func (it *Dynamic) ItemUsingKey(key any) any {
	return it.ReflectValue().MapIndex(reflect.ValueOf(key)).Interface()
}

func (it *Dynamic) ReflectSetTo(toPointer any) error {
	if it == nil {
		return defaulterr.NilResult
	}

	return ReflectSetFromTo(
		it.innerData,
		toPointer,
	)
}

func (it *Dynamic) ConvertUsingFunc(
	converter SimpleInOutConverter,
	expectedType reflect.Type,
) *SimpleResult {
	return converter(it.innerData, expectedType)
}

// =============================================================================
// Iteration
// =============================================================================

func (it *Dynamic) Loop(
	loopProcessorFunc func(index int, item any) (isBreak bool),
) (isCalled bool) {
	if it.IsInvalid() || it.IsNull() || it.Length() <= 0 {
		return false
	}

	length := it.Length()
	rv := *it.ReflectValue()

	for i := 0; i < length; i++ {
		isBreak := loopProcessorFunc(
			i,
			rv.Index(i).Interface(),
		)

		if isBreak {
			return true
		}
	}

	return true
}

func (it *Dynamic) FilterAsDynamicCollection(
	filterFunc func(index int, itemAsDynamic Dynamic) (isTake, isBreak bool),
) *DynamicCollection {
	if it.IsInvalid() || it.IsNull() || it.Length() <= 0 {
		return EmptyDynamicCollection()
	}

	length := it.Length()
	rv := *it.ReflectValue()
	dynamicCollection := NewDynamicCollection(length / 2)

	for i := 0; i < length; i++ {
		currentRv := rv.Index(i)
		valInf := currentRv.Interface()
		currentDynamic := NewDynamic(valInf, currentRv.IsValid())

		isTake, isBreak := filterFunc(
			i,
			currentDynamic,
		)

		if isTake {
			dynamicCollection.Add(currentDynamic)
		}

		if isBreak {
			return dynamicCollection
		}
	}

	return dynamicCollection
}

func (it *Dynamic) LoopMap(
	mapLoopProcessorFunc func(index int, key, value any) (isBreak bool),
) (isCalled bool) {
	if it.IsInvalid() || it.IsNull() || it.Length() <= 0 {
		return false
	}

	rv := *it.ReflectValue()
	mapIterator := rv.MapRange()
	index := 0

	for mapIterator.Next() {
		k := mapIterator.Key()
		v := mapIterator.Value()
		isBreak := mapLoopProcessorFunc(index, k.Interface(), v.Interface())

		if isBreak {
			return true
		}

		index++
	}

	return true
}
