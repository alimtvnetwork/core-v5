package regexnewtests

import (
	"regexp"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ══════════════════════════════════════════════════════════════════════════════
// Migrated from regexnew/unexported_test.go — public-API-only tests.
// Raises integrated coverage from 87.4% toward 100%.
// ══════════════════════════════════════════════════════════════════════════════

// ── Create ──

func Test_Src11_Create_New(t *testing.T) {
	// Arrange & Act
	r, err := regexnew.Create(`^src11-new$`)

	// Assert
	actual := args.Map{
		"notNil":  r != nil,
		"noError": err == nil,
	}
	expected := args.Map{
		"notNil":  true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "Create returns regex -- new pattern", actual)
}

func Test_Src11_Create_Cached(t *testing.T) {
	// Arrange
	r1, _ := regexnew.Create(`^src11-cached$`)

	// Act
	r2, _ := regexnew.Create(`^src11-cached$`)

	// Assert
	actual := args.Map{"samePointer": r1 == r2}
	expected := args.Map{"samePointer": true}
	expected.ShouldBeEqual(t, 0, "Create returns same pointer -- cached", actual)
}

func Test_Src11_Create_Invalid(t *testing.T) {
	// Arrange & Act
	_, err := regexnew.Create(`[invalid`)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Create returns error -- invalid pattern", actual)
}

// ── CreateLock ──

func Test_Src11_CreateLock_Valid(t *testing.T) {
	// Arrange & Act
	r, err := regexnew.CreateLock(`^src11-lock$`)

	// Assert
	actual := args.Map{
		"notNil":  r != nil,
		"noError": err == nil,
	}
	expected := args.Map{
		"notNil":  true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "CreateLock returns regex -- valid", actual)
}

// ── CreateMust ──

func Test_Src11_CreateMust_Valid(t *testing.T) {
	// Arrange & Act
	r := regexnew.CreateMust(`^src11-must$`)

	// Assert
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMust returns regex -- valid", actual)
}

func Test_Src11_CreateMust_Cached(t *testing.T) {
	// Arrange
	r1 := regexnew.CreateMust(`^src11-mustcached$`)

	// Act
	r2 := regexnew.CreateMust(`^src11-mustcached$`)

	// Assert
	actual := args.Map{"samePointer": r1 == r2}
	expected := args.Map{"samePointer": true}
	expected.ShouldBeEqual(t, 0, "CreateMust returns same pointer -- cached", actual)
}

func Test_Src11_CreateMust_Panic(t *testing.T) {
	// Arrange
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		regexnew.CreateMust(`[invalid`)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "CreateMust panics -- invalid pattern", actual)
}

// ── MatchError (no lock) ──

func Test_Src11_MatchError_Match(t *testing.T) {
	// Arrange & Act
	err := regexnew.MatchError(`^\d+$`, "123")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "MatchError returns nil -- match", actual)
}

func Test_Src11_MatchError_NoMatch(t *testing.T) {
	// Arrange & Act
	err := regexnew.MatchError(`^\d+$`, "abc")

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns error -- no match", actual)
}

func Test_Src11_MatchError_InvalidPattern(t *testing.T) {
	// Arrange & Act
	err := regexnew.MatchError(`[invalid`, "abc")

	// Assert
	actual := args.Map{
		"hasError":       err != nil,
		"containsFailed": strings.Contains(err.Error(), "compile failed"),
	}
	expected := args.Map{
		"hasError":       true,
		"containsFailed": true,
	}
	expected.ShouldBeEqual(t, 0, "MatchError returns compile error -- invalid pattern", actual)
}

// ── IsMatchLock invalid regex ──

func Test_Src11_IsMatchLock_InvalidRegex(t *testing.T) {
	// Arrange & Act
	result := regexnew.IsMatchLock(`[invalid`, "abc")

	// Assert
	actual := args.Map{"isMatch": result}
	expected := args.Map{"isMatch": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns false -- invalid regex", actual)
}

// ── LazyRegex: nil receiver methods ──

func Test_Src11_LazyRegex_IsNull_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act & Assert
	actual := args.Map{"isNull": lr.IsNull()}
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsNull returns true -- nil", actual)
}

func Test_Src11_LazyRegex_IsUndefined_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act & Assert
	actual := args.Map{"isUndefined": lr.IsUndefined()}
	expected := args.Map{"isUndefined": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsUndefined returns true -- nil", actual)
}

