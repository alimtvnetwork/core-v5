package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ============================================================================
// And / Or / WildcardApply — additional branches
// ============================================================================

func Test_Cov4_And_TrueFalse(t *testing.T) {
	actual := args.Map{
		"trueAndFalse":  issetter.True.And(issetter.False) == issetter.False,
		"falseAndTrue":  issetter.False.And(issetter.True) == issetter.False,
		"falseAndFalse": issetter.False.And(issetter.False) == issetter.False,
		"wildAndTrue":   issetter.Wildcard.And(issetter.True) == issetter.True,
	}
	expected := args.Map{
		"trueAndFalse": true, "falseAndTrue": true,
		"falseAndFalse": true, "wildAndTrue": true,
	}
	expected.ShouldBeEqual(t, 0, "And returns expected -- all combinations", actual)
}

func Test_Cov4_OrBool_Branches(t *testing.T) {
	actual := args.Map{
		"falseOrTrue":  issetter.False.OrBool(true),
		"falseOrFalse": issetter.False.OrBool(false),
		"wildOrFalse":  issetter.Wildcard.OrBool(false),
	}
	expected := args.Map{
		"falseOrTrue": true, "falseOrFalse": false, "wildOrFalse": false,
	}
	expected.ShouldBeEqual(t, 0, "OrBool returns expected -- all branches", actual)
}

func Test_Cov4_WildcardApply_False(t *testing.T) {
	actual := args.Map{
		"wildFalse":  issetter.Wildcard.WildcardApply(false),
		"falseFalse": issetter.False.WildcardApply(true),
	}
	expected := args.Map{"wildFalse": false, "falseFalse": false}
	expected.ShouldBeEqual(t, 0, "WildcardApply returns expected -- false paths", actual)
}

func Test_Cov4_WildcardValueApply(t *testing.T) {
	actual := args.Map{
		"wildTrue":  issetter.Wildcard.WildcardValueApply(issetter.True),
		"wildFalse": issetter.Wildcard.WildcardValueApply(issetter.False),
		"trueAny":   issetter.True.WildcardValueApply(issetter.False),
		"falseAny":  issetter.False.WildcardValueApply(issetter.True),
	}
	expected := args.Map{
		"wildTrue": true, "wildFalse": false,
		"trueAny": true, "falseAny": false,
	}
	expected.ShouldBeEqual(t, 0, "WildcardValueApply returns expected -- all paths", actual)
}

// ============================================================================
// OrValue
// ============================================================================

func Test_Cov4_OrValue(t *testing.T) {
	actual := args.Map{
		"trueOrFalse":  issetter.True.OrValue(issetter.False),
		"falseOrTrue":  issetter.False.OrValue(issetter.True),
		"falseOrFalse": issetter.False.OrValue(issetter.False),
	}
	expected := args.Map{
		"trueOrFalse": true, "falseOrTrue": true, "falseOrFalse": false,
	}
	expected.ShouldBeEqual(t, 0, "OrValue returns expected -- all combinations", actual)
}

// ============================================================================
// AndBool edge cases
// ============================================================================

func Test_Cov4_AndBool_WildcardTrue(t *testing.T) {
	actual := args.Map{
		"wildTrue":  issetter.Wildcard.AndBool(true),
		"wildFalse": issetter.Wildcard.AndBool(false),
	}
	expected := args.Map{"wildTrue": true, "wildFalse": false}
	expected.ShouldBeEqual(t, 0, "AndBool wildcard passes through -- both", actual)
}

// ============================================================================
// IsWildcardOrBool with True
// ============================================================================

func Test_Cov4_IsWildcardOrBool(t *testing.T) {
	actual := args.Map{
		"wildTrue":  issetter.Wildcard.IsWildcardOrBool(true),
		"wildFalse": issetter.Wildcard.IsWildcardOrBool(false),
		"trueTrue":  issetter.True.IsWildcardOrBool(true),
		"trueFalse": issetter.True.IsWildcardOrBool(false),
		"falseFalse": issetter.False.IsWildcardOrBool(false),
	}
	expected := args.Map{
		"wildTrue": true, "wildFalse": true,
		"trueTrue": true, "trueFalse": false,
		"falseFalse": false,
	}
	expected.ShouldBeEqual(t, 0, "IsWildcardOrBool returns expected -- all combos", actual)
}

// ============================================================================
// ToByteCondition — Uninitialized
// ============================================================================

func Test_Cov4_ToByteCondition_Uninit(t *testing.T) {
	actual := args.Map{
		"result": int(issetter.Uninitialized.ToByteCondition(1, 0, 255)),
	}
	expected := args.Map{"result": 255}
	expected.ShouldBeEqual(t, 0, "ToByteCondition returns invalid -- Uninitialized", actual)
}

func Test_Cov4_ToByteConditionWithWildcard_Uninit(t *testing.T) {
	actual := args.Map{
		"result": int(issetter.Uninitialized.ToByteConditionWithWildcard(99, 1, 0, 255)),
	}
	expected := args.Map{"result": 255}
	expected.ShouldBeEqual(t, 0, "ToByteConditionWithWildcard returns invalid -- Uninitialized", actual)
}

// ============================================================================
// ToBooleanValue / ToSetUnsetValue — edge cases
// ============================================================================

