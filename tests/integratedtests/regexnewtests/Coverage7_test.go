package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex — Compile edge cases (undefined receiver)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_LazyRegex_Compile_Undefined(t *testing.T) {
	// Create a LazyRegex with empty pattern → IsUndefined() == true
	lr := regexnew.New.Lazy("")
	re, err := lr.Compile()
	actual := args.Map{"nil": re == nil, "hasErr": err != nil}
	expected := args.Map{"nil": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "Compile undefined", actual)
}

func Test_Cov7_LazyRegex_CompileMust_Panic(t *testing.T) {
	lr := regexnew.New.Lazy("")
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		lr.CompileMust()
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "CompileMust panic", actual)
}

func Test_Cov7_LazyRegex_OnRequiredCompiledMust_Panic(t *testing.T) {
	lr := regexnew.New.Lazy("")
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		lr.OnRequiredCompiledMust()
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiledMust panic", actual)
}

func Test_Cov7_LazyRegex_MustBeSafe_Panic(t *testing.T) {
	lr := regexnew.New.Lazy("")
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		lr.MustBeSafe()
	}()
	actual := args.Map{"panicked": didPanic}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "MustBeSafe panic", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex — Match methods with invalid/undefined regex
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_LazyRegex_MatchError_Undefined(t *testing.T) {
	lr := regexnew.New.Lazy("")
	err := lr.MatchError("test")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError undefined", actual)
}

func Test_Cov7_LazyRegex_MatchUsingFuncError_Undefined(t *testing.T) {
	lr := regexnew.New.Lazy("")
	err := lr.MatchUsingFuncError("test", func(re interface{ MatchString(string) bool }, s string) bool {
		return false
	})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncError undefined", actual)
}

func Test_Cov7_LazyRegex_IsMatch_Undefined(t *testing.T) {
	lr := regexnew.New.Lazy("")
	actual := args.Map{"v": lr.IsMatch("test")}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsMatch undefined", actual)
}

func Test_Cov7_LazyRegex_IsMatchBytes_Undefined(t *testing.T) {
	lr := regexnew.New.Lazy("")
	actual := args.Map{"v": lr.IsMatchBytes([]byte("test"))}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsMatchBytes undefined", actual)
}

func Test_Cov7_LazyRegex_IsFailedMatch_Undefined(t *testing.T) {
	lr := regexnew.New.Lazy("")
	actual := args.Map{"v": lr.IsFailedMatch("test")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsFailedMatch undefined", actual)
}

func Test_Cov7_LazyRegex_IsFailedMatchBytes_Undefined(t *testing.T) {
	lr := regexnew.New.Lazy("")
	actual := args.Map{"v": lr.IsFailedMatchBytes([]byte("test"))}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsFailedMatchBytes undefined", actual)
}

func Test_Cov7_LazyRegex_FirstMatchLine_Undefined(t *testing.T) {
	lr := regexnew.New.Lazy("")
	line, invalid := lr.FirstMatchLine("test")
	actual := args.Map{"line": line, "invalid": invalid}
	expected := args.Map{"line": "", "invalid": true}
	expected.ShouldBeEqual(t, 0, "FirstMatchLine undefined", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex — HasError with error from invalid pattern
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_LazyRegex_HasError_InvalidPattern(t *testing.T) {
	lr := regexnew.New.LazyLock(`[invalid`)
	actual := args.Map{"v": lr.HasError()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HasError invalid pattern", actual)
}

func Test_Cov7_LazyRegex_IsApplicable_Undefined(t *testing.T) {
	lr := regexnew.New.Lazy("")
	actual := args.Map{"v": lr.IsApplicable()}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsApplicable undefined", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Precompiled regex vars — match testing to exercise compiled branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_PrecompiledRegex_WhitespaceFinder_Match(t *testing.T) {
	actual := args.Map{"v": regexnew.WhitespaceFinderRegex.IsMatch(" ")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "WhitespaceFinderRegex match", actual)
}

func Test_Cov7_PrecompiledRegex_HashComment_Match(t *testing.T) {
	actual := args.Map{"v": regexnew.HashCommentWithSpaceOptionalRegex.IsMatch("# comment")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HashComment match", actual)
}

func Test_Cov7_PrecompiledRegex_PrettyName_Match(t *testing.T) {
	actual := args.Map{"v": regexnew.PrettyNameRegex.IsMatch("hello-world")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "PrettyNameRegex match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MatchUsingCustomizeErrorFuncLock — with invalid regex + custom err
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_MatchUsingCustomizeErrorFuncLock_InvalidRegex_CustomErr(t *testing.T) {
	matchFn := func(re interface{ MatchString(string) bool }, s string) bool { return false }
	customErr := func(pattern, term string, err error, re interface{ MatchString(string) bool }) error {
		return err
	}
	// Use a type-compatible wrapper
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`[invalid`, "abc",
		func(re interface{ MatchString(string) bool }, s string) bool { return false },
		nil,
	)
	_ = matchFn
	_ = customErr
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CustomErrLock invalid regex", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MatchError with invalid regex pattern — exercises regExMatchValidationError nil-regex branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_MatchErrorLock_InvalidRegex(t *testing.T) {
	err := regexnew.MatchErrorLock(`[invalid`, "abc")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchErrorLock invalid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// newCreator — Default / DefaultLock with invalid patterns
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_NewCreator_Default_Invalid(t *testing.T) {
	_, err := regexnew.New.Default(`[invalid`)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "New.Default invalid", actual)
}

func Test_Cov7_NewCreator_DefaultLock_Invalid(t *testing.T) {
	_, err := regexnew.New.DefaultLock(`[invalid`)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "New.DefaultLock invalid", actual)
}

func Test_Cov7_NewCreator_DefaultLockIf_NoLock(t *testing.T) {
	re, err := regexnew.New.DefaultLockIf(false, `^cov7\d+$`)
	actual := args.Map{"notNil": re != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "New.DefaultLockIf no lock", actual)
}

func Test_Cov7_NewCreator_DefaultApplicableLock_Invalid(t *testing.T) {
	_, err, ok := regexnew.New.DefaultApplicableLock(`[invalid`)
	actual := args.Map{"hasErr": err != nil, "ok": ok}
	expected := args.Map{"hasErr": true, "ok": false}
	expected.ShouldBeEqual(t, 0, "New.DefaultApplicableLock invalid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex — FullString with error
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_LazyRegex_FullString_Invalid(t *testing.T) {
	lr := regexnew.New.LazyLock(`[invalid`)
	s := lr.FullString()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullString invalid pattern", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex MatchUsingFuncError with match func returning false
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_LazyRegex_MatchUsingFuncError_FuncReturnsFalse(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov7mfe\d+$`)
	err := lr.MatchUsingFuncError("cov7mfe123", func(re interface{ MatchString(string) bool }, s string) bool {
		return false // always fails
	})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncError func returns false", actual)
}