func Test_Src11_LazyRegex_IsApplicable_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act & Assert
	actual := args.Map{"isApplicable": lr.IsApplicable()}
	expected := args.Map{"isApplicable": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsApplicable returns false -- nil", actual)
}

func Test_Src11_LazyRegex_String_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act & Assert
	actual := args.Map{"value": lr.String()}
	expected := args.Map{"value": ""}
	expected.ShouldBeEqual(t, 0, "LazyRegex.String returns empty -- nil", actual)
}

func Test_Src11_LazyRegex_FullString_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act & Assert
	actual := args.Map{"value": lr.FullString()}
	expected := args.Map{"value": ""}
	expected.ShouldBeEqual(t, 0, "LazyRegex.FullString returns empty -- nil", actual)
}

func Test_Src11_LazyRegex_Pattern_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act & Assert
	actual := args.Map{"value": lr.Pattern()}
	expected := args.Map{"value": ""}
	expected.ShouldBeEqual(t, 0, "LazyRegex.Pattern returns empty -- nil", actual)
}

func Test_Src11_LazyRegex_OnRequiredCompiled_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act
	err := lr.OnRequiredCompiled()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.OnRequiredCompiled returns error -- nil", actual)
}

func Test_Src11_LazyRegex_OnRequiredCompiledMust_Panic_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		lr.OnRequiredCompiledMust()
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.OnRequiredCompiledMust panics -- nil", actual)
}

func Test_Src11_LazyRegex_HasAnyIssues_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act & Assert
	actual := args.Map{"hasIssues": lr.HasAnyIssues()}
	expected := args.Map{"hasIssues": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.HasAnyIssues returns true -- nil", actual)
}

func Test_Src11_LazyRegex_IsInvalid_Nil(t *testing.T) {
	// Arrange
	var lr *regexnew.LazyRegex

	// Act & Assert
	actual := args.Map{"isInvalid": lr.IsInvalid()}
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsInvalid returns true -- nil", actual)
}

// ── LazyRegex: valid pattern methods ──

func Test_Src11_LazyRegex_IsDefined_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-defined$`)

	// Act & Assert
	actual := args.Map{"isDefined": lr.IsDefined()}
	expected := args.Map{"isDefined": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsDefined returns true -- valid", actual)
}

func Test_Src11_LazyRegex_IsApplicable_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-applicable$`)

	// Act & Assert
	actual := args.Map{"isApplicable": lr.IsApplicable()}
	expected := args.Map{"isApplicable": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsApplicable returns true -- valid", actual)
}

func Test_Src11_LazyRegex_IsApplicable_Cached(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-appl-cached$`)
	lr.Compile()

	// Act & Assert
	actual := args.Map{"isApplicable": lr.IsApplicable()}
	expected := args.Map{"isApplicable": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsApplicable returns true -- cached", actual)
}

func Test_Src11_LazyRegex_IsApplicable_InvalidPattern(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid-src11`)

	// Act & Assert
	actual := args.Map{"isApplicable": lr.IsApplicable()}
	expected := args.Map{"isApplicable": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsApplicable returns false -- invalid pattern", actual)
}

func Test_Src11_LazyRegex_Compile_Cached(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-compile-cached$`)
	lr.Compile()

	// Act
	r, err := lr.Compile()

	// Assert
	actual := args.Map{
		"notNil":  r != nil,
		"noError": err == nil,
	}
	expected := args.Map{
		"notNil":  true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.Compile returns cached -- already compiled", actual)
}

func Test_Src11_LazyRegex_CompileMust_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-compilemust$`)

	// Act
	r := lr.CompileMust()

	// Assert
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.CompileMust returns regex -- valid", actual)
}

func Test_Src11_LazyRegex_CompileMust_Panic(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid-src11-panic`)
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		lr.CompileMust()
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.CompileMust panics -- invalid", actual)
}

func Test_Src11_LazyRegex_OnRequiredCompiled_AlreadyCompiled(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-onreq$`)
	lr.Compile()

	// Act
	err := lr.OnRequiredCompiled()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.OnRequiredCompiled returns nil -- already compiled", actual)
}

func Test_Src11_LazyRegex_OnRequiredCompiledMust_Success(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-onreqmust$`)

	// Act (should not panic)
	lr.OnRequiredCompiledMust()

	// Assert
	actual := args.Map{"isCompiled": lr.IsCompiled()}
	expected := args.Map{"isCompiled": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.OnRequiredCompiledMust succeeds -- valid", actual)
}

func Test_Src11_LazyRegex_HasError_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-haserror$`)

	// Act & Assert
	actual := args.Map{"hasError": lr.HasError()}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.HasError returns false -- valid", actual)
}

