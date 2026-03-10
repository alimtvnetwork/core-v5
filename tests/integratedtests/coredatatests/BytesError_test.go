package coredatatests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata"
)

// ===== BytesError Tests =====

// Note: String nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_String_WithBytes(t *testing.T) {
	be := &coredata.BytesError{
		Bytes: []byte("hello"),
	}

	got := be.String()
	if got != "hello" {
		t.Errorf("BytesError.String() = %q, want %q", got, "hello")
	}
}

func Test_BytesError_String_EmptyBytes(t *testing.T) {
	be := &coredata.BytesError{}

	got := be.String()
	if got != "" {
		t.Errorf("BytesError.String() on empty = %q, want empty", got)
	}
}

func Test_BytesError_String_CachesResult(t *testing.T) {
	be := &coredata.BytesError{
		Bytes: []byte("cached"),
	}

	first := be.String()
	second := be.String()

	if first != second {
		t.Error("expected cached string to be identical on second call")
	}
}

func Test_BytesError_HasError_True(t *testing.T) {
	be := &coredata.BytesError{
		Error: errors.New("some error"),
	}

	if !be.HasError() {
		t.Error("expected HasError() = true")
	}
}

func Test_BytesError_HasError_False(t *testing.T) {
	be := &coredata.BytesError{}

	if be.HasError() {
		t.Error("expected HasError() = false")
	}
}

// Note: HasError nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_IsEmptyError_True(t *testing.T) {
	be := &coredata.BytesError{}

	if !be.IsEmptyError() {
		t.Error("expected IsEmptyError() = true")
	}
}

// Note: IsEmptyError nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_IsEmptyError_False(t *testing.T) {
	be := &coredata.BytesError{
		Error: errors.New("err"),
	}

	if be.IsEmptyError() {
		t.Error("expected IsEmptyError() = false")
	}
}

func Test_BytesError_HasBytes_True(t *testing.T) {
	be := &coredata.BytesError{
		Bytes: []byte("data"),
	}

	if !be.HasBytes() {
		t.Error("expected HasBytes() = true")
	}
}

func Test_BytesError_HasBytes_False_NilBytes(t *testing.T) {
	be := &coredata.BytesError{}

	if be.HasBytes() {
		t.Error("expected HasBytes() = false for nil bytes")
	}
}

func Test_BytesError_HasBytes_False_EmptyBytes(t *testing.T) {
	be := &coredata.BytesError{
		Bytes: []byte{},
	}

	if be.HasBytes() {
		t.Error("expected HasBytes() = false for empty bytes")
	}
}

func Test_BytesError_HasBytes_False_EmptyJson(t *testing.T) {
	be := &coredata.BytesError{
		Bytes: []byte("{}"),
	}

	if be.HasBytes() {
		t.Error("expected HasBytes() = false for empty JSON {}")
	}
}

func Test_BytesError_HasBytes_False_WithError(t *testing.T) {
	be := &coredata.BytesError{
		Bytes: []byte("data"),
		Error: errors.New("err"),
	}

	if be.HasBytes() {
		t.Error("expected HasBytes() = false when error is present")
	}
}

// Note: Length nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_Length_WithBytes(t *testing.T) {
	be := &coredata.BytesError{
		Bytes: []byte("hello"),
	}

	got := be.Length()
	if got != 5 {
		t.Errorf("BytesError.Length() = %d, want 5", got)
	}
}

// Note: IsEmpty nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_IsEmpty_EmptyBytes(t *testing.T) {
	be := &coredata.BytesError{
		Bytes: []byte{},
	}

	if !be.IsEmpty() {
		t.Error("expected IsEmpty() = true for empty bytes")
	}
}

func Test_BytesError_IsEmpty_False(t *testing.T) {
	be := &coredata.BytesError{
		Bytes: []byte("data"),
	}

	if be.IsEmpty() {
		t.Error("expected IsEmpty() = false for non-empty bytes")
	}
}

func Test_BytesError_HandleError_NoError(t *testing.T) {
	be := &coredata.BytesError{}

	// Should not panic
	be.HandleError()
}

// Note: HandleError nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_NilReceiver(t *testing.T) {
	for caseIndex, tc := range bytesErrorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

func Test_BytesError_HandleError_Panics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected HandleError to panic")
		}
	}()

	be := &coredata.BytesError{
		Error: errors.New("boom"),
	}
	be.HandleError()
}

func Test_BytesError_HandleErrorWithMsg_NoError(t *testing.T) {
	be := &coredata.BytesError{}

	// Should not panic
	be.HandleErrorWithMsg("prefix: ")
}

func Test_BytesError_HandleErrorWithMsg_PanicsWithMsg(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected HandleErrorWithMsg to panic")
			return
		}

		msg, ok := r.(string)
		if !ok {
			t.Error("expected panic value to be string")
			return
		}

		if msg != "prefix: boom" {
			t.Errorf("panic message = %q, want %q", msg, "prefix: boom")
		}
	}()

	be := &coredata.BytesError{
		Error: errors.New("boom"),
	}
	be.HandleErrorWithMsg("prefix: ")
}

func Test_BytesError_HandleErrorWithMsg_PanicsEmptyMsg(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected HandleErrorWithMsg to panic")
		}
	}()

	be := &coredata.BytesError{
		Error: errors.New("boom"),
	}
	be.HandleErrorWithMsg("")
}

func Test_BytesError_CombineErrorWithRef_NoError(t *testing.T) {
	be := &coredata.BytesError{}

	got := be.CombineErrorWithRef("ref1")
	if got != "" {
		t.Errorf("CombineErrorWithRef on no error = %q, want empty", got)
	}
}

func Test_BytesError_CombineErrorWithRefError_NoError(t *testing.T) {
	be := &coredata.BytesError{}

	got := be.CombineErrorWithRefError("ref1")
	if got != nil {
		t.Errorf("CombineErrorWithRefError on no error = %v, want nil", got)
	}
}

func Test_BytesError_CombineErrorWithRefError_WithError(t *testing.T) {
	be := &coredata.BytesError{
		Error: errors.New("something failed"),
	}

	got := be.CombineErrorWithRefError("ref1", "ref2")
	if got == nil {
		t.Error("expected non-nil error from CombineErrorWithRefError")
	}
}
