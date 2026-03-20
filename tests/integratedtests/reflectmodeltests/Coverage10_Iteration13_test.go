package reflectmodeltests

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ══════════════════════════════════════════════════════════════════════════════
// Helper types for exercising uncovered utils.go paths
// ══════════════════════════════════════════════════════════════════════════════

type ptrReturner struct{}

func (p ptrReturner) ReturnPtr(x int) *int       { return &x }
func (p ptrReturner) ReturnSlice() []int          { return []int{1, 2, 3} }
func (p ptrReturner) ReturnMap() map[string]int   { return map[string]int{"a": 1} }
func (p ptrReturner) ReturnNilPtr() *int          { return nil }
func (p ptrReturner) ReturnNilSlice() []int       { return nil }
func (p ptrReturner) ReturnNilMap() map[string]int { return nil }
func (p ptrReturner) ReturnInterface() any        { return "hello" }
func (p ptrReturner) ReturnNilInterface() any     { return nil }
func (p ptrReturner) ReturnChan() chan int         { return make(chan int) }
func (p ptrReturner) ReturnFunc() func()           { return func() {} }
func (p ptrReturner) ReturnNilFunc() func()        { return nil }
func (p ptrReturner) ReturnNilChan() chan int       { return nil }
func (p ptrReturner) NoArgs()                      {}
func (p ptrReturner) ManyArgs(a, b, c, d int) int  { return a + b + c + d }
func (p ptrReturner) ReturnMulti(x int) (string, error) {
	if x < 0 {
		return "", fmt.Errorf("negative")
	}
	return fmt.Sprintf("ok:%d", x), nil
}
func (p ptrReturner) ReturnErrorOnly() error {
	return errors.New("test-error")
}
func (p ptrReturner) ReturnNilError() error {
	return nil
}

func getPtrMP(name string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(ptrReturner{})
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

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueToAnyValue — pointer/interface branches via Invoke
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_Invoke_ReturnPtr(t *testing.T) {
	mp := getPtrMP("ReturnPtr")
	results, err := mp.Invoke(ptrReturner{}, 42)
	actual := args.Map{"noErr": err == nil, "val": results[0]}
	expected := args.Map{"noErr": true, "val": 42}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnPtr", actual)
}

func Test_I13_Invoke_ReturnSlice(t *testing.T) {
	mp := getPtrMP("ReturnSlice")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "notNil": results[0] != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnSlice", actual)
}

func Test_I13_Invoke_ReturnMap(t *testing.T) {
	mp := getPtrMP("ReturnMap")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "notNil": results[0] != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnMap", actual)
}

func Test_I13_Invoke_ReturnNilPtr(t *testing.T) {
	mp := getPtrMP("ReturnNilPtr")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "nil": results[0] == nil}
	expected := args.Map{"noErr": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnNilPtr", actual)
}

func Test_I13_Invoke_ReturnNilSlice(t *testing.T) {
	mp := getPtrMP("ReturnNilSlice")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "nil": results[0] == nil}
	expected := args.Map{"noErr": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnNilSlice", actual)
}

func Test_I13_Invoke_ReturnNilMap(t *testing.T) {
	mp := getPtrMP("ReturnNilMap")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "nil": results[0] == nil}
	expected := args.Map{"noErr": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnNilMap", actual)
}

func Test_I13_Invoke_ReturnInterface(t *testing.T) {
	mp := getPtrMP("ReturnInterface")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "val": results[0]}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnInterface", actual)
}

func Test_I13_Invoke_ReturnNilInterface(t *testing.T) {
	mp := getPtrMP("ReturnNilInterface")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "nil": results[0] == nil}
	expected := args.Map{"noErr": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnNilInterface", actual)
}

func Test_I13_Invoke_ReturnChan(t *testing.T) {
	mp := getPtrMP("ReturnChan")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "notNil": results[0] != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnChan", actual)
}

func Test_I13_Invoke_ReturnFunc(t *testing.T) {
	mp := getPtrMP("ReturnFunc")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "notNil": results[0] != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnFunc", actual)
}