func Test_Src11_LazyRegex_HasError_Invalid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid-src11-haserr`)

	// Act & Assert
	actual := args.Map{"hasError": lr.HasError()}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.HasError returns true -- invalid", actual)
}

func Test_Src11_LazyRegex_HasAnyIssues_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-issues$`)

	// Act & Assert
	actual := args.Map{"hasIssues": lr.HasAnyIssues()}
	expected := args.Map{"hasIssues": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.HasAnyIssues returns false -- valid", actual)
}

func Test_Src11_LazyRegex_IsInvalid_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-invalid$`)

	// Act & Assert
	actual := args.Map{"isInvalid": lr.IsInvalid()}
	expected := args.Map{"isInvalid": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsInvalid returns false -- valid", actual)
}

func Test_Src11_LazyRegex_CompiledError_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-comperr$`)

	// Act
	err := lr.CompiledError()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.CompiledError returns nil -- valid", actual)
}

func Test_Src11_LazyRegex_Error_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-err$`)

	// Act
	err := lr.Error()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.Error returns nil -- valid", actual)
}

func Test_Src11_LazyRegex_MustBeSafe_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-mustbesafe$`)

	// Act (should not panic)
	lr.MustBeSafe()

	// Assert
	actual := args.Map{"isApplicable": lr.IsApplicable()}
	expected := args.Map{"isApplicable": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.MustBeSafe succeeds -- valid", actual)
}

func Test_Src11_LazyRegex_MustBeSafe_Panic(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid-src11-safe`)
	didPanic := false

	// Act
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		lr.MustBeSafe()
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.MustBeSafe panics -- invalid", actual)
}

func Test_Src11_LazyRegex_String_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-str$`)

	// Act & Assert
	actual := args.Map{"value": lr.String()}
	expected := args.Map{"value": `^src11-str$`}
	expected.ShouldBeEqual(t, 0, "LazyRegex.String returns pattern -- valid", actual)
}

func Test_Src11_LazyRegex_FullString_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-fullstr$`)

	// Act
	s := lr.FullString()

	// Assert
	actual := args.Map{"nonEmpty": s != ""}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.FullString returns non-empty -- valid", actual)
}

func Test_Src11_LazyRegex_Pattern_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-pattern$`)

	// Act & Assert
	actual := args.Map{"value": lr.Pattern()}
	expected := args.Map{"value": `^src11-pattern$`}
	expected.ShouldBeEqual(t, 0, "LazyRegex.Pattern returns pattern -- valid", actual)
}

func Test_Src11_LazyRegex_IsCompiled_BeforeAndAfter(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^src11-iscompiled$`)

	// Act & Assert (before)
	beforeCompiled := lr.IsCompiled()

	lr.Compile()
	afterCompiled := lr.IsCompiled()

	actual := args.Map{
		"before": beforeCompiled,
		"after":  afterCompiled,
	}
	expected := args.Map{
		"before": false,
		"after":  true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsCompiled transitions -- compile", actual)
}

// ── LazyRegex: match methods ──

func Test_Src11_LazyRegex_IsMatch_Match(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act & Assert
	actual := args.Map{"isMatch": lr.IsMatch("123")}
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsMatch returns true -- match", actual)
}

func Test_Src11_LazyRegex_IsMatch_NoMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act & Assert
	actual := args.Map{"isMatch": lr.IsMatch("abc")}
	expected := args.Map{"isMatch": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsMatch returns false -- no match", actual)
}

func Test_Src11_LazyRegex_IsMatch_Invalid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid-src11-match`)

	// Act & Assert
	actual := args.Map{"isMatch": lr.IsMatch("abc")}
	expected := args.Map{"isMatch": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsMatch returns false -- invalid", actual)
}

func Test_Src11_LazyRegex_IsMatchBytes_Match(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act & Assert
	actual := args.Map{"isMatch": lr.IsMatchBytes([]byte("123"))}
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsMatchBytes returns true -- match", actual)
}

