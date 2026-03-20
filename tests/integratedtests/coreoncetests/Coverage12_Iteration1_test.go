package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
)

// Test_Cov12_IntegersOnce_IsEqual_NilReceiver tests IsEqual on nil *IntegersOnce.
func Test_Cov12_IntegersOnce_IsEqual_NilReceiver(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.IntegersOnce

	// Act
	actual := nilOnce.IsEqual(nil...)

	// Assert
	if !actual {
		t.Fatalf("nil IntegersOnce.IsEqual(nil) should return true, got false")
	}
}

// Test_Cov12_IntegersOnce_IsEqual_NilReceiverEmpty tests IsEqual nil receiver with empty variadic.
func Test_Cov12_IntegersOnce_IsEqual_NilReceiverEmpty(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.IntegersOnce

	// Act
	actual := nilOnce.IsEqual()

	// Assert
	if !actual {
		t.Fatalf("nil IntegersOnce.IsEqual() should return true, got false")
	}
}

// Test_Cov12_MapStringStringOnce_IsEqual_NilReceiver tests IsEqual on nil *MapStringStringOnce.
func Test_Cov12_MapStringStringOnce_IsEqual_NilReceiver(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.MapStringStringOnce

	// Act
	actual := nilOnce.IsEqual(nil)

	// Assert
	if !actual {
		t.Fatalf("nil MapStringStringOnce.IsEqual(nil) should return true, got false")
	}
}

// Test_Cov12_MapStringStringOnce_JsonStringMust_Success tests JsonStringMust on valid data.
func Test_Cov12_MapStringStringOnce_JsonStringMust_Success(t *testing.T) {
	// Arrange
	once := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{"key": "value"}
	})

	// Act
	actual := once.JsonStringMust()

	// Assert
	expected := `{"key":"value"}`
	if actual != expected {
		t.Fatalf("MapStringStringOnce.JsonStringMust: got %q, want %q", actual, expected)
	}
}

// Test_Cov12_StringsOnce_IsEqual_NilReceiver tests IsEqual on nil *StringsOnce.
func Test_Cov12_StringsOnce_IsEqual_NilReceiver(t *testing.T) {
	// Arrange
	var nilOnce *coreonce.StringsOnce

	// Act
	actual := nilOnce.IsEqual()

	// Assert
	if !actual {
		t.Fatalf("nil StringsOnce.IsEqual() should return true, got false")
	}
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
	if actual != expected {
		t.Fatalf("StringsOnce.JsonStringMust: got %q, want %q", actual, expected)
	}
}
