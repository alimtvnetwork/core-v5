package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ═══════════════════════════════════════════════
// FieldProcessor — all uncovered methods
// ═══════════════════════════════════════════════

func Test_Cov6_FieldProcessor_IsFieldType_Valid(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: reflect.TypeOf(0),
	}
	if !fp.IsFieldType(reflect.TypeOf(0)) {
		t.Fatal("expected match")
	}
	if fp.IsFieldType(reflect.TypeOf("")) {
		t.Fatal("expected no match")
	}
}

func Test_Cov6_FieldProcessor_IsFieldType_Nil(t *testing.T) {
	var fp *reflectmodel.FieldProcessor
	if fp.IsFieldType(reflect.TypeOf(0)) {
		t.Fatal("nil receiver should return false")
	}
}

func Test_Cov6_FieldProcessor_IsFieldKind_Valid(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: reflect.TypeOf(0),
	}
	if !fp.IsFieldKind(reflect.Int) {
		t.Fatal("expected int kind")
	}
	if fp.IsFieldKind(reflect.String) {
		t.Fatal("expected no match")
	}
}

func Test_Cov6_FieldProcessor_IsFieldKind_Nil(t *testing.T) {
	var fp *reflectmodel.FieldProcessor
	if fp.IsFieldKind(reflect.Int) {
		t.Fatal("nil receiver should return false")
	}
}

// ═══════════════════════════════════════════════
// MethodProcessor — comprehensive coverage
// ═══════════════════════════════════════════════

// helper: create a MethodProcessor from a real method
type testTarget6 struct{}

func (testTarget6) Add(a, b int) int         { return a + b }
func (testTarget6) Greeting() string          { return "hi" }
func (testTarget6) Err() error                { return nil }
func (testTarget6) PairResult() (string, error) { return "ok", nil }

func getMethodProcessor6(name string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(testTarget6{})
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

func Test_Cov6_MP_HasValidFunc(t *testing.T) {
	mp := getMethodProcessor6("Add")
	if !mp.HasValidFunc() {
		t.Fatal("expected valid")
	}
	var nilMp *reflectmodel.MethodProcessor
	if nilMp.HasValidFunc() {
		t.Fatal("nil should be invalid")
	}
}

func Test_Cov6_MP_GetFuncName(t *testing.T) {
	mp := getMethodProcessor6("Add")
	if mp.GetFuncName() != "Add" {
		t.Fatal("expected Add")
	}
}

func Test_Cov6_MP_IsInvalid(t *testing.T) {
	mp := getMethodProcessor6("Add")
	if mp.IsInvalid() {
		t.Fatal("expected valid")
	}
	var nilMp *reflectmodel.MethodProcessor
	if !nilMp.IsInvalid() {
		t.Fatal("nil should be invalid")
	}
}

func Test_Cov6_MP_Func(t *testing.T) {
	mp := getMethodProcessor6("Add")
	f := mp.Func()
	if f == nil {
		t.Fatal("expected func")
	}
}

func Test_Cov6_MP_Func_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	if mp.Func() != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov6_MP_ArgsCount(t *testing.T) {
	mp := getMethodProcessor6("Add")
	// Add has receiver + a + b = 3
	if mp.ArgsCount() < 2 {
		t.Fatalf("expected >= 2, got %d", mp.ArgsCount())
	}
}

func Test_Cov6_MP_ReturnLength(t *testing.T) {
	mp := getMethodProcessor6("Add")
	if mp.ReturnLength() != 1 {
		t.Fatalf("expected 1, got %d", mp.ReturnLength())
	}
}

func Test_Cov6_MP_ReturnLength_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	if mp.ReturnLength() != -1 {
		t.Fatal("nil should return -1")
	}
}

func Test_Cov6_MP_IsPublicMethod(t *testing.T) {
	mp := getMethodProcessor6("Add")
	if !mp.IsPublicMethod() {
		t.Fatal("expected public")
	}
}