func Test_Src11_LazyRegex_IsMatchBytes_Invalid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid-src11-bytes`)

	// Act & Assert
	actual := args.Map{"isMatch": lr.IsMatchBytes([]byte("abc"))}
	expected := args.Map{"isMatch": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsMatchBytes returns false -- invalid", actual)
}

func Test_Src11_LazyRegex_IsFailedMatch_Match(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act & Assert
	actual := args.Map{"isFailed": lr.IsFailedMatch("123")}
	expected := args.Map{"isFailed": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsFailedMatch returns false -- match", actual)
}

func Test_Src11_LazyRegex_IsFailedMatch_NoMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act & Assert
	actual := args.Map{"isFailed": lr.IsFailedMatch("abc")}
	expected := args.Map{"isFailed": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsFailedMatch returns true -- no match", actual)
}

func Test_Src11_LazyRegex_IsFailedMatch_Invalid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid-src11-fail`)

	// Act & Assert
	actual := args.Map{"isFailed": lr.IsFailedMatch("abc")}
	expected := args.Map{"isFailed": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsFailedMatch returns true -- invalid", actual)
}

func Test_Src11_LazyRegex_IsFailedMatchBytes_Match(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act & Assert
	actual := args.Map{"isFailed": lr.IsFailedMatchBytes([]byte("123"))}
	expected := args.Map{"isFailed": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsFailedMatchBytes returns false -- match", actual)
}

func Test_Src11_LazyRegex_IsFailedMatchBytes_NoMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)

	// Act & Assert
	actual := args.Map{"isFailed": lr.IsFailedMatchBytes([]byte("abc"))}
	expected := args.Map{"isFailed": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsFailedMatchBytes returns true -- no match", actual)
}

func Test_Src11_LazyRegex_IsFailedMatchBytes_Invalid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid-src11-failbytes`)

	// Act & Assert
	actual := args.Map{"isFailed": lr.IsFailedMatchBytes([]byte("abc"))}
	expected := args.Map{"isFailed": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsFailedMatchBytes returns true -- invalid", actual)
}

func Test_Src11_LazyRegex_FirstMatchLine_Match(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`(\d+)`)

	// Act
	match, isInvalid := lr.FirstMatchLine("abc123def")

	// Assert
	actual := args.Map{
		"match":     match,
		"isInvalid": isInvalid,
	}
	expected := args.Map{
		"match":     "123",
		"isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.FirstMatchLine returns match -- found", actual)
}

func Test_Src11_LazyRegex_FirstMatchLine_NoMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^zzz$`)

	// Act
	match, isInvalid := lr.FirstMatchLine("abc")

	// Assert
	actual := args.Map{
		"match":     match,
		"isInvalid": isInvalid,
	}
	expected := args.Map{
		"match":     "",
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.FirstMatchLine returns invalid -- no match", actual)
}

func Test_Src11_LazyRegex_FirstMatchLine_Invalid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid-src11-fml`)

	// Act
	match, isInvalid := lr.FirstMatchLine("abc")

	// Assert
	actual := args.Map{
		"match":     match,
		"isInvalid": isInvalid,
	}
	expected := args.Map{
		"match":     "",
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.FirstMatchLine returns invalid -- bad pattern", actual)
}

func Test_Src11_LazyRegex_MatchUsingFuncError_Match(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)
	fn := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }

	// Act
	err := lr.MatchUsingFuncError("123", fn)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.MatchUsingFuncError returns nil -- match", actual)
}

func Test_Src11_LazyRegex_MatchUsingFuncError_NoMatch(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^\d+$`)
	fn := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }

	// Act
	err := lr.MatchUsingFuncError("abc", fn)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.MatchUsingFuncError returns error -- no match", actual)
}

// ── newCreator methods ──

func Test_Src11_NewCreator_Lazy(t *testing.T) {
	// Arrange & Act
	lr := regexnew.New.Lazy(`^src11-lazy$`)

	// Assert
	actual := args.Map{
		"notNil":    lr != nil,
		"isDefined": lr.IsDefined(),
	}
	expected := args.Map{
		"notNil":    true,
		"isDefined": true,
	}
	expected.ShouldBeEqual(t, 0, "New.Lazy returns defined -- valid pattern", actual)
}

func Test_Src11_NewCreator_Default(t *testing.T) {
	// Arrange & Act
	r, err := regexnew.New.Default(`^src11-default$`)

	// Assert
	actual := args.Map{
		"notNil":  r != nil,
		"noError": err == nil,
	}
	expected := args.Map{
		"notNil":  true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "New.Default returns regex -- valid", actual)
}

func Test_Src11_NewCreator_DefaultLock(t *testing.T) {
	// Arrange & Act
	r, err := regexnew.New.DefaultLock(`^src11-defaultlock$`)

	// Assert
	actual := args.Map{
		"notNil":  r != nil,
		"noError": err == nil,
	}
	expected := args.Map{
		"notNil":  true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "New.DefaultLock returns regex -- valid", actual)
}

