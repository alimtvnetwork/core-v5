package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// ReflectSetFromTo
// =============================================================================

func Test_Cov48_ReflectSetFromTo_BothNil(t *testing.T) {
	err := coredynamic.ReflectSetFromTo(nil, nil)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo both nil", actual)
}

func Test_Cov48_ReflectSetFromTo_ToNil(t *testing.T) {
	err := coredynamic.ReflectSetFromTo("hello", nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo to nil", actual)
}

func Test_Cov48_ReflectSetFromTo_ToNotPointer(t *testing.T) {
	err := coredynamic.ReflectSetFromTo("hello", "world")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo to not pointer", actual)
}

func Test_Cov48_ReflectSetFromTo_SameNonPointerToPointer(t *testing.T) {
	var dest string
	err := coredynamic.ReflectSetFromTo("hello", &dest)
	actual := args.Map{"noErr": err == nil, "dest": dest}
	expected := args.Map{"noErr": true, "dest": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo non-ptr to ptr", actual)
}

func Test_Cov48_ReflectSetFromTo_SamePointerTypes(t *testing.T) {
	src := "hello"
	var dest string
	err := coredynamic.ReflectSetFromTo(&src, &dest)
	actual := args.Map{"noErr": err == nil, "dest": dest}
	expected := args.Map{"noErr": true, "dest": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo same ptr types", actual)
}

func Test_Cov48_ReflectSetFromTo_BytesToStruct(t *testing.T) {
	type Simple struct{ Name string }
	b := []byte(`{"Name":"test"}`)
	var dest Simple
	err := coredynamic.ReflectSetFromTo(b, &dest)
	actual := args.Map{"noErr": err == nil, "name": dest.Name}
	expected := args.Map{"noErr": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo bytes to struct", actual)
}

func Test_Cov48_ReflectSetFromTo_StructToBytes(t *testing.T) {
	type Simple struct{ Name string }
	src := Simple{Name: "test"}
	var dest []byte
	err := coredynamic.ReflectSetFromTo(src, &dest)
	actual := args.Map{"noErr": err == nil, "hasBytes": len(dest) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo struct to bytes", actual)
}

func Test_Cov48_ReflectSetFromTo_TypeMismatch(t *testing.T) {
	var dest int
	err := coredynamic.ReflectSetFromTo("hello", &dest)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo type mismatch", actual)
}

// =============================================================================
// ReflectTypeValidation
// =============================================================================

func Test_Cov48_ReflectTypeValidation_NilNotExpected(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation nil not expected", actual)
}

func Test_Cov48_ReflectTypeValidation_Match(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation match", actual)
}

func Test_Cov48_ReflectTypeValidation_Mismatch(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), 42)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation mismatch", actual)
}

// =============================================================================
// ReflectKindValidation
// =============================================================================

func Test_Cov48_ReflectKindValidation_Match(t *testing.T) {
	err := coredynamic.ReflectKindValidation(reflect.String, "hello")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation match", actual)
}

func Test_Cov48_ReflectKindValidation_Mismatch(t *testing.T) {
	err := coredynamic.ReflectKindValidation(reflect.Int, "hello")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation mismatch", actual)
}

// =============================================================================
// ReflectInterfaceVal
// =============================================================================

func Test_Cov48_ReflectInterfaceVal_Value(t *testing.T) {
	r := coredynamic.ReflectInterfaceVal("hello")
	actual := args.Map{"r": r}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal value", actual)
}

func Test_Cov48_ReflectInterfaceVal_Pointer(t *testing.T) {
	s := "hello"
	r := coredynamic.ReflectInterfaceVal(&s)
	actual := args.Map{"r": r}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal pointer", actual)
}

// =============================================================================
// PointerOrNonPointer
// =============================================================================

func Test_Cov48_PointerOrNonPointer_NonPointerOutput(t *testing.T) {
	s := "hello"
	out, _ := coredynamic.PointerOrNonPointer(false, &s)
	actual := args.Map{"r": out}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer non-ptr output", actual)
}

func Test_Cov48_PointerOrNonPointer_ValuePassthrough(t *testing.T) {
	out, _ := coredynamic.PointerOrNonPointer(false, "hello")
	actual := args.Map{"r": out}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer value passthrough", actual)
}

// =============================================================================
// AnyToReflectVal
// =============================================================================

func Test_Cov48_AnyToReflectVal(t *testing.T) {
	rv := coredynamic.AnyToReflectVal("hello")
	actual := args.Map{"valid": rv.IsValid(), "kind": rv.Kind().String()}
	expected := args.Map{"valid": true, "kind": "string"}
	expected.ShouldBeEqual(t, 0, "AnyToReflectVal", actual)
}

// =============================================================================
// CastTo
// =============================================================================

func Test_Cov48_CastTo_Matching(t *testing.T) {
	r := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))
	actual := args.Map{"valid": r.IsValid, "matching": r.IsMatchingAcceptedType}
	expected := args.Map{"valid": true, "matching": true}
	expected.ShouldBeEqual(t, 0, "CastTo matching", actual)
}

func Test_Cov48_CastTo_NotMatching(t *testing.T) {
	r := coredynamic.CastTo(false, "hello", reflect.TypeOf(0))
	actual := args.Map{"matching": r.IsMatchingAcceptedType, "hasErr": r.HasError()}
	expected := args.Map{"matching": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "CastTo not matching", actual)
}

// =============================================================================
// NotAcceptedTypesErr / MustBeAcceptedTypes
// =============================================================================

func Test_Cov48_NotAcceptedTypesErr_Accepted(t *testing.T) {
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(""))
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr accepted", actual)
}

