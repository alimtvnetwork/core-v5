package issettertests

import (
	"testing"

	"github.com/smarty/assertions/should"
	convey "github.com/smartystreets/goconvey/convey"

	"gitlab.com/auk-go/core/issetter"
)

// =============================================================================
// IsOnLogically / IsOffLogically — compound checks: IsInitialized() && trueMap/falseMap
// =============================================================================

func Test_Value_IsOnLogically_Uninitialized(t *testing.T) {
	// Arrange & Act & Assert
	convey.Convey("IsOnLogically - Uninitialized should return false", t, func() {
		convey.So(issetter.Uninitialized.IsOnLogically(), should.BeFalse)
	})
}

func Test_Value_IsOnLogically_True(t *testing.T) {
	convey.Convey("IsOnLogically - True should return true", t, func() {
		convey.So(issetter.True.IsOnLogically(), should.BeTrue)
	})
}

func Test_Value_IsOnLogically_False(t *testing.T) {
	convey.Convey("IsOnLogically - False should return false", t, func() {
		convey.So(issetter.False.IsOnLogically(), should.BeFalse)
	})
}

func Test_Value_IsOnLogically_Unset(t *testing.T) {
	convey.Convey("IsOnLogically - Unset should return false", t, func() {
		convey.So(issetter.Unset.IsOnLogically(), should.BeFalse)
	})
}

func Test_Value_IsOnLogically_Set(t *testing.T) {
	convey.Convey("IsOnLogically - Set should return true", t, func() {
		convey.So(issetter.Set.IsOnLogically(), should.BeTrue)
	})
}

func Test_Value_IsOnLogically_Wildcard(t *testing.T) {
	convey.Convey("IsOnLogically - Wildcard should return false (not in trueMap)", t, func() {
		convey.So(issetter.Wildcard.IsOnLogically(), should.BeFalse)
	})
}

func Test_Value_IsOffLogically_Uninitialized(t *testing.T) {
	convey.Convey("IsOffLogically - Uninitialized should return false (not initialized)", t, func() {
		convey.So(issetter.Uninitialized.IsOffLogically(), should.BeFalse)
	})
}

func Test_Value_IsOffLogically_True(t *testing.T) {
	convey.Convey("IsOffLogically - True should return false", t, func() {
		convey.So(issetter.True.IsOffLogically(), should.BeFalse)
	})
}

func Test_Value_IsOffLogically_False(t *testing.T) {
	convey.Convey("IsOffLogically - False should return true", t, func() {
		convey.So(issetter.False.IsOffLogically(), should.BeTrue)
	})
}

func Test_Value_IsOffLogically_Unset(t *testing.T) {
	convey.Convey("IsOffLogically - Unset should return true", t, func() {
		convey.So(issetter.Unset.IsOffLogically(), should.BeTrue)
	})
}

func Test_Value_IsOffLogically_Set(t *testing.T) {
	convey.Convey("IsOffLogically - Set should return false", t, func() {
		convey.So(issetter.Set.IsOffLogically(), should.BeFalse)
	})
}

func Test_Value_IsOffLogically_Wildcard(t *testing.T) {
	convey.Convey("IsOffLogically - Wildcard should return false (not in falseMap)", t, func() {
		convey.So(issetter.Wildcard.IsOffLogically(), should.BeFalse)
	})
}

// =============================================================================
// WildcardApply — wildcard/unset/uninit pass through input, others use IsTrue
// =============================================================================

func Test_Value_WildcardApply_Wildcard_True(t *testing.T) {
	convey.Convey("WildcardApply - Wildcard passes through true", t, func() {
		convey.So(issetter.Wildcard.WildcardApply(true), should.BeTrue)
	})
}

func Test_Value_WildcardApply_Wildcard_False(t *testing.T) {
	convey.Convey("WildcardApply - Wildcard passes through false", t, func() {
		convey.So(issetter.Wildcard.WildcardApply(false), should.BeFalse)
	})
}

func Test_Value_WildcardApply_Uninitialized_True(t *testing.T) {
	convey.Convey("WildcardApply - Uninitialized passes through true", t, func() {
		convey.So(issetter.Uninitialized.WildcardApply(true), should.BeTrue)
	})
}

