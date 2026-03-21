package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ============================================================================
// AllNull
// ============================================================================

func Test_Cov5_AllNull_AllNil(t *testing.T) {
	actual := args.Map{"result": isany.AllNull(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull returns nil -- all nil", actual)
}

func Test_Cov5_AllNull_Mixed(t *testing.T) {
	actual := args.Map{"result": isany.AllNull(nil, "hello")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNull returns correct value -- mixed", actual)
}

func Test_Cov5_AllNull_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AllNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull returns empty -- empty", actual)
}

// ============================================================================
// AnyNull
// ============================================================================

func Test_Cov5_AnyNull_HasNil(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull("a", nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyNull returns nil -- has nil", actual)
}

func Test_Cov5_AnyNull_NoNil(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull("a", "b")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull returns nil -- no nil", actual)
}

func Test_Cov5_AnyNull_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull returns empty -- empty", actual)
}

// ============================================================================
// AllZero
// ============================================================================

func Test_Cov5_AllZero_AllZero(t *testing.T) {
	actual := args.Map{"result": isany.AllZero(0, "", false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero returns correct value -- all zero", actual)
}

func Test_Cov5_AllZero_Mixed(t *testing.T) {
	actual := args.Map{"result": isany.AllZero(0, 1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllZero returns correct value -- mixed", actual)
}

func Test_Cov5_AllZero_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AllZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllZero returns empty -- empty", actual)
}

// ============================================================================
// AnyZero
// ============================================================================

func Test_Cov5_AnyZero_HasZero(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero(1, 0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero returns correct value -- has zero", actual)
}

func Test_Cov5_AnyZero_NoZero(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero(1, 2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyZero returns empty -- no zero", actual)
}

func Test_Cov5_AnyZero_Empty(t *testing.T) {
	actual := args.Map{"result": isany.AnyZero()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyZero returns empty -- empty", actual)
}

// ============================================================================
// DefinedBoth
// ============================================================================

func Test_Cov5_DefinedBoth_BothDefined(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth("a", "b")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedBoth returns correct value -- both defined", actual)
}

func Test_Cov5_DefinedBoth_OneNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth("a", nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth returns nil -- one nil", actual)
}

func Test_Cov5_DefinedBoth_BothNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth(nil, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth returns nil -- both nil", actual)
}

// ============================================================================
// NotDeepEqual
// ============================================================================

func Test_Cov5_NotDeepEqual_Same(t *testing.T) {
	actual := args.Map{"result": isany.NotDeepEqual(42, 42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotDeepEqual returns correct value -- same", actual)
}

func Test_Cov5_NotDeepEqual_Diff(t *testing.T) {
	actual := args.Map{"result": isany.NotDeepEqual(42, 43)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotDeepEqual returns correct value -- diff", actual)
}

// ============================================================================
// DeepEqualAllItems
// ============================================================================

func Test_Cov5_DeepEqualAllItems_AllSame(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 1, 1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems returns correct value -- all same", actual)
}

func Test_Cov5_DeepEqualAllItems_OneDiff(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqualAllItems(1, 1, 2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems returns correct value -- one diff", actual)
}

func Test_Cov5_DeepEqualAllItems_Single(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqualAllItems(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems returns correct value -- single", actual)
}

func Test_Cov5_DeepEqualAllItems_Empty(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqualAllItems()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems returns empty -- empty", actual)
}

func Test_Cov5_DeepEqualAllItems_Two(t *testing.T) {
	actual := args.Map{"result": isany.DeepEqualAllItems("a", "a")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DeepEqualAllItems returns correct value -- two same", actual)
}

// ============================================================================
// NullBoth
// ============================================================================

func Test_Cov5_NullBoth_BothNil(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NullBoth returns nil -- both nil", actual)
}

func Test_Cov5_NullBoth_OneNil(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth(nil, "a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth returns nil -- one nil", actual)
}

// ============================================================================
// NullLeftRight
// ============================================================================

func Test_Cov5_NullLeftRight(t *testing.T) {
	l, r := isany.NullLeftRight(nil, "a")
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": true, "right": false}
	expected.ShouldBeEqual(t, 0, "NullLeftRight returns nil -- nil and string", actual)
}

// ============================================================================
// NotNull
// ============================================================================

func Test_Cov5_NotNull_Nil(t *testing.T) {
	actual := args.Map{"result": isany.NotNull(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotNull returns nil -- nil", actual)
}

func Test_Cov5_NotNull_Value(t *testing.T) {
	actual := args.Map{"result": isany.NotNull("a")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotNull returns correct value -- value", actual)
}

// ============================================================================
// StringEqual
// ============================================================================

func Test_Cov5_StringEqual_Same(t *testing.T) {
	actual := args.Map{"result": isany.StringEqual(42, 42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StringEqual returns non-empty -- same values", actual)
}

func Test_Cov5_StringEqual_Diff(t *testing.T) {
	actual := args.Map{"result": isany.StringEqual(42, 43)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringEqual returns non-empty -- diff values", actual)
}

// ============================================================================
// Function
// ============================================================================

func Test_Cov5_Function_Func(t *testing.T) {
	isFunc, name := isany.Function(func() {})
	actual := args.Map{"isFunc": isFunc, "hasName": name != ""}
	expected := args.Map{"isFunc": true, "hasName": true}
	expected.ShouldBeEqual(t, 0, "Function returns correct value -- func", actual)
}

func Test_Cov5_Function_NotFunc(t *testing.T) {
	isFunc, name := isany.Function("hello")
	actual := args.Map{"isFunc": isFunc, "name": name}
	expected := args.Map{"isFunc": false, "name": ""}
	expected.ShouldBeEqual(t, 0, "Function returns correct value -- not func", actual)
}

func Test_Cov5_Function_Nil(t *testing.T) {
	isFunc, name := isany.Function(nil)
	actual := args.Map{"isFunc": isFunc, "name": name}
	expected := args.Map{"isFunc": false, "name": ""}
	expected.ShouldBeEqual(t, 0, "Function returns nil -- nil", actual)
}

// ============================================================================
// Pointer
// ============================================================================

func Test_Cov5_Pointer_Ptr(t *testing.T) {
	s := "hello"
	actual := args.Map{"result": isany.Pointer(&s)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Pointer returns correct value -- ptr", actual)
}

func Test_Cov5_Pointer_NotPtr(t *testing.T) {
	actual := args.Map{"result": isany.Pointer("hello")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Pointer returns correct value -- not ptr", actual)
}

// ============================================================================
// Null — typed nil channels, maps, slices
// ============================================================================

func Test_Cov5_Null_NilSlice(t *testing.T) {
	var s []string
	actual := args.Map{"result": isany.Null(s)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- nil slice", actual)
}

func Test_Cov5_Null_NilMap(t *testing.T) {
	var m map[string]int
	actual := args.Map{"result": isany.Null(m)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- nil map", actual)
}

func Test_Cov5_Null_NilFunc(t *testing.T) {
	var f func()
	actual := args.Map{"result": isany.Null(f)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- nil func", actual)
}

func Test_Cov5_Null_NonNilSlice(t *testing.T) {
	actual := args.Map{"result": isany.Null([]string{})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Null returns nil -- non-nil slice", actual)
}
