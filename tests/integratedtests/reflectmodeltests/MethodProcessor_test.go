package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ===== MethodProcessor Tests =====

// --- Validity & Identity ---

func Test_MethodProcessor_HasValidFunc_True(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")
	if mp == nil {
		t.Fatal("failed to create MethodProcessor for PublicMethod")
	}

	if !mp.HasValidFunc() {
		t.Error("expected HasValidFunc() = true")
	}
}

// Note: HasValidFunc nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

// Note: IsInvalid nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_IsInvalid_Valid(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	if mp.IsInvalid() {
		t.Error("expected IsInvalid() = false for valid method")
	}
}

func Test_MethodProcessor_GetFuncName(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	if mp.GetFuncName() != "PublicMethod" {
		t.Errorf("GetFuncName() = %q, want %q", mp.GetFuncName(), "PublicMethod")
	}
}

// --- Func ---

func Test_MethodProcessor_Func_Valid(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	fn := mp.Func()
	if fn == nil {
		t.Error("expected Func() to return non-nil for valid method")
	}
}

// Note: Func nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

// --- Args & Return Counts ---

func Test_MethodProcessor_ArgsCount(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	// reflect.Method includes receiver as first arg
	// sampleStruct.PublicMethod(a string, b int) => 3 args (receiver + 2)
	got := mp.ArgsCount()
	if got != 3 {
		t.Errorf("ArgsCount() = %d, want 3 (receiver + 2 params)", got)
	}
}

func Test_MethodProcessor_ArgsLength(t *testing.T) {
	mp := newMethodProcessor("NoArgsMethod")

	// NoArgsMethod() => 1 arg (receiver only)
	got := mp.ArgsLength()
	if got != 1 {
		t.Errorf("ArgsLength() = %d, want 1 (receiver only)", got)
	}
}

func Test_MethodProcessor_ReturnLength(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	got := mp.ReturnLength()
	if got != 2 {
		t.Errorf("ReturnLength() = %d, want 2 (string, error)", got)
	}
}

// Note: ReturnLength nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_ReturnLength_MultiReturn(t *testing.T) {
	mp := newMethodProcessor("MultiReturn")

	got := mp.ReturnLength()
	if got != 3 {
		t.Errorf("ReturnLength() = %d, want 3 (int, string, error)", got)
	}
}

// --- Public/Private ---

func Test_MethodProcessor_IsPublicMethod(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	if !mp.IsPublicMethod() {
		t.Error("expected IsPublicMethod() = true for PublicMethod")
	}
}

// Note: IsPublicMethod nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_IsPrivateMethod_False(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	if mp.IsPrivateMethod() {
		t.Error("expected IsPrivateMethod() = false for PublicMethod")
	}
}

// --- GetType ---

func Test_MethodProcessor_GetType_Valid(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	got := mp.GetType()
	if got == nil {
		t.Error("expected GetType() to return non-nil for valid method")
	}
}

// Note: GetType nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

// --- InArgs & OutArgs Types ---

func Test_MethodProcessor_GetInArgsTypes(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	types := mp.GetInArgsTypes()
	// receiver + 2 params = 3
	if len(types) != 3 {
		t.Errorf("GetInArgsTypes() len = %d, want 3", len(types))
	}
}

// Note: GetInArgsTypes nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_GetInArgsTypes_Cached(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	first := mp.GetInArgsTypes()
	second := mp.GetInArgsTypes()

	if len(first) != len(second) {
		t.Error("expected cached GetInArgsTypes to return same length")
	}
}

func Test_MethodProcessor_GetOutArgsTypes(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	types := mp.GetOutArgsTypes()
	if len(types) != 2 {
		t.Errorf("GetOutArgsTypes() len = %d, want 2", len(types))
	}
}

// Note: GetOutArgsTypes nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_GetOutArgsTypes_NoArgs(t *testing.T) {
	mp := newMethodProcessor("NoArgsMethod")

	// NoArgsMethod returns string => 1 out type
	types := mp.GetOutArgsTypes()
	if len(types) != 1 {
		t.Errorf("GetOutArgsTypes() len = %d, want 1", len(types))
	}
}

func Test_MethodProcessor_GetInArgsTypesNames(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	names := mp.GetInArgsTypesNames()
	// receiver type + string + int = 3
	if len(names) != 3 {
		t.Errorf("GetInArgsTypesNames() len = %d, want 3", len(names))
	}
}

// Note: GetInArgsTypesNames nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

// --- IsEqual / IsNotEqual ---

