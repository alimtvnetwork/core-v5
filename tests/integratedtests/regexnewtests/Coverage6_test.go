package regexnewtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ══════════════════════════════════════════════════════════════════════════════
// Create / CreateLock / CreateLockIf / CreateMust / CreateMustLockIf
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_Create_Valid(t *testing.T) {
	re, err := regexnew.Create(`^\d+$`)
	actual := args.Map{"notNil": re != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Create valid", actual)
}

func Test_Cov6_Create_Invalid(t *testing.T) {
	_, err := regexnew.Create(`[invalid`)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Create invalid", actual)
}

func Test_Cov6_Create_Cached(t *testing.T) {
	re1, _ := regexnew.Create(`^cov6cached\d+$`)
	re2, _ := regexnew.Create(`^cov6cached\d+$`)
	actual := args.Map{"same": re1 == re2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "Create cached", actual)
}

func Test_Cov6_CreateLock_Valid(t *testing.T) {
	re, err := regexnew.CreateLock(`^cov6lock\d+$`)
	actual := args.Map{"notNil": re != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "CreateLock valid", actual)
}

func Test_Cov6_CreateLockIf_WithLock(t *testing.T) {
	re, err := regexnew.CreateLockIf(true, `^cov6lockif\d+$`)
	actual := args.Map{"notNil": re != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "CreateLockIf lock", actual)
}

func Test_Cov6_CreateLockIf_WithoutLock(t *testing.T) {
	re, err := regexnew.CreateLockIf(false, `^cov6lockifno\d+$`)
	actual := args.Map{"notNil": re != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "CreateLockIf no lock", actual)
}

func Test_Cov6_CreateMust_Valid(t *testing.T) {
	re := regexnew.CreateMust(`^cov6must\d+$`)
	actual := args.Map{"notNil": re != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMust valid", actual)
}

func Test_Cov6_CreateMust_Cached(t *testing.T) {
	re1 := regexnew.CreateMust(`^cov6mustcache\d+$`)
	re2 := regexnew.CreateMust(`^cov6mustcache\d+$`)
	actual := args.Map{"same": re1 == re2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "CreateMust cached", actual)
}

func Test_Cov6_CreateMustLockIf_WithLock(t *testing.T) {
	re := regexnew.CreateMustLockIf(true, `^cov6mustlock\d+$`)
	actual := args.Map{"notNil": re != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMustLockIf lock", actual)
}

func Test_Cov6_CreateMustLockIf_WithoutLock(t *testing.T) {
	re := regexnew.CreateMustLockIf(false, `^cov6mustlockno\d+$`)
	actual := args.Map{"notNil": re != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMustLockIf no lock", actual)
}

func Test_Cov6_CreateApplicableLock_Valid(t *testing.T) {
	re, err, ok := regexnew.CreateApplicableLock(`^cov6applock\d+$`)
	actual := args.Map{"notNil": re != nil, "noErr": err == nil, "ok": ok}
	expected := args.Map{"notNil": true, "noErr": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "CreateApplicableLock valid", actual)
}

func Test_Cov6_CreateApplicableLock_Invalid(t *testing.T) {
	_, err, ok := regexnew.CreateApplicableLock(`[invalid`)
	actual := args.Map{"hasErr": err != nil, "ok": ok}
	expected := args.Map{"hasErr": true, "ok": false}
	expected.ShouldBeEqual(t, 0, "CreateApplicableLock invalid", actual)
}

func Test_Cov6_NewMustLock(t *testing.T) {
	re := regexnew.NewMustLock(`^cov6newmust\d+$`)
	actual := args.Map{"notNil": re != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewMustLock", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsMatchLock / IsMatchFailed
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_IsMatchLock_Match(t *testing.T) {
	actual := args.Map{"v": regexnew.IsMatchLock(`^\d+$`, "123")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsMatchLock match", actual)
}

func Test_Cov6_IsMatchLock_NoMatch(t *testing.T) {
	actual := args.Map{"v": regexnew.IsMatchLock(`^\d+$`, "abc")}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock no match", actual)
}

func Test_Cov6_IsMatchLock_InvalidRegex(t *testing.T) {
	actual := args.Map{"v": regexnew.IsMatchLock(`[invalid`, "abc")}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock invalid regex", actual)
}

func Test_Cov6_IsMatchFailed_Match(t *testing.T) {
	actual := args.Map{"v": regexnew.IsMatchFailed(`^\d+$`, "123")}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsMatchFailed match", actual)
}

func Test_Cov6_IsMatchFailed_NoMatch(t *testing.T) {
	actual := args.Map{"v": regexnew.IsMatchFailed(`^\d+$`, "abc")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsMatchFailed no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MatchError / MatchErrorLock / MatchUsingFuncErrorLock / MatchUsingCustomizeErrorFuncLock
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_MatchError_Match(t *testing.T) {
	actual := args.Map{"noErr": regexnew.MatchError(`^\d+$`, "123") == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError match", actual)
}

func Test_Cov6_MatchError_NoMatch(t *testing.T) {
	actual := args.Map{"hasErr": regexnew.MatchError(`^\d+$`, "abc") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError no match", actual)
}

func Test_Cov6_MatchError_InvalidRegex(t *testing.T) {
	actual := args.Map{"hasErr": regexnew.MatchError(`[invalid`, "abc") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError invalid", actual)
}

func Test_Cov6_MatchErrorLock_Match(t *testing.T) {
	actual := args.Map{"noErr": regexnew.MatchErrorLock(`^\d+$`, "123") == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchErrorLock match", actual)
}

func Test_Cov6_MatchErrorLock_NoMatch(t *testing.T) {
	actual := args.Map{"hasErr": regexnew.MatchErrorLock(`^\d+$`, "abc") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchErrorLock no match", actual)
}

func Test_Cov6_MatchUsingFuncErrorLock_Match(t *testing.T) {
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }
	actual := args.Map{"noErr": regexnew.MatchUsingFuncErrorLock(`^\d+$`, "123", matchFn) == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock match", actual)
}

func Test_Cov6_MatchUsingFuncErrorLock_NoMatch(t *testing.T) {
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }
	actual := args.Map{"hasErr": regexnew.MatchUsingFuncErrorLock(`^\d+$`, "abc", matchFn) != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock no match", actual)
}

func Test_Cov6_MatchUsingCustomizeErrorFuncLock_Match(t *testing.T) {
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }
	actual := args.Map{"noErr": regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "123", matchFn, nil) == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CustomizeErrLock match", actual)
}

func Test_Cov6_MatchUsingCustomizeErrorFuncLock_NoMatch_NilCustomize(t *testing.T) {
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }
	actual := args.Map{"hasErr": regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFn, nil) != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CustomizeErrLock no match nil customize", actual)
}

func Test_Cov6_MatchUsingCustomizeErrorFuncLock_NoMatch_CustomErr(t *testing.T) {
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }
	customErr := func(pattern, term string, err error, re *regexp.Regexp) error {
		return fmt.Errorf("custom error for %s", term)
	}
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFn, customErr)
	actual := args.Map{"hasErr": err != nil, "custom": err.Error() == "custom error for abc"}
	expected := args.Map{"hasErr": true, "custom": true}
	expected.ShouldBeEqual(t, 0, "CustomizeErrLock custom err", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex — all methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_LazyRegex_IsNull(t *testing.T) {
	var lr *regexnew.LazyRegex
	actual := args.Map{"nil": lr.IsNull()}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsNull", actual)
}

func Test_Cov6_LazyRegex_IsDefined(t *testing.T) {
	lr := regexnew.New.Lazy(`^cov6lazy\d+$`)
	actual := args.Map{"defined": lr.IsDefined()}
	expected := args.Map{"defined": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsDefined", actual)
}

func Test_Cov6_LazyRegex_IsUndefined_Nil(t *testing.T) {
	var lr *regexnew.LazyRegex
	actual := args.Map{"v": lr.IsUndefined()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsUndefined nil", actual)
}

func Test_Cov6_LazyRegex_IsApplicable(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6lazyapp\d+$`)
	actual := args.Map{"v": lr.IsApplicable()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsApplicable", actual)
}

func Test_Cov6_LazyRegex_IsApplicable_Nil(t *testing.T) {
	var lr *regexnew.LazyRegex
	actual := args.Map{"v": lr.IsApplicable()}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsApplicable nil", actual)
}

func Test_Cov6_LazyRegex_IsApplicable_Cached(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6lazycached\d+$`)
	_ = lr.IsApplicable() // first call compiles
	actual := args.Map{"v": lr.IsApplicable()} // second call returns cached
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsApplicable cached", actual)
}

func Test_Cov6_LazyRegex_Compile_Valid(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6comp\d+$`)
	re, err := lr.Compile()
	actual := args.Map{"notNil": re != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex Compile valid", actual)
}

func Test_Cov6_LazyRegex_Compile_Cached(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6compcache\d+$`)
	re1, _ := lr.Compile()
	re2, _ := lr.Compile()
	actual := args.Map{"same": re1 == re2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex Compile cached", actual)
}

func Test_Cov6_LazyRegex_IsCompiled(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6iscomp\d+$`)
	actual := args.Map{"before": lr.IsCompiled()}
	expected := args.Map{"before": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsCompiled before", actual)
	lr.Compile()
	actual2 := args.Map{"after": lr.IsCompiled()}
	expected2 := args.Map{"after": true}
	expected2.ShouldBeEqual(t, 0, "LazyRegex IsCompiled after", actual2)
}

func Test_Cov6_LazyRegex_IsCompiled_Nil(t *testing.T) {
	var lr *regexnew.LazyRegex
	actual := args.Map{"v": lr.IsCompiled()}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsCompiled nil", actual)
}

func Test_Cov6_LazyRegex_OnRequiredCompiled_Nil(t *testing.T) {
	var lr *regexnew.LazyRegex
	err := lr.OnRequiredCompiled()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiled nil", actual)
}

func Test_Cov6_LazyRegex_OnRequiredCompiled_Valid(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6reqcomp\d+$`)
	err := lr.OnRequiredCompiled()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiled valid", actual)
}

func Test_Cov6_LazyRegex_OnRequiredCompiled_AlreadyCompiled(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6reqcomp2\d+$`)
	lr.Compile()
	err := lr.OnRequiredCompiled()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiled already compiled", actual)
}

func Test_Cov6_LazyRegex_HasError(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6haserr\d+$`)
	actual := args.Map{"v": lr.HasError()}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex HasError", actual)
}

func Test_Cov6_LazyRegex_HasAnyIssues_Nil(t *testing.T) {
	var lr *regexnew.LazyRegex
	actual := args.Map{"v": lr.HasAnyIssues()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "HasAnyIssues nil", actual)
}

func Test_Cov6_LazyRegex_HasAnyIssues_Valid(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6anyiss\d+$`)
	actual := args.Map{"v": lr.HasAnyIssues()}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "HasAnyIssues valid", actual)
}

func Test_Cov6_LazyRegex_IsInvalid_Nil(t *testing.T) {
	var lr *regexnew.LazyRegex
	actual := args.Map{"v": lr.IsInvalid()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsInvalid nil", actual)
}

func Test_Cov6_LazyRegex_IsInvalid_Valid(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6inv\d+$`)
	actual := args.Map{"v": lr.IsInvalid()}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsInvalid valid", actual)
}

func Test_Cov6_LazyRegex_CompiledError(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6comperr\d+$`)
	actual := args.Map{"noErr": lr.CompiledError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CompiledError", actual)
}

func Test_Cov6_LazyRegex_Error(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6err\d+$`)
	actual := args.Map{"noErr": lr.Error() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "Error", actual)
}

func Test_Cov6_LazyRegex_MustBeSafe(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6safe\d+$`)
	// should not panic
	lr.MustBeSafe()
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeSafe", actual)
}

func Test_Cov6_LazyRegex_String_Nil(t *testing.T) {
	var lr *regexnew.LazyRegex
	actual := args.Map{"v": lr.String()}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "String nil", actual)
}

func Test_Cov6_LazyRegex_String_Valid(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6str\d+$`)
	actual := args.Map{"v": lr.String()}
	expected := args.Map{"v": `^cov6str\d+$`}
	expected.ShouldBeEqual(t, 0, "String valid", actual)
}

func Test_Cov6_LazyRegex_FullString_Nil(t *testing.T) {
	var lr *regexnew.LazyRegex
	actual := args.Map{"v": lr.FullString()}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "FullString nil", actual)
}

func Test_Cov6_LazyRegex_FullString_Valid(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6full\d+$`)
	actual := args.Map{"notEmpty": lr.FullString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullString valid", actual)
}

func Test_Cov6_LazyRegex_Pattern_Nil(t *testing.T) {
	var lr *regexnew.LazyRegex
	actual := args.Map{"v": lr.Pattern()}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "Pattern nil", actual)
}

func Test_Cov6_LazyRegex_Pattern_Valid(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6pat\d+$`)
	actual := args.Map{"v": lr.Pattern()}
	expected := args.Map{"v": `^cov6pat\d+$`}
	expected.ShouldBeEqual(t, 0, "Pattern valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex — Match methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_LazyRegex_MatchError_Match(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6me\d+$`)
	actual := args.Map{"noErr": lr.MatchError("cov6me123") == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError match", actual)
}

func Test_Cov6_LazyRegex_MatchError_NoMatch(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6me2\d+$`)
	actual := args.Map{"hasErr": lr.MatchError("abc") != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError no match", actual)
}

func Test_Cov6_LazyRegex_MatchUsingFuncError_Match(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6mfe\d+$`)
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }
	actual := args.Map{"noErr": lr.MatchUsingFuncError("cov6mfe123", matchFn) == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncError match", actual)
}

func Test_Cov6_LazyRegex_MatchUsingFuncError_NoMatch(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6mfe2\d+$`)
	matchFn := func(re *regexp.Regexp, s string) bool { return re.MatchString(s) }
	actual := args.Map{"hasErr": lr.MatchUsingFuncError("abc", matchFn) != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncError no match", actual)
}

func Test_Cov6_LazyRegex_IsMatch(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6im\d+$`)
	actual := args.Map{"match": lr.IsMatch("cov6im123"), "noMatch": lr.IsMatch("abc")}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsMatch", actual)
}

func Test_Cov6_LazyRegex_IsMatchBytes(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6imb\d+$`)
	actual := args.Map{"match": lr.IsMatchBytes([]byte("cov6imb123")), "noMatch": lr.IsMatchBytes([]byte("abc"))}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsMatchBytes", actual)
}

func Test_Cov6_LazyRegex_IsFailedMatch(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6ifm\d+$`)
	actual := args.Map{"match": lr.IsFailedMatch("cov6ifm123"), "noMatch": lr.IsFailedMatch("abc")}
	expected := args.Map{"match": false, "noMatch": true}
	expected.ShouldBeEqual(t, 0, "IsFailedMatch", actual)
}

func Test_Cov6_LazyRegex_IsFailedMatchBytes(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6ifmb\d+$`)
	actual := args.Map{"match": lr.IsFailedMatchBytes([]byte("cov6ifmb123")), "noMatch": lr.IsFailedMatchBytes([]byte("abc"))}
	expected := args.Map{"match": false, "noMatch": true}
	expected.ShouldBeEqual(t, 0, "IsFailedMatchBytes", actual)
}

func Test_Cov6_LazyRegex_FirstMatchLine_Match(t *testing.T) {
	lr := regexnew.New.LazyLock(`cov6fml(\d+)`)
	line, invalid := lr.FirstMatchLine("cov6fml123")
	actual := args.Map{"line": line, "invalid": invalid}
	expected := args.Map{"line": "cov6fml123", "invalid": false}
	expected.ShouldBeEqual(t, 0, "FirstMatchLine match", actual)
}

func Test_Cov6_LazyRegex_FirstMatchLine_NoMatch(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6fml2\d+$`)
	line, invalid := lr.FirstMatchLine("abc")
	actual := args.Map{"line": line, "invalid": invalid}
	expected := args.Map{"line": "", "invalid": true}
	expected.ShouldBeEqual(t, 0, "FirstMatchLine no match", actual)
}

func Test_Cov6_LazyRegex_CompileMust(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6cm\d+$`)
	re := lr.CompileMust()
	actual := args.Map{"notNil": re != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CompileMust", actual)
}

func Test_Cov6_LazyRegex_OnRequiredCompiledMust(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6orcm\d+$`)
	lr.OnRequiredCompiledMust() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "OnRequiredCompiledMust", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// lazyRegexMap — via New.LazyRegex methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_NewCreator_Lazy(t *testing.T) {
	lr := regexnew.New.Lazy(`^cov6nclazy\d+$`)
	actual := args.Map{"notNil": lr != nil, "defined": lr.IsDefined()}
	expected := args.Map{"notNil": true, "defined": true}
	expected.ShouldBeEqual(t, 0, "New.Lazy", actual)
}

func Test_Cov6_NewCreator_LazyLock(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6nclazylock\d+$`)
	actual := args.Map{"notNil": lr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.LazyLock", actual)
}

func Test_Cov6_NewCreator_Default(t *testing.T) {
	re, err := regexnew.New.Default(`^cov6ncdef\d+$`)
	actual := args.Map{"notNil": re != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "New.Default", actual)
}

func Test_Cov6_NewCreator_DefaultLock(t *testing.T) {
	re, err := regexnew.New.DefaultLock(`^cov6ncdefl\d+$`)
	actual := args.Map{"notNil": re != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "New.DefaultLock", actual)
}

func Test_Cov6_NewCreator_DefaultLockIf(t *testing.T) {
	re, err := regexnew.New.DefaultLockIf(true, `^cov6ncdefli\d+$`)
	actual := args.Map{"notNil": re != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "New.DefaultLockIf", actual)
}

func Test_Cov6_NewCreator_DefaultApplicableLock(t *testing.T) {
	re, err, ok := regexnew.New.DefaultApplicableLock(`^cov6ncal\d+$`)
	actual := args.Map{"notNil": re != nil, "noErr": err == nil, "ok": ok}
	expected := args.Map{"notNil": true, "noErr": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "New.DefaultApplicableLock", actual)
}

func Test_Cov6_NewLazyRegexCreator_TwoLock(t *testing.T) {
	first, second := regexnew.New.LazyRegex.TwoLock(`^cov6two1\d+$`, `^cov6two2\d+$`)
	actual := args.Map{"f": first != nil, "s": second != nil}
	expected := args.Map{"f": true, "s": true}
	expected.ShouldBeEqual(t, 0, "TwoLock", actual)
}

func Test_Cov6_NewLazyRegexCreator_ManyUsingLock_Empty(t *testing.T) {
	result := regexnew.New.LazyRegex.ManyUsingLock()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ManyUsingLock empty", actual)
}

func Test_Cov6_NewLazyRegexCreator_ManyUsingLock(t *testing.T) {
	result := regexnew.New.LazyRegex.ManyUsingLock(`^cov6many1\d+$`, `^cov6many2\d+$`)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ManyUsingLock", actual)
}

func Test_Cov6_NewLazyRegexCreator_AllPatternsMap(t *testing.T) {
	result := regexnew.New.LazyRegex.AllPatternsMap()
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "AllPatternsMap", actual)
}

func Test_Cov6_NewLazyRegexCreator_NewLockIf_Lock(t *testing.T) {
	lr := regexnew.New.LazyRegex.NewLockIf(true, `^cov6nli1\d+$`)
	actual := args.Map{"notNil": lr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewLockIf lock", actual)
}

func Test_Cov6_NewLazyRegexCreator_NewLockIf_NoLock(t *testing.T) {
	lr := regexnew.New.LazyRegex.NewLockIf(false, `^cov6nli2\d+$`)
	actual := args.Map{"notNil": lr != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewLockIf no lock", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// regExMatchValidationError — all 3 branches (via MatchError)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_RegExMatchValidationError_CompileError(t *testing.T) {
	err := regexnew.MatchError(`[invalid`, "abc")
	actual := args.Map{"hasErr": err != nil, "hasCompile": err.Error() != ""}
	expected := args.Map{"hasErr": true, "hasCompile": true}
	expected.ShouldBeEqual(t, 0, "regExMatchValidationError compile err", actual)
}

func Test_Cov6_RegExMatchValidationError_NoMatch(t *testing.T) {
	err := regexnew.MatchError(`^\d+$`, "abc")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "regExMatchValidationError no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// prettyJson — via FullString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_PrettyJson_ViaFullString(t *testing.T) {
	lr := regexnew.New.LazyLock(`^cov6pj\d+$`)
	s := lr.FullString()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "prettyJson via FullString", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// regexes-compiled.go — pre-compiled vars
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_PrecompiledRegexes(t *testing.T) {
	actual := args.Map{
		"ws":   regexnew.WhitespaceFinderRegex != nil,
		"hash": regexnew.HashCommentWithSpaceOptionalRegex != nil,
		"pipe": regexnew.WhitespaceOrPipeFinderRegex != nil,
		"dol":  regexnew.DollarIdentifierRegex != nil,
		"pct":  regexnew.PercentIdentifierRegex != nil,
		"pn":   regexnew.PrettyNameRegex != nil,
		"id":   regexnew.ExactIdFieldMatchingRegex != nil,
		"vid":  regexnew.ExactVersionIdFieldMatchingRegex != nil,
		"ubu":  regexnew.UbuntuNameCheckerRegex != nil,
		"cent": regexnew.CentOsNameCheckerRegex != nil,
		"rh":   regexnew.RedHatNameCheckerRegex != nil,
		"num":  regexnew.FirstNumberAnyWhereCheckerRegex != nil,
		"win":  regexnew.WindowsVersionNumberCheckerRegex != nil,
	}
	expected := args.Map{
		"ws": true, "hash": true, "pipe": true, "dol": true, "pct": true,
		"pn": true, "id": true, "vid": true, "ubu": true, "cent": true,
		"rh": true, "num": true, "win": true,
	}
	expected.ShouldBeEqual(t, 0, "pre-compiled regexes", actual)
}
