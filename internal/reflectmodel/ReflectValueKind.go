package reflectmodel

import (
	"errors"
	"reflect"
)

type ReflectValueKind struct {
	IsValid         bool
	FinalReflectVal reflect.Value
	Kind            reflect.Kind
	Error           error
}

func InvalidReflectValueKindModel(err string) *ReflectValueKind {
	return &ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(nil),
		Kind:            0,
		Error:           errors.New(err),
	}
}

func (it *ReflectValueKind) IsInvalid() bool {
	return it == nil || !it.IsValid
}

func (it *ReflectValueKind) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *ReflectValueKind) IsEmptyError() bool {
	return it == nil || it.Error == nil
}

func (it *ReflectValueKind) ActualInstance() interface{} {
	if it == nil {
		return nil
	}

	return it.FinalReflectVal.Interface()
}
