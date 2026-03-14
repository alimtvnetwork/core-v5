package regexnewtests

import (
	"errors"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ── LazyRegex ──

func Test_Cov4_LazyRegex_Nil(t *testing.T) {
	var lr *regexnew.LazyRegex
	actual := args.Map{
		"isNull":      lr.IsNull(),
		"isDefined":   lr.IsDefined(),
		"isUndefined": lr.IsUndefined(),
		"isApplicable": lr.IsApplicable(),
		"isCompiled":  lr.IsCompiled(),
		"hasError":    lr.HasError(),
		"hasIssues":   lr.HasAnyIssues(),
		"isInvalid":   lr.IsInvalid(),
		"string":      lr.String(),
		"fullString":  lr.FullString(),
		"pattern":     lr.Pattern(),
	}
	expected := args.Map{
		"isNull": true, "isDefined": false, "isUndefined": true,
		"isApplicable": false, "isCompiled": false, "hasError": false,
		"hasIssues": true, "isInvalid": true,
		"string": "", "fullString": "", "pattern": "",
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex nil", actual)
}

func Test_Cov4_LazyRegex_Valid(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	actual := args.Map{
		"isDefined":    lr.IsDefined(),
		"isUndefined":  lr.IsUndefined(),
		"isApplicable": lr.IsApplicable(),
		"isCompiled":   lr.IsCompiled(),
		"hasError":     lr.HasError(),
		"hasIssues":    lr.HasAnyIssues(),
		"isInvalid":    lr.IsInvalid(),
		"pattern":      lr.Pattern(),
		"string":       lr.String(),
	}
	expected := args.Map{
		"isDefined": true, "isUndefined": false, "isApplicable": true,
		"isCompiled": true, "hasError": false, "hasIssues": false,
		"isInvalid": false, "pattern": `^\d+$`, "string": `^\d+$`,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex valid", actual)
}

func Test_Cov4_LazyRegex_Invalid(t *testing.T) {
	lr := regexnew.New.LazyLock(`[invalid`)
	actual := args.Map{
		"isApplicable": lr.IsApplicable(),
		"hasError":     lr.HasError(),
		"hasIssues":    lr.HasAnyIssues(),
		"isInvalid":    lr.IsInvalid(),
	}
	expected := args.Map{
		"isApplicable": false, "hasError": true,
		"hasIssues": true, "isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex invalid", actual)
}

func Test_Cov4_LazyRegex_Compile(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	r, err := lr.Compile()
	actual := args.Map{"notNil": r != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex Compile", actual)
}

func Test_Cov4_LazyRegex_CompileMust(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	r := lr.CompileMust()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex CompileMust", actual)
}

func Test_Cov4_LazyRegex_CompileMust_Panic(t *testing.T) {
	lr := regexnew.New.LazyLock(`[invalid`)
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "LazyRegex CompileMust panic", actual)
	}()
	lr.CompileMust()
}

func Test_Cov4_LazyRegex_OnRequiredCompiled(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.OnRequiredCompiled()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex OnRequiredCompiled", actual)
}

func Test_Cov4_LazyRegex_OnRequiredCompiled_Nil(t *testing.T) {
	var lr *regexnew.LazyRegex
	err := lr.OnRequiredCompiled()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex OnRequiredCompiled nil", actual)
}

func Test_Cov4_LazyRegex_OnRequiredCompiledMust(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	lr.OnRequiredCompiledMust() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex OnRequiredCompiledMust", actual)
}

func Test_Cov4_LazyRegex_OnRequiredCompiledMust_Panic(t *testing.T) {
	lr := regexnew.New.LazyLock(`[invalid`)
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "LazyRegex OnRequiredCompiledMust panic", actual)
	}()
	lr.OnRequiredCompiledMust()
}

func Test_Cov4_LazyRegex_CompiledError(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.CompiledError()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex CompiledError", actual)
}

func Test_Cov4_LazyRegex_Error(t *testing.T) {
	lr := regexnew.New.LazyLock(`[invalid`)
	err := lr.Error()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex Error", actual)
}

func Test_Cov4_LazyRegex_MustBeSafe(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	lr.MustBeSafe() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex MustBeSafe", actual)
}

func Test_Cov4_LazyRegex_MustBeSafe_Panic(t *testing.T) {
	lr := regexnew.New.LazyLock(`[invalid`)
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "LazyRegex MustBeSafe panic", actual)
	}()
	lr.MustBeSafe()
}

func Test_Cov4_LazyRegex_FullString(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	result := lr.FullString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex FullString", actual)
}

func Test_Cov4_LazyRegex_MatchError_Success(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.MatchError("123")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex MatchError success", actual)
}

func Test_Cov4_LazyRegex_MatchError_Fail(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.MatchError("abc")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex MatchError fail", actual)
}

func Test_Cov4_LazyRegex_MatchUsingFuncError_Success(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.MatchUsingFuncError("123", func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	})
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex MatchUsingFuncError success", actual)
}

func Test_Cov4_LazyRegex_MatchUsingFuncError_Fail(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	err := lr.MatchUsingFuncError("abc", func(r *regexp.Regexp, s string) bool {
		return r.MatchString(s)
	})
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex MatchUsingFuncError fail", actual)
}

