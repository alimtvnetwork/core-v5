package reflectmodeltests

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ══════════════════════════════════════════════════════════════════════════════
// rvUtils.IsNull — UnsafePointer kind via internal test exercises
// These test ReflectValueKind with various pointer/interface kinds
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_RVK_UnsafePointerKind(t *testing.T) {
	x := 42
	up := unsafe.Pointer(&x)
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(up),
		Kind:            reflect.UnsafePointer,
	}
	actual := args.Map{
		"isInvalid": rvk.IsInvalid(),
		"actNotNil": rvk.ActualInstance() != nil,
		"typeName":  rvk.TypeName() != "",
	}
	expected := args.Map{"isInvalid": false, "actNotNil": true, "typeName": true}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- UnsafePointer kind", actual)
}

func Test_I15_RVK_FuncKind(t *testing.T) {
	fn := func() string { return "hello" }
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(fn),
		Kind:            reflect.Func,
	}
	actual := args.Map{
		"isInvalid": rvk.IsInvalid(),
		"actNotNil": rvk.ActualInstance() != nil,
		"typeName":  rvk.TypeName() != "",
	}
	expected := args.Map{"isInvalid": false, "actNotNil": true, "typeName": true}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Func kind", actual)
}

func Test_I15_RVK_ChanKind(t *testing.T) {
	ch := make(chan string, 1)
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(ch),
		Kind:            reflect.Chan,
	}
	actual := args.Map{
		"isInvalid": rvk.IsInvalid(),
		"actNotNil": rvk.ActualInstance() != nil,
	}
	expected := args.Map{"isInvalid": false, "actNotNil": true}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Chan kind", actual)
}

func Test_I15_RVK_Float64Kind(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(3.14),
		Kind:            reflect.Float64,
	}
	actual := args.Map{
		"actInst":  rvk.ActualInstance(),
		"typeName": rvk.TypeName() != "",
		"pkgPath":  rvk.PkgPath(),
	}
	expected := args.Map{"actInst": 3.14, "typeName": true, "pkgPath": ""}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- Float64 kind", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueKind — PointerRv and PointerInterface with valid struct
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_RVK_Valid_PointerRv_Struct(t *testing.T) {
	type testStruct struct{ X int }
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(testStruct{X: 99}),
		Kind:            reflect.Struct,
	}
	ptr := rvk.PointerRv()
	ptrIface := rvk.PointerInterface()
	actual := args.Map{"ptrNotNil": ptr != nil, "ptrIfaceNotNil": ptrIface != nil}
	expected := args.Map{"ptrNotNil": true, "ptrIfaceNotNil": true}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- PointerRv struct", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueKind — HasError with actual error set
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_RVK_HasError_True(t *testing.T) {
	rvk := reflectmodel.InvalidReflectValueKindModel("some error message")
	actual := args.Map{
		"hasErr":   rvk.HasError(),
		"emptyErr": rvk.IsEmptyError(),
		"errMsg":   rvk.Error.Error(),
	}
	expected := args.Map{"hasErr": true, "emptyErr": false, "errMsg": "some error message"}
	expected.ShouldBeEqual(t, 0, "RVK returns error -- HasError with message", actual)
}

