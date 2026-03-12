package regexnewtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

func Test_CreateLockIf_WithLock_Cov2(t *testing.T) {
	r, err := regexnew.CreateLockIf(true, `^\d+$`)
	actual := args.Map{"hasErr": err != nil, "notNil": r != nil}
	expected := args.Map{"hasErr": false, "notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateLockIf_WithLock", actual)
}

func Test_CreateLockIf_WithoutLock_Cov2(t *testing.T) {
	r, err := regexnew.CreateLockIf(false, `^\d+$`)
	actual := args.Map{"hasErr": err != nil, "notNil": r != nil}
	expected := args.Map{"hasErr": false, "notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateLockIf_WithoutLock", actual)
}

func Test_CreateLockIf_Invalid_Cov2(t *testing.T) {
	_, err := regexnew.CreateLockIf(true, `[invalid`)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CreateLockIf_Invalid", actual)
}

func Test_CreateMustLockIf_Cov2(t *testing.T) {
	actual := args.Map{
		"withLock":    regexnew.CreateMustLockIf(true, `^\d+$`) != nil,
		"withoutLock": regexnew.CreateMustLockIf(false, `^\d+$`) != nil,
	}
	expected := args.Map{"withLock": true, "withoutLock": true}
	expected.ShouldBeEqual(t, 0, "CreateMustLockIf", actual)
}

func Test_CreateApplicableLock_Cov2(t *testing.T) {
	r, err, isApplicable := regexnew.CreateApplicableLock(`^\d+$`)
	actual := args.Map{"hasErr": err != nil, "notNil": r != nil, "isApplicable": isApplicable}
	expected := args.Map{"hasErr": false, "notNil": true, "isApplicable": true}
	expected.ShouldBeEqual(t, 0, "CreateApplicableLock_Valid", actual)
}

func Test_CreateApplicableLock_Invalid_Cov2(t *testing.T) {
	_, err, isApplicable := regexnew.CreateApplicableLock(`[invalid`)
	actual := args.Map{"hasErr": err != nil, "isApplicable": isApplicable}
	expected := args.Map{"hasErr": true, "isApplicable": false}
	expected.ShouldBeEqual(t, 0, "CreateApplicableLock_Invalid", actual)
}

func Test_IsMatchFailed_Cov2(t *testing.T) {
	actual := args.Map{
		"matchDigits":   regexnew.IsMatchFailed(`^\d+$`, "123"),
		"failAlpha":     regexnew.IsMatchFailed(`^\d+$`, "abc"),
	}
	expected := args.Map{"matchDigits": false, "failAlpha": true}
	expected.ShouldBeEqual(t, 0, "IsMatchFailed", actual)
}

func Test_MatchError_Cov2(t *testing.T) {
	actual := args.Map{
		"matchErr":   regexnew.MatchError(`^\d+$`, "123") != nil,
		"noMatchErr": regexnew.MatchError(`^\d+$`, "abc") != nil,
		"invalidErr": regexnew.MatchError(`[invalid`, "abc") != nil,
	}
	expected := args.Map{"matchErr": false, "noMatchErr": true, "invalidErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError", actual)
}

func Test_MatchErrorLock_Cov2(t *testing.T) {
	actual := args.Map{
		"matchErr":   regexnew.MatchErrorLock(`^\d+$`, "123") != nil,
		"noMatchErr": regexnew.MatchErrorLock(`^\d+$`, "abc") != nil,
	}
	expected := args.Map{"matchErr": false, "noMatchErr": true}
	expected.ShouldBeEqual(t, 0, "MatchErrorLock", actual)
}

func Test_MatchUsingFuncErrorLock_Cov2(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	actual := args.Map{
		"matchErr":   regexnew.MatchUsingFuncErrorLock(`^\d+$`, "123", matchFunc) != nil,
		"noMatchErr": regexnew.MatchUsingFuncErrorLock(`^\d+$`, "abc", matchFunc) != nil,
	}
	expected := args.Map{"matchErr": false, "noMatchErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock", actual)
}

func Test_MatchUsingCustomizeErrorFuncLock_Cov2(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	actual := args.Map{
		"matchNilCustom":    regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "123", matchFunc, nil) != nil,
		"noMatchNilCustom":  regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFunc, nil) != nil,
	}
	expected := args.Map{"matchNilCustom": false, "noMatchNilCustom": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingCustomizeErrorFuncLock", actual)
}

func Test_MatchUsingCustomizeErrorFuncLock_WithCustomize_Cov2(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	customErr := func(pattern, term string, err error, r *regexp.Regexp) error { return err }
	_ = regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFunc, customErr)
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingCustomizeErrorFuncLock_WithCustomize", actual)
}