func Test_Cov4_ToBooleanValue_Wildcard(t *testing.T) {
	result := issetter.Wildcard.ToBooleanValue()
	actual := args.Map{"isWild": result == issetter.Wildcard}
	expected := args.Map{"isWild": true}
	expected.ShouldBeEqual(t, 0, "ToBooleanValue Wildcard stays Wildcard", actual)
}

func Test_Cov4_ToSetUnsetValue_Wildcard(t *testing.T) {
	result := issetter.Wildcard.ToSetUnsetValue()
	actual := args.Map{"isWild": result == issetter.Wildcard}
	expected := args.Map{"isWild": true}
	expected.ShouldBeEqual(t, 0, "ToSetUnsetValue Wildcard stays Wildcard", actual)
}

// ============================================================================
// IsOnLogically / IsOffLogically — more values
// ============================================================================

func Test_Cov4_IsOnLogically_Set(t *testing.T) {
	actual := args.Map{
		"setOn":    issetter.Set.IsOnLogically(),
		"unsetOff": issetter.Unset.IsOffLogically(),
		"wildOff":  issetter.Wildcard.IsOffLogically(),
	}
	expected := args.Map{
		"setOn": true, "unsetOff": true, "wildOff": false,
	}
	expected.ShouldBeEqual(t, 0, "IsOnLogically/IsOffLogically returns expected -- Set/Unset/Wild", actual)
}

// ============================================================================
// GetSetBoolOnInvalid — already initialized
// ============================================================================

func Test_Cov4_GetSetBoolOnInvalid_AlreadyInit(t *testing.T) {
	v := issetter.True
	result := v.GetSetBoolOnInvalid(false)
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetSetBoolOnInvalid returns existing -- already True", actual)
}

func Test_Cov4_GetSetBoolOnInvalidFunc_AlreadyInit(t *testing.T) {
	v := issetter.False
	result := v.GetSetBoolOnInvalidFunc(func() bool { return true })
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetSetBoolOnInvalidFunc returns existing -- already False", actual)
}

// ============================================================================
// LazyEvaluateBool — True value
// ============================================================================

func Test_Cov4_LazyEvaluateBool_True(t *testing.T) {
	v := issetter.True
	called := v.LazyEvaluateBool(func() {})
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "LazyEvaluateBool not called -- already True", actual)
}

// ============================================================================
// LazyEvaluateSet — Set value
// ============================================================================

func Test_Cov4_LazyEvaluateSet_Set(t *testing.T) {
	v := issetter.Set
	called := v.LazyEvaluateSet(func() {})
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "LazyEvaluateSet not called -- already Set", actual)
}

// ============================================================================
// YesNoMappedValue edge values
// ============================================================================

func Test_Cov4_YesNoMappedValue_SetUnset(t *testing.T) {
	actual := args.Map{
		"set":   issetter.Set.YesNoMappedValue(),
		"unset": issetter.Unset.YesNoMappedValue(),
	}
	expected := args.Map{"set": "yes", "unset": "no"}
	expected.ShouldBeEqual(t, 0, "YesNoMappedValue returns expected -- Set/Unset", actual)
}

// ============================================================================
// Name methods — False/Unset/Wildcard/Uninitialized
// ============================================================================

func Test_Cov4_NameMethods_False(t *testing.T) {
	actual := args.Map{
		"yesNo":        issetter.False.YesNoName(),
		"trueFalse":    issetter.False.TrueFalseName(),
		"onOff":        issetter.False.OnOffName(),
		"yesNoLower":   issetter.False.YesNoLowercaseName(),
		"trFaLower":    issetter.False.TrueFalseLowercaseName(),
		"onOffLower":   issetter.False.OnOffLowercaseName(),
		"setUnsetLow":  issetter.False.SetUnsetLowercaseName(),
	}
	expected := args.Map{
		"yesNo": "No", "trueFalse": "False", "onOff": "Off",
		"yesNoLower": "no", "trFaLower": "false",
		"onOffLower": "off", "setUnsetLow": "unset",
	}
	expected.ShouldBeEqual(t, 0, "Name methods False returns expected -- all variants", actual)
}

func Test_Cov4_NameMethods_Wildcard(t *testing.T) {
	actual := args.Map{
		"yesNo":      issetter.Wildcard.YesNoMappedValue(),
		"trueFalse":  issetter.Wildcard.TrueFalseName(),
		"onOff":      issetter.Wildcard.OnOffName(),
	}
	expected := args.Map{
		"yesNo": issetter.Wildcard.YesNoMappedValue(), "trueFalse": issetter.Wildcard.TrueFalseName(), "onOff": issetter.Wildcard.OnOffName(),
	}
	expected.ShouldBeEqual(t, 0, "Name methods Wildcard returns empty -- undefined", actual)
}

// ============================================================================
// IsEqual / IsBetween additional edge cases
// ============================================================================

func Test_Cov4_IsBetween_OutOfRange(t *testing.T) {
	actual := args.Map{
		"below": issetter.Uninitialized.IsBetween(1, 5),
		"above": issetter.Wildcard.IsBetween(0, 2),
	}
	expected := args.Map{"below": false, "above": false}
	expected.ShouldBeEqual(t, 0, "IsBetween out of range returns false -- both edges", actual)
}
