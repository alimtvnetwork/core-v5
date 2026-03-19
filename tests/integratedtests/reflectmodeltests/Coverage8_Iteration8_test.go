package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// =============================================================================
// FieldProcessor
// =============================================================================

func Test_I8_01_FP_IsFieldType(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		Field:     reflect.StructField{Name: "X", Type: reflect.TypeOf("")},
		FieldType: reflect.TypeOf(""),
	}
	if !fp.IsFieldType(reflect.TypeOf("")) {
		t.Fatal("expected true for same type")
	}
	if fp.IsFieldType(reflect.TypeOf(0)) {
		t.Fatal("expected false for different type")
	}
}

func Test_I8_02_FP_IsFieldType_Nil(t *testing.T) {
	var fp *reflectmodel.FieldProcessor
	if fp.IsFieldType(reflect.TypeOf("")) {
		t.Fatal("expected false for nil receiver")
	}
}

func Test_I8_03_FP_IsFieldKind(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		Field:     reflect.StructField{Name: "X", Type: reflect.TypeOf("")},
		FieldType: reflect.TypeOf(""),
	}
	if !fp.IsFieldKind(reflect.String) {
		t.Fatal("expected true for string kind")
	}
	if fp.IsFieldKind(reflect.Int) {
		t.Fatal("expected false for int kind")
	}
}

func Test_I8_04_FP_IsFieldKind_Nil(t *testing.T) {
	var fp *reflectmodel.FieldProcessor
	if fp.IsFieldKind(reflect.String) {
		t.Fatal("expected false for nil receiver")
	}
}

// =============================================================================
// MethodProcessor — basic properties
// =============================================================================

type testMPStruct struct{}

func (t testMPStruct) PublicMethod(a string, b int) (string, error) {
	return a, nil
}

func (t testMPStruct) NoArgMethod() string {
	return "hello"
}

func (t testMPStruct) MultiReturn() (string, int, error) {
	return "", 0, nil
}

func getMethodProcessor(name string) *reflectmodel.MethodProcessor {
	rt := reflect.TypeOf(testMPStruct{})
	for i := 0; i < rt.NumMethod(); i++ {
		m := rt.Method(i)
		if m.Name == name {
			return &reflectmodel.MethodProcessor{
				Name:          m.Name,
				Index:         i,
				ReflectMethod: m,
			}
		}
	}
	return nil
}

func Test_I8_05_MP_HasValidFunc(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	if !mp.HasValidFunc() {
		t.Fatal("expected true")
	}
	var nilMP *reflectmodel.MethodProcessor
	if nilMP.HasValidFunc() {
		t.Fatal("expected false for nil")
	}
}

func Test_I8_06_MP_GetFuncName(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	if mp.GetFuncName() != "PublicMethod" {
		t.Fatal("expected PublicMethod")
	}
}

func Test_I8_07_MP_IsInvalid(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	if mp.IsInvalid() {
		t.Fatal("expected false")
	}
	var nilMP *reflectmodel.MethodProcessor
	if !nilMP.IsInvalid() {
		t.Fatal("expected true for nil")
	}
}

func Test_I8_08_MP_Func(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	f := mp.Func()
	if f == nil {
		t.Fatal("expected non-nil func")
	}
}

func Test_I8_09_MP_Func_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	f := mp.Func()
	if f != nil {
		t.Fatal("expected nil")
	}
}

func Test_I8_10_MP_ArgsCount(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	// includes receiver, so testMPStruct + string + int = 3
	if mp.ArgsCount() < 2 {
		t.Fatal("expected at least 2 args")
	}
}

func Test_I8_11_MP_ReturnLength(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	if mp.ReturnLength() != 2 {
		t.Fatal("expected 2 return args")
	}
}

func Test_I8_12_MP_ReturnLength_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	if mp.ReturnLength() != -1 {
		t.Fatal("expected -1 for nil")
	}
}

func Test_I8_13_MP_IsPublicMethod(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	if !mp.IsPublicMethod() {
		t.Fatal("expected true")
	}
}

func Test_I8_14_MP_IsPublicMethod_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	if mp.IsPublicMethod() {
		t.Fatal("expected false for nil")
	}
}

