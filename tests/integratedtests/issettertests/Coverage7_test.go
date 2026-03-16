package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ── New ──

func Test_Cov7_New_ValidName(t *testing.T) {
	v, err := issetter.New("True")
	actual := args.Map{"val": v, "isNilErr": err == nil}
	expected := args.Map{"val": issetter.True, "isNilErr": true}
	expected.ShouldBeEqual(t, 0, "New valid name", actual)
}

func Test_Cov7_New_InvalidName(t *testing.T) {
	_, err := issetter.New("bogus")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "New invalid name", actual)
}

// ── NewBool ──

func Test_Cov7_NewBool_True(t *testing.T) {
	v := issetter.NewBool(true)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "NewBool true", actual)
}

func Test_Cov7_NewBool_False(t *testing.T) {
	v := issetter.NewBool(false)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "NewBool false", actual)
}

// ── NewBooleans ──

func Test_Cov7_NewBooleans_AllTrue(t *testing.T) {
	v := issetter.NewBooleans(true, true)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "NewBooleans all true", actual)
}

func Test_Cov7_NewBooleans_AnyFalse(t *testing.T) {
	v := issetter.NewBooleans(true, false)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "NewBooleans any false", actual)
}

// ── CombinedBooleans ──

func Test_Cov7_CombinedBooleans_AllTrue(t *testing.T) {
	v := issetter.CombinedBooleans(true, true, true)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "CombinedBooleans all true", actual)
}

func Test_Cov7_CombinedBooleans_HasFalse(t *testing.T) {
	v := issetter.CombinedBooleans(true, false, true)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "CombinedBooleans has false", actual)
}

// ── GetBool ──

func Test_Cov7_GetBool_True(t *testing.T) {
	v := issetter.GetBool(true)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "GetBool true", actual)
}

func Test_Cov7_GetBool_False(t *testing.T) {
	v := issetter.GetBool(false)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "GetBool false", actual)
}

// ── GetSet ──

func Test_Cov7_GetSet_True(t *testing.T) {
	v := issetter.GetSet(true, issetter.True, issetter.False)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "GetSet true", actual)
}

func Test_Cov7_GetSet_False(t *testing.T) {
	v := issetter.GetSet(false, issetter.True, issetter.False)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "GetSet false", actual)
}

// ── GetSetUnset ──

func Test_Cov7_GetSetUnset_True(t *testing.T) {
	v := issetter.GetSetUnset(true)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.Set}
	expected.ShouldBeEqual(t, 0, "GetSetUnset true", actual)
}

func Test_Cov7_GetSetUnset_False(t *testing.T) {
	v := issetter.GetSetUnset(false)
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.Unset}
	expected.ShouldBeEqual(t, 0, "GetSetUnset false", actual)
}

// ── GetSetterByComparing ──

func Test_Cov7_GetSetterByComparing_Match(t *testing.T) {
	v := issetter.GetSetterByComparing(issetter.True, issetter.False, "a", "a", "b")
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing match", actual)
}

func Test_Cov7_GetSetterByComparing_NoMatch(t *testing.T) {
	v := issetter.GetSetterByComparing(issetter.True, issetter.False, "x", "a", "b")
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing no match", actual)
}

// ── Value methods ──

func Test_Cov7_Value_IsTrue(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsTrue()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsTrue", actual)
}

