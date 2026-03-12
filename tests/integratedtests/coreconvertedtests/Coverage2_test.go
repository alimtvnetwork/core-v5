package coreconvertedtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/converters/coreconverted"
)

func Test_Cov2_Bytes_HandleWithPanic_NoError(t *testing.T) {
	b := &coreconverted.Bytes{Values: []byte{1}, CombinedError: nil}
	// Should not panic
	b.HandleWithPanic()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "HandleWithPanic_NoError", actual)
}

func Test_Cov2_Bytes_HandleWithPanic_WithError(t *testing.T) {
	b := &coreconverted.Bytes{Values: nil, CombinedError: errors.New("err")}
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		b.HandleWithPanic()
	}()
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "HandleWithPanic_WithError", actual)
}

func Test_Cov2_Integers_HandleWithPanic_NoError(t *testing.T) {
	i := &coreconverted.Integers{Values: []int{1}, CombinedError: nil}
	i.HandleWithPanic()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "Integers_HandleWithPanic_NoError", actual)
}

func Test_Cov2_Integers_HandleWithPanic_WithError(t *testing.T) {
	i := &coreconverted.Integers{Values: nil, CombinedError: errors.New("err")}
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		i.HandleWithPanic()
	}()
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "Integers_HandleWithPanic_WithError", actual)
}

func Test_Cov2_Bytes_HasAnyItem_Empty(t *testing.T) {
	b := &coreconverted.Bytes{Values: []byte{}}
	actual := args.Map{"hasAny": b.HasAnyItem(), "isEmpty": b.IsEmpty()}
	expected := args.Map{"hasAny": false, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Bytes_HasAnyItem_Empty", actual)
}

func Test_Cov2_Integers_HasAnyItem_Empty(t *testing.T) {
	i := &coreconverted.Integers{Values: []int{}}
	actual := args.Map{"hasAny": i.HasAnyItem(), "isEmpty": i.IsEmpty()}
	expected := args.Map{"hasAny": false, "isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Integers_HasAnyItem_Empty", actual)
}