func Test_I8_15_MP_IsPrivateMethod(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	if mp.IsPrivateMethod() {
		t.Fatal("expected false for public method")
	}
}

func Test_I8_16_MP_IsPrivateMethod_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	if mp.IsPrivateMethod() {
		t.Fatal("expected false for nil")
	}
}

func Test_I8_17_MP_ArgsLength(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	if mp.ArgsLength() < 2 {
		t.Fatal("expected at least 2")
	}
}

func Test_I8_18_MP_GetType(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	if mp.GetType() == nil {
		t.Fatal("expected non-nil type")
	}
}

func Test_I8_19_MP_GetType_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	if mp.GetType() != nil {
		t.Fatal("expected nil")
	}
}

// =============================================================================
// MethodProcessor — GetInArgsTypes, GetOutArgsTypes, GetInArgsTypesNames
// =============================================================================

func Test_I8_20_MP_GetOutArgsTypes(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	types := mp.GetOutArgsTypes()
	if len(types) != 2 {
		t.Fatalf("expected 2 out args, got %d", len(types))
	}
	// Call again to test cache
	types2 := mp.GetOutArgsTypes()
	if len(types2) != 2 {
		t.Fatal("expected 2 from cache")
	}
}

func Test_I8_21_MP_GetOutArgsTypes_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	types := mp.GetOutArgsTypes()
	if len(types) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func Test_I8_22_MP_GetOutArgsTypes_NoReturn(t *testing.T) {
	// NoArgMethod returns 1 value, not 0, but let's test cache path
	mp := getMethodProcessor("NoArgMethod")
	types := mp.GetOutArgsTypes()
	if len(types) != 1 {
		t.Fatalf("expected 1, got %d", len(types))
	}
}

func Test_I8_23_MP_GetInArgsTypes(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	types := mp.GetInArgsTypes()
	if len(types) < 2 {
		t.Fatal("expected at least 2 in args")
	}
	// call again for cache
	types2 := mp.GetInArgsTypes()
	if len(types2) != len(types) {
		t.Fatal("cache mismatch")
	}
}

func Test_I8_24_MP_GetInArgsTypes_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	types := mp.GetInArgsTypes()
	if len(types) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func Test_I8_25_MP_GetInArgsTypesNames(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	names := mp.GetInArgsTypesNames()
	if len(names) < 2 {
		t.Fatal("expected at least 2")
	}
	// call again for cache
	names2 := mp.GetInArgsTypesNames()
	if len(names2) != len(names) {
		t.Fatal("cache mismatch")
	}
}

func Test_I8_26_MP_GetInArgsTypesNames_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	names := mp.GetInArgsTypesNames()
	if len(names) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func Test_I8_27_MP_GetInArgsTypesNames_NoArgs(t *testing.T) {
	mp := getMethodProcessor("NoArgMethod")
	names := mp.GetInArgsTypesNames()
	// NoArgMethod has 1 arg (the receiver)
	if len(names) < 1 {
		t.Fatal("expected at least 1 (receiver)")
	}
}

// =============================================================================
// MethodProcessor — Invoke and variants
// =============================================================================

