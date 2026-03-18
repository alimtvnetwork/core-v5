package errcore

import (
	"errors"
	"testing"
)

func TestHandleErr_Nil(t *testing.T) {
	HandleErr(nil) // should not panic
}

func TestHandleErr_WithErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	HandleErr(errors.New("e"))
}

func TestHandleErrMessage_Empty(t *testing.T) {
	HandleErrMessage("") // should not panic
}

func TestHandleErrMessage_WithMsg(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	HandleErrMessage("msg")
}

func TestSimpleHandleErr_Nil(t *testing.T) {
	SimpleHandleErr(nil, "msg") // should not panic
}

func TestSimpleHandleErr_WithErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	SimpleHandleErr(errors.New("e"), "msg")
}

func TestSimpleHandleErrMany_Nil(t *testing.T) {
	SimpleHandleErrMany("msg")      // nil
	SimpleHandleErrMany("msg", nil) // only nil errors
}

func TestSimpleHandleErrMany_WithErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	SimpleHandleErrMany("msg", errors.New("e"))
}

func TestMustBeEmpty_Nil(t *testing.T) {
	MustBeEmpty(nil) // should not panic
}

func TestMustBeEmpty_WithErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	MustBeEmpty(errors.New("e"))
}

func TestHandleCompiledErrorGetter_Nil(t *testing.T) {
	HandleCompiledErrorGetter(nil) // should not panic
}

func TestHandleCompiledErrorWithTracesGetter_Nil(t *testing.T) {
	HandleCompiledErrorWithTracesGetter(nil)
}

func TestHandleErrorGetter_Nil(t *testing.T) {
	HandleErrorGetter(nil)
}

func TestHandleFullStringsWithTracesGetter_Nil(t *testing.T) {
	HandleFullStringsWithTracesGetter(nil)
}

func TestPrintError_Nil(t *testing.T) {
	PrintError(nil)
}

func TestPrintError_WithErr(t *testing.T) {
	PrintError(errors.New("e"))
}

func TestPrintErrorWithTestIndex_Nil(t *testing.T) {
	PrintErrorWithTestIndex(0, "test", nil)
}

func TestPrintErrorWithTestIndex_WithErr(t *testing.T) {
	PrintErrorWithTestIndex(0, "test", errors.New("e"))
}

func TestPanicOnIndexOutOfRange_Valid(t *testing.T) {
	PanicOnIndexOutOfRange(5, []int{0, 1, 2})
}

func TestPanicOnIndexOutOfRange_OutOfRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	PanicOnIndexOutOfRange(2, []int{5})
}
