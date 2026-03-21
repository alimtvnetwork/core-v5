package reflectmodeltests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ══════════════════════════════════════════════════════════════════════════════
// MethodProcessor — nil receiver paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_MethodProcessor_NilReceiver_HasValidFunc(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{"hasValid": mp.HasValidFunc()}
	expected := args.Map{"hasValid": false}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- HasValidFunc", actual)
}

func Test_I14_MethodProcessor_NilReceiver_IsInvalid(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{"isInvalid": mp.IsInvalid()}
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- IsInvalid", actual)
}

func Test_I14_MethodProcessor_NilReceiver_Func(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{"funcNil": mp.Func() == nil}
	expected := args.Map{"funcNil": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- Func", actual)
}

func Test_I14_MethodProcessor_NilReceiver_ReturnLength(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{"retLen": mp.ReturnLength()}
	expected := args.Map{"retLen": -1}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- ReturnLength", actual)
}

func Test_I14_MethodProcessor_NilReceiver_GetType(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{"typeNil": mp.GetType() == nil}
	expected := args.Map{"typeNil": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- GetType", actual)
}

func Test_I14_MethodProcessor_NilReceiver_GetOutArgsTypes(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	out := mp.GetOutArgsTypes()
	actual := args.Map{"len": len(out)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- GetOutArgsTypes", actual)
}

func Test_I14_MethodProcessor_NilReceiver_GetInArgsTypes(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	in := mp.GetInArgsTypes()
	actual := args.Map{"len": len(in)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- GetInArgsTypes", actual)
}

func Test_I14_MethodProcessor_NilReceiver_GetInArgsTypesNames(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	names := mp.GetInArgsTypesNames()
	actual := args.Map{"len": len(names)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- GetInArgsTypesNames", actual)
}

func Test_I14_MethodProcessor_NilReceiver_Invoke(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	_, err := mp.Invoke()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- Invoke", actual)
}

func Test_I14_MethodProcessor_NilReceiver_IsPublic(t *testing.T) {
	var mp *reflectmodel.MethodProcessor
	actual := args.Map{"isPublic": mp.IsPublicMethod(), "isPrivate": mp.IsPrivateMethod()}
	expected := args.Map{"isPublic": false, "isPrivate": false}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- IsPublicMethod/IsPrivateMethod", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsEqual — nil and same-pointer paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_IsEqual_BothNil(t *testing.T) {
	var a, b *reflectmodel.MethodProcessor
	actual := args.Map{"eq": a.IsEqual(b)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns nil -- both nil", actual)
}

func Test_I14_IsEqual_LeftNil(t *testing.T) {
	var a *reflectmodel.MethodProcessor
	b := newMethodProcessor("PublicMethod")
	actual := args.Map{"eq": a.IsEqual(b)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns nil -- left nil", actual)
}

func Test_I14_IsEqual_RightNil(t *testing.T) {
	a := newMethodProcessor("PublicMethod")
	actual := args.Map{"eq": a.IsEqual(nil)}
	expected := args.Map{"eq": false}
	expected.ShouldBeEqual(t, 0, "IsEqual returns nil -- right nil", actual)
}

func Test_I14_IsEqual_SamePointer(t *testing.T) {
	a := newMethodProcessor("PublicMethod")
	actual := args.Map{"eq": a.IsEqual(a)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- same pointer", actual)
}

func Test_I14_IsEqual_SameSignature(t *testing.T) {
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("PublicMethod")
	actual := args.Map{"eq": a.IsEqual(b)}
	expected := args.Map{"eq": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- same signature", actual)
}

func Test_I14_IsNotEqual_DiffSignature(t *testing.T) {
	a := newMethodProcessor("PublicMethod")
	b := newMethodProcessor("NoArgsMethod")
	actual := args.Map{"notEq": a.IsNotEqual(b)}
	expected := args.Map{"notEq": true}
	expected.ShouldBeEqual(t, 0, "IsNotEqual returns correct value -- different signature", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// InvokeFirstAndError — single return (error path)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_InvokeFirstAndError_SingleReturn(t *testing.T) {
	mp := newMethodProcessor("NoArgsMethod")
	_, _, procErr := mp.InvokeFirstAndError(sampleStruct{})
	actual := args.Map{"hasErr": procErr != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- single return", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// InvokeError — with non-error return type (should panic or error)
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_InvokeError_NoArgsError(t *testing.T) {
	mp := newMethodProcessor("NoArgsMethod")
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "InvokeError panics -- non-error return panics", actual)
	}()
	mp.InvokeError(sampleStruct{})
}

// ══════════════════════════════════════════════════════════════════════════════
// InvalidReflectValueKindModel — constructor
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_InvalidReflectValueKindModel(t *testing.T) {
	rvk := reflectmodel.InvalidReflectValueKindModel("test error msg")
	actual := args.Map{
		"isValid":   rvk.IsValid,
		"hasErr":    rvk.HasError(),
		"isInvalid": rvk.IsInvalid(),
		"emptyErr":  rvk.IsEmptyError(),
		"typeName":  rvk.TypeName(),
		"pkgPath":   rvk.PkgPath(),
	}
	expected := args.Map{
		"isValid":   false,
		"hasErr":    true,
		"isInvalid": true,
		"emptyErr":  false,
		"typeName":  "",
		"pkgPath":   "",
	}
	expected.ShouldBeEqual(t, 0, "InvalidReflectValueKindModel returns error -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueKind — nil receiver paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_RVK_NilReceiver_IsInvalid(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{"isInvalid": rvk.IsInvalid(), "hasErr": rvk.HasError(), "emptyErr": rvk.IsEmptyError()}
	expected := args.Map{"isInvalid": true, "hasErr": false, "emptyErr": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK IsInvalid/HasError/IsEmptyError", actual)
}

func Test_I14_RVK_NilReceiver_ActualInstance(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{"actNil": rvk.ActualInstance() == nil}
	expected := args.Map{"actNil": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK ActualInstance", actual)
}

func Test_I14_RVK_NilReceiver_PkgPath(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{"pkgPath": rvk.PkgPath()}
	expected := args.Map{"pkgPath": ""}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK PkgPath", actual)
}

func Test_I14_RVK_NilReceiver_PointerRv(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{"nil": rvk.PointerRv() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK PointerRv", actual)
}

func Test_I14_RVK_NilReceiver_PointerInterface(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{"nil": rvk.PointerInterface() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK PointerInterface", actual)
}

func Test_I14_RVK_NilReceiver_TypeName(t *testing.T) {
	var rvk *reflectmodel.ReflectValueKind
	actual := args.Map{"typeName": rvk.TypeName()}
	expected := args.Map{"typeName": ""}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- RVK TypeName", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectValueKind — invalid (IsValid=false) with non-nil receiver
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_RVK_InvalidNotNil_PointerRv(t *testing.T) {
	rvk := &reflectmodel.ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(42),
		Kind:            reflect.Int,
	}
	ptr := rvk.PointerRv()
	actual := args.Map{"notNil": ptr != nil, "pkgPath": rvk.PkgPath(), "typeName": rvk.TypeName()}
	expected := args.Map{"notNil": true, "pkgPath": "", "typeName": ""}
	expected.ShouldBeEqual(t, 0, "RVK returns nil -- invalid non-nil PointerRv", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// FieldProcessor — nil receiver paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_FieldProcessor_NilReceiver_IsFieldType(t *testing.T) {
	var fp *reflectmodel.FieldProcessor
	actual := args.Map{"isType": fp.IsFieldType(reflect.TypeOf(0))}
	expected := args.Map{"isType": false}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- FieldProcessor IsFieldType", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// argsCountMismatchErrorMessage — triggered via ValidateMethodArgs
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_ValidateMethodArgs_TooFewArgs(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")
	// PublicMethod expects (sampleStruct, string, int) = 3 args; give 1
	err := mp.ValidateMethodArgs([]any{sampleStruct{}})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns non-empty -- too few args", actual)
}

func Test_I14_ValidateMethodArgs_TooManyArgs(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")
	err := mp.ValidateMethodArgs([]any{sampleStruct{}, "a", 1, "extra", "extra2"})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns non-empty -- too many args", actual)
}

func Test_I14_ValidateMethodArgs_EmptyArgs(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")
	err := mp.ValidateMethodArgs([]any{})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidateMethodArgs returns empty -- empty args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MultiReturn method — InvokeFirstAndError exercises multi-return parsing
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_InvokeFirstAndError_MultiReturn(t *testing.T) {
	mp := newMethodProcessor("MultiReturn")
	first, funcErr, procErr := mp.InvokeFirstAndError(sampleStruct{})
	actual := args.Map{"procErr": procErr == nil, "funcErr": funcErr == nil, "first": first}
	expected := args.Map{"procErr": true, "funcErr": true, "first": 0}
	expected.ShouldBeEqual(t, 0, "InvokeFirstAndError returns error -- MultiReturn", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// GetFirstResponseOfInvoke — success path
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_GetFirstResponseOfInvoke_Success(t *testing.T) {
	mp := newMethodProcessor("NoArgsMethod")
	first, err := mp.GetFirstResponseOfInvoke(sampleStruct{})
	actual := args.Map{"noErr": err == nil, "val": first}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "GetFirstResponseOfInvoke returns correct value -- success", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// InvokeResultOfIndex — success path index 1
// ══════════════════════════════════════════════════════════════════════════════

func Test_I14_InvokeResultOfIndex_SecondResult(t *testing.T) {
	mp := newMethodProcessor("PublicMethod")
	second, err := mp.InvokeResultOfIndex(1, sampleStruct{}, "test", 42)
	actual := args.Map{"noErr": err == nil, "nilErr": second == nil}
	expected := args.Map{"noErr": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "InvokeResultOfIndex returns correct value -- second result", actual)
}