func Test_I15_RVK_HasError_False(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(42),
		Kind:            reflect.Int,
		Error:           nil,
	}
	actual := args.Map{"hasErr": rvk.HasError(), "emptyErr": rvk.IsEmptyError()}
	expected := args.Map{"hasErr": false, "emptyErr": true}
	expected.ShouldBeEqual(t, 0, "RVK returns correct value -- no error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — InvokeFirstAndError with nil error interface (panic recovery)
// ══════════════════════════════════════════════════════════════════════════════

type helperI15 struct{}

func (h helperI15) ReturnTwoNoError() (string, error) {
	return "ok", nil
}

func (h helperI15) ReturnStringOnly() string {
	return "hello"
}

func (h helperI15) ReturnBoolInt(x int) (bool, int) {
	return x > 0, x * 2
}

func (h helperI15) ReturnThree() (int, string, bool) {
	return 1, "two", true
}

func getI15MP(name string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(helperI15{})
	m, ok := t.MethodByName(name)
	if !ok {
		return nil
	}
	return &reflectmodel.MethodProcessor{
		Name:          m.Name,
		Index:         m.Index,
		ReflectMethod: m,
	}
}

func Test_I15_InvokeFirstAndError_NilErrorInterface(t *testing.T) {
	mp := getI15MP("ReturnTwoNoError")
	first, funcErr, procErr := mp.InvokeFirstAndError(helperI15{})
	actual := args.Map{"procErr": procErr == nil, "funcErr": funcErr == nil, "first": first}
	expected := args.Map{"procErr": true, "funcErr": true, "first": "ok"}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns correct value -- nil error interface", actual)
}

func Test_I15_InvokeFirstAndError_NonErrorSecondReturn_Panics(t *testing.T) {
	mp := getI15MP("ReturnBoolInt")
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "InvokeFirstAndError panics -- non-error second return", actual)
	}()
	mp.InvokeFirstAndError(helperI15{}, 5)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — Invoke success with zero-return method
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_Invoke_ZeroReturn(t *testing.T) {
	mp := newMethodProcessor("NoArgsMethod")
	results, err := mp.Invoke(sampleStruct{})
	actual := args.Map{"noErr": err == nil, "len": len(results)}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- single return", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — VerifyInArgs with wrong arg count
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_VerifyInArgs_WrongCount(t *testing.T) {
	mp := getI15MP("ReturnBoolInt")
	ok, err := mp.VerifyInArgs([]any{helperI15{}}) // needs receiver + int
	actual := args.Map{"ok": ok, "hasErr": err != nil}
	expected := args.Map{"ok": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyInArgs returns error -- wrong count", actual)
}

func Test_I15_VerifyOutArgs_WrongCount(t *testing.T) {
	mp := getI15MP("ReturnBoolInt")
	ok, err := mp.VerifyOutArgs([]any{true}) // needs 2 out
	actual := args.Map{"ok": ok, "hasErr": err != nil}
	expected := args.Map{"ok": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyOutArgs returns error -- wrong count", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — argsCountMismatchErrorMessage exercises with 0 args
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_ValidateMethodArgs_ZeroGiven(t *testing.T) {
	mp := getI15MP("ReturnBoolInt")
	err := mp.ValidateMethodArgs([]any{})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns error -- zero args given", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// FieldProcessor — with various types
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_FieldProcessor_StringField(t *testing.T) {
	fp := newFieldProcessor("Name", 0)
	actual := args.Map{
		"isStr":  fp.IsFieldKind(reflect.String),
		"isInt":  fp.IsFieldKind(reflect.Int),
		"typeOk": fp.IsFieldType(reflect.TypeOf("")),
	}
	expected := args.Map{"isStr": true, "isInt": false, "typeOk": true}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns correct value -- String field", actual)
}

func Test_I15_FieldProcessor_NilReceiver_IsFieldKind(t *testing.T) {
	var fp *reflectmodel.FieldProcessor
	actual := args.Map{"kind": fp.IsFieldKind(reflect.Int)}
	expected := args.Map{"kind": false}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns false -- nil IsFieldKind", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValue — nil RawData
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_ReflectValue_NilRawData(t *testing.T) {
	rv := reflectmodel.ReflectValue{
		TypeName:     "NilType",
		FieldsNames:  nil,
		MethodsNames: nil,
		RawData:      nil,
	}
	actual := args.Map{
		"typeName":   rv.TypeName,
		"fieldsNil":  rv.FieldsNames == nil,
		"methodsNil": rv.MethodsNames == nil,
		"rawNil":     rv.RawData == nil,
	}
	expected := args.Map{"typeName": "NilType", "fieldsNil": true, "methodsNil": true, "rawNil": true}
	expected.ShouldBeEqual(t, 0, "ReflectValue returns nil -- nil raw data", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — GetOutArgsTypes with 3 returns
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_GetOutArgsTypes_ThreeReturns(t *testing.T) {
	mp := getI15MP("ReturnThree")
	out := mp.GetOutArgsTypes()
	out2 := mp.GetOutArgsTypes() // cached
	actual := args.Map{"len": len(out), "cached": len(out2)}
	expected := args.Map{"len": 3, "cached": 3}
	expected.ShouldBeEqual(t, 0, "GetOutArgsTypes returns correct value -- three returns", actual)
}

func Test_I15_GetInArgsTypesNames_ThreeReturns(t *testing.T) {
	mp := getI15MP("ReturnThree")
	names := mp.GetInArgsTypesNames()
	names2 := mp.GetInArgsTypesNames() // cached
	actual := args.Map{"len": len(names), "cached": len(names2)}
	expected := args.Map{"len": 1, "cached": 1}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypesNames returns correct value -- receiver only", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — IsEqual with different out args
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_IsEqual_DiffOutArgs(t *testing.T) {
	mp1 := getI15MP("ReturnStringOnly")
	mp2 := getI15MP("ReturnThree")
	actual := args.Map{"eq": mp1.IsEqual(mp2)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns false -- diff out args", actual)
}

func Test_I15_IsEqual_SameSig(t *testing.T) {
	mp1 := getI15MP("ReturnStringOnly")
	mp2 := getI15MP("ReturnStringOnly")
	actual := args.Map{"eq": mp1.IsEqual(mp2)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- same sig", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — InvokeError with nil processing error
// ══════════════════════════════════════════════════════════════════════════════

func Test_I15_InvokeError_NilFunc(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, procErr := mp.InvokeError()
	actual := args.Map{"hasErr": procErr != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeError returns error -- nil receiver", actual)
}

func Test_I15_GetFirstResponseOfInvoke_NilReceiver(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, err := mp.GetFirstResponseOfInvoke()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetFirstResponseOfInvoke returns error -- nil receiver", actual)
}

func Test_I15_InvokeResultOfIndex_NilReceiver(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, err := mp.InvokeResultOfIndex(0)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex returns error -- nil receiver", actual)
}

func Test_I15_InvokeFirstAndError_NilReceiver(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, _, procErr := mp.InvokeFirstAndError()
	actual := args.Map{"hasErr": procErr != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- nil receiver", actual)
}
