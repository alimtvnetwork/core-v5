package isanytests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

// ── ReflectNull — reflect.Value input ──

func Test_Cov8_ReflectNull_WithReflectValue(t *testing.T) {
	rv := reflect.ValueOf(42)
	actual := args.Map{"result": isany.ReflectNull(rv)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull with reflect.Value -- false", actual)
}

func Test_Cov8_ReflectNull_WithInvalidReflectValue(t *testing.T) {
	rv := reflect.Value{}
	actual := args.Map{"result": isany.ReflectNull(rv)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull with invalid reflect.Value -- true", actual)
}

func Test_Cov8_ReflectNull_NilMap(t *testing.T) {
	var m map[string]string
	actual := args.Map{"result": isany.ReflectNull(m)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectNull nil map -- true", actual)
}

func Test_Cov8_ReflectNull_NonNilValue(t *testing.T) {
	actual := args.Map{"result": isany.ReflectNull(42)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectNull non-nil value -- false", actual)
}

// ── ReflectValueNull ──

func Test_Cov8_ReflectValueNull_InvalidValue(t *testing.T) {
	actual := args.Map{"result": isany.ReflectValueNull(reflect.Value{})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull invalid -- true", actual)
}

func Test_Cov8_ReflectValueNull_NilSlice(t *testing.T) {
	var s []int
	actual := args.Map{"result": isany.ReflectValueNull(reflect.ValueOf(s))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull nil slice -- true", actual)
}

func Test_Cov8_ReflectValueNull_NonNilInt(t *testing.T) {
	actual := args.Map{"result": isany.ReflectValueNull(reflect.ValueOf(42))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReflectValueNull non-nil int -- false", actual)
}

// ── NullBoth ──

func Test_Cov8_NullBoth_BothNil(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NullBoth both nil -- true", actual)
}

func Test_Cov8_NullBoth_OneDefined(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth(nil, "x")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth one defined -- false", actual)
}

func Test_Cov8_NullBoth_BothDefined(t *testing.T) {
	actual := args.Map{"result": isany.NullBoth("a", "b")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NullBoth both defined -- false", actual)
}

// ── DefinedBoth ──

func Test_Cov8_DefinedBoth_BothDefined(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth("a", "b")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedBoth both defined -- true", actual)
}

func Test_Cov8_DefinedBoth_OneNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth("a", nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth one nil -- false", actual)
}

func Test_Cov8_DefinedBoth_BothNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedBoth(nil, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedBoth both nil -- false", actual)
}

// ── DefinedAllOf ──

func Test_Cov8_DefinedAllOf_Empty(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAllOf()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf empty -- false", actual)
}

func Test_Cov8_DefinedAllOf_AllDefined(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAllOf("a", 1, true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf all defined -- true", actual)
}

func Test_Cov8_DefinedAllOf_OneNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAllOf("a", nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAllOf one nil -- false", actual)
}

// ── DefinedAnyOf ──

func Test_Cov8_DefinedAnyOf_Empty(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAnyOf()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf empty -- false", actual)
}

func Test_Cov8_DefinedAnyOf_OneDefined(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAnyOf(nil, "a")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf one defined -- true", actual)
}

func Test_Cov8_DefinedAnyOf_AllNil(t *testing.T) {
	actual := args.Map{"result": isany.DefinedAnyOf(nil, nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DefinedAnyOf all nil -- false", actual)
}

// ── AllNull ──

func Test_Cov8_AllNull_AllNil(t *testing.T) {
	actual := args.Map{"result": isany.AllNull(nil, nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllNull all nil -- true", actual)
}

func Test_Cov8_AllNull_OneDefined(t *testing.T) {
	actual := args.Map{"result": isany.AllNull(nil, "a")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNull one defined -- false", actual)
}

// ── AnyNull ──

func Test_Cov8_AnyNull_OneNil(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull("a", nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyNull one nil -- true", actual)
}

func Test_Cov8_AnyNull_NoneNil(t *testing.T) {
	actual := args.Map{"result": isany.AnyNull("a", "b")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyNull none nil -- false", actual)
}

// ── NotDeepEqual ──

func Test_Cov8_NotDeepEqual(t *testing.T) {
	actual := args.Map{
		"diff": isany.NotDeepEqual(1, 2),
		"same": isany.NotDeepEqual(1, 1),
	}
	expected := args.Map{"diff": true, "same": false}
	expected.ShouldBeEqual(t, 0, "NotDeepEqual returns correct value -- with args", actual)
}

// ── NullLeftRight ──

func Test_Cov8_NullLeftRight_BothNil(t *testing.T) {
	l, r := isany.NullLeftRight(nil, nil)
	actual := args.Map{"left": l, "right": r}
	expected := args.Map{"left": true, "right": true}
	expected.ShouldBeEqual(t, 0, "NullLeftRight returns nil -- both nil", actual)
}

// ── Null with typed nil variants ──

func Test_Cov8_Null_NilChannel(t *testing.T) {
	var ch chan int
	actual := args.Map{"result": isany.Null(ch)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null nil channel -- true", actual)
}

func Test_Cov8_Null_NilFunc(t *testing.T) {
	var fn func()
	actual := args.Map{"result": isany.Null(fn)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Null nil func -- true", actual)
}
