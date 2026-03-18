package regexnewtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex — basic operations
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_LazyRegex_Compile_Valid(t *testing.T) {
	lr := regexnew.New.Lazy(`^cov7\d+$`)
	re, err := lr.Compile()
	actual := args.Map{"notNil": re != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Compile valid", actual)
}

func Test_Cov7_LazyRegex_Compile_Invalid(t *testing.T) {
	lr := regexnew.New.Lazy(`[invalid`)
	re, err := lr.Compile()
	actual := args.Map{"nil": re == nil, "hasErr": err != nil}
	expected := args.Map{"nil": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "Compile invalid", actual)
}

func Test_Cov7_LazyRegex_IsMatch_True(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov7m\d+$`)
	actual := args.Map{"v": lr.IsMatch("cov7m123")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsMatch true", actual)
}

func Test_Cov7_LazyRegex_IsMatch_False(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov7m\d+$`)
	actual := args.Map{"v": lr.IsMatch("notmatch")}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsMatch false", actual)
}

func Test_Cov7_LazyRegex_MatchError_Valid(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov7me\d+$`)
	err := lr.MatchError("cov7me123")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError valid", actual)
}

func Test_Cov7_LazyRegex_MatchError_NoMatch(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov7me\d+$`)
	err := lr.MatchError("notmatch")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError no match", actual)
}

func Test_Cov7_LazyRegex_MatchError_Undefined(t *testing.T) {
	lr := regexnew.New.Lazy("")
	err := lr.MatchError("test")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError undefined", actual)
}

func Test_Cov7_LazyRegex_MatchUsingFuncError_Undefined(t *testing.T) {
	lr := regexnew.New.Lazy("")
	err := lr.MatchUsingFuncError("test", func(re *regexp.Regexp, s string) bool {
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

func Test_Cov7_LazyRegex_CompileMust_Valid(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov7cm\d+$`)
	re := lr.CompileMust()
	actual := args.Map{"notNil": re != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CompileMust valid", actual)
}

func Test_Cov7_LazyRegex_CompileMust_Panic(t *testing.T) {
	lr := regexnew.New.Lazy(`[invalid`)
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		lr.CompileMust()
	}()
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "CompileMust panic", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MatchUsingCustomizeErrorFuncLock
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_MatchUsingCustomizeErrorFuncLock_InvalidRegex(t *testing.T) {
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`[invalid`, "abc",
		func(re *regexp.Regexp, s string) bool { return false },
		nil,
	)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CustomErrLock invalid regex", actual)
}

func Test_Cov7_MatchUsingCustomizeErrorFuncLock_Valid(t *testing.T) {
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^cov7cust\d+$`, "cov7cust123",
		func(re *regexp.Regexp, s string) bool { return re.MatchString(s) },
		nil,
	)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CustomErrLock valid", actual)
}

func Test_Cov7_MatchUsingCustomizeErrorFuncLock_NoMatch(t *testing.T) {
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^cov7cust\d+$`, "nomatch",
		func(re *regexp.Regexp, s string) bool { return re.MatchString(s) },
		nil,
	)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CustomErrLock no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MatchUsingFuncErrorLock
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_MatchUsingFuncErrorLock_Valid(t *testing.T) {
	err := regexnew.MatchUsingFuncErrorLock(`^cov7fl\d+$`, "cov7fl123",
		func(re *regexp.Regexp, s string) bool { return re.MatchString(s) },
	)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock valid", actual)
}

func Test_Cov7_MatchUsingFuncErrorLock_Invalid(t *testing.T) {
	err := regexnew.MatchUsingFuncErrorLock(`[invalid`, "abc",
		func(re *regexp.Regexp, s string) bool { return false },
	)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock invalid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex FullString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_LazyRegex_FullString_Valid(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov7fs\d+$`)
	_ = lr.CompileMust()
	s := lr.FullString()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullString valid", actual)
}

func Test_Cov7_LazyRegex_FullString_Invalid(t *testing.T) {
	lr := regexnew.New.Lazy(`[invalid`)
	s := lr.FullString()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullString invalid pattern", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex MatchUsingFuncError with func returning false
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov7_LazyRegex_MatchUsingFuncError_FuncReturnsFalse(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov7mfe\d+$`)
	err := lr.MatchUsingFuncError("cov7mfe123", func(re *regexp.Regexp, s string) bool {
		return false
	})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncError func returns false", actual)
}