func Test_Cov4_LazyRegex_IsMatch(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	actual := args.Map{
		"match":   lr.IsMatch("123"),
		"noMatch": lr.IsMatch("abc"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsMatch", actual)
}

func Test_Cov4_LazyRegex_IsMatchBytes(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	actual := args.Map{
		"match":   lr.IsMatchBytes([]byte("123")),
		"noMatch": lr.IsMatchBytes([]byte("abc")),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsMatchBytes", actual)
}

func Test_Cov4_LazyRegex_IsFailedMatch(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	actual := args.Map{
		"failOnAlpha":  lr.IsFailedMatch("abc"),
		"failOnDigits": lr.IsFailedMatch("123"),
	}
	expected := args.Map{"failOnAlpha": true, "failOnDigits": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsFailedMatch", actual)
}

func Test_Cov4_LazyRegex_IsFailedMatchBytes(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	actual := args.Map{
		"failOnAlpha":  lr.IsFailedMatchBytes([]byte("abc")),
		"failOnDigits": lr.IsFailedMatchBytes([]byte("123")),
	}
	expected := args.Map{"failOnAlpha": true, "failOnDigits": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsFailedMatchBytes", actual)
}

func Test_Cov4_LazyRegex_FirstMatchLine(t *testing.T) {
	lr := regexnew.New.LazyLock(`\d+`)
	match, isInvalid := lr.FirstMatchLine("abc 123 def")
	actual := args.Map{"match": match, "isInvalid": isInvalid}
	expected := args.Map{"match": "123", "isInvalid": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex FirstMatchLine", actual)
}

func Test_Cov4_LazyRegex_FirstMatchLine_NoMatch(t *testing.T) {
	lr := regexnew.New.LazyLock(`\d+`)
	match, isInvalid := lr.FirstMatchLine("abc def")
	actual := args.Map{"match": match, "isInvalid": isInvalid}
	expected := args.Map{"match": "", "isInvalid": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex FirstMatchLine no match", actual)
}

func Test_Cov4_LazyRegex_FirstMatchLine_Invalid(t *testing.T) {
	lr := regexnew.New.LazyLock(`[invalid`)
	match, isInvalid := lr.FirstMatchLine("abc")
	actual := args.Map{"match": match, "isInvalid": isInvalid}
	expected := args.Map{"match": "", "isInvalid": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex FirstMatchLine invalid regex", actual)
}

// ── MatchUsingCustomizeErrorFuncLock with custom err ──

func Test_Cov4_MatchCustomErr_WithCustom(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	customErr := func(pattern, match string, compileErr error, r *regexp.Regexp) error {
		return errors.New("custom: " + match)
	}
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFunc, customErr)
	actual := args.Map{"hasErr": err != nil, "isCustom": err.Error() == "custom: abc"}
	expected := args.Map{"hasErr": true, "isCustom": true}
	expected.ShouldBeEqual(t, 0, "MatchCustomErr with custom func", actual)
}

// ── NewMustLock ──

func Test_Cov4_NewMustLock(t *testing.T) {
	r := regexnew.NewMustLock(`^\d+$`)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewMustLock", actual)
}

// ── Create / CreateLock / CreateMust ──

func Test_Cov4_Create_Valid(t *testing.T) {
	r, err := regexnew.Create(`^[a-z]+$`)
	actual := args.Map{"notNil": r != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Create valid", actual)
}

func Test_Cov4_Create_Invalid(t *testing.T) {
	_, err := regexnew.Create(`[invalid`)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Create invalid", actual)
}

func Test_Cov4_CreateLock_Valid(t *testing.T) {
	r, err := regexnew.CreateLock(`^[A-Z]+$`)
	actual := args.Map{"notNil": r != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "CreateLock valid", actual)
}

func Test_Cov4_CreateMust(t *testing.T) {
	r := regexnew.CreateMust(`^test\d+$`)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMust", actual)
}

func Test_Cov4_IsMatchLock_Valid(t *testing.T) {
	actual := args.Map{
		"match":   regexnew.IsMatchLock(`^\d+$`, "456"),
		"noMatch": regexnew.IsMatchLock(`^\d+$`, "abc"),
		"invalid": regexnew.IsMatchLock(`[invalid`, "abc"),
	}
	expected := args.Map{"match": true, "noMatch": false, "invalid": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock", actual)
}

// ── New.Lazy creators ──

func Test_Cov4_NewLazy_NoLock(t *testing.T) {
	lr := regexnew.New.Lazy.NoLock(`^\d+$`)
	actual := args.Map{"isDefined": lr.IsDefined(), "isApplicable": lr.IsApplicable()}
	expected := args.Map{"isDefined": true, "isApplicable": true}
	expected.ShouldBeEqual(t, 0, "NewLazy NoLock", actual)
}

func Test_Cov4_NewLazy_LockIf_True(t *testing.T) {
	lr := regexnew.New.LazyLockIf(true, `^\d+$`)
	actual := args.Map{"isDefined": lr.IsDefined()}
	expected := args.Map{"isDefined": true}
	expected.ShouldBeEqual(t, 0, "NewLazy LockIf true", actual)
}

func Test_Cov4_NewLazy_LockIf_False(t *testing.T) {
	lr := regexnew.New.LazyLockIf(false, `^\d+$`)
	actual := args.Map{"isDefined": lr.IsDefined()}
	expected := args.Map{"isDefined": true}
	expected.ShouldBeEqual(t, 0, "NewLazy LockIf false", actual)
}