func Test_Value_WildcardApply_Unset_False(t *testing.T) {
	convey.Convey("WildcardApply - Unset passes through false", t, func() {
		convey.So(issetter.Unset.WildcardApply(false), should.BeFalse)
	})
}

func Test_Value_WildcardApply_True_IgnoresInput(t *testing.T) {
	convey.Convey("WildcardApply - True ignores input and returns true", t, func() {
		convey.So(issetter.True.WildcardApply(false), should.BeTrue)
	})
}

func Test_Value_WildcardApply_False_IgnoresInput(t *testing.T) {
	convey.Convey("WildcardApply - False ignores input and returns false", t, func() {
		convey.So(issetter.False.WildcardApply(true), should.BeFalse)
	})
}

func Test_Value_WildcardApply_Set_IgnoresInput(t *testing.T) {
	convey.Convey("WildcardApply - Set ignores input and returns false (Set != True)", t, func() {
		convey.So(issetter.Set.WildcardApply(true), should.BeFalse)
	})
}

// =============================================================================
// IsWildcardOrBool — Wildcard short-circuits to true, others return isBool
// =============================================================================

func Test_Value_IsWildcardOrBool_Wildcard(t *testing.T) {
	convey.Convey("IsWildcardOrBool - Wildcard always returns true", t, func() {
		convey.So(issetter.Wildcard.IsWildcardOrBool(false), should.BeTrue)
	})
}

func Test_Value_IsWildcardOrBool_True_WithTrue(t *testing.T) {
	convey.Convey("IsWildcardOrBool - non-Wildcard returns isBool=true", t, func() {
		convey.So(issetter.True.IsWildcardOrBool(true), should.BeTrue)
	})
}

func Test_Value_IsWildcardOrBool_False_WithFalse(t *testing.T) {
	convey.Convey("IsWildcardOrBool - non-Wildcard returns isBool=false", t, func() {
		convey.So(issetter.False.IsWildcardOrBool(false), should.BeFalse)
	})
}

// =============================================================================
// ToByteCondition — 3-way branch: True, False, other
// =============================================================================

func Test_Value_ToByteCondition_True(t *testing.T) {
	convey.Convey("ToByteCondition - True returns trueVal", t, func() {
		convey.So(issetter.True.ToByteCondition(10, 20, 255), should.Equal, byte(10))
	})
}

func Test_Value_ToByteCondition_False(t *testing.T) {
	convey.Convey("ToByteCondition - False returns falseVal", t, func() {
		convey.So(issetter.False.ToByteCondition(10, 20, 255), should.Equal, byte(20))
	})
}

func Test_Value_ToByteCondition_Uninitialized(t *testing.T) {
	convey.Convey("ToByteCondition - Uninitialized returns invalid", t, func() {
		convey.So(issetter.Uninitialized.ToByteCondition(10, 20, 255), should.Equal, byte(255))
	})
}

func Test_Value_ToByteCondition_Set(t *testing.T) {
	convey.Convey("ToByteCondition - Set returns invalid (not True/False)", t, func() {
		convey.So(issetter.Set.ToByteCondition(10, 20, 255), should.Equal, byte(255))
	})
}

func Test_Value_ToByteCondition_Wildcard(t *testing.T) {
	convey.Convey("ToByteCondition - Wildcard returns invalid", t, func() {
		convey.So(issetter.Wildcard.ToByteCondition(10, 20, 255), should.Equal, byte(255))
	})
}

// =============================================================================
// ToByteConditionWithWildcard — 4-way branch: Wildcard, True, False, other
// =============================================================================

func Test_Value_ToByteConditionWithWildcard_Wildcard(t *testing.T) {
	convey.Convey("ToByteConditionWithWildcard - Wildcard returns wildcard byte", t, func() {
		convey.So(issetter.Wildcard.ToByteConditionWithWildcard(99, 10, 20, 255), should.Equal, byte(99))
	})
}

func Test_Value_ToByteConditionWithWildcard_True(t *testing.T) {
	convey.Convey("ToByteConditionWithWildcard - True returns trueVal", t, func() {
		convey.So(issetter.True.ToByteConditionWithWildcard(99, 10, 20, 255), should.Equal, byte(10))
	})
}

