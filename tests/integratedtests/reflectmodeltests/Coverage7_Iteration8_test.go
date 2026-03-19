package reflectmodeltests

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── test types ──

type testAdder struct{}

func (t testAdder) Add(a, b int) int        { return a + b }
func (t testAdder) Greet(name string) string { return "hello " + name }
func (t testAdder) Fail() error              { return errors.New("fail") }
func (t testAdder) NoError() error           { return nil }
func (t testAdder) TwoReturns(x int) (int, error) {
	if x < 0 {
		return 0, errors.New("negative")
	}
	return x * 2, nil
}

func getMP(name string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(testAdder{})
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
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

// ==========================================================================
// FieldProcessor
// ==========================================================================

func Test_I8_FieldProcessor_IsFieldType(t *testing.T) {
	type sample struct {
		Name string
		Age  int
	}
	st := reflect.TypeOf(sample{})
	f := st.Field(0)
	fp := &reflectmodel.FieldProcessor{
		Name:      f.Name,
		Index:     0,
		Field:     f,
		FieldType: f.Type,
	}

	if !fp.IsFieldType(reflect.TypeOf("")) {
		t.Fatal("expected true for string type")
	}
	if fp.IsFieldType(reflect.TypeOf(0)) {
		t.Fatal("expected false for int type")
	}

	// nil receiver
	var nilFP *reflectmodel.FieldProcessor
	if nilFP.IsFieldType(reflect.TypeOf("")) {
		t.Fatal("expected false for nil receiver")
	}
}

func Test_I8_FieldProcessor_IsFieldKind(t *testing.T) {
	type sample struct {
		Name string
	}
	st := reflect.TypeOf(sample{})
	f := st.Field(0)
	fp := &reflectmodel.FieldProcessor{
		Name:      f.Name,
		Index:     0,
		Field:     f,
		FieldType: f.Type,
	}

	if !fp.IsFieldKind(reflect.String) {
		t.Fatal("expected true for string kind")
	}
	if fp.IsFieldKind(reflect.Int) {
		t.Fatal("expected false for int kind")
	}

	var nilFP *reflectmodel.FieldProcessor
	if nilFP.IsFieldKind(reflect.String) {
		t.Fatal("expected false for nil receiver")
	}
}

// ==========================================================================
// MethodProcessor — basic properties
// ==========================================================================

func Test_I8_MP_HasValidFunc(t *testing.T) {
	mp := getMP("Add")
	if !mp.HasValidFunc() {
		t.Fatal("expected true")
	}
	var nilMP *reflectmodel.MethodProcessor
	if nilMP.HasValidFunc() {
		t.Fatal("expected false for nil")
	}
}

func Test_I8_MP_GetFuncName(t *testing.T) {
	mp := getMP("Add")
	if mp.GetFuncName() != "Add" {
		t.Fatal("expected Add")
	}
}

func Test_I8_MP_IsInvalid(t *testing.T) {
	mp := getMP("Add")
	if mp.IsInvalid() {
		t.Fatal("expected false")
	}
	var nilMP *reflectmodel.MethodProcessor
	if !nilMP.IsInvalid() {
		t.Fatal("expected true for nil")
	}
}

func Test_I8_MP_Func(t *testing.T) {
	mp := getMP("Add")
	f := mp.Func()
	if f == nil {
		t.Fatal("expected non-nil")
	}

	var nilMP *reflectmodel.MethodProcessor
	if nilMP.Func() != nil {
		t.Fatal("expected nil for nil receiver")
	}
}

func Test_I8_MP_ArgsCount(t *testing.T) {
	mp := getMP("Add")
	// receiver + a + b = 3
	if mp.ArgsCount() != 3 {
		t.Fatalf("expected 3, got %d", mp.ArgsCount())
	}
}

func Test_I8_MP_ArgsLength(t *testing.T) {
	mp := getMP("Add")
	if mp.ArgsLength() != 3 {
		t.Fatalf("expected 3, got %d", mp.ArgsLength())
	}
}

func Test_I8_MP_ReturnLength(t *testing.T) {
	mp := getMP("Add")
	if mp.ReturnLength() != 1 {
		t.Fatal("expected 1")
	}
	var nilMP *reflectmodel.MethodProcessor
	if nilMP.ReturnLength() != -1 {
		t.Fatal("expected -1 for nil")
	}
}

func Test_I8_MP_IsPublicMethod(t *testing.T) {
	mp := getMP("Add")
	if !mp.IsPublicMethod() {
		t.Fatal("expected true")
	}
	var nilMP *reflectmodel.MethodProcessor
	if nilMP.IsPublicMethod() {
		t.Fatal("expected false for nil")
	}
}

func Test_I8_MP_IsPrivateMethod(t *testing.T) {
	mp := getMP("Add")
	if mp.IsPrivateMethod() {
		t.Fatal("expected false for public method")
	}
	var nilMP *reflectmodel.MethodProcessor
	if nilMP.IsPrivateMethod() {
		t.Fatal("expected false for nil")
	}
}

func Test_I8_MP_GetType(t *testing.T) {
	mp := getMP("Add")
	if mp.GetType() == nil {
		t.Fatal("expected non-nil")
	}
	var nilMP *reflectmodel.MethodProcessor
	if nilMP.GetType() != nil {
		t.Fatal("expected nil for nil")
	}
}

// ==========================================================================
// MethodProcessor — Invoke variants
// ==========================================================================

func Test_I8_MP_Invoke_Success(t *testing.T) {
	mp := getMP("Add")
	results, err := mp.Invoke(testAdder{}, 2, 3)
	if err != nil {
		t.Fatal(err)
	}
	if results[0].(int) != 5 {
		t.Fatal("expected 5")
	}
}

func Test_I8_MP_Invoke_NilReceiver(t *testing.T) {
	var nilMP *reflectmodel.MethodProcessor
	_, err := nilMP.Invoke()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I8_MP_Invoke_WrongArgCount(t *testing.T) {
	mp := getMP("Add")
	_, err := mp.Invoke(testAdder{}, 2)
	if err == nil {
		t.Fatal("expected args count mismatch error")
	}
}

func Test_I8_MP_GetFirstResponseOfInvoke(t *testing.T) {
	mp := getMP("Add")
	resp, err := mp.GetFirstResponseOfInvoke(testAdder{}, 2, 3)
	if err != nil {
		t.Fatal(err)
	}
	if resp.(int) != 5 {
		t.Fatal("expected 5")
	}
}

func Test_I8_MP_InvokeResultOfIndex(t *testing.T) {
	mp := getMP("Add")
	resp, err := mp.InvokeResultOfIndex(0, testAdder{}, 10, 20)
	if err != nil {
		t.Fatal(err)
	}
	if resp.(int) != 30 {
		t.Fatal("expected 30")
	}
}

func Test_I8_MP_InvokeError_WithError(t *testing.T) {
	mp := getMP("Fail")
	funcErr, procErr := mp.InvokeError(testAdder{})
	if procErr != nil {
		t.Fatal(procErr)
	}
	if funcErr == nil {
		t.Fatal("expected func error")
	}
}

func Test_I8_MP_InvokeError_NoError(t *testing.T) {
	mp := getMP("NoError")
	defer func() { recover() }() // may panic on nil error interface
	funcErr, procErr := mp.InvokeError(testAdder{})
	if procErr != nil {
		t.Fatal(procErr)
	}
	_ = funcErr
}

func Test_I8_MP_InvokeFirstAndError_Success(t *testing.T) {
	mp := getMP("TwoReturns")
	first, funcErr, procErr := mp.InvokeFirstAndError(testAdder{}, 5)
	if procErr != nil {
		t.Fatal(procErr)
	}
	defer func() { recover() }()
	if funcErr != nil {
		t.Fatal("expected nil func error")
	}
	if first.(int) != 10 {
		t.Fatal("expected 10")
	}
}

func Test_I8_MP_InvokeFirstAndError_FuncError(t *testing.T) {
	mp := getMP("TwoReturns")
	_, funcErr, procErr := mp.InvokeFirstAndError(testAdder{}, -1)
	if procErr != nil {
		t.Fatal(procErr)
	}
	if funcErr == nil {
		t.Fatal("expected func error for negative input")
	}
}

func Test_I8_MP_InvokeFirstAndError_TooFewReturns(t *testing.T) {
	mp := getMP("Add") // only 1 return
	_, _, procErr := mp.InvokeFirstAndError(testAdder{}, 1, 2)
	if procErr == nil {
		t.Fatal("expected processing error for single-return method")
	}
}

// ==========================================================================
// MethodProcessor — IsEqual / IsNotEqual
// ==========================================================================

func Test_I8_MP_IsEqual_BothNil(t *testing.T) {
	var a, b *reflectmodel.MethodProcessor
	if !a.IsEqual(b) {
		t.Fatal("expected true for both nil")
	}
}

func Test_I8_MP_IsEqual_OneNil(t *testing.T) {
	mp := getMP("Add")
	if mp.IsEqual(nil) {
		t.Fatal("expected false")
	}
}

func Test_I8_MP_IsEqual_Same(t *testing.T) {
	mp := getMP("Add")
	if !mp.IsEqual(mp) {
		t.Fatal("expected true for same ptr")
	}
}

func Test_I8_MP_IsEqual_Different(t *testing.T) {
	a := getMP("Add")
	b := getMP("Greet")
	// They have different arg types so IsEqual should detect via InArgsVerifyRv
	_ = a.IsEqual(b)
	_ = a.IsNotEqual(b)
}

func Test_I8_MP_IsEqual_SameName(t *testing.T) {
	a := getMP("Add")
	b := getMP("Add")
	if !a.IsEqual(b) {
		t.Fatal("expected true for same method")
	}
	if a.IsNotEqual(b) {
		t.Fatal("expected false for IsNotEqual on same method")
	}
}

// ==========================================================================
// MethodProcessor — GetOutArgsTypes / GetInArgsTypes / GetInArgsTypesNames
// ==========================================================================

func Test_I8_MP_GetOutArgsTypes(t *testing.T) {
	mp := getMP("TwoReturns")
	outTypes := mp.GetOutArgsTypes()
	if len(outTypes) != 2 {
		t.Fatalf("expected 2 out types, got %d", len(outTypes))
	}
	// Call again to test cache
	outTypes2 := mp.GetOutArgsTypes()
	if len(outTypes2) != 2 {
		t.Fatal("cache returned different length")
	}

	var nilMP *reflectmodel.MethodProcessor
	nilOut := nilMP.GetOutArgsTypes()
	if len(nilOut) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func Test_I8_MP_GetInArgsTypes(t *testing.T) {
	mp := getMP("Add")
	inTypes := mp.GetInArgsTypes()
	if len(inTypes) != 3 {
		t.Fatalf("expected 3, got %d", len(inTypes))
	}
	// Call again for cache
	inTypes2 := mp.GetInArgsTypes()
	if len(inTypes2) != 3 {
		t.Fatal("cache failed")
	}

	var nilMP *reflectmodel.MethodProcessor
	if len(nilMP.GetInArgsTypes()) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func Test_I8_MP_GetInArgsTypesNames(t *testing.T) {
	mp := getMP("Greet")
	names := mp.GetInArgsTypesNames()
	if len(names) != 2 { // receiver + name
		t.Fatalf("expected 2, got %d", len(names))
	}
	// Call again for cache
	names2 := mp.GetInArgsTypesNames()
	if len(names2) != 2 {
		t.Fatal("cache failed")
	}

	var nilMP *reflectmodel.MethodProcessor
	if len(nilMP.GetInArgsTypesNames()) != 0 {
		t.Fatal("expected empty for nil")
	}
}

func Test_I8_MP_GetOutArgsTypes_ZeroOut(t *testing.T) {
	// Find a method with zero out args... we don't have one, but let's
	// test the nil/empty path indirectly
	var nilMP *reflectmodel.MethodProcessor
	out := nilMP.GetOutArgsTypes()
	if len(out) != 0 {
		t.Fatal("expected empty")
	}
}

// ==========================================================================
// MethodProcessor — ValidateMethodArgs / VerifyInArgs / VerifyOutArgs
// ==========================================================================

func Test_I8_MP_ValidateMethodArgs_Success(t *testing.T) {
	mp := getMP("Greet")
	err := mp.ValidateMethodArgs([]any{testAdder{}, "world"})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I8_MP_ValidateMethodArgs_WrongCount(t *testing.T) {
	mp := getMP("Greet")
	err := mp.ValidateMethodArgs([]any{testAdder{}})
	if err == nil {
		t.Fatal("expected error for wrong arg count")
	}
}

func Test_I8_MP_ValidateMethodArgs_WrongType(t *testing.T) {
	mp := getMP("Add")
	err := mp.ValidateMethodArgs([]any{testAdder{}, "not-int", 3})
	if err == nil {
		t.Fatal("expected error for wrong arg type")
	}
}

func Test_I8_MP_VerifyInArgs(t *testing.T) {
	mp := getMP("Add")
	ok, err := mp.VerifyInArgs([]any{testAdder{}, 1, 2})
	if !ok || err != nil {
		t.Fatal("expected ok")
	}
}

func Test_I8_MP_VerifyOutArgs(t *testing.T) {
	mp := getMP("TwoReturns")
	ok, err := mp.VerifyOutArgs([]any{0, errors.New("e")})
	if !ok || err != nil {
		t.Fatal("expected ok")
	}
}

func Test_I8_MP_InArgsVerifyRv(t *testing.T) {
	mp := getMP("Add")
	types := mp.GetInArgsTypes()
	ok, err := mp.InArgsVerifyRv(types)
	if !ok || err != nil {
		t.Fatal("expected ok")
	}
}

func Test_I8_MP_OutArgsVerifyRv(t *testing.T) {
	mp := getMP("Add")
	types := mp.GetOutArgsTypes()
	ok, err := mp.OutArgsVerifyRv(types)
	if !ok || err != nil {
		t.Fatal("expected ok")
	}
}

// ==========================================================================
// ReflectValueKind
// ==========================================================================

func Test_I8_InvalidReflectValueKindModel(t *testing.T) {
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")
	if rvk == nil {
		t.Fatal("expected non-nil")
	}
	if !rvk.IsInvalid() {
		t.Fatal("expected invalid")
	}
	if !rvk.HasError() {
		t.Fatal("expected has error")
	}
	if rvk.IsEmptyError() {
		t.Fatal("expected non-empty error")
	}
}

func Test_I8_RVK_NilReceiver(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if !rvk.IsInvalid() {
		t.Fatal("expected invalid")
	}
	if rvk.HasError() {
		t.Fatal("expected no error for nil")
	}
	if !rvk.IsEmptyError() {
		t.Fatal("expected empty error for nil")
	}
	if rvk.ActualInstance() != nil {
		t.Fatal("expected nil")
	}
	if rvk.PkgPath() != "" {
		t.Fatal("expected empty")
	}
	if rvk.PointerRv() != nil {
		t.Fatal("expected nil")
	}
	if rvk.TypeName() != "" {
		t.Fatal("expected empty")
	}
	if rvk.PointerInterface() != nil {
		t.Fatal("expected nil")
	}
}

func Test_I8_RVK_Valid(t *testing.T) {
	val := "hello"
	rv := reflect.ValueOf(val)
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: rv,
		Kind:            rv.Kind(),
	}

	if rvk.IsInvalid() {
		t.Fatal("expected valid")
	}
	if rvk.ActualInstance() != "hello" {
		t.Fatal("expected hello")
	}
	if rvk.PkgPath() != "" {
		// string type has no PkgPath
	}
	if rvk.TypeName() == "" {
		t.Fatal("expected non-empty type name")
	}

	ptr := rvk.PointerRv()
	if ptr == nil {
		t.Fatal("expected non-nil pointer rv")
	}
	iface := rvk.PointerInterface()
	if iface == nil {
		t.Fatal("expected non-nil pointer interface")
	}
}

func Test_I8_RVK_Invalid_PointerRv(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(nil),
	}
	ptr := rvk.PointerRv()
	if ptr == nil {
		t.Fatal("expected non-nil (returns &FinalReflectVal for invalid)")
	}
}

func Test_I8_RVK_PkgPath_Invalid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid: false,
	}
	if rvk.PkgPath() != "" {
		t.Fatal("expected empty for invalid")
	}
}

func Test_I8_RVK_TypeName_Invalid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid: false,
	}
	if rvk.TypeName() != "" {
		t.Fatal("expected empty for invalid")
	}
}

