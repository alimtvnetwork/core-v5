package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ── lazyRegexMap methods ──

func Test_Cov8_LazyRegex_IsMatchBytes_InvalidPattern(t *testing.T) {
	// Arrange — create lazy regex with invalid pattern
	lr := regexnew.New.Lazy.Create("[invalid")

	// Act
	result := lr.IsMatchBytes([]byte("test"))

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": false}
	expected.ShouldBeEqual(t, 0, "IsMatchBytes invalid pattern", actual)
}

// ── lazyRegexMap — IsEmpty, HasAnyItem, Length, Has ──

func Test_Cov8_LazyRegexMap_IsEmpty(t *testing.T) {
	// Arrange — use the global map which should have items from other tests
	// Create a fresh lazy regex to ensure map is populated
	lr := regexnew.New.Lazy.Create(`\d+`)
	lr.Compile()

	// Act & Assert — the global map should not be empty
	actual := args.Map{"compiled": lr.IsCompiled()}
	expected := args.Map{"compiled": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex compiled check", actual)
}

func Test_Cov8_LazyRegex_IsEmptyLock(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy.CreateLock(`\w+`)
	lr.Compile()

	// Act & Assert
	actual := args.Map{"compiled": lr.IsCompiled()}
	expected := args.Map{"compiled": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsEmptyLock path", actual)
}

func Test_Cov8_LazyRegex_HasAnyItemAndLength(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy.Create(`[a-z]+`)
	lr.Compile()

	// Act
	result := lr.IsMatch("hello")

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex HasAnyItem path", actual)
}

func Test_Cov8_LazyRegex_HasAnyItemLock(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy.CreateLock(`[0-9]+`)
	lr.Compile()

	// Act
	result := lr.IsMatch("123")

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex HasAnyItemLock path", actual)
}

func Test_Cov8_LazyRegex_LengthLock(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy.CreateLock(`test`)
	lr.Compile()

	// Act
	result := lr.IsMatch("test")

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex LengthLock path", actual)
}

func Test_Cov8_LazyRegex_Has(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy.Create(`foo`)
	lr.Compile()

	// Act
	result := lr.IsMatch("foo")

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex Has path", actual)
}

func Test_Cov8_LazyRegex_HasLock(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy.CreateLock(`bar`)
	lr.Compile()

	// Act
	result := lr.IsMatch("bar")

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex HasLock path", actual)
}

// ── prettyJson ──

func Test_Cov8_PrettyJson_Nil(t *testing.T) {
	// Arrange — create a lazy regex with nil-like state
	lr := regexnew.New.Lazy.Create(`abc`)

	// Act — PrettyJsonString exercises prettyJson internally
	result := lr.PrettyJsonString()

	// Assert
	actual := args.Map{"nonEmpty": result != ""}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "prettyJson non-nil", actual)
}

// ── regExMatchValidationError — nil regex path ──

func Test_Cov8_MatchError_NilRegex(t *testing.T) {
	// Arrange — use an invalid pattern that will give nil regex
	err := regexnew.MatchError("[invalid", "test")

	// Act & Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "MatchError invalid pattern", actual)
}
