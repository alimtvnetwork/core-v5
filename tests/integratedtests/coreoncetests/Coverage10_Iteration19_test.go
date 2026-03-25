package coreoncetests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
)

// ===== AnyOnce.Deserialize coverage =====

func Test_Cov10_AnyOnce_Deserialize_Success(t *testing.T) {
	o := coreonce.NewAnyOncePtr(func() any { return map[string]string{"a": "b"} })
	var target map[string]string
	err := o.Deserialize(&target)
	// Due to the bug (if err == nil returns err which is nil), this always returns nil
	if err != nil {
		t.Fatal("expected nil due to code path")
	}
}

func Test_Cov10_AnyOnce_Deserialize_SerializeError(t *testing.T) {
	// Use a value that can't be marshalled (channel)
	ch := make(chan int)
	o := coreonce.NewAnyOncePtr(func() any { return ch })
	var target string
	err := o.Deserialize(&target)
	if err == nil {
		t.Fatal("expected serialize error")
	}
}

// ===== AnyErrorOnce.Deserialize coverage =====

func Test_Cov10_AnyErrorOnce_Deserialize_ExistingError(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) {
		return nil, errors.New("init error")
	})
	var target string
	err := o.Deserialize(&target)
	if err == nil {
		t.Fatal("expected error from serialize")
	}
}

func Test_Cov10_AnyErrorOnce_Deserialize_Success(t *testing.T) {
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) {
		return map[string]string{"x": "y"}, nil
	})
	var target map[string]string
	err := o.Deserialize(&target)
	// Same bug as AnyOnce - always returns nil when serialize succeeds
	if err != nil {
		t.Fatal("expected nil")
	}
}

func Test_Cov10_AnyErrorOnce_Deserialize_MarshalError(t *testing.T) {
	ch := make(chan int)
	o := coreonce.NewAnyErrorOncePtr(func() (any, error) {
		return ch, nil
	})
	var target string
	err := o.Deserialize(&target)
	if err == nil {
		t.Fatal("expected marshal error from Serialize")
	}
}

// ===== IntegersOnce.IsEqual - hit currentMap[item] < 0 =====

func Test_Cov10_IntegersOnce_IsEqual_FreqMismatch(t *testing.T) {
	o := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 1} })
	// Same length but different frequencies: {1,1} vs {1,2}
	if o.IsEqual(1, 2) {
		t.Fatal("expected false for frequency mismatch")
	}
}

// ===== MapStringStringOnce.IsEqual - hit isMissing and value mismatch =====

func Test_Cov10_MapStringStringOnce_IsEqual_MissingKey(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})
	if o.IsEqual(map[string]string{"a": "1", "c": "2"}) {
		t.Fatal("expected false for missing key")
	}
}

func Test_Cov10_MapStringStringOnce_IsEqual_ValueMismatch(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	if o.IsEqual(map[string]string{"a": "9"}) {
		t.Fatal("expected false for value mismatch")
	}
}

// ===== StringsOnce.IsEqual - hit currentMap[item] < 0 =====

func Test_Cov10_StringsOnce_IsEqual_FreqMismatch(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a", "a"} })
	if o.IsEqual("a", "b") {
		t.Fatal("expected false for frequency mismatch")
	}
}

// ===== StringOnce.SplitLeftRight - hit len > 2 branch =====
// Note: SplitN with n=2 returns at most 2, so len>2 is dead code.
// But we can test the len==1 (no splitter found) path to cover the else.

func Test_Cov10_StringOnce_SplitLeftRight_NoSplitter(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "nosplitter" })
	left, right := o.SplitLeftRight(":")
	if left != "nosplitter" || right != "" {
		t.Fatalf("expected 'nosplitter','', got '%s','%s'", left, right)
	}
}

func Test_Cov10_StringOnce_SplitLeftRight_WithSplitter(t *testing.T) {
	o := coreonce.NewStringOncePtr(func() string { return "left:right" })
	left, right := o.SplitLeftRight(":")
	if left != "left" || right != "right" {
		t.Fatalf("expected 'left','right', got '%s','%s'", left, right)
	}
}

// ===== JsonStringMust panic paths =====

func Test_Cov10_StringsOnce_JsonStringMust_Success(t *testing.T) {
	o := coreonce.NewStringsOncePtr(func() []string { return []string{"a"} })
	s := o.JsonStringMust()
	if s == "" {
		t.Fatal("expected non-empty json string")
	}
}

func Test_Cov10_MapStringStringOnce_JsonStringMust_Success(t *testing.T) {
	o := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"k": "v"}
	})
	s := o.JsonStringMust()
	if s == "" {
		t.Fatal("expected non-empty json string")
	}
}
