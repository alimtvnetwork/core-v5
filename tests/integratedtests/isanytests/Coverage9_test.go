package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── Null ──

func Test_Cov9_Null_NilInterface(t *testing.T) {
	actual := args.Map{"result": isany.Null(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- nil interface", actual)
}

func Test_Cov9_Null_NilSlice(t *testing.T) {
	var s []string
	actual := args.Map{"result": isany.Null(s)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- nil slice", actual)
}

func Test_Cov9_Null_NilMap(t *testing.T) {
	var m map[string]string
	actual := args.Map{"result": isany.Null(m)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- nil map", actual)
}

func Test_Cov9_Null_NilPtr(t *testing.T) {
	var p *int
	actual := args.Map{"result": isany.Null(p)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- nil ptr", actual)
}

func Test_Cov9_Null_NilFunc(t *testing.T) {
	var f func()
	actual := args.Map{"result": isany.Null(f)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- nil func", actual)
}

func Test_Cov9_Null_NonNilValue(t *testing.T) {
	actual := args.Map{"result": isany.Null(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- non-nil value", actual)
}

func Test_Cov9_Null_NonNilString(t *testing.T) {
	actual := args.Map{"result": isany.Null("hello")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- non-nil string", actual)
}

// ── NotNull ──

func Test_Cov9_NotNull_Nil(t *testing.T) {
	actual := args.Map{"result": isany.NotNull(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotNull returns nil -- nil", actual)
}

func Test_Cov9_NotNull_NonNil(t *testing.T) {
	actual := args.Map{"result": isany.NotNull(42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotNull returns nil -- non-nil", actual)
}

// ── Defined ──

func Test_Cov9_Defined_Nil(t *testing.T) {
	actual := args.Map{"result": isany.Defined(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Defined returns nil -- nil", actual)
}

func Test_Cov9_Defined_NonNil(t *testing.T) {
	actual := args.Map{"result": isany.Defined("x")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Defined returns nil -- non-nil", actual)
}

// ── Zero ──

func Test_Cov9_Zero_ZeroInt(t *testing.T) {
	actual := args.Map{"result": isany.Zero(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero returns correct value -- zero int", actual)
}

func Test_Cov9_Zero_NonZero(t *testing.T) {
	actual := args.Map{"result": isany.Zero(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Zero returns non-empty -- non-zero", actual)
}

func Test_Cov9_Zero_EmptyString(t *testing.T) {
	actual := args.Map{"result": isany.Zero("")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero returns empty -- empty string", actual)
}

// ── AllNull ──

func Test_Cov9_AllNull_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AllNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull returns empty -- empty", actual)
}

func Test_Cov9_AllNull_AllNil(t *testing.T) {
	actual := args.Map{"result": isany.AllNull(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull returns nil -- all nil", actual)
}

func Test_Cov9_AllNull_Mixed(t *testing.T) {
	actual := args.Map{"result": isany.AllNull(nil, "a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNull returns correct value -- mixed", actual)
}

// ── AnyNull ──

func Test_Cov9_AnyNull_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull returns empty -- empty", actual)
}

func Test_Cov9_AnyNull_HasNil(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull("a", nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyNull returns nil -- has nil", actual)
}

func Test_Cov9_AnyNull_NoNil(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull("a", "b")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull returns nil -- no nil", actual)
}

// ── AllZero ──

func Test_Cov9_AllZero_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AllZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero returns empty -- empty", actual)
}

func Test_Cov9_AllZero_AllZeros(t *testing.T) {
	actual := args.Map{"result": isany.AllZero(0, "", false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero returns correct value -- all zeros", actual)
}

func Test_Cov9_AllZero_Mixed(t *testing.T) {
	actual := args.Map{"result": isany.AllZero(0, "x")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllZero returns correct value -- mixed", actual)
}

// ── AnyZero ──

func Test_Cov9_AnyZero_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero returns empty -- empty", actual)
}

func Test_Cov9_AnyZero_HasZero(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero("x", 0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero returns correct value -- has zero", actual)
}

func Test_Cov9_AnyZero_NoZero(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero("x", 1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyZero returns empty -- no zero", actual)
}

// ── DefinedBoth ──

func Test_Cov9_DefinedBoth_BothDefined(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth("a", "b")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedBoth returns correct value -- both defined", actual)
}

func Test_Cov9_DefinedBoth_OneNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth("a", nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth returns nil -- one nil", actual)
}

func Test_Cov9_DefinedBoth_BothNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth(nil, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth returns nil -- both nil", actual)
}

// ── NullBoth ──

func Test_Cov9_NullBoth_BothNil(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NullBoth returns nil -- both nil", actual)
}

func Test_Cov9_NullBoth_OneNil(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth(nil, "a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth returns nil -- one nil", actual)
}

func Test_Cov9_NullBoth_BothDefined(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth("a", "b")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth returns correct value -- both defined", actual)
}

// ── DefinedAllOf ──

func Test_Cov9_DefinedAllOf_Empty(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAllOf()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf returns empty -- empty", actual)
}

func Test_Cov9_DefinedAllOf_AllDefined(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAllOf("a", 1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf returns correct value -- all defined", actual)
}

func Test_Cov9_DefinedAllOf_HasNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAllOf("a", nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf returns nil -- has nil", actual)
}

// ── DefinedAnyOf ──

func Test_Cov9_DefinedAnyOf_Empty(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAnyOf()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf returns empty -- empty", actual)
}

func Test_Cov9_DefinedAnyOf_HasDefined(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAnyOf(nil, "a")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf returns correct value -- has defined", actual)
}

func Test_Cov9_DefinedAnyOf_AllNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAnyOf(nil, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf returns nil -- all nil", actual)
}

// ── DeepEqual ──

func Test_Cov9_DeepEqual_SameInt(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqual(1, 1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqual returns correct value -- same int", actual)
}

func Test_Cov9_DeepEqual_DiffInt(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqual(1, 2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DeepEqual returns correct value -- diff int", actual)
}

// ── NotDeepEqual ──

func Test_Cov9_NotDeepEqual(t *testing.T) {
	actual := args.Map{"result": isany.NotDeepEqual(1, 2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotDeepEqual returns correct value -- with args", actual)
}

// ── Pointer ──

func Test_Cov9_Pointer_IsPointer(t *testing.T) {
	v := 42
	actual := args.Map{"result": isany.Pointer(&v)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Pointer returns correct value -- is pointer", actual)
}

func Test_Cov9_Pointer_NotPointer(t *testing.T) {
	actual := args.Map{"result": isany.Pointer(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Pointer returns correct value -- not pointer", actual)
}

// ── StringEqual ──

func Test_Cov9_StringEqual_Same(t *testing.T) {
	actual := args.Map{"result": isany.StringEqual("abc", "abc")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringEqual returns correct value -- same", actual)
}

func Test_Cov9_StringEqual_Different(t *testing.T) {
	actual := args.Map{"result": isany.StringEqual("abc", "xyz")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringEqual returns correct value -- different", actual)
}

// ── Conclusive ──

func Test_Cov9_Conclusive_BothNil(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(nil, nil)
	actual := args.Map{"isEqual": isEqual, "isConclusive": isConclusive}
	expected := args.Map{"isEqual": true, "isConclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive returns nil -- both nil", actual)
}

func Test_Cov9_Conclusive_LeftNil(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(nil, "a")
	actual := args.Map{"isEqual": isEqual, "isConclusive": isConclusive}
	expected := args.Map{"isEqual": false, "isConclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive returns nil -- left nil", actual)
}

func Test_Cov9_Conclusive_SameValue(t *testing.T) {
	v := 42
	isEqual, isConclusive := isany.Conclusive(v, v)
	actual := args.Map{"isEqual": isEqual, "isConclusive": isConclusive}
	expected := args.Map{"isEqual": true, "isConclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive returns correct value -- same value", actual)
}

func Test_Cov9_Conclusive_DifferentTypes(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(42, "42")
	actual := args.Map{"isEqual": isEqual, "isConclusive": isConclusive}
	expected := args.Map{"isEqual": false, "isConclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive returns correct value -- different types", actual)
}

func Test_Cov9_Conclusive_Inconclusive(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(1, 2)
	actual := args.Map{"isEqual": isEqual, "isConclusive": isConclusive}
	expected := args.Map{"isEqual": false, "isConclusive": false}
	expected.ShouldBeEqual(t, 0, "Conclusive returns correct value -- inconclusive", actual)
}