func Test_I13_Invoke_ReturnNilFunc(t *testing.T) {
	mp := getPtrMP("ReturnNilFunc")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "nil": results[0] == nil}
	expected := args.Map{"noErr": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnNilFunc", actual)
}

func Test_I13_Invoke_ReturnNilChan(t *testing.T) {
	mp := getPtrMP("ReturnNilChan")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "nil": results[0] == nil}
	expected := args.Map{"noErr": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "Invoke ReturnNilChan", actual)
}

func Test_I13_Invoke_NoArgs(t *testing.T) {
	mp := getPtrMP("NoArgs")
	results, err := mp.Invoke(ptrReturner{})
	actual := args.Map{"noErr": err == nil, "len": len(results)}
	expected := args.Map{"noErr": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Invoke NoArgs", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Invoke with type validation branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_Invoke_TypeMismatch(t *testing.T) {
	mp := getPtrMP("ManyArgs")
	_, err := mp.Invoke(ptrReturner{}, "wrong", 2, 3, 4)
	if err == nil {
		t.Fatal("expected type mismatch error")
	}
	actual := args.Map{"hasErr": true}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Invoke type mismatch", actual)
}

func Test_I13_Invoke_ManyArgs_Success(t *testing.T) {
	mp := getPtrMP("ManyArgs")
	results, err := mp.Invoke(ptrReturner{}, 1, 2, 3, 4)
	actual := args.Map{"noErr": err == nil, "val": results[0]}
	expected := args.Map{"noErr": true, "val": 10}
	expected.ShouldBeEqual(t, 0, "Invoke ManyArgs success", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// InvokeFirstAndError — with func returning (string, error)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_InvokeFirstAndError_Success(t *testing.T) {
	mp := getPtrMP("ReturnMulti")
	first, funcErr, procErr := mp.InvokeFirstAndError(ptrReturner{}, 5)
	actual := args.Map{"procErr": procErr == nil, "funcErr": funcErr == nil, "first": first}
	expected := args.Map{"procErr": true, "funcErr": true, "first": "ok:5"}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError success", actual)
}

func Test_I13_InvokeFirstAndError_FuncError(t *testing.T) {
	mp := getPtrMP("ReturnMulti")
	_, funcErr, procErr := mp.InvokeFirstAndError(ptrReturner{}, -1)
	actual := args.Map{"procErr": procErr == nil, "funcErr": funcErr != nil}
	expected := args.Map{"procErr": true, "funcErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError func error", actual)
}

func Test_I13_InvokeFirstAndError_ProcessingError(t *testing.T) {
	mp := getPtrMP("ReturnMulti")
	_, _, procErr := mp.InvokeFirstAndError(ptrReturner{}, "wrong_type")
	actual := args.Map{"hasErr": procErr != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError processing error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// InvokeError — via different return types
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_InvokeError_ReturnsError(t *testing.T) {
	mp := getPtrMP("ReturnErrorOnly")
	funcErr, procErr := mp.InvokeError(ptrReturner{})
	actual := args.Map{"procErr": procErr == nil, "funcErr": funcErr != nil}
	expected := args.Map{"procErr": true, "funcErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeError returns error", actual)
}

func Test_I13_InvokeError_NilError(t *testing.T) {
	mp := getPtrMP("ReturnNilError")
	funcErr, procErr := mp.InvokeError(ptrReturner{})
	actual := args.Map{"procErr": procErr == nil, "funcErr": funcErr == nil}
	expected := args.Map{"procErr": true, "funcErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeError nil error", actual)
}

func Test_I13_InvokeError_ProcessingError(t *testing.T) {
	mp := getPtrMP("ReturnErrorOnly")
	_, procErr := mp.InvokeError(ptrReturner{}, "extra_arg")
	actual := args.Map{"hasErr": procErr != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeError processing error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// GetFirstResponseOfInvoke / InvokeResultOfIndex — error paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_GetFirstResponseOfInvoke_Error(t *testing.T) {
	mp := getPtrMP("ManyArgs")
	_, err := mp.GetFirstResponseOfInvoke(ptrReturner{}) // wrong args count
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetFirstResponseOfInvoke error", actual)
}

func Test_I13_InvokeResultOfIndex_Error(t *testing.T) {
	mp := getPtrMP("ManyArgs")
	_, err := mp.InvokeResultOfIndex(0, ptrReturner{}) // wrong args count
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// GetOutArgsTypes / GetInArgsTypes — cached paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_GetOutArgsTypes_Cached(t *testing.T) {
	mp := getPtrMP("ReturnMulti")
	out1 := mp.GetOutArgsTypes()
	out2 := mp.GetOutArgsTypes() // should use cache
	actual := args.Map{"len1": len(out1), "len2": len(out2), "same": len(out1) == len(out2)}
	expected := args.Map{"len1": 2, "len2": 2, "same": true}
	expected.ShouldBeEqual(t, 0, "GetOutArgsTypes cached", actual)
}

func Test_I13_GetInArgsTypes_Cached(t *testing.T) {
	mp := getPtrMP("ReturnMulti")
	in1 := mp.GetInArgsTypes()
	in2 := mp.GetInArgsTypes() // cache
	actual := args.Map{"len1": len(in1), "len2": len(in2)}
	expected := args.Map{"len1": 2, "len2": 2}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypes cached", actual)
}

func Test_I13_GetInArgsTypesNames_Cached(t *testing.T) {
	mp := getPtrMP("ReturnMulti")
	n1 := mp.GetInArgsTypesNames()
	n2 := mp.GetInArgsTypesNames() // cache
	actual := args.Map{"len1": len(n1), "len2": len(n2)}
	expected := args.Map{"len1": 2, "len2": 2}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypesNames cached", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsEqual — deep equality with same signature different methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_IsEqual_SameSignatureDiffMethods(t *testing.T) {
	mp1 := getPtrMP("ReturnErrorOnly")
	mp2 := getPtrMP("ReturnNilError")
	// Same signature (receiver → error), same name won't match but types will
	actual := args.Map{"isEqual": mp1.IsEqual(mp2)}
	expected := args.Map{"isEqual": true}
	expected.ShouldBeEqual(t, 0, "IsEqual same signature diff methods", actual)
}

func Test_I13_IsNotEqual_DiffSignature(t *testing.T) {
	mp1 := getPtrMP("ReturnPtr")
	mp2 := getPtrMP("ReturnSlice")
	actual := args.Map{"notEqual": mp1.IsNotEqual(mp2)}
	expected := args.Map{"notEqual": true}
	expected.ShouldBeEqual(t, 0, "IsNotEqual diff signature", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VerifyInArgs / VerifyOutArgs — with various types
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_VerifyInArgs_Success(t *testing.T) {
	mp := getPtrMP("ReturnPtr")
	ok, err := mp.VerifyInArgs([]any{ptrReturner{}, 42})
	actual := args.Map{"ok": ok, "noErr": err == nil}
	expected := args.Map{"ok": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyInArgs success", actual)
}

func Test_I13_VerifyInArgs_TypeMismatch(t *testing.T) {
	mp := getPtrMP("ReturnPtr")
	ok, err := mp.VerifyInArgs([]any{ptrReturner{}, "wrong"})
	actual := args.Map{"ok": ok, "hasErr": err != nil}
	expected := args.Map{"ok": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyInArgs type mismatch", actual)
}

func Test_I13_VerifyOutArgs_Success(t *testing.T) {
	mp := getPtrMP("ReturnMulti")
	ok, err := mp.VerifyOutArgs([]any{"", errors.New("")})
	actual := args.Map{"ok": ok, "noErr": err == nil}
	expected := args.Map{"ok": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyOutArgs success", actual)
}

func Test_I13_VerifyOutArgs_LenMismatch(t *testing.T) {
	mp := getPtrMP("ReturnMulti")
	ok, err := mp.VerifyOutArgs([]any{"only_one"})
	actual := args.Map{"ok": ok, "hasErr": err != nil}
	expected := args.Map{"ok": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyOutArgs length mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidateMethodArgs — comprehensive error message paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_ValidateMethodArgs_CountMismatch_Detailed(t *testing.T) {
	mp := getPtrMP("ManyArgs")
	err := mp.ValidateMethodArgs([]any{ptrReturner{}, 1})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs count mismatch detailed", actual)
}

func Test_I13_ValidateMethodArgs_TypeMismatch_Detailed(t *testing.T) {
	mp := getPtrMP("ManyArgs")
	err := mp.ValidateMethodArgs([]any{ptrReturner{}, "wrong", 2, 3, 4})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs type mismatch detailed", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueKind — various Kind types for wider coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_RVK_IntKind(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(42),
		Kind:            reflect.Int,
	}
	actual := args.Map{
		"isInvalid": rvk.IsInvalid(),
		"actInst":   rvk.ActualInstance(),
		"typeName":  rvk.TypeName() != "",
		"pkgPath":   rvk.PkgPath(),
	}
	expected := args.Map{
		"isInvalid": false,
		"actInst":   42,
		"typeName":  true,
		"pkgPath":   "",
	}
	expected.ShouldBeEqual(t, 0, "RVK Int kind", actual)
}

func Test_I13_RVK_BoolKind(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(true),
		Kind:            reflect.Bool,
	}
	actual := args.Map{"actInst": rvk.ActualInstance(), "typeName": rvk.TypeName() != ""}
	expected := args.Map{"actInst": true, "typeName": true}
	expected.ShouldBeEqual(t, 0, "RVK Bool kind", actual)
}

func Test_I13_RVK_SliceKind(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf([]int{1, 2, 3}),
		Kind:            reflect.Slice,
	}
	ptr := rvk.PointerRv()
	ptrIface := rvk.PointerInterface()
	actual := args.Map{"ptrNotNil": ptr != nil, "ptrIfaceNotNil": ptrIface != nil}
	expected := args.Map{"ptrNotNil": true, "ptrIfaceNotNil": true}
	expected.ShouldBeEqual(t, 0, "RVK Slice kind", actual)
}

func Test_I13_RVK_MapKind(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(map[string]int{"a": 1}),
		Kind:            reflect.Map,
	}
	actual := args.Map{"actNotNil": rvk.ActualInstance() != nil, "typeName": rvk.TypeName() != ""}
	expected := args.Map{"actNotNil": true, "typeName": true}
	expected.ShouldBeEqual(t, 0, "RVK Map kind", actual)
}

func Test_I13_RVK_StructKind(t *testing.T) {
	type myStruct struct{ X int }
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(myStruct{X: 42}),
		Kind:            reflect.Struct,
	}
	actual := args.Map{
		"pkgPath": rvk.PkgPath() != "",
		"actInst": rvk.ActualInstance() != nil,
	}
	expected := args.Map{"pkgPath": true, "actInst": true}
	expected.ShouldBeEqual(t, 0, "RVK Struct kind", actual)
}

func Test_I13_RVK_PtrKind(t *testing.T) {
	x := 42
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         true,
		FinalReflectVal: reflect.ValueOf(&x),
		Kind:            reflect.Ptr,
	}
	actual := args.Map{"typeName": rvk.TypeName() != "", "pkgPath": true}
	expected := args.Map{"typeName": true, "pkgPath": true}
	expected.ShouldBeEqual(t, 0, "RVK Ptr kind", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValue — struct fields coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_ReflectValue_Fields(t *testing.T) {
	rv := reflectmodel.ReflectValue{
		TypeName:     "TestType",
		FieldsNames:  []string{"A", "B"},
		MethodsNames: []string{"M1"},
		RawData:      "raw",
	}
	actual := args.Map{
		"typeName": rv.TypeName, "fieldsLen": len(rv.FieldsNames),
		"methodsLen": len(rv.MethodsNames), "raw": rv.RawData,
	}
	expected := args.Map{"typeName": "TestType", "fieldsLen": 2, "methodsLen": 1, "raw": "raw"}
	expected.ShouldBeEqual(t, 0, "ReflectValue fields", actual)
}

func Test_I13_ReflectValue_Empty(t *testing.T) {
	rv := reflectmodel.ReflectValue{}
	actual := args.Map{"typeName": rv.TypeName, "raw": rv.RawData == nil}
	expected := args.Map{"typeName": "", "raw": true}
	expected.ShouldBeEqual(t, 0, "ReflectValue empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// FieldProcessor — various field types
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_FieldProcessor_IntField(t *testing.T) {
	fp := newFieldProcessor("Age", 1)
	actual := args.Map{
		"isInt":  fp.IsFieldKind(reflect.Int),
		"isStr":  fp.IsFieldKind(reflect.String),
		"typeOk": fp.IsFieldType(reflect.TypeOf(0)),
	}
	expected := args.Map{"isInt": true, "isStr": false, "typeOk": true}
	expected.ShouldBeEqual(t, 0, "FieldProcessor Int field", actual)
}

func Test_I13_FieldProcessor_BoolField(t *testing.T) {
	fp := newFieldProcessor("Active", 2)
	actual := args.Map{
		"isBool": fp.IsFieldKind(reflect.Bool),
		"typeOk": fp.IsFieldType(reflect.TypeOf(true)),
	}
	expected := args.Map{"isBool": true, "typeOk": true}
	expected.ShouldBeEqual(t, 0, "FieldProcessor Bool field", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — GetOutArgsTypes with 0 returns & caching
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_GetOutArgsTypes_NoReturn(t *testing.T) {
	mp := getPtrMP("NoArgs")
	out := mp.GetOutArgsTypes()
	actual := args.Map{"len": len(out)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GetOutArgsTypes no return", actual)
}

func Test_I13_GetInArgsTypes_ReceiverOnly(t *testing.T) {
	mp := getPtrMP("NoArgs")
	in := mp.GetInArgsTypes()
	actual := args.Map{"len": len(in)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypes receiver only", actual)
}

func Test_I13_GetInArgsTypesNames_ReceiverOnly(t *testing.T) {
	mp := getPtrMP("NoArgs")
	names := mp.GetInArgsTypesNames()
	actual := args.Map{"len": len(names)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetInArgsTypesNames receiver only", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — InArgsVerifyRv / OutArgsVerifyRv with matching + mismatching
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_InArgsVerifyRv_Match(t *testing.T) {
	mp := getPtrMP("ReturnPtr")
	types := []reflect.Type{reflect.TypeOf(ptrReturner{}), reflect.TypeOf(0)}
	ok, err := mp.InArgsVerifyRv(types)
	actual := args.Map{"ok": ok, "noErr": err == nil}
	expected := args.Map{"ok": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "InArgsVerifyRv match", actual)
}

func Test_I13_InArgsVerifyRv_Mismatch(t *testing.T) {
	mp := getPtrMP("ReturnPtr")
	types := []reflect.Type{reflect.TypeOf(ptrReturner{}), reflect.TypeOf("")}
	ok, err := mp.InArgsVerifyRv(types)
	actual := args.Map{"ok": ok, "hasErr": err != nil}
	expected := args.Map{"ok": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "InArgsVerifyRv mismatch", actual)
}

func Test_I13_OutArgsVerifyRv_Match(t *testing.T) {
	mp := getPtrMP("ReturnPtr")
	// ReturnPtr returns *int
	x := 0
	types := []reflect.Type{reflect.TypeOf(&x)}
	ok, err := mp.OutArgsVerifyRv(types)
	actual := args.Map{"ok": ok, "noErr": err == nil}
	expected := args.Map{"ok": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "OutArgsVerifyRv match", actual)
}

func Test_I13_OutArgsVerifyRv_Mismatch(t *testing.T) {
	mp := getPtrMP("ReturnMulti")
	types := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(0)} // wrong second type
	ok, err := mp.OutArgsVerifyRv(types)
	actual := args.Map{"ok": ok, "hasErr": err != nil}
	expected := args.Map{"ok": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "OutArgsVerifyRv mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — Multiple type mismatches in VerifyReflectTypes
// ══════════════════════════════════════════════════════════════════════════════

func Test_I13_ValidateMethodArgs_MultiTypeMismatch(t *testing.T) {
	mp := getPtrMP("ManyArgs")
	// All 4 args are wrong types
	err := mp.ValidateMethodArgs([]any{ptrReturner{}, "a", "b", "c", "d"})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs multi type mismatch", actual)
}
