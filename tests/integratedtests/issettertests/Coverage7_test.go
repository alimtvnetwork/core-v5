package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ── New ──

func Test_Cov7_New_True(t *testing.T) {
	v := issetter.New(true, 42)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": true, "value": 42}
	expected.ShouldBeEqual(t, 0, "New true", actual)
}

func Test_Cov7_New_False(t *testing.T) {
	v := issetter.New(false, 42)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": false, "value": 42}
	expected.ShouldBeEqual(t, 0, "New false", actual)
}

// ── NewBool ──

func Test_Cov7_NewBool_True(t *testing.T) {
	v := issetter.NewBool(true)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": true, "value": true}
	expected.ShouldBeEqual(t, 0, "NewBool true", actual)
}

func Test_Cov7_NewBool_False(t *testing.T) {
	v := issetter.NewBool(false)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": false, "value": false}
	expected.ShouldBeEqual(t, 0, "NewBool false", actual)
}

// ── NewMust ──

func Test_Cov7_NewMust(t *testing.T) {
	v := issetter.NewMust(99)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": true, "value": 99}
	expected.ShouldBeEqual(t, 0, "NewMust", actual)
}

// ── NewBooleans ──

func Test_Cov7_NewBooleans(t *testing.T) {
	v := issetter.NewBooleans(true, false)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": true, "value": false}
	expected.ShouldBeEqual(t, 0, "NewBooleans", actual)
}

// ── GetSet ──

func Test_Cov7_GetSet_Set(t *testing.T) {
	v := issetter.New(true, 42)
	actual := args.Map{"isSet": v.IsSet}
	expected := args.Map{"isSet": true}
	expected.ShouldBeEqual(t, 0, "GetSet set", actual)
}

// ── GetBool ──

func Test_Cov7_GetBool_True(t *testing.T) {
	v := issetter.NewBool(true)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": true, "value": true}
	expected.ShouldBeEqual(t, 0, "GetBool true", actual)
}

// ── CombinedBooleans ──

func Test_Cov7_CombinedBooleans_BothTrue(t *testing.T) {
	v := issetter.CombinedBooleans(true, true)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": true, "value": true}
	expected.ShouldBeEqual(t, 0, "CombinedBooleans both true", actual)
}

func Test_Cov7_CombinedBooleans_Mixed(t *testing.T) {
	v := issetter.CombinedBooleans(true, false)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": true, "value": false}
	expected.ShouldBeEqual(t, 0, "CombinedBooleans mixed", actual)
}

// ── GetSetByte ──

func Test_Cov7_GetSetByte_Set(t *testing.T) {
	v := issetter.New(true, byte(10))
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": true, "value": byte(10)}
	expected.ShouldBeEqual(t, 0, "GetSetByte set", actual)
}

// ── Value ──

func Test_Cov7_Value_IsSet(t *testing.T) {
	v := issetter.New(true, 100)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": true, "value": 100}
	expected.ShouldBeEqual(t, 0, "Value isSet", actual)
}

func Test_Cov7_Value_NotSet(t *testing.T) {
	v := issetter.New(false, 0)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": false, "value": 0}
	expected.ShouldBeEqual(t, 0, "Value not set", actual)
}

// ── GetSetterByComparing ──

func Test_Cov7_GetSetterByComparing_Match(t *testing.T) {
	v := issetter.GetSetterByComparing(10, 10, 99)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": true, "value": 10}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing match", actual)
}

func Test_Cov7_GetSetterByComparing_NoMatch(t *testing.T) {
	v := issetter.GetSetterByComparing(10, 20, 99)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": false, "value": 99}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing no match", actual)
}

// ── GetSetUnset ──

func Test_Cov7_GetSetUnset_Set(t *testing.T) {
	v := issetter.GetSetUnset(true, 10, 20)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": true, "value": 10}
	expected.ShouldBeEqual(t, 0, "GetSetUnset set", actual)
}

func Test_Cov7_GetSetUnset_Unset(t *testing.T) {
	v := issetter.GetSetUnset(false, 10, 20)
	actual := args.Map{"isSet": v.IsSet, "value": v.Value}
	expected := args.Map{"isSet": false, "value": 20}
	expected.ShouldBeEqual(t, 0, "GetSetUnset unset", actual)
}
