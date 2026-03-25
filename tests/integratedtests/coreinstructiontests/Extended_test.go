package coreinstructiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// TestStringSearch_IsMatch verifies match logic.
func TestStringSearch_IsMatch(t *testing.T) {
	for _, tc := range stringSearchIsMatchCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := tc.search.IsMatch(tc.content)

			// Assert
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestStringSearch_NilIsMatch verifies nil receiver returns true.
func TestStringSearch_NilIsMatch(t *testing.T) {
	// Arrange
	var s *coreinstruction.StringSearch

	// Act & Assert
	if !s.IsMatch("anything") {
		t.Error("nil StringSearch.IsMatch should return true")
	}
}

// TestStringSearch_IsEmpty verifies nil check.
func TestStringSearch_IsEmpty(t *testing.T) {
	var s *coreinstruction.StringSearch
	if !s.IsEmpty() { t.Error("nil should be empty") }
	if s.IsExist() { t.Error("nil should not exist") }
	if s.Has() { t.Error("nil should not have") }
}

// TestStringSearch_IsAllMatch verifies all-match logic.
func TestStringSearch_IsAllMatch(t *testing.T) {
	// Arrange
	s := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}

	// Act & Assert
	if !s.IsAllMatch("hello", "hello") {
		t.Error("all 'hello' should match")
	}
	if s.IsAllMatch("hello", "world") {
		t.Error("mixed should fail")
	}
	if !s.IsAllMatch() {
		t.Error("empty contents should return true")
	}
}

// TestStringSearch_IsAnyMatchFailed verifies any-fail logic.
func TestStringSearch_IsAnyMatchFailed(t *testing.T) {
	s := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}
	if !s.IsAnyMatchFailed("hello", "world") {
		t.Error("expected any match failed")
	}
}

// TestStringSearch_IsMatchFailed verifies match failure.
func TestStringSearch_IsMatchFailed(t *testing.T) {
	s := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}
	if !s.IsMatchFailed("world") {
		t.Error("expected match failed")
	}
}

// TestStringSearch_VerifyError verifies error on mismatch.
func TestStringSearch_VerifyError(t *testing.T) {
	// Nil returns nil
	var nilS *coreinstruction.StringSearch
	if nilS.VerifyError("x") != nil {
		t.Error("nil should return nil error")
	}

	// Equal match → nil error
	s := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Equal,
		Search:        "hello",
	}
	if s.VerifyError("hello") != nil {
		t.Error("matching content should return nil error")
	}
	if s.VerifyError("world") == nil {
		t.Error("mismatched content should return error")
	}
}

// TestStringCompare_IsMatch verifies compare match.
func TestStringCompare_IsMatch(t *testing.T) {
	// Arrange
	sc := coreinstruction.NewStringCompare(
		stringcompareas.Equal,
		false,
		"test",
		"test",
	)

	// Act & Assert
	if !sc.IsMatch() {
		t.Error("equal content should match")
	}
	if sc.IsMatchFailed() {
		t.Error("should not be match failed")
	}
}

// TestStringCompare_Nil verifies nil receiver.
func TestStringCompare_Nil(t *testing.T) {
	var sc *coreinstruction.StringCompare
	if !sc.IsMatch() { t.Error("nil should match") }
	if sc.IsDefined() { t.Error("nil should not be defined") }
	if !sc.IsInvalid() { t.Error("nil should be invalid") }
	if sc.VerifyError() != nil { t.Error("nil should return nil error") }
}

// TestStringCompare_VerifyError verifies verify error.
func TestStringCompare_VerifyError(t *testing.T) {
	sc := coreinstruction.NewStringCompare(
		stringcompareas.Equal,
		false,
		"expected",
		"actual",
	)
	if sc.VerifyError() == nil {
		t.Error("mismatched should return error")
	}
}

// TestNewStringCompareEqual verifies constructor.
func TestNewStringCompareEqual(t *testing.T) {
	sc := coreinstruction.NewStringCompareEqual("a", "a")
	if !sc.IsMatch() {
		t.Error("equal should match")
	}
}

// TestNewStringCompareStartsWith verifies constructor.
func TestNewStringCompareStartsWith(t *testing.T) {
	sc := coreinstruction.NewStringCompareStartsWith(false, "hel", "hello")
	if !sc.IsMatch() {
		t.Error("should start with 'hel'")
	}
}

// TestNewStringCompareEndsWith verifies constructor.
func TestNewStringCompareEndsWith(t *testing.T) {
	sc := coreinstruction.NewStringCompareEndsWith(false, "llo", "hello")
	if !sc.IsMatch() {
		t.Error("should end with 'llo'")
	}
}

// TestNewStringCompareContains verifies constructor.
func TestNewStringCompareContains(t *testing.T) {
	sc := coreinstruction.NewStringCompareContains(false, "ell", "hello")
	if !sc.IsMatch() {
		t.Error("should contain 'ell'")
	}
}

// TestNewStringCompareRegex verifies regex constructor.
func TestNewStringCompareRegex(t *testing.T) {
	sc := coreinstruction.NewStringCompareRegex("^he.*o$", "hello")
	if !sc.IsMatch() {
		t.Error("regex should match")
	}
}

// TestNewStringCompare_IgnoreCase verifies case-insensitive.
func TestNewStringCompare_IgnoreCase(t *testing.T) {
	sc := coreinstruction.NewStringCompare(
		stringcompareas.Equal,
		true,
		"HELLO",
		"hello",
	)
	if !sc.IsMatch() {
		t.Error("ignore case equal should match")
	}
}

// TestStringCompare_VerifyError_Regex verifies regex verify error.
func TestStringCompare_VerifyError_Regex(t *testing.T) {
	sc := coreinstruction.NewStringCompareRegex("^abc$", "abc")
	if sc.VerifyError() != nil {
		t.Error("matching regex should return nil error")
	}
	sc2 := coreinstruction.NewStringCompareRegex("^abc$", "xyz")
	if sc2.VerifyError() == nil {
		t.Error("non-matching regex should return error")
	}
}

// TestStringSearch_VerifyError_Regex verifies regex verify on StringSearch.
func TestStringSearch_VerifyError_Regex(t *testing.T) {
	s := &coreinstruction.StringSearch{
		CompareMethod: stringcompareas.Regex,
		Search:        "^test$",
		BaseIsIgnoreCase: corecomparator.BaseIsIgnoreCase{
			IsIgnoreCase: false,
		},
	}
	if s.VerifyError("test") != nil {
		t.Error("matching regex should return nil error")
	}
	if s.VerifyError("nope") == nil {
		t.Error("non-matching regex should return error")
	}
}
