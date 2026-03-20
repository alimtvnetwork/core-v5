package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coredata/coreonce"
)

// Test_Cov12_IntegersOnce_IsEqual_NilReceiver tests IsEqual on nil *IntegersOnce.
func Test_Cov12_IntegersOnce_IsEqual_NilReceiver(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.IntegersOnce

	// Act
	actual := nilOnce.IsEqual(nil...)

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"nil IntegersOnce.IsEqual(nil) should return true",
		actual,
		true,
	)
}

// Test_Cov12_IntegersOnce_IsEqual_NilReceiverWithArgs tests IsEqual nil receiver with non-nil args.
func Test_Cov12_IntegersOnce_IsEqual_NilReceiverWithArgs(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.IntegersOnce

	// Act — this should hit the nil-receiver-with-non-nil-args branch
	// Since nil receiver and non-nil comparingItems, it goes past the first guard
	// and calls it.Value() which will panic on nil receiver.
	// Actually, let's check: line 130 is `if it == nil && integerItems == nil`
	// So passing nil variadic triggers the true branch (line 131 return true).
	// To cover line 130 false branch (nil receiver, non-nil args), we'd need
	// it.Value() on nil receiver which would panic.
	// The uncovered line IS line 130.38,132.3 = the body of `if it == nil && integerItems == nil`.
	// So we just need to pass nil to both.
	actual := nilOnce.IsEqual()

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"nil IntegersOnce.IsEqual() should return true",
		actual,
		true,
	)
}

// Test_Cov12_MapStringStringOnce_IsEqual_NilReceiver tests IsEqual on nil *MapStringStringOnce.
func Test_Cov12_MapStringStringOnce_IsEqual_NilReceiver(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.MapStringStringOnce

	// Act
	actual := nilOnce.IsEqual(nil)

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"nil MapStringStringOnce.IsEqual(nil) should return true",
		actual,
		true,
	)
}

// Test_Cov12_MapStringStringOnce_JsonStringMust_Success tests JsonStringMust on valid data.
// The error branch (lines 309-314) is a panic path triggered by marshal failure.
// For normal data, JsonStringMust should succeed.
func Test_Cov12_MapStringStringOnce_JsonStringMust_Success(t *testing.T) {
	// Arrange
	once := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"key": "value"}
	})

	// Act
	actual := once.JsonStringMust()

	// Assert
	expected := `{"key":"value"}`
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"MapStringStringOnce.JsonStringMust should return valid JSON",
		actual,
		expected,
	)
}

// Test_Cov12_StringsOnce_IsEqual_NilReceiver tests IsEqual on nil *StringsOnce.
func Test_Cov12_StringsOnce_IsEqual_NilReceiver(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.StringsOnce

	// Act
	actual := nilOnce.IsEqual()

	// Assert
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"nil StringsOnce.IsEqual() should return true",
		actual,
		true,
	)
}

// Test_Cov12_StringsOnce_JsonStringMust_Success tests JsonStringMust on valid data.
func Test_Cov12_StringsOnce_JsonStringMust_Success(t *testing.T) {
	// Arrange
	once := coreonce.NewStringsOnce(func() []string {
		return []string{"a", "b"}
	})

	// Act
	actual := once.JsonStringMust()

	// Assert
	expected := `["a","b"]`
	coretests.GetAssert.ShouldBeEqual(
		t,
		0,
		"StringsOnce.JsonStringMust should return valid JSON",
		actual,
		expected,
	)
}