func Test_Cov7_Value_IsFalse(t *testing.T) {
	actual := args.Map{"result": issetter.False.IsFalse()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsFalse", actual)
}

func Test_Cov7_Value_IsSet(t *testing.T) {
	actual := args.Map{"result": issetter.Set.IsSet()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsSet", actual)
}

func Test_Cov7_Value_IsUnset(t *testing.T) {
	actual := args.Map{"result": issetter.Unset.IsUnset()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsUnset", actual)
}

func Test_Cov7_Value_IsWildcard(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.IsWildcard()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsWildcard", actual)
}

func Test_Cov7_Value_HasInitialized(t *testing.T) {
	actual := args.Map{"result": issetter.True.HasInitialized()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.HasInitialized", actual)
}

func Test_Cov7_Value_IsInvalid(t *testing.T) {
	actual := args.Map{"result": issetter.Uninitialized.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsInvalid", actual)
}

func Test_Cov7_Value_IsValid(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsValid", actual)
}

func Test_Cov7_Value_IsOn(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsOn()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsOn", actual)
}

func Test_Cov7_Value_IsOff(t *testing.T) {
	actual := args.Map{"result": issetter.False.IsOff()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsOff", actual)
}

func Test_Cov7_Value_Boolean(t *testing.T) {
	actual := args.Map{"result": issetter.True.Boolean()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.Boolean", actual)
}

func Test_Cov7_Value_String(t *testing.T) {
	result := issetter.True.String()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Value.String", actual)
}

func Test_Cov7_Value_ValueByte(t *testing.T) {
	actual := args.Map{"result": issetter.True.ValueByte()}
	expected := args.Map{"result": byte(1)}
	expected.ShouldBeEqual(t, 0, "Value.ValueByte", actual)
}

func Test_Cov7_Value_ValueInt(t *testing.T) {
	actual := args.Map{"result": issetter.True.ValueInt()}
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "Value.ValueInt", actual)
}

func Test_Cov7_Value_IsNot(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsNot(issetter.False)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsNot", actual)
}

func Test_Cov7_Value_IsAccept(t *testing.T) {
	actual := args.Map{"result": issetter.Set.IsAccept()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsAccept", actual)
}

func Test_Cov7_Value_IsReject(t *testing.T) {
	actual := args.Map{"result": issetter.Unset.IsReject()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsReject", actual)
}

func Test_Cov7_Value_IsYes(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsYes()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsYes", actual)
}

func Test_Cov7_Value_IsTrueOrSet(t *testing.T) {
	actual := args.Map{"true": issetter.True.IsTrueOrSet(), "set": issetter.Set.IsTrueOrSet(), "false": issetter.False.IsTrueOrSet()}
	expected := args.Map{"true": true, "set": true, "false": false}
	expected.ShouldBeEqual(t, 0, "Value.IsTrueOrSet", actual)
}

func Test_Cov7_Value_IsInitBoolean(t *testing.T) {
	actual := args.Map{"true": issetter.True.IsInitBoolean(), "set": issetter.Set.IsInitBoolean()}
	expected := args.Map{"true": true, "set": false}
	expected.ShouldBeEqual(t, 0, "Value.IsInitBoolean", actual)
}

func Test_Cov7_Value_IsDefinedLogically(t *testing.T) {
	actual := args.Map{"true": issetter.True.IsDefinedLogically(), "wild": issetter.Wildcard.IsDefinedLogically()}
	expected := args.Map{"true": true, "wild": false}
	expected.ShouldBeEqual(t, 0, "Value.IsDefinedLogically", actual)
}

func Test_Cov7_Value_IsUndefinedLogically(t *testing.T) {
	actual := args.Map{"wild": issetter.Wildcard.IsUndefinedLogically(), "true": issetter.True.IsUndefinedLogically()}
	expected := args.Map{"wild": true, "true": false}
	expected.ShouldBeEqual(t, 0, "Value.IsUndefinedLogically", actual)
}

func Test_Cov7_Value_ToBooleanValue(t *testing.T) {
	v := issetter.Set.ToBooleanValue()
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "Value.ToBooleanValue", actual)
}

func Test_Cov7_Value_ToSetUnsetValue(t *testing.T) {
	v := issetter.True.ToSetUnsetValue()
	actual := args.Map{"val": v}
	expected := args.Map{"val": issetter.Set}
	expected.ShouldBeEqual(t, 0, "Value.ToSetUnsetValue", actual)
}

func Test_Cov7_Value_NameValue(t *testing.T) {
	result := issetter.True.NameValue()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Value.NameValue", actual)
}

func Test_Cov7_Value_ToNumberString(t *testing.T) {
	result := issetter.True.ToNumberString()
	actual := args.Map{"result": result}
	expected := args.Map{"result": "1"}
	expected.ShouldBeEqual(t, 0, "Value.ToNumberString", actual)
}

func Test_Cov7_Value_ValueString(t *testing.T) {
	result := issetter.True.ValueString()
	actual := args.Map{"result": result}
	expected := args.Map{"result": "1"}
	expected.ShouldBeEqual(t, 0, "Value.ValueString", actual)
}

func Test_Cov7_Value_IsWildcardOrBool_Wildcard(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.IsWildcardOrBool(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsWildcardOrBool wildcard", actual)
}

func Test_Cov7_Value_IsWildcardOrBool_True(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsWildcardOrBool(true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsWildcardOrBool true", actual)
}

func Test_Cov7_Value_IsWildcardOrBool_False(t *testing.T) {
	actual := args.Map{"result": issetter.False.IsWildcardOrBool(false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsWildcardOrBool false", actual)
}

func Test_Cov7_Value_ToByteCondition(t *testing.T) {
	actual := args.Map{"true": issetter.True.ToByteCondition(1, 0, 255), "false": issetter.False.ToByteCondition(1, 0, 255), "uninit": issetter.Uninitialized.ToByteCondition(1, 0, 255)}
	expected := args.Map{"true": byte(1), "false": byte(0), "uninit": byte(255)}
	expected.ShouldBeEqual(t, 0, "Value.ToByteCondition", actual)
}

func Test_Cov7_Value_Format(t *testing.T) {
	result := issetter.True.Format("{name}={value}")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Value.Format", actual)
}

func Test_Cov7_Value_IsNameEqual(t *testing.T) {
	name := issetter.True.String()
	actual := args.Map{"result": issetter.True.IsNameEqual(name)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsNameEqual", actual)
}

func Test_Cov7_Value_IsAnyNamesOf_Match(t *testing.T) {
	name := issetter.True.String()
	actual := args.Map{"result": issetter.True.IsAnyNamesOf(name, "other")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsAnyNamesOf match", actual)
}

func Test_Cov7_Value_IsAnyNamesOf_NoMatch(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsAnyNamesOf("bogus")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Value.IsAnyNamesOf no match", actual)
}

func Test_Cov7_Value_AllNameValues(t *testing.T) {
	result := issetter.True.AllNameValues()
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Value.AllNameValues", actual)
}

func Test_Cov7_Value_RangeNamesCsv(t *testing.T) {
	result := issetter.True.RangeNamesCsv()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Value.RangeNamesCsv", actual)
}
