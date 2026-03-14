package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/issetter"
)

func Test_Value_Methods_Ext2(t *testing.T) {
	// Arrange
	v, err := issetter.New("Set")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Assert
	if v.IsUnset() {
		t.Error("should not be unset")
	}
	if !v.IsSet() {
		t.Error("should be set")
	}
}

func Test_NewBool_Ext2(t *testing.T) {
	// Arrange
	v := issetter.NewBool(true)

	// Assert
	if v.IsUnset() {
		t.Error("should not be unset")
	}
	if !v.Boolean() {
		t.Error("should be true")
	}
}

func Test_NewMust_Ext2(t *testing.T) {
	// Act
	v := issetter.NewMust("True")

	// Assert
	if v.IsUnset() {
		t.Error("should not be unset")
	}
	if !v.IsTrue() {
		t.Error("should be True")
	}
}

func Test_Max_Ext2(t *testing.T) {
	// Act
	result := issetter.Max()

	// Assert
	if result != issetter.Wildcard {
		t.Error("expected Wildcard to be max")
	}
}

func Test_Min_Ext2(t *testing.T) {
	// Act
	result := issetter.Min()

	// Assert
	if result != issetter.Uninitialized {
		t.Error("expected Uninitialized to be min")
	}
}