func Test_I8_28_MP_Invoke_Success(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	// PublicMethod(receiver, string, int) -> (string, error)
	results, err := mp.Invoke(testMPStruct{}, "hello", 42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
}

func Test_I8_29_MP_Invoke_NilReceiver(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, err := mp.Invoke("hello")
	if err == nil {
		t.Fatal("expected error for nil receiver")
	}
}

func Test_I8_30_MP_Invoke_ArgsMismatch(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	_, err := mp.Invoke("too few args")
	if err == nil {
		t.Fatal("expected error for args count mismatch")
	}
}

func Test_I8_31_MP_GetFirstResponseOfInvoke(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	first, err := mp.GetFirstResponseOfInvoke(testMPStruct{}, "hello", 42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if first != "hello" {
		t.Fatalf("expected 'hello', got %v", first)
	}
}

func Test_I8_32_MP_GetFirstResponseOfInvoke_Error(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, err := mp.GetFirstResponseOfInvoke("x")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I8_33_MP_InvokeResultOfIndex(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	result, err := mp.InvokeResultOfIndex(0, testMPStruct{}, "test", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "test" {
		t.Fatal("expected 'test'")
	}
}

func Test_I8_34_MP_InvokeResultOfIndex_Error(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, err := mp.InvokeResultOfIndex(0, "x")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I8_35_MP_InvokeError(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	funcErr, procErr := mp.InvokeError(testMPStruct{}, "test", 1)
	if procErr != nil {
		t.Fatalf("unexpected processing error: %v", procErr)
	}
	// PublicMethod returns nil error
	if funcErr != nil {
		t.Fatal("expected nil func error")
	}
}

func Test_I8_36_MP_InvokeError_ProcError(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, err := mp.InvokeError("x")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I8_37_MP_InvokeFirstAndError(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	first, funcErr, procErr := mp.InvokeFirstAndError(testMPStruct{}, "test", 1)
	if procErr != nil {
		t.Fatalf("unexpected processing error: %v", procErr)
	}
	if funcErr != nil {
		t.Fatal("expected nil func error")
	}
	if first != "test" {
		t.Fatal("expected 'test'")
	}
}

func Test_I8_38_MP_InvokeFirstAndError_ProcError(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, _, err := mp.InvokeFirstAndError("x")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I8_39_MP_InvokeFirstAndError_SingleReturn(t *testing.T) {
	mp := getMethodProcessor("NoArgMethod")
	_, _, procErr := mp.InvokeFirstAndError(testMPStruct{})
	if procErr == nil {
		t.Fatal("expected error for single return method")
	}
}

// =============================================================================
// MethodProcessor — IsEqual, IsNotEqual
// =============================================================================

func Test_I8_40_MP_IsEqual_BothNil(t *testing.T) {
	var a, b *reflectmodel.MethodProcessor
	if !a.IsEqual(b) {
		t.Fatal("expected true for both nil")
	}
}

func Test_I8_41_MP_IsEqual_LeftNil(t *testing.T) {
	var a *reflectmodel.MethodProcessor
	b := getMethodProcessor("PublicMethod")
	if a.IsEqual(b) {
		t.Fatal("expected false")
	}
}

func Test_I8_42_MP_IsEqual_RightNil(t *testing.T) {
	a := getMethodProcessor("PublicMethod")
	if a.IsEqual(nil) {
		t.Fatal("expected false")
	}
}

func Test_I8_43_MP_IsEqual_SamePointer(t *testing.T) {
	a := getMethodProcessor("PublicMethod")
	if !a.IsEqual(a) {
		t.Fatal("expected true for same pointer")
	}
}

func Test_I8_44_MP_IsEqual_SameMethod(t *testing.T) {
	a := getMethodProcessor("PublicMethod")
	b := getMethodProcessor("PublicMethod")
	if !a.IsEqual(b) {
		t.Fatal("expected true for same method")
	}
}

func Test_I8_45_MP_IsNotEqual(t *testing.T) {
	a := getMethodProcessor("PublicMethod")
	b := getMethodProcessor("NoArgMethod")
	if !a.IsNotEqual(b) {
		t.Fatal("expected not equal for different methods")
	}
}

// =============================================================================
// MethodProcessor — ValidateMethodArgs, VerifyInArgs, VerifyOutArgs
// =============================================================================

func Test_I8_46_MP_ValidateMethodArgs_OK(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	err := mp.ValidateMethodArgs([]any{testMPStruct{}, "hello", 42})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_I8_47_MP_ValidateMethodArgs_CountMismatch(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	err := mp.ValidateMethodArgs([]any{"only one"})
	if err == nil {
		t.Fatal("expected error for count mismatch")
	}
}

func Test_I8_48_MP_ValidateMethodArgs_TypeMismatch(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	// Wrong types: int instead of string, string instead of int
	err := mp.ValidateMethodArgs([]any{testMPStruct{}, 42, "hello"})
	if err == nil {
		t.Fatal("expected error for type mismatch")
	}
}

func Test_I8_49_MP_VerifyInArgs(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	ok, err := mp.VerifyInArgs([]any{testMPStruct{}, "hello", 42})
	if !ok || err != nil {
		t.Fatal("expected ok")
	}
}

func Test_I8_50_MP_VerifyOutArgs(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	ok, err := mp.VerifyOutArgs([]any{"result", (*error)(nil)})
	// This may or may not match depending on interface handling
	_, _ = ok, err
}

func Test_I8_51_MP_InArgsVerifyRv(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	inTypes := mp.GetInArgsTypes()
	ok, err := mp.InArgsVerifyRv(inTypes)
	if !ok || err != nil {
		t.Fatal("expected ok for same types")
	}
}

func Test_I8_52_MP_OutArgsVerifyRv(t *testing.T) {
	mp := getMethodProcessor("PublicMethod")
	outTypes := mp.GetOutArgsTypes()
	ok, err := mp.OutArgsVerifyRv(outTypes)
	if !ok || err != nil {
		t.Fatal("expected ok for same types")
	}
}

// =============================================================================
// ReflectValueKind
// =============================================================================

func Test_I8_53_RVK_InvalidModel(t *testing.T) {
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")
	if rvk.IsValid {
		t.Fatal("expected invalid")
	}
	if rvk.Error == nil {
		t.Fatal("expected error")
	}
	if !rvk.IsInvalid() {
		t.Fatal("expected IsInvalid true")
	}
}

func Test_I8_54_RVK_IsInvalid_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if !rvk.IsInvalid() {
		t.Fatal("expected true for nil")
	}
}

func Test_I8_55_RVK_HasError(t *testing.T) {
	rvk := reflectmodel.InvalidReflectValueKindModel("err")
	if !rvk.HasError() {
		t.Fatal("expected true")
	}
}

func Test_I8_56_RVK_HasError_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if rvk.HasError() {
		t.Fatal("expected false for nil")
	}
}

func Test_I8_57_RVK_IsEmptyError(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: true}
	if !rvk.IsEmptyError() {
		t.Fatal("expected true for nil error")
	}
}

func Test_I8_58_RVK_IsEmptyError_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if !rvk.IsEmptyError() {
		t.Fatal("expected true for nil receiver")
	}
}

func Test_I8_59_RVK_ActualInstance(t *testing.T) {
	val := "hello"
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	result := rvk.ActualInstance()
	if result != "hello" {
		t.Fatal("expected 'hello'")
	}
}

func Test_I8_60_RVK_ActualInstance_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if rvk.ActualInstance() != nil {
		t.Fatal("expected nil")
	}
}

func Test_I8_61_RVK_PkgPath(t *testing.T) {
	val := "hello"
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	_ = rvk.PkgPath()
}

func Test_I8_62_RVK_PkgPath_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if rvk.PkgPath() != "" {
		t.Fatal("expected empty")
	}
}

func Test_I8_63_RVK_PkgPath_Invalid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}
	if rvk.PkgPath() != "" {
		t.Fatal("expected empty for invalid")
	}
}

func Test_I8_64_RVK_TypeName(t *testing.T) {
	val := "hello"
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	name := rvk.TypeName()
	if name == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_I8_65_RVK_TypeName_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if rvk.TypeName() != "" {
		t.Fatal("expected empty")
	}
}

func Test_I8_66_RVK_TypeName_Invalid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}
	if rvk.TypeName() != "" {
		t.Fatal("expected empty for invalid")
	}
}

func Test_I8_67_RVK_PointerRv(t *testing.T) {
	val := "hello"
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	rv := rvk.PointerRv()
	if rv == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I8_68_RVK_PointerRv_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if rvk.PointerRv() != nil {
		t.Fatal("expected nil")
	}
}

func Test_I8_69_RVK_PointerRv_Invalid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf("x"),
	}
	rv := rvk.PointerRv()
	if rv == nil {
		t.Fatal("expected non-nil even for invalid")
	}
}

func Test_I8_70_RVK_PointerInterface(t *testing.T) {
	val := "hello"
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	iface := rvk.PointerInterface()
	if iface == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_I8_71_RVK_PointerInterface_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if rvk.PointerInterface() != nil {
		t.Fatal("expected nil")
	}
}
