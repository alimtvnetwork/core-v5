package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ── LazyRegex.IsMatchBytes — failure path ──
// Covers LazyRegex.go L244-246

func Test_Cov8_LazyRegex_IsMatchBytes_InvalidPattern(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy("[invalid")

	// Act
	result := lr.IsMatchBytes([]byte("test"))

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": false}
	expected.ShouldBeEqual(t, 0, "IsMatchBytes invalid pattern", actual)
}

// ── lazyRegexMap — IsEmpty, IsEmptyLock ──
// Covers lazyRegexMap.go L9-11, L13-18

func Test_Cov8_LazyRegex_IsEmpty_Via_Compile(t *testing.T) {
	// Arrange — creating and compiling a lazy regex populates the internal map,
	// exercising IsEmpty/HasAnyItem paths internally
	lr := regexnew.New.LazyLock(`\d+`)

	// Act
	_, err := lr.Compile()

	// Assert
	actual := args.Map{"hasError": err != nil, "compiled": lr.IsCompiled()}
	expected := args.Map{"hasError": false, "compiled": true}
	expected.ShouldBeEqual(t, 0, "LazyLock compile populates map", actual)
}

// ── lazyRegexMap — HasAnyItem, HasAnyItemLock ──
// Covers lazyRegexMap.go L20-22, L24-29

func Test_Cov8_LazyRegex_HasAnyItem_Via_Match(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[a-z]+`)

	// Act
	result := lr.IsMatch("hello")

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": true}
	expected.ShouldBeEqual(t, 0, "Lazy match exercises map paths", actual)
}

// ── lazyRegexMap — Length, LengthLock ──
// Covers lazyRegexMap.go L31-36, L39-44

func Test_Cov8_LazyRegex_Length_Via_CreateLockIf(t *testing.T) {
	// Arrange — CreateLockIf exercises the LockIf path
	lr := regexnew.New.LazyLock(`unique-pattern-length-test`)

	// Act
	_, err := lr.Compile()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "LazyLock length path", actual)
}

// ── lazyRegexMap — Has, HasLock ──
// Covers lazyRegexMap.go L46-50, L52-59

func Test_Cov8_LazyRegex_Has_Via_RepeatedCompile(t *testing.T) {
	// Arrange — compile same pattern twice, second time hits Has/existing path
	pattern := `unique-pattern-has-test`
	lr1 := regexnew.New.Lazy(pattern)
	lr1.Compile()

	lr2 := regexnew.New.Lazy(pattern)

	// Act
	result := lr2.IsMatch("unique-pattern-has-test")

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": true}
	expected.ShouldBeEqual(t, 0, "Lazy repeated pattern uses Has", actual)
}

// ── lazyRegexMap — CreateOrExistingLockIf ──
// Covers lazyRegexMap.go L92-99

func Test_Cov8_LazyRegex_CreateOrExistingLockIf(t *testing.T) {
	// Arrange — use LazyLock which internally calls CreateOrExistingLock
	lr := regexnew.New.LazyLock(`lockif-pattern-test`)

	// Act
	_, err := lr.Compile()
	result := lr.IsMatch("lockif-pattern-test")

	// Assert
	actual := args.Map{"hasError": err != nil, "matches": result}
	expected := args.Map{"hasError": false, "matches": true}
	expected.ShouldBeEqual(t, 0, "CreateOrExistingLockIf path", actual)
}

// ── lazyRegexMap — createLazyRegex ──
// Covers lazyRegexMap.go L114-119

func Test_Cov8_LazyRegex_CreateLazyRegex_Via_NewLazyCreator(t *testing.T) {
	// Arrange & Act — New.LazyRegex.New exercises createDefaultLazyRegex
	lr := regexnew.New.Lazy(`create-lazy-regex-test`)
	result := lr.IsMatch("create-lazy-regex-test")

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": true}
	expected.ShouldBeEqual(t, 0, "createLazyRegex path", actual)
}

// ── prettyJson — nil and error paths ──
// Covers prettyJson.go L16-18, L22-24

func Test_Cov8_PrettyJson_NonNil(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`abc`)

	// Act — PrettyJsonString exercises prettyJson internally
	result := lr.PrettyJsonString()

	// Assert
	actual := args.Map{"nonEmpty": result != ""}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "prettyJson non-nil", actual)
}

// ── regExMatchValidationError — nil regex path ──
// Covers regExMatchValidationError.go L21-26

func Test_Cov8_MatchError_InvalidPattern(t *testing.T) {
	// Arrange & Act — invalid pattern gives compile error + nil regex
	err := regexnew.MatchError("[invalid", "test")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "MatchError invalid pattern", actual)
}
