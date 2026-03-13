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
	v1, _ := issetter.New("True")
	v2, _ := issetter.New("Set")
	result := issetter.Max(v1, v2)

	// Assert
	if result != v2 {
		t.Error("expected Set to be max")
	}
}

func Test_Min_Ext2(t *testing.T) {
	// Act
	v1, _ := issetter.New("True")
	v2, _ := issetter.New("Set")
	result := issetter.Min(v1, v2)

	// Assert
	if result != v1 {
		t.Error("expected True to be min")
	}
}