func Test_MethodProcessor_IsEqual_BothNil(t *testing.T) {
	var a, b *reflectmodel.MethodProcessor

	if !a.IsEqual(b) {
		t.Error("expected IsEqual(nil, nil) = true")
	}
}

func Test_MethodProcessor_IsEqual_OneNil(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")
	var nilMp *reflectmodel.MethodProcessor

	if mp.IsEqual(nilMp) {
		t.Error("expected IsEqual(valid, nil) = false")
	}
}

func Test_MethodProcessor_IsEqual_Same(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	if !mp.IsEqual(mp) {
		t.Error("expected IsEqual with itself = true")
	}
}

func Test_MethodProcessor_IsEqual_SameMethod(t *testing.T) {
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("PublicMethod")

	if !a.IsEqual(b) {
		t.Error("expected IsEqual for same method = true")
	}
}

func Test_MethodProcessor_IsNotEqual(t *testing.T) {
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("NoArgsMethod")

	if !a.IsNotEqual(b) {
		t.Error("expected IsNotEqual for different methods = true")
	}
}

// --- ValidateMethodArgs ---

func Test_MethodProcessor_ValidateMethodArgs_WrongCount(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	// PublicMethod expects receiver + string + int = 3 args
	err := mp.ValidateMethodArgs([]any{"a"})
	if err == nil {
		t.Error("expected error for wrong arg count")
	}
}

func Test_MethodProcessor_ValidateMethodArgs_WrongType(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	// receiver + string + int, but we give receiver + int + int
	err := mp.ValidateMethodArgs([]any{sampleStruct{}, 42, 42})
	if err == nil {
		t.Error("expected error for wrong arg type")
	}
}

func Test_MethodProcessor_ValidateMethodArgs_Correct(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	err := mp.ValidateMethodArgs([]any{sampleStruct{}, "hello", 42})
	if err != nil {
		t.Errorf("expected no error for correct args, got: %v", err)
	}
}

// --- VerifyInArgs / VerifyOutArgs ---

func Test_MethodProcessor_VerifyInArgs_Match(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	ok, err := mp.VerifyInArgs([]any{sampleStruct{}, "s", 1})
	if !ok || err != nil {
		t.Errorf("expected VerifyInArgs match, got ok=%v err=%v", ok, err)
	}
}

func Test_MethodProcessor_VerifyOutArgs_Match(t *testing.T) {
	mp := newMethodProcessor("NoArgsMethod")

	ok, err := mp.VerifyOutArgs([]any{""})
	if !ok || err != nil {
		t.Errorf("expected VerifyOutArgs match, got ok=%v err=%v", ok, err)
	}
}

func Test_MethodProcessor_InArgsVerifyRv_LengthMismatch(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	ok, err := mp.InArgsVerifyRv([]reflect.Type{reflect.TypeOf("")})
	if ok {
		t.Error("expected InArgsVerifyRv = false for length mismatch")
	}
	if err == nil {
		t.Error("expected error for length mismatch")
	}
}

// --- Invoke ---

func Test_MethodProcessor_Invoke_Success(t *testing.T) {
	mp := newMethodProcessor("NoArgsMethod")

	results, err := mp.Invoke(sampleStruct{})
	if err != nil {
		t.Errorf("Invoke error: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("Invoke results len = %d, want 1", len(results))
	}

	if results[0] != "hello" {
		t.Errorf("Invoke result = %v, want %q", results[0], "hello")
	}
}

// Note: Invoke nil receiver test migrated to MethodProcessor_NilReceiver_testcases.go

func Test_MethodProcessor_NilReceiver(t *testing.T) {
	for caseIndex, tc := range methodProcessorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

func Test_MethodProcessor_Invoke_ArgsMismatch(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")

	_, err := mp.Invoke(sampleStruct{}, "only one arg")
	if err == nil {
		t.Error("expected error for args count mismatch")
	}
}

func Test_MethodProcessor_GetFirstResponseOfInvoke(t *testing.T) {
	mp := newMethodProcessor("NoArgsMethod")

	first, err := mp.GetFirstResponseOfInvoke(sampleStruct{})
	if err != nil {
		t.Errorf("GetFirstResponseOfInvoke error: %v", err)
	}

	if first != "hello" {
		t.Errorf("first response = %v, want %q", first, "hello")
	}
}

func Test_MethodProcessor_InvokeResultOfIndex(t *testing.T) {
	mp := newMethodProcessor("NoArgsMethod")

	result, err := mp.InvokeResultOfIndex(0, sampleStruct{})
	if err != nil {
		t.Errorf("InvokeResultOfIndex error: %v", err)
	}

	if result != "hello" {
		t.Errorf("result = %v, want %q", result, "hello")
	}
}