func Test_Value_ToByteConditionWithWildcard_False(t *testing.T) {
	convey.Convey("ToByteConditionWithWildcard - False returns falseVal", t, func() {
		convey.So(issetter.False.ToByteConditionWithWildcard(99, 10, 20, 255), should.Equal, byte(20))
	})
}

func Test_Value_ToByteConditionWithWildcard_Uninitialized(t *testing.T) {
	convey.Convey("ToByteConditionWithWildcard - Uninitialized returns invalid", t, func() {
		convey.So(issetter.Uninitialized.ToByteConditionWithWildcard(99, 10, 20, 255), should.Equal, byte(255))
	})
}

// =============================================================================
// GetSetBoolOnInvalid — mutates receiver if not a defined boolean
// =============================================================================

func Test_Value_GetSetBoolOnInvalid_AlreadyTrue(t *testing.T) {
	// Arrange
	v := issetter.True

	// Act
	result := v.GetSetBoolOnInvalid(false)

	// Assert
	convey.Convey("GetSetBoolOnInvalid - already True returns true, ignores setter", t, func() {
		convey.So(result, should.BeTrue)
		convey.So(v.IsTrue(), should.BeTrue)
	})
}

func Test_Value_GetSetBoolOnInvalid_AlreadyFalse(t *testing.T) {
	// Arrange
	v := issetter.False

	// Act
	result := v.GetSetBoolOnInvalid(true)

	// Assert
	convey.Convey("GetSetBoolOnInvalid - already False returns false, ignores setter", t, func() {
		convey.So(result, should.BeFalse)
		convey.So(v.IsFalse(), should.BeTrue)
	})
}

func Test_Value_GetSetBoolOnInvalid_Uninitialized_SetsTrue(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized

	// Act
	result := v.GetSetBoolOnInvalid(true)

	// Assert
	convey.Convey("GetSetBoolOnInvalid - Uninitialized with true sets to True", t, func() {
		convey.So(result, should.BeTrue)
		convey.So(v.IsTrue(), should.BeTrue)
	})
}

func Test_Value_GetSetBoolOnInvalid_Uninitialized_SetsFalse(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized

	// Act
	result := v.GetSetBoolOnInvalid(false)

	// Assert
	convey.Convey("GetSetBoolOnInvalid - Uninitialized with false sets to False", t, func() {
		convey.So(result, should.BeFalse)
		convey.So(v.IsFalse(), should.BeTrue)
	})
}

func Test_Value_GetSetBoolOnInvalid_Set_SetsValue(t *testing.T) {
	// Arrange
	v := issetter.Set

	// Act
	result := v.GetSetBoolOnInvalid(true)

	// Assert
	convey.Convey("GetSetBoolOnInvalid - Set (not a defined boolean) triggers setter", t, func() {
		convey.So(result, should.BeTrue)
		convey.So(v.IsTrue(), should.BeTrue)
	})
}

// =============================================================================
// LazyEvaluateBool — executes func only on first call when not defined boolean
// =============================================================================

func Test_Value_LazyEvaluateBool_Uninitialized_CallsFunc(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	called := false

	// Act
	result := v.LazyEvaluateBool(func() { called = true })

	// Assert
	convey.Convey("LazyEvaluateBool - Uninitialized calls func and sets True", t, func() {
		convey.So(called, should.BeTrue)
		convey.So(result, should.BeTrue)
		convey.So(v.IsTrue(), should.BeTrue)
	})
}

func Test_Value_LazyEvaluateBool_AlreadyTrue_SkipsFunc(t *testing.T) {
	// Arrange
	v := issetter.True
	called := false

	// Act
	result := v.LazyEvaluateBool(func() { called = true })

	// Assert
	convey.Convey("LazyEvaluateBool - already True skips func", t, func() {
		convey.So(called, should.BeFalse)
		convey.So(result, should.BeFalse)
	})
}

func Test_Value_LazyEvaluateBool_AlreadyFalse_SkipsFunc(t *testing.T) {
	// Arrange
	v := issetter.False
	called := false

	// Act
	result := v.LazyEvaluateBool(func() { called = true })

	// Assert
	convey.Convey("LazyEvaluateBool - already False skips func", t, func() {
		convey.So(called, should.BeFalse)
		convey.So(result, should.BeFalse)
	})
}