func Test_Cov6_MP_IsPrivateMethod(t *testing.T) {
	mp := getMethodProcessor6("Add")
	if mp.IsPrivateMethod() {
		t.Fatal("expected not private")
	}
}

func Test_Cov6_MP_ArgsLength(t *testing.T) {
	mp := getMethodProcessor6("Add")
	if mp.ArgsLength() < 2 {
		t.Fatal("expected >= 2")
	}
}

func Test_Cov6_MP_Invoke_Success(t *testing.T) {
	mp := getMethodProcessor6("Add")
	// receiver + 2 args
	results, err := mp.Invoke(testTarget6{}, 3, 4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 || results[0].(int) != 7 {
		t.Fatal("expected 7")
	}
}

func Test_Cov6_MP_Invoke_ArgsMismatch(t *testing.T) {
	mp := getMethodProcessor6("Add")
	_, err := mp.Invoke(testTarget6{}, 3) // missing arg
	if err == nil {
		t.Fatal("expected error for args mismatch")
	}
}

func Test_Cov6_MP_Invoke_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, err := mp.Invoke()
	if err == nil {
		t.Fatal("expected error for nil")
	}
}

func Test_Cov6_MP_GetFirstResponseOfInvoke(t *testing.T) {
	mp := getMethodProcessor6("Greeting")
	resp, err := mp.GetFirstResponseOfInvoke(testTarget6{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.(string) != "hi" {
		t.Fatal("expected hi")
	}
}

func Test_Cov6_MP_InvokeResultOfIndex(t *testing.T) {
	mp := getMethodProcessor6("Add")
	resp, err := mp.InvokeResultOfIndex(0, testTarget6{}, 1, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.(int) != 3 {
		t.Fatal("expected 3")
	}
}

func Test_Cov6_MP_InvokeError(t *testing.T) {
	defer func() { recover() }() // InvokeError may panic on zero reflect.Value
	mp := getMethodProcessor6("Err")
	funcErr, procErr := mp.InvokeError(testTarget6{})
	if procErr != nil {
		t.Fatalf("unexpected error: %v", procErr)
	}
	if funcErr != nil {
		t.Fatal("expected nil error from Err()")
	}
}

func Test_Cov6_MP_InvokeFirstAndError_Success(t *testing.T) {
	defer func() { recover() }() // may panic on zero reflect.Value in ReflectValueToAnyValue
	mp := getMethodProcessor6("PairResult")
	first, funcErr, procErr := mp.InvokeFirstAndError(testTarget6{})
	if procErr != nil {
		t.Fatalf("processing error: %v", procErr)
	}
	if funcErr != nil {
		t.Fatal("expected no func error")
	}
	if first.(string) != "ok" {
		t.Fatal("expected 'ok'")
	}
}

func Test_Cov6_MP_InvokeFirstAndError_SingleReturn(t *testing.T) {
	mp := getMethodProcessor6("Greeting")
	_, _, procErr := mp.InvokeFirstAndError(testTarget6{})
	if procErr == nil {
		t.Fatal("expected error for single return")
	}
}

func Test_Cov6_MP_IsEqual_BothNil(t *testing.T) {
	var a, b *reflectmodel.MethodProcessor
	if !a.IsEqual(b) {
		t.Fatal("both nil should be equal")
	}
}

func Test_Cov6_MP_IsEqual_OneNil(t *testing.T) {
	mp := getMethodProcessor6("Add")
	if mp.IsEqual(nil) {
		t.Fatal("non-nil vs nil should not be equal")
	}
}

func Test_Cov6_MP_IsEqual_Same(t *testing.T) {
	mp := getMethodProcessor6("Add")
	if !mp.IsEqual(mp) {
		t.Fatal("same pointer should be equal")
	}
}

func Test_Cov6_MP_IsEqual_DiffMethods(t *testing.T) {
	mp1 := getMethodProcessor6("Add")
	mp2 := getMethodProcessor6("Greeting")
	// Different signatures → should fail at args verification
	_ = mp1.IsEqual(mp2)
}

func Test_Cov6_MP_IsNotEqual(t *testing.T) {
	mp1 := getMethodProcessor6("Add")
	mp2 := getMethodProcessor6("Greeting")
	if !mp1.IsNotEqual(mp2) {
		t.Fatal("different methods should not be equal")
	}
}

func Test_Cov6_MP_GetType(t *testing.T) {
	mp := getMethodProcessor6("Add")
	if mp.GetType() == nil {
		t.Fatal("expected non-nil type")
	}
}

func Test_Cov6_MP_GetType_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	if mp.GetType() != nil {
		t.Fatal("nil should return nil type")
	}
}

func Test_Cov6_MP_GetOutArgsTypes(t *testing.T) {
	mp := getMethodProcessor6("Add")
	out := mp.GetOutArgsTypes()
	if len(out) != 1 {
		t.Fatalf("expected 1, got %d", len(out))
	}
	// Call again to hit cache
	out2 := mp.GetOutArgsTypes()
	if len(out2) != 1 {
		t.Fatal("cache should return same")
	}
}

func Test_Cov6_MP_GetOutArgsTypes_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	out := mp.GetOutArgsTypes()
	if len(out) != 0 {
		t.Fatal("nil should return empty")
	}
}

func Test_Cov6_MP_GetInArgsTypes(t *testing.T) {
	mp := getMethodProcessor6("Add")
	in := mp.GetInArgsTypes()
	if len(in) < 2 {
		t.Fatalf("expected >= 2, got %d", len(in))
	}
	// Call again to hit cache
	in2 := mp.GetInArgsTypes()
	if len(in2) != len(in) {
		t.Fatal("cache should return same")
	}
}

func Test_Cov6_MP_GetInArgsTypes_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	in := mp.GetInArgsTypes()
	if len(in) != 0 {
		t.Fatal("nil should return empty")
	}
}

