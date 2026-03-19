package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
)

// === anyItemConverter (converters.AnyTo) uncovered branches ===

func Test_Cov_AnyTo_ToStringsUsingProcessor_Break(t *testing.T) {
	result := converters.AnyTo.ToStringsUsingProcessor(
		false,
		func(index int, in any) (string, bool, bool) {
			return "x", true, true // take + break
		},
		[]string{"a", "b"},
	)
	if len(result) != 1 {
		t.Errorf("expected 1 got %d", len(result))
	}
}

func Test_Cov_AnyTo_ToStringsUsingSimpleProcessor_Empty(t *testing.T) {
	result := converters.AnyTo.ToStringsUsingSimpleProcessor(
		false,
		func(index int, in any) string { return "x" },
		[]string{},
	)
	if len(result) != 0 {
		t.Errorf("expected 0 got %d", len(result))
	}
}

func Test_Cov_AnyTo_ToPrettyJson_Error(t *testing.T) {
	// channels can't be marshaled
	ch := make(chan int)
	result := converters.AnyTo.ToPrettyJson(ch)
	if result != "" {
		t.Error("expected empty for unmarshalable")
	}
}

func Test_Cov_AnyTo_Bytes_Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for unmarshalable")
		}
	}()
	ch := make(chan int)
	converters.AnyTo.Bytes(ch)
}

// === stringTo uncovered branches ===

func Test_Cov_StringTo_Float64Must_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()
	converters.StringTo.Float64Must("notanumber")
}

// === stringsTo uncovered branches ===

func Test_Cov_StringsTo_IntegersOptionPanic_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()
	converters.StringsTo.IntegersOptionPanic(true, "not_int")
}

func Test_Cov_StringsTo_BytesConditional_Break(t *testing.T) {
	result := converters.StringsTo.BytesConditional(
		func(in string) (byte, bool, bool) {
			return 0, true, true
		},
		[]string{"1", "2"},
	)
	if len(result) != 1 {
		t.Errorf("expected 1 got %d", len(result))
	}
}

func Test_Cov_StringsTo_BytesMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()
	converters.StringsTo.BytesMust("not_byte")
}

func Test_Cov_StringsTo_Float64sMust_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()
	converters.StringsTo.Float64sMust("not_float")
}

func Test_Cov_StringsTo_Float64sConditional_Break(t *testing.T) {
	result := converters.StringsTo.Float64sConditional(
		func(in string) (float64, bool, bool) {
			return 0, true, true
		},
		[]string{"1.0", "2.0"},
	)
	if len(result) != 1 {
		t.Errorf("expected 1 got %d", len(result))
	}
}