// =============================================================================
// LazyEvaluateSet — same as LazyEvaluateBool but for Set/Unset
// =============================================================================

func Test_Value_LazyEvaluateSet_Uninitialized_CallsFunc(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	called := false

	// Act
	result := v.LazyEvaluateSet(func() { called = true })

	// Assert
	convey.Convey("LazyEvaluateSet - Uninitialized calls func and sets Set", t, func() {
		convey.So(called, should.BeTrue)
		convey.So(result, should.BeTrue)
		convey.So(v.IsSet(), should.BeTrue)
	})
}

func Test_Value_LazyEvaluateSet_AlreadySet_SkipsFunc(t *testing.T) {
	// Arrange
	v := issetter.Set
	called := false

	// Act
	result := v.LazyEvaluateSet(func() { called = true })

	// Assert
	convey.Convey("LazyEvaluateSet - already Set skips func", t, func() {
		convey.So(called, should.BeFalse)
		convey.So(result, should.BeFalse)
	})
}

func Test_Value_LazyEvaluateSet_AlreadyUnset_SkipsFunc(t *testing.T) {
	// Arrange
	v := issetter.Unset
	called := false

	// Act
	result := v.LazyEvaluateSet(func() { called = true })

	// Assert
	convey.Convey("LazyEvaluateSet - already Unset skips func", t, func() {
		convey.So(called, should.BeFalse)
		convey.So(result, should.BeFalse)
	})
}

// =============================================================================
// IsDefinedLogically / IsUndefinedLogically
// =============================================================================

func Test_Value_IsDefinedLogically_AllValues(t *testing.T) {
	convey.Convey("IsDefinedLogically - covers all 6 values", t, func() {
		convey.So(issetter.Uninitialized.IsDefinedLogically(), should.BeFalse)
		convey.So(issetter.True.IsDefinedLogically(), should.BeTrue)
		convey.So(issetter.False.IsDefinedLogically(), should.BeTrue)
		convey.So(issetter.Unset.IsDefinedLogically(), should.BeTrue)
		convey.So(issetter.Set.IsDefinedLogically(), should.BeTrue)
		convey.So(issetter.Wildcard.IsDefinedLogically(), should.BeFalse)
	})
}

func Test_Value_IsUndefinedLogically_AllValues(t *testing.T) {
	convey.Convey("IsUndefinedLogically - covers all 6 values", t, func() {
		convey.So(issetter.Uninitialized.IsUndefinedLogically(), should.BeTrue)
		convey.So(issetter.True.IsUndefinedLogically(), should.BeFalse)
		convey.So(issetter.False.IsUndefinedLogically(), should.BeFalse)
		convey.So(issetter.Unset.IsUndefinedLogically(), should.BeFalse)
		convey.So(issetter.Set.IsUndefinedLogically(), should.BeFalse)
		convey.So(issetter.Wildcard.IsUndefinedLogically(), should.BeTrue)
	})
}

// =============================================================================
// IsPositive / IsNegative
// =============================================================================

func Test_Value_IsPositive_AllValues(t *testing.T) {
	convey.Convey("IsPositive - True and Set are positive", t, func() {
		convey.So(issetter.Uninitialized.IsPositive(), should.BeFalse)
		convey.So(issetter.True.IsPositive(), should.BeTrue)
		convey.So(issetter.False.IsPositive(), should.BeFalse)
		convey.So(issetter.Unset.IsPositive(), should.BeFalse)
		convey.So(issetter.Set.IsPositive(), should.BeTrue)
		convey.So(issetter.Wildcard.IsPositive(), should.BeFalse)
	})
}

func Test_Value_IsNegative_AllValues(t *testing.T) {
	convey.Convey("IsNegative - Uninitialized, False, Unset are negative", t, func() {
		convey.So(issetter.Uninitialized.IsNegative(), should.BeTrue)
		convey.So(issetter.True.IsNegative(), should.BeFalse)
		convey.So(issetter.False.IsNegative(), should.BeTrue)
		convey.So(issetter.Unset.IsNegative(), should.BeTrue)
		convey.So(issetter.Set.IsNegative(), should.BeFalse)
		convey.So(issetter.Wildcard.IsNegative(), should.BeFalse)
	})
}