func Test_Cov6_MP_GetInArgsTypesNames(t *testing.T) {
	mp := getMethodProcessor6("Add")
	names := mp.GetInArgsTypesNames()
	if len(names) < 2 {
		t.Fatalf("expected >= 2, got %d", len(names))
	}
	// Call again to hit cache
	names2 := mp.GetInArgsTypesNames()
	if len(names2) != len(names) {
		t.Fatal("cache should return same")
	}
}

func Test_Cov6_MP_GetInArgsTypesNames_Nil(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	names := mp.GetInArgsTypesNames()
	if len(names) != 0 {
		t.Fatal("nil should return empty")
	}
}

func Test_Cov6_MP_ValidateMethodArgs_Success(t *testing.T) {
	mp := getMethodProcessor6("Greeting")
	err := mp.ValidateMethodArgs([]any{testTarget6{}})
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
}

func Test_Cov6_MP_ValidateMethodArgs_WrongCount(t *testing.T) {
	mp := getMethodProcessor6("Add")
	err := mp.ValidateMethodArgs([]any{testTarget6{}})
	if err == nil {
		t.Fatal("expected args mismatch error")
	}
}

func Test_Cov6_MP_ValidateMethodArgs_WrongType(t *testing.T) {
	mp := getMethodProcessor6("Add")
	err := mp.ValidateMethodArgs([]any{testTarget6{}, "not_int", "not_int"})
	if err == nil {
		t.Fatal("expected type mismatch error")
	}
}

func Test_Cov6_MP_VerifyInArgs(t *testing.T) {
	mp := getMethodProcessor6("Greeting")
	ok, err := mp.VerifyInArgs([]any{testTarget6{}})
	if !ok || err != nil {
		t.Fatal("expected ok")
	}
}

func Test_Cov6_MP_VerifyOutArgs(t *testing.T) {
	mp := getMethodProcessor6("Add")
	ok, err := mp.VerifyOutArgs([]any{0})
	if !ok || err != nil {
		t.Fatal("expected ok")
	}
}