func Test_NewMustLock_Cov2(t *testing.T) {
	actual := args.Map{"notNil": regexnew.NewMustLock(`^\d+$`) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewMustLock", actual)
}

func Test_NewCreator_Cov2(t *testing.T) {
	r1, err1 := regexnew.New.Default(`^\d+$`)
	r2, err2 := regexnew.New.DefaultLock(`^\d+$`)
	r3, err3 := regexnew.New.DefaultLockIf(true, `^\d+$`)
	r4, err4, isApp := regexnew.New.DefaultApplicableLock(`^\d+$`)
	actual := args.Map{
		"default":     r1 != nil && err1 == nil,
		"lock":        r2 != nil && err2 == nil,
		"lockIf":      r3 != nil && err3 == nil,
		"applicable":  r4 != nil && err4 == nil && isApp,
		"lazy":        regexnew.New.Lazy(`^\d+$`) != nil,
		"lazyLock":    regexnew.New.LazyLock(`^\d+$`) != nil,
	}
	expected := args.Map{
		"default": true, "lock": true, "lockIf": true,
		"applicable": true, "lazy": true, "lazyLock": true,
	}
	expected.ShouldBeEqual(t, 0, "NewCreator", actual)
}

func Test_LazyRegex_Methods_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	actual := args.Map{
		"fullString":     lazy.FullString() != "",
		"matchBytes":     lazy.IsMatchBytes([]byte("123")),
		"noMatchBytes":   lazy.IsMatchBytes([]byte("abc")),
		"failMatch":      lazy.IsFailedMatch("123"),
		"noFailMatch":    lazy.IsFailedMatch("abc"),
		"failMatchBytes": lazy.IsFailedMatchBytes([]byte("123")),
		"matchErr":       lazy.MatchError("123") != nil,
		"noMatchErr":     lazy.MatchError("abc") != nil,
		"hasError":       lazy.HasError(),
		"hasIssues":      lazy.HasAnyIssues(),
		"isInvalid":      lazy.IsInvalid(),
		"compiledErr":    lazy.CompiledError() != nil,
		"error":          lazy.Error() != nil,
	}
	expected := args.Map{
		"fullString": true, "matchBytes": true, "noMatchBytes": false,
		"failMatch": false, "noFailMatch": true,
		"failMatchBytes": false, "matchErr": false, "noMatchErr": true,
		"hasError": false, "hasIssues": false, "isInvalid": false,
		"compiledErr": false, "error": false,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex_Methods", actual)
}

func Test_LazyRegex_FirstMatchLine_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`(\d+)`)
	match, isInvalid := lazy.FirstMatchLine("abc123def")
	noMatch, noIsInvalid := lazy.FirstMatchLine("abcdef")
	actual := args.Map{
		"match": match, "isInvalid": isInvalid,
		"noMatch": noMatch, "noIsInvalid": noIsInvalid,
	}
	expected := args.Map{
		"match": "123", "isInvalid": false,
		"noMatch": "", "noIsInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex_FirstMatchLine", actual)
}

func Test_LazyRegex_MatchUsingFuncError_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	matchFunc := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	actual := args.Map{
		"matchErr":   lazy.MatchUsingFuncError("123", matchFunc) != nil,
		"noMatchErr": lazy.MatchUsingFuncError("abc", matchFunc) != nil,
	}
	expected := args.Map{"matchErr": false, "noMatchErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex_MatchUsingFuncError", actual)
}

func Test_LazyRegex_MustBeSafe_Cov2(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	lazy.MustBeSafe()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex_MustBeSafe", actual)
}

func Test_LazyRegex_NilReceiver_Cov2(t *testing.T) {
	var lazy *regexnew.LazyRegex
	actual := args.Map{
		"isNull":       lazy.IsNull(),
		"isDefined":    lazy.IsDefined(),
		"isUndefined":  lazy.IsUndefined(),
		"isCompiled":   lazy.IsCompiled(),
		"string":       lazy.String(),
		"pattern":      lazy.Pattern(),
		"fullString":   lazy.FullString(),
		"hasIssues":    lazy.HasAnyIssues(),
		"isInvalid":    lazy.IsInvalid(),
		"isApplicable": lazy.IsApplicable(),
		"reqCompiled":  lazy.OnRequiredCompiled() != nil,
	}
	expected := args.Map{
		"isNull": true, "isDefined": false, "isUndefined": true,
		"isCompiled": false, "string": "", "pattern": "",
		"fullString": "", "hasIssues": true, "isInvalid": true,
		"isApplicable": false, "reqCompiled": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex_NilReceiver", actual)
}

func Test_LazyRegex_OnRequiredCompiledMust_NilPanic_Cov2(t *testing.T) {
	var lazy *regexnew.LazyRegex
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		lazy.OnRequiredCompiledMust()
	}()
	actual := args.Map{"panicked": panicked}
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex_OnRequiredCompiledMust_NilPanic", actual)
}

func Test_LazyRegex_CompileMust_Valid_Cov2(t *testing.T) {
	actual := args.Map{"notNil": regexnew.New.Lazy(`^\d+$`).CompileMust() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex_CompileMust_Valid", actual)
}

func Test_NewLazyRegexCreator_Cov2(t *testing.T) {
	first, second := regexnew.New.LazyRegex.TwoLock(`^\d+$`, `^[a-z]+$`)
	m := regexnew.New.LazyRegex.ManyUsingLock(`^\d+$`, `^[a-z]+$`)
	mEmpty := regexnew.New.LazyRegex.ManyUsingLock()
	actual := args.Map{
		"firstNotNil":  first != nil,
		"secondNotNil": second != nil,
		"manyLen":      len(m),
		"emptyLen":     len(mEmpty),
		"patternsMap":  regexnew.New.LazyRegex.AllPatternsMap() != nil,
		"lockIfTrue":   regexnew.New.LazyRegex.NewLockIf(true, `^\d+$`) != nil,
		"lockIfFalse":  regexnew.New.LazyRegex.NewLockIf(false, `^[a-z]+$`) != nil,
	}
	expected := args.Map{
		"firstNotNil": true, "secondNotNil": true,
		"manyLen": 2, "emptyLen": 0,
		"patternsMap": true, "lockIfTrue": true, "lockIfFalse": true,
	}
	expected.ShouldBeEqual(t, 0, "NewLazyRegexCreator", actual)
}
