package reflectmodel

import (
	"reflect"
)

type MethodProcessor struct {
	Name          string
	Index         int
	ReflectMethod reflect.Method
}

// ArgsCount is same as ArgsLength
//
// Reference:
//
//	https://stackoverflow.com/a/47626214
func (it *MethodProcessor) ArgsCount() int {
	return it.ReflectMethod.Type.NumIn()
}

func (it *MethodProcessor) IsPublicMethod() bool {
	return it != nil && it.ReflectMethod.PkgPath == ""
}

func (it *MethodProcessor) IsPrivateMethod() bool {
	return it != nil && it.ReflectMethod.PkgPath != ""
}

// ArgsLength
//
// https://stackoverflow.com/a/47626214
// It is same as ArgsCount
func (it *MethodProcessor) ArgsLength() int {
	return it.ReflectMethod.Type.NumIn()
}

func (it *MethodProcessor) Invoke(
	args ...interface{},
) (
	responses []interface{},
) {
	rvs := it.ReflectMethod.Func.Call(argsReflectValues(args))

	return reflectValuesToInterfaces(rvs)
}

func (it *MethodProcessor) GetFirstResponseOfInvoke(
	args ...interface{},
) (
	firstResponse interface{},
) {
	return it.GetResponseOfIndexFromInvoke(0, args...)
}

func (it *MethodProcessor) GetResponseOfIndexFromInvoke(
	index int,
	args ...interface{},
) (
	firstResponse interface{},
) {
	rvs := it.ReflectMethod.Func.Call(argsReflectValues(args))

	return reflectValueToAnyValue(rvs[index])
}

func (it *MethodProcessor) InvokeError(args ...interface{}) (err error) {
	return it.GetFirstResponseOfInvoke(args...).(error)
}

func (it *MethodProcessor) InvokeVoid(args ...interface{}) {
	it.ReflectMethod.Func.Call(argsReflectValues(args))
}