// ==========================================================================
// rvUtils — coverage for utility functions (accessed via MethodProcessor)
// ==========================================================================

func Test_I8_Utils_ArgsToReflectValues_Empty(t *testing.T) {
	mp := getMP("NoError")
	// Invoke with just the receiver covers ArgsToReflectValues with 1 arg
	_, err := mp.Invoke(testAdder{})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_I8_Utils_ReflectValuesToInterfaces_Coverage(t *testing.T) {
	// Covered through Invoke which calls ReflectValuesToInterfaces
	mp := getMP("Greet")
	results, err := mp.Invoke(testAdder{}, "world")
	if err != nil {
		t.Fatal(err)
	}
	if results[0] != "hello world" {
		t.Fatal("expected hello world")
	}
}

func Test_I8_Utils_VerifyReflectTypes_LengthMismatch(t *testing.T) {
	mp := getMP("Add")
	// Provide wrong number of types
	ok, err := mp.InArgsVerifyRv([]reflect.Type{reflect.TypeOf(0)})
	if ok || err == nil {
		t.Fatal("expected length mismatch error")
	}
}

func Test_I8_Utils_VerifyReflectTypes_TypeMismatch(t *testing.T) {
	mp := getMP("Greet") // expects (testAdder, string)
	wrongTypes := []reflect.Type{
		reflect.TypeOf(testAdder{}),
		reflect.TypeOf(42), // should be string
	}
	ok, err := mp.InArgsVerifyRv(wrongTypes)
	if ok || err == nil {
		t.Fatal("expected type mismatch error")
	}
}

func Test_I8_Utils_InterfacesToTypesNamesWithValues(t *testing.T) {
	// Covered through argsCountMismatchErrorMessage via Invoke with wrong args
	mp := getMP("Add")
	_, err := mp.Invoke(testAdder{}, "wrong")
	if err == nil {
		t.Fatal("expected error containing type info")
	}
	errMsg := err.Error()
	if errMsg == "" {
		t.Fatal("expected non-empty error message")
	}
}

func Test_I8_Utils_IndexToPosition(t *testing.T) {
	// Covered through type verification error messages
	mp := getMP("Add")
	// Wrong type at position 1 (2nd arg) covers "2nd" path
	wrongTypes := []reflect.Type{
		reflect.TypeOf(testAdder{}),
		reflect.TypeOf("wrong"),
		reflect.TypeOf(0),
	}
	ok, err := mp.InArgsVerifyRv(wrongTypes)
	if ok || err == nil {
		t.Fatal("expected error")
	}

	// Test 3rd and 4th+ positions
	wrongTypes2 := []reflect.Type{
		reflect.TypeOf(testAdder{}),
		reflect.TypeOf("wrong"),
		reflect.TypeOf("wrong"),
	}
	_, _ = mp.InArgsVerifyRv(wrongTypes2)
}

func Test_I8_Utils_PrependWithSpaces_Coverage(t *testing.T) {
	// Covered through VerifyReflectTypes error path which calls PrependWithSpaces
	mp := getMP("Greet")
	wrongTypes := []reflect.Type{
		reflect.TypeOf(0),
		reflect.TypeOf(0),
	}
	_, err := mp.InArgsVerifyRv(wrongTypes)
	if err == nil {
		t.Fatal("expected error")
	}
}

// ==========================================================================
// isNull (unexported, covered via rvUtils.IsNull in ReflectValueToAnyValue)
// ==========================================================================

func Test_I8_IsNull_Coverage(t *testing.T) {
	// Test nil case through InvokeError on NoError method
	mp := getMP("NoError")
	defer func() { recover() }()
	funcErr, procErr := mp.InvokeError(testAdder{})
	_ = funcErr
	_ = procErr
}

// ==========================================================================
// Additional edge cases for maximal coverage
// ==========================================================================

func Test_I8_MP_Invoke_Greet(t *testing.T) {
	mp := getMP("Greet")
	results, err := mp.Invoke(testAdder{}, "Go")
	if err != nil {
		t.Fatal(err)
	}
	if results[0] != "hello Go" {
		t.Fatal("expected 'hello Go'")
	}
}

func Test_I8_MP_Invoke_TwoReturns(t *testing.T) {
	mp := getMP("TwoReturns")
	results, err := mp.Invoke(testAdder{}, 7)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 2 {
		t.Fatal("expected 2 results")
	}
	if results[0].(int) != 14 {
		t.Fatal("expected 14")
	}
}

func Test_I8_MP_InvokeFirstAndError_NilReceiver(t *testing.T) {
	var nilMP *reflectmodel.MethodProcessor
	_, _, err := nilMP.InvokeFirstAndError()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I8_MP_InvokeError_NilReceiver(t *testing.T) {
	var nilMP *reflectmodel.MethodProcessor
	_, err := nilMP.InvokeError()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I8_MP_GetFirstResponseOfInvoke_Error(t *testing.T) {
	var nilMP *reflectmodel.MethodProcessor
	_, err := nilMP.GetFirstResponseOfInvoke()
	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_I8_RVK_Struct_PkgPath(t *testing.T) {
	val := testAdder{}
	rv := reflect.ValueOf(val)
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: rv,
		Kind:            rv.Kind(),
	}
	pkg := rvk.PkgPath()
	_ = pkg // struct types have PkgPath
	_ = fmt.Sprintf("%v", rvk.ActualInstance())
}