func Test_Cov48_NotAcceptedTypesErr_NotAccepted(t *testing.T) {
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(0))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr not accepted", actual)
}

func Test_Cov48_MustBeAcceptedTypes_Valid(t *testing.T) {
	coredynamic.MustBeAcceptedTypes("hello", reflect.TypeOf(""))
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeAcceptedTypes valid", actual)
}

func Test_Cov48_MustBeAcceptedTypes_Panics(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "MustBeAcceptedTypes panics", actual)
	}()
	coredynamic.MustBeAcceptedTypes("hello", reflect.TypeOf(0))
}

// =============================================================================
// TypeNotEqualErr / TypeMustBeSame
// =============================================================================

func Test_Cov48_TypeNotEqualErr_Same(t *testing.T) {
	err := coredynamic.TypeNotEqualErr("a", "b")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr same", actual)
}

func Test_Cov48_TypeNotEqualErr_Different(t *testing.T) {
	err := coredynamic.TypeNotEqualErr("a", 42)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr different", actual)
}

func Test_Cov48_TypeMustBeSame_Same(t *testing.T) {
	coredynamic.TypeMustBeSame("a", "b")
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeMustBeSame same", actual)
}

func Test_Cov48_TypeMustBeSame_Panics(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeMustBeSame panics", actual)
	}()
	coredynamic.TypeMustBeSame("a", 42)
}

// =============================================================================
// TypesIndexOf
// =============================================================================

func Test_Cov48_TypesIndexOf_Found(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	actual := args.Map{"r": coredynamic.TypesIndexOf(strType, intType, strType)}
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf found", actual)
}

func Test_Cov48_TypesIndexOf_NotFound(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	actual := args.Map{"r": coredynamic.TypesIndexOf(strType, intType)}
	expected := args.Map{"r": -1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf not found", actual)
}

// =============================================================================
// Type
// =============================================================================

func Test_Cov48_Type(t *testing.T) {
	rt := coredynamic.Type("hello")
	actual := args.Map{"name": rt.Name()}
	expected := args.Map{"name": "string"}
	expected.ShouldBeEqual(t, 0, "Type", actual)
}

// =============================================================================
// ZeroSet / ZeroSetAny / SafeZeroSet
// =============================================================================

func Test_Cov48_ZeroSet(t *testing.T) {
	type S struct{ Name string }
	s := S{Name: "hello"}
	coredynamic.ZeroSet(reflect.ValueOf(&s))
	actual := args.Map{"name": s.Name}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "ZeroSet", actual)
}

func Test_Cov48_ZeroSetAny_Valid(t *testing.T) {
	type S struct{ Name string }
	s := S{Name: "hello"}
	coredynamic.ZeroSetAny(&s)
	actual := args.Map{"name": s.Name}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny valid", actual)
}

func Test_Cov48_ZeroSetAny_Nil(t *testing.T) {
	coredynamic.ZeroSetAny(nil) // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny nil", actual)
}

func Test_Cov48_SafeZeroSet(t *testing.T) {
	type S struct{ Name string }
	s := S{Name: "hello"}
	coredynamic.SafeZeroSet(reflect.ValueOf(&s))
	actual := args.Map{"name": s.Name}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet", actual)
}

// =============================================================================
// LengthOfReflect — pointer case
// =============================================================================

func Test_Cov48_LengthOfReflect_Ptr(t *testing.T) {
	s := []int{1, 2, 3}
	rv := reflect.ValueOf(&s)
	actual := args.Map{"r": coredynamic.LengthOfReflect(rv)}
	expected := args.Map{"r": 3}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect ptr", actual)
}