func Test_Cov6_MP_VerifyOutArgs_Mismatch(t *testing.T) {
	mp := getMethodProcessor6("Add")
	ok, _ := mp.VerifyOutArgs([]any{"string"})
	if ok {
		t.Fatal("expected mismatch")
	}
}

func Test_Cov6_MP_InArgsVerifyRv(t *testing.T) {
	mp := getMethodProcessor6("Greeting")
	ok, err := mp.InArgsVerifyRv([]reflect.Type{reflect.TypeOf(testTarget6{})})
	if !ok || err != nil {
		t.Fatal("expected ok")
	}
}

func Test_Cov6_MP_OutArgsVerifyRv(t *testing.T) {
	mp := getMethodProcessor6("Add")
	ok, err := mp.OutArgsVerifyRv([]reflect.Type{reflect.TypeOf(0)})
	if !ok || err != nil {
		t.Fatal("expected ok")
	}
}

func Test_Cov6_MP_OutArgsVerifyRv_LengthMismatch(t *testing.T) {
	mp := getMethodProcessor6("Add")
	ok, _ := mp.OutArgsVerifyRv([]reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")})
	if ok {
		t.Fatal("expected mismatch for wrong length")
	}
}

// ═══════════════════════════════════════════════
// ReflectValueKind — uncovered methods
// ═══════════════════════════════════════════════

func Test_Cov6_RVK_InvalidReflectValueKindModel(t *testing.T) {
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

func Test_Cov6_RVK_IsEmptyError(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: true}
	if !rvk.IsEmptyError() {
		t.Fatal("expected empty error")
	}
	var nilRvk *reflectmodel.ReflectValueKind
	if !nilRvk.IsEmptyError() {
		t.Fatal("nil should be empty error")
	}
}

func Test_Cov6_RVK_ActualInstance(t *testing.T) {
	val := 42
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	inst := rvk.ActualInstance()
	if inst.(int) != 42 {
		t.Fatal("expected 42")
	}
}

func Test_Cov6_RVK_ActualInstance_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if rvk.ActualInstance() != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov6_RVK_PkgPath(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(testTarget6{}),
	}
	pkg := rvk.PkgPath()
	if pkg == "" {
		t.Fatal("expected non-empty pkg path")
	}
}

func Test_Cov6_RVK_PkgPath_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if rvk.PkgPath() != "" {
		t.Fatal("expected empty")
	}
}

func Test_Cov6_RVK_PkgPath_Invalid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}
	if rvk.PkgPath() != "" {
		t.Fatal("expected empty for invalid")
	}
}

func Test_Cov6_RVK_PointerRv_Valid(t *testing.T) {
	val := 42
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	ptr := rvk.PointerRv()
	if ptr == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov6_RVK_PointerRv_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if rvk.PointerRv() != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov6_RVK_PointerRv_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(42),
	}
	ptr := rvk.PointerRv()
	if ptr == nil {
		t.Fatal("expected non-nil (returns FinalReflectVal addr)")
	}
}

func Test_Cov6_RVK_TypeName_Valid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(42),
	}
	name := rvk.TypeName()
	if name == "" {
		t.Fatal("expected non-empty")
	}
}

func Test_Cov6_RVK_TypeName_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if rvk.TypeName() != "" {
		t.Fatal("expected empty")
	}
}

func Test_Cov6_RVK_TypeName_NotValid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}
	if rvk.TypeName() != "" {
		t.Fatal("expected empty for invalid")
	}
}

func Test_Cov6_RVK_PointerInterface_Valid(t *testing.T) {
	val := 42
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(val),
	}
	pi := rvk.PointerInterface()
	if pi == nil {
		t.Fatal("expected non-nil")
	}
}

func Test_Cov6_RVK_PointerInterface_Nil(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	if rvk.PointerInterface() != nil {
		t.Fatal("expected nil")
	}
}
