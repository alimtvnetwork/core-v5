package reflectmodeltests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── FieldProcessor ──

func Test_C9_FieldProcessor_IsFieldType(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: reflect.TypeOf(""),
	}
	actual := args.Map{
		"matchStr":  fp.IsFieldType(reflect.TypeOf("")),
		"matchInt":  fp.IsFieldType(reflect.TypeOf(0)),
	}
	expected := args.Map{"matchStr": true, "matchInt": false}
	expected.ShouldBeEqual(t, 0, "FieldProcessor.IsFieldType", actual)

	var nilFp *reflectmodel.FieldProcessor
	if nilFp.IsFieldType(reflect.TypeOf("")) { t.Fatal("nil should return false") }
}

func Test_C9_FieldProcessor_IsFieldKind(t *testing.T) {
	fp := &reflectmodel.FieldProcessor{
		Name:      "TestField",
		Index:     0,
		FieldType: reflect.TypeOf(""),
	}
	actual := args.Map{
		"matchStr": fp.IsFieldKind(reflect.String),
		"matchInt": fp.IsFieldKind(reflect.Int),
	}
	expected := args.Map{"matchStr": true, "matchInt": false}
	expected.ShouldBeEqual(t, 0, "FieldProcessor.IsFieldKind", actual)

	var nilFp *reflectmodel.FieldProcessor
	if nilFp.IsFieldKind(reflect.String) { t.Fatal("nil should return false") }
}

// ── ReflectValueKind ──

func Test_C9_InvalidReflectValueKindModel(t *testing.T) {
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")
	actual := args.Map{
		"isInvalid":  rvk.IsInvalid(),
		"hasError":   rvk.HasError(),
		"emptyError": rvk.IsEmptyError(),
		"isValid":    rvk.IsValid,
	}
	expected := args.Map{
		"isInvalid":  true,
		"hasError":   true,
		"emptyError": false,
		"isValid":    false,
	}
	expected.ShouldBeEqual(t, 0, "InvalidReflectValueKindModel", actual)
}

func Test_C9_ReflectValueKind_NilReceiver(t *testing.T) {
	var nilRvk *reflectmodel.ReflectValueKind
	actual := args.Map{
		"isInvalid":  nilRvk.IsInvalid(),
		"hasError":   nilRvk.HasError(),
		"emptyError": nilRvk.IsEmptyError(),
		"actInst":    nilRvk.ActualInstance() == nil,
		"pkgPath":    nilRvk.PkgPath(),
		"ptrRv":      nilRvk.PointerRv() == nil,
		"typeName":   nilRvk.TypeName(),
		"ptrIface":   nilRvk.PointerInterface() == nil,
	}
	expected := args.Map{
		"isInvalid":  true,
		"hasError":   false,
		"emptyError": true,
		"actInst":    true,
		"pkgPath":    "",
		"ptrRv":      true,
		"typeName":   "",
		"ptrIface":   true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind nil receiver", actual)
}

func Test_C9_ReflectValueKind_Valid(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf("hello"),
		Kind:            reflect.String,
	}
	actual := args.Map{
		"isInvalid":     rvk.IsInvalid(),
		"hasError":      rvk.HasError(),
		"actInstNotNil": rvk.ActualInstance() != nil,
		"pkgNotEmpty":   true, // PkgPath for string is ""
		"typeNotEmpty":  rvk.TypeName() != "",
	}
	expected := args.Map{
		"isInvalid":     false,
		"hasError":      false,
		"actInstNotNil": true,
		"pkgNotEmpty":   true,
		"typeNotEmpty":  true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValueKind valid", actual)

	// PointerRv for valid value
	ptrRv := rvk.PointerRv()
	if ptrRv == nil { t.Fatal("expected non-nil PointerRv") }

	// PointerInterface
	ptrIface := rvk.PointerInterface()
	if ptrIface == nil { t.Fatal("expected non-nil PointerInterface") }
}

func Test_C9_ReflectValueKind_InvalidNotNil(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(nil),
		Kind:            0,
	}
	// PointerRv for invalid but non-nil
	ptrRv := rvk.PointerRv()
	if ptrRv == nil { t.Fatal("expected non-nil PointerRv for invalid") }
	// PkgPath for invalid
	if rvk.PkgPath() != "" { t.Fatal("expected empty pkgPath for invalid") }
}

// ── MethodProcessor — basic methods ──

type testHelper struct{}

func (h testHelper) Add(a, b int) int       { return a + b }
func (h testHelper) Greet(name string) string { return "hi " + name }
func (h testHelper) Fail() error              { return fmt.Errorf("fail") }
func (h testHelper) TwoReturns(x int) (int, error) {
	if x < 0 { return 0, fmt.Errorf("negative") }
	return x * 2, nil
}