// ── newLazyRegexCreator methods ──

func Test_Src11_LazyRegexCreator_NewLock(t *testing.T) {
	// Arrange & Act
	lr := regexnew.New.LazyRegex.NewLock(`^src11-newlock$`)

	// Assert
	actual := args.Map{"notNil": lr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.NewLock returns lazy -- valid", actual)
}

func Test_Src11_LazyRegexCreator_New(t *testing.T) {
	// Arrange & Act
	lr := regexnew.New.LazyRegex.New(`^src11-new$`)

	// Assert
	actual := args.Map{"notNil": lr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.New returns lazy -- valid", actual)
}

// ── Compiled regex vars ──

func Test_Src11_CompiledRegexVars(t *testing.T) {
	// Arrange & Act & Assert
	actual := args.Map{
		"WhitespaceFinderRegex":             regexnew.WhitespaceFinderRegex != nil,
		"HashCommentWithSpaceOptionalRegex": regexnew.HashCommentWithSpaceOptionalRegex != nil,
		"WhitespaceOrPipeFinderRegex":       regexnew.WhitespaceOrPipeFinderRegex != nil,
		"DollarIdentifierRegex":             regexnew.DollarIdentifierRegex != nil,
		"PercentIdentifierRegex":            regexnew.PercentIdentifierRegex != nil,
		"PrettyNameRegex":                   regexnew.PrettyNameRegex != nil,
		"ExactIdFieldMatchingRegex":         regexnew.ExactIdFieldMatchingRegex != nil,
		"ExactVersionIdFieldMatchingRegex":  regexnew.ExactVersionIdFieldMatchingRegex != nil,
		"UbuntuNameCheckerRegex":            regexnew.UbuntuNameCheckerRegex != nil,
		"CentOsNameCheckerRegex":            regexnew.CentOsNameCheckerRegex != nil,
		"RedHatNameCheckerRegex":            regexnew.RedHatNameCheckerRegex != nil,
		"FirstNumberAnyWhereCheckerRegex":   regexnew.FirstNumberAnyWhereCheckerRegex != nil,
		"WindowsVersionNumberCheckerRegex":  regexnew.WindowsVersionNumberCheckerRegex != nil,
	}
	expected := args.Map{
		"WhitespaceFinderRegex":             true,
		"HashCommentWithSpaceOptionalRegex": true,
		"WhitespaceOrPipeFinderRegex":       true,
		"DollarIdentifierRegex":             true,
		"PercentIdentifierRegex":            true,
		"PrettyNameRegex":                   true,
		"ExactIdFieldMatchingRegex":         true,
		"ExactVersionIdFieldMatchingRegex":  true,
		"UbuntuNameCheckerRegex":            true,
		"CentOsNameCheckerRegex":            true,
		"RedHatNameCheckerRegex":            true,
		"FirstNumberAnyWhereCheckerRegex":   true,
		"WindowsVersionNumberCheckerRegex":  true,
	}
	expected.ShouldBeEqual(t, 0, "All compiled regex vars are non-nil -- package init", actual)
}

// ── Compiled regex vars: IsApplicable ──

func Test_Src11_CompiledRegexVars_Applicable(t *testing.T) {
	// Arrange & Act & Assert
	actual := args.Map{
		"WhitespaceFinderRegex":             regexnew.WhitespaceFinderRegex.IsApplicable(),
		"HashCommentWithSpaceOptionalRegex": regexnew.HashCommentWithSpaceOptionalRegex.IsApplicable(),
		"WhitespaceOrPipeFinderRegex":       regexnew.WhitespaceOrPipeFinderRegex.IsApplicable(),
		"DollarIdentifierRegex":             regexnew.DollarIdentifierRegex.IsApplicable(),
		"PercentIdentifierRegex":            regexnew.PercentIdentifierRegex.IsApplicable(),
		"PrettyNameRegex":                   regexnew.PrettyNameRegex.IsApplicable(),
	}
	expected := args.Map{
		"WhitespaceFinderRegex":             true,
		"HashCommentWithSpaceOptionalRegex": true,
		"WhitespaceOrPipeFinderRegex":       true,
		"DollarIdentifierRegex":             true,
		"PercentIdentifierRegex":            true,
		"PrettyNameRegex":                   true,
	}
	expected.ShouldBeEqual(t, 0, "Compiled regex vars are applicable -- compile on demand", actual)
}
