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
	expected.ShouldBeEqual(t, 0, "Null nil interface", actual)
}

func Test_Cov9_Null_NilSlice(t *testing.T) {
	var s []string
	actual := args.Map{"result": isany.Null(s)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null nil slice", actual)
}

func Test_Cov9_Null_NilMap(t *testing.T) {
	var m map[string]string
	actual := args.Map{"result": isany.Null(m)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null nil map", actual)
}

func Test_Cov9_Null_NilPtr(t *testing.T) {
	var p *int
	actual := args.Map{"result": isany.Null(p)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null nil ptr", actual)
}

func Test_Cov9_Null_NilFunc(t *testing.T) {
	var f func()
	actual := args.Map{"result": isany.Null(f)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null nil func", actual)
}

func Test_Cov9_Null_NonNilValue(t *testing.T) {
	actual := args.Map{"result": isany.Null(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Null non-nil value", actual)
}

func Test_Cov9_Null_NonNilString(t *testing.T) {
	actual := args.Map{"result": isany.Null("hello")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Null non-nil string", actual)
}

// ── NotNull ──

func Test_Cov9_NotNull_Nil(t *testing.T) {
	actual := args.Map{"result": isany.NotNull(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotNull nil", actual)
}

func Test_Cov9_NotNull_NonNil(t *testing.T) {
	actual := args.Map{"result": isany.NotNull(42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotNull non-nil", actual)
}

// ── Defined ──

func Test_Cov9_Defined_Nil(t *testing.T) {
	actual := args.Map{"result": isany.Defined(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Defined nil", actual)
}

func Test_Cov9_Defined_NonNil(t *testing.T) {
	actual := args.Map{"result": isany.Defined("x")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Defined non-nil", actual)
}

// ── Zero ──

func Test_Cov9_Zero_ZeroInt(t *testing.T) {
	actual := args.Map{"result": isany.Zero(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero zero int", actual)
}

func Test_Cov9_Zero_NonZero(t *testing.T) {
	actual := args.Map{"result": isany.Zero(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Zero non-zero", actual)
}

func Test_Cov9_Zero_EmptyString(t *testing.T) {
	actual := args.Map{"result": isany.Zero("")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Zero empty string", actual)
}

// ── AllNull ──

func Test_Cov9_AllNull_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AllNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull empty", actual)
}

func Test_Cov9_AllNull_AllNil(t *testing.T) {
	actual := args.Map{"result": isany.AllNull(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull all nil", actual)
}

func Test_Cov9_AllNull_Mixed(t *testing.T) {
	actual := args.Map{"result": isany.AllNull(nil, "a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNull mixed", actual)
}

// ── AnyNull ──

func Test_Cov9_AnyNull_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull empty", actual)
}

func Test_Cov9_AnyNull_HasNil(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull("a", nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyNull has nil", actual)
}

func Test_Cov9_AnyNull_NoNil(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull("a", "b")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull no nil", actual)
}

// ── AllZero ──

func Test_Cov9_AllZero_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AllZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero empty", actual)
}

func Test_Cov9_AllZero_AllZeros(t *testing.T) {
	actual := args.Map{"result": isany.AllZero(0, "", false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero all zeros", actual)
}

func Test_Cov9_AllZero_Mixed(t *testing.T) {
	actual := args.Map{"result": isany.AllZero(0, "x")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllZero mixed", actual)
}

// ── AnyZero ──

func Test_Cov9_AnyZero_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero empty", actual)
}

func Test_Cov9_AnyZero_HasZero(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero("x", 0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero has zero", actual)
}

func Test_Cov9_AnyZero_NoZero(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero("x", 1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyZero no zero", actual)
}

// ── DefinedBoth ──

func Test_Cov9_DefinedBoth_BothDefined(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth("a", "b")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedBoth both defined", actual)
}

func Test_Cov9_DefinedBoth_OneNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth("a", nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth one nil", actual)
}

func Test_Cov9_DefinedBoth_BothNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth(nil, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth both nil", actual)
}

// ── NullBoth ──

func Test_Cov9_NullBoth_BothNil(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NullBoth both nil", actual)
}

func Test_Cov9_NullBoth_OneNil(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth(nil, "a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth one nil", actual)
}

func Test_Cov9_NullBoth_BothDefined(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth("a", "b")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth both defined", actual)
}

// ── DefinedAllOf ──

func Test_Cov9_DefinedAllOf_Empty(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAllOf()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf empty", actual)
}

func Test_Cov9_DefinedAllOf_AllDefined(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAllOf("a", 1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf all defined", actual)
}

func Test_Cov9_DefinedAllOf_HasNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAllOf("a", nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf has nil", actual)
}

// ── DefinedAnyOf ──

func Test_Cov9_DefinedAnyOf_Empty(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAnyOf()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf empty", actual)
}

func Test_Cov9_DefinedAnyOf_HasDefined(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAnyOf(nil, "a")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf has defined", actual)
}

func Test_Cov9_DefinedAnyOf_AllNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAnyOf(nil, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf all nil", actual)
}

// ── DeepEqual ──

func Test_Cov9_DeepEqual_SameInt(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqual(1, 1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqual same int", actual)
}

func Test_Cov9_DeepEqual_DiffInt(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqual(1, 2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DeepEqual diff int", actual)
}

// ── NotDeepEqual ──

func Test_Cov9_NotDeepEqual(t *testing.T) {
	actual := args.Map{"result": isany.NotDeepEqual(1, 2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotDeepEqual", actual)
}

// ── Pointer ──

func Test_Cov9_Pointer_IsPointer(t *testing.T) {
	v := 42
	actual := args.Map{"result": isany.Pointer(&v)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Pointer is pointer", actual)
}

func Test_Cov9_Pointer_NotPointer(t *testing.T) {
	actual := args.Map{"result": isany.Pointer(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Pointer not pointer", actual)
}

// ── StringEqual ──

func Test_Cov9_StringEqual_Same(t *testing.T) {
	actual := args.Map{"result": isany.StringEqual("abc", "abc")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringEqual same", actual)
}

func Test_Cov9_StringEqual_Different(t *testing.T) {
	actual := args.Map{"result": isany.StringEqual("abc", "xyz")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringEqual different", actual)
}

// ── Conclusive ──

func Test_Cov9_Conclusive_BothNil(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(nil, nil)
	actual := args.Map{"isEqual": isEqual, "isConclusive": isConclusive}
	expected := args.Map{"isEqual": true, "isConclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive both nil", actual)
}

func Test_Cov9_Conclusive_LeftNil(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(nil, "a")
	actual := args.Map{"isEqual": isEqual, "isConclusive": isConclusive}
	expected := args.Map{"isEqual": false, "isConclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive left nil", actual)
}

func Test_Cov9_Conclusive_SameValue(t *testing.T) {
	v := 42
	isEqual, isConclusive := isany.Conclusive(v, v)
	actual := args.Map{"isEqual": isEqual, "isConclusive": isConclusive}
	expected := args.Map{"isEqual": true, "isConclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive same value", actual)
}

func Test_Cov9_Conclusive_DifferentTypes(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(42, "42")
	actual := args.Map{"isEqual": isEqual, "isConclusive": isConclusive}
	expected := args.Map{"isEqual": false, "isConclusive": true}
	expected.ShouldBeEqual(t, 0, "Conclusive different types", actual)
}

func Test_Cov9_Conclusive_Inconclusive(t *testing.T) {
	isEqual, isConclusive := isany.Conclusive(1, 2)
	actual := args.Map{"isEqual": isEqual, "isConclusive": isConclusive}
	expected := args.Map{"isEqual": false, "isConclusive": false}
	expected.ShouldBeEqual(t, 0, "Conclusive inconclusive", actual)
}