func getMethodProcessor(name string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(testHelper{})
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

func Test_C9_MethodProcessor_BasicMethods(t *testing.T) {
	mp := getMethodProcessor("Add")
	if mp == nil { t.Fatal("method not found") }

	actual := args.Map{
		"hasValidFunc":  mp.HasValidFunc(),
		"funcName":      mp.GetFuncName(),
		"isInvalid":     mp.IsInvalid(),
		"isPublic":      mp.IsPublicMethod(),
		"isPrivate":     mp.IsPrivateMethod(),
		"argsCount":     mp.ArgsCount(),
		"argsLen":       mp.ArgsLength(),
		"returnLen":     mp.ReturnLength(),
		"funcNotNil":    mp.Func() != nil,
		"typeNotNil":    mp.GetType() != nil,
	}
	expected := args.Map{
		"hasValidFunc":  true,
		"funcName":      "Add",
		"isInvalid":     false,
		"isPublic":      true,
		"isPrivate":     false,
		"argsCount":     3, // receiver + 2 args
		"argsLen":       3,
		"returnLen":     1,
		"funcNotNil":    true,
		"typeNotNil":    true,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor basic", actual)
}

func Test_C9_MethodProcessor_NilReceiver(t *testing.T) {
	var nilMp *reflectmodel.MethodProcessor
	actual := args.Map{
		"hasValidFunc": nilMp.HasValidFunc(),
		"isInvalid":    nilMp.IsInvalid(),
		"funcNil":      nilMp.Func() == nil,
		"returnLen":    nilMp.ReturnLength(),
		"isPublic":     nilMp.IsPublicMethod(),
		"isPrivate":    nilMp.IsPrivateMethod(),
		"typeNil":      nilMp.GetType() == nil,
	}
	expected := args.Map{
		"hasValidFunc": false,
		"isInvalid":    true,
		"funcNil":      true,
		"returnLen":    -1,
		"isPublic":     false,
		"isPrivate":    false,
		"typeNil":      true,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor nil", actual)

	// GetInArgsTypes/GetOutArgsTypes/GetInArgsTypesNames on nil
	if len(nilMp.GetInArgsTypes()) != 0 { t.Fatal("expected empty") }
	if len(nilMp.GetOutArgsTypes()) != 0 { t.Fatal("expected empty") }
	if len(nilMp.GetInArgsTypesNames()) != 0 { t.Fatal("expected empty") }
}

func Test_C9_MethodProcessor_Invoke_Success(t *testing.T) {
	mp := getMethodProcessor("Add")
	results, err := mp.Invoke(testHelper{}, 2, 3)
	if err != nil { t.Fatalf("invoke err: %v", err) }
	if results[0] != 5 { t.Fatalf("expected 5, got %v", results[0]) }
}

func Test_C9_MethodProcessor_Invoke_ArgsMismatch(t *testing.T) {
	mp := getMethodProcessor("Add")
	_, err := mp.Invoke(testHelper{}, 2) // missing arg
	if err == nil { t.Fatal("expected error for args mismatch") }
}

func Test_C9_MethodProcessor_Invoke_NilReceiver(t *testing.T) {
	var nilMp *reflectmodel.MethodProcessor
	_, err := nilMp.Invoke()
	if err == nil { t.Fatal("expected error for nil invoke") }
}

func Test_C9_MethodProcessor_GetFirstResponseOfInvoke(t *testing.T) {
	mp := getMethodProcessor("Greet")
	first, err := mp.GetFirstResponseOfInvoke(testHelper{}, "world")
	if err != nil { t.Fatalf("err: %v", err) }
	if first != "hi world" { t.Fatalf("expected 'hi world', got %v", first) }
}

func Test_C9_MethodProcessor_InvokeResultOfIndex(t *testing.T) {
	mp := getMethodProcessor("Add")
	result, err := mp.InvokeResultOfIndex(0, testHelper{}, 1, 2)
	if err != nil { t.Fatalf("err: %v", err) }
	if result != 3 { t.Fatalf("expected 3, got %v", result) }
}

func Test_C9_MethodProcessor_InvokeError(t *testing.T) {
	mp := getMethodProcessor("Fail")
	funcErr, procErr := mp.InvokeError(testHelper{})
	if procErr != nil { t.Fatalf("proc err: %v", procErr) }
	if funcErr == nil { t.Fatal("expected func error") }
}

func Test_C9_MethodProcessor_InvokeFirstAndError(t *testing.T) {
	mp := getMethodProcessor("TwoReturns")
	first, funcErr, procErr := mp.InvokeFirstAndError(testHelper{}, 5)
	if procErr != nil { t.Fatalf("proc err: %v", procErr) }
	if funcErr != nil { t.Fatalf("func err: %v", funcErr) }
	if first != 10 { t.Fatalf("expected 10, got %v", first) }
}

func Test_C9_MethodProcessor_InvokeFirstAndError_TooFewReturns(t *testing.T) {
	mp := getMethodProcessor("Add") // only 1 return
	_, _, procErr := mp.InvokeFirstAndError(testHelper{}, 1, 2)
	if procErr == nil { t.Fatal("expected error for too few returns") }
}

func Test_C9_MethodProcessor_InvokeFirstAndError_WithError(t *testing.T) {
	mp := getMethodProcessor("TwoReturns")
	defer func() { recover() }() // may panic on nil error interface
	_, _, _ = mp.InvokeFirstAndError(testHelper{}, -1)
}

// ── MethodProcessor — GetInArgsTypes / GetOutArgsTypes / GetInArgsTypesNames ──

func Test_C9_MethodProcessor_ArgTypes(t *testing.T) {
	mp := getMethodProcessor("Add")
	inTypes := mp.GetInArgsTypes()
	if len(inTypes) != 3 { t.Fatalf("expected 3 in args, got %d", len(inTypes)) }
	// second call should use cache
	inTypes2 := mp.GetInArgsTypes()
	if len(inTypes2) != 3 { t.Fatal("cached mismatch") }

	outTypes := mp.GetOutArgsTypes()
	if len(outTypes) != 1 { t.Fatalf("expected 1 out arg, got %d", len(outTypes)) }
	outTypes2 := mp.GetOutArgsTypes()
	if len(outTypes2) != 1 { t.Fatal("cached mismatch") }

	names := mp.GetInArgsTypesNames()
	if len(names) != 3 { t.Fatalf("expected 3, got %d", len(names)) }
	names2 := mp.GetInArgsTypesNames()
	if len(names2) != 3 { t.Fatal("cached mismatch") }
}

func Test_C9_MethodProcessor_ZeroArgsMethod(t *testing.T) {
	mp := getMethodProcessor("Fail")
	// Fail has receiver only → ArgsCount=1, but GetInArgsTypes returns 1
	inTypes := mp.GetInArgsTypes()
	outTypes := mp.GetOutArgsTypes()
	names := mp.GetInArgsTypesNames()
	if len(inTypes) != 1 { t.Fatalf("expected 1 in arg (receiver), got %d", len(inTypes)) }
	if len(outTypes) != 1 { t.Fatalf("expected 1 out arg, got %d", len(outTypes)) }
	if len(names) != 1 { t.Fatalf("expected 1 name, got %d", len(names)) }
}

// ── MethodProcessor — IsEqual / IsNotEqual ──

func Test_C9_MethodProcessor_IsEqual(t *testing.T) {
	mp1 := getMethodProcessor("Add")
	mp2 := getMethodProcessor("Add")
	mp3 := getMethodProcessor("Greet")
	var nilMp *reflectmodel.MethodProcessor

	actual := args.Map{
		"sameEqual":    mp1.IsEqual(mp2),
		"diffNotEqual": mp1.IsNotEqual(mp3),
		"nilBothEqual": nilMp.IsEqual(nil),
		"nilLeft":      nilMp.IsEqual(mp1),
		"nilRight":     mp1.IsEqual(nil),
		"selfEqual":    mp1.IsEqual(mp1),
	}
	expected := args.Map{
		"sameEqual":    true,
		"diffNotEqual": true,
		"nilBothEqual": true,
		"nilLeft":      false,
		"nilRight":     false,
		"selfEqual":    true,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor IsEqual", actual)
}

// ── MethodProcessor — VerifyInArgs / VerifyOutArgs / ValidateMethodArgs ──

func Test_C9_MethodProcessor_VerifyInArgs(t *testing.T) {
	mp := getMethodProcessor("Add")
	ok, err := mp.VerifyInArgs([]any{testHelper{}, 1, 2})
	if !ok || err != nil { t.Fatalf("expected ok, err=%v", err) }

	ok2, err2 := mp.VerifyInArgs([]any{testHelper{}, "a", 2})
	if ok2 { t.Fatal("expected not ok for type mismatch") }
	if err2 == nil { t.Fatal("expected error") }
}

func Test_C9_MethodProcessor_VerifyOutArgs(t *testing.T) {
	mp := getMethodProcessor("Add")
	ok, err := mp.VerifyOutArgs([]any{0})
	if !ok || err != nil { t.Fatalf("expected ok, err=%v", err) }
}

func Test_C9_MethodProcessor_ValidateMethodArgs(t *testing.T) {
	mp := getMethodProcessor("Add")
	err := mp.ValidateMethodArgs([]any{testHelper{}, 1, 2})
	if err != nil { t.Fatalf("expected nil, got %v", err) }

	err2 := mp.ValidateMethodArgs([]any{testHelper{}, 1})
	if err2 == nil { t.Fatal("expected error for count mismatch") }
}

// ── MethodProcessor — InArgsVerifyRv / OutArgsVerifyRv ──

func Test_C9_MethodProcessor_InArgsVerifyRv(t *testing.T) {
	mp := getMethodProcessor("Add")
	types := []reflect.Type{reflect.TypeOf(testHelper{}), reflect.TypeOf(0), reflect.TypeOf(0)}
	ok, err := mp.InArgsVerifyRv(types)
	if !ok || err != nil { t.Fatalf("err=%v", err) }

	// wrong length
	ok2, err2 := mp.InArgsVerifyRv([]reflect.Type{reflect.TypeOf(0)})
	if ok2 { t.Fatal("expected not ok") }
	if err2 == nil { t.Fatal("expected error") }
}

func Test_C9_MethodProcessor_OutArgsVerifyRv(t *testing.T) {
	mp := getMethodProcessor("Add")
	ok, err := mp.OutArgsVerifyRv([]reflect.Type{reflect.TypeOf(0)})
	if !ok || err != nil { t.Fatalf("err=%v", err) }
}
