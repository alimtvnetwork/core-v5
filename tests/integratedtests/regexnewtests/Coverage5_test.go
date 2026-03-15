package regexnewtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ── CreateMustLockIf ──

func Test_Cov5_CreateMustLockIf_WithLock(t *testing.T) {
	r := regexnew.CreateMustLockIf(true, `^\d+$`)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMustLockIf with lock", actual)
}

func Test_Cov5_CreateMustLockIf_WithoutLock(t *testing.T) {
	r := regexnew.CreateMustLockIf(false, `^\d+$`)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMustLockIf without lock", actual)
}

// ── CreateApplicableLock ──

func Test_Cov5_CreateApplicableLock_Valid(t *testing.T) {
	r, err, ok := regexnew.CreateApplicableLock(`^\d+$`)
	actual := args.Map{"notNil": r != nil, "noErr": err == nil, "ok": ok}
	expected := args.Map{"notNil": true, "noErr": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "CreateApplicableLock valid", actual)
}

func Test_Cov5_CreateApplicableLock_Invalid(t *testing.T) {
	r, err, ok := regexnew.CreateApplicableLock(`[invalid`)
	actual := args.Map{"isNil": r == nil, "hasErr": err != nil, "ok": ok}
	expected := args.Map{"isNil": true, "hasErr": true, "ok": false}
	expected.ShouldBeEqual(t, 0, "CreateApplicableLock invalid", actual)
}

// ── IsMatchFailed ──

func Test_Cov5_IsMatchFailed(t *testing.T) {
	actual := args.Map{
		"fail":    regexnew.IsMatchFailed(`^\d+$`, "abc"),
		"noFail":  regexnew.IsMatchFailed(`^\d+$`, "123"),
		"invalid": regexnew.IsMatchFailed(`[invalid`, "abc"),
	}
	expected := args.Map{"fail": true, "noFail": false, "invalid": true}
	expected.ShouldBeEqual(t, 0, "IsMatchFailed", actual)
}

// ── MatchError ──

func Test_Cov5_MatchError_Success(t *testing.T) {
	err := regexnew.MatchError(`^\d+$`, "123")
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError success", actual)
}

func Test_Cov5_MatchError_Fail(t *testing.T) {
	err := regexnew.MatchError(`^\d+$`, "abc")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError fail", actual)
}

// ── MatchErrorLock ──

func Test_Cov5_MatchErrorLock(t *testing.T) {
	err := regexnew.MatchErrorLock(`^\d+$`, "123")
	errFail := regexnew.MatchErrorLock(`^\d+$`, "abc")
	actual := args.Map{"noErr": err == nil, "hasErr": errFail != nil}
	expected := args.Map{"noErr": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchErrorLock", actual)
}

// ── MatchUsingFuncErrorLock ──

func Test_Cov5_MatchUsingFuncErrorLock(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	err := regexnew.MatchUsingFuncErrorLock(`^\d+$`, "123", matchFunc)
	errFail := regexnew.MatchUsingFuncErrorLock(`^\d+$`, "abc", matchFunc)
	actual := args.Map{"noErr": err == nil, "hasErr": errFail != nil}
	expected := args.Map{"noErr": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock", actual)
}

// ── MatchUsingCustomizeErrorFuncLock — nil custom error func ──

func Test_Cov5_MatchCustomErr_NilCustom(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool { return r.MatchString(s) }
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^\d+$`, "abc", matchFunc, nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchCustomErr nil custom func", actual)
}

func Test_Cov5_MatchCustomErr_InvalidRegex(t *testing.T) {
	matchFunc := func(r *regexp.Regexp, s string) bool { return false }
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`[invalid`, "abc", matchFunc, nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchCustomErr invalid regex", actual)
}

// ── PrettyJson ──

func Test_Cov5_LazyRegex_FullString(t *testing.T) {
	lr := regexnew.New.LazyLock(`^\d+$`)
	result := lr.FullString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex FullString", actual)
}

// ── newCreator — All creator methods ──

func Test_Cov5_New_Must(t *testing.T) {
	r := regexnew.CreateMust(`^\d+$`)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.Must", actual)
}

func Test_Cov5_New_MustLock(t *testing.T) {
	r := regexnew.NewMustLock(`^\d+$`)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "New.MustLock", actual)
}

func Test_Cov5_New_Create(t *testing.T) {
	r, err := regexnew.New.Default(`^\d+$`)
	actual := args.Map{"notNil": r != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "New.Create", actual)
}

func Test_Cov5_New_CreateLock(t *testing.T) {
	r, err := regexnew.New.DefaultLock(`^\d+$`)
	actual := args.Map{"notNil": r != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "New.CreateLock", actual)
}

func Test_Cov5_New_DefaultLockIf(t *testing.T) {
	r, err := regexnew.New.DefaultLockIf(true, `^\d+$`)
	actual := args.Map{"notNil": r != nil, "noErr": err == nil}
	expected := args.Map{"notNil": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "New.DefaultLockIf", actual)
}

func Test_Cov5_New_DefaultApplicableLock(t *testing.T) {
	r, err, ok := regexnew.New.DefaultApplicableLock(`^\d+$`)
	actual := args.Map{"notNil": r != nil, "noErr": err == nil, "ok": ok}
	expected := args.Map{"notNil": true, "noErr": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "New.DefaultApplicableLock", actual)
}

// ── regExMatchValidationError ──

func Test_Cov5_RegExMatchValidationError(t *testing.T) {
	// regExMatchValidationError is unexported; test via MatchError instead
	err := regexnew.MatchError(`^\d+$`, "abc")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RegExMatchValidationError via MatchError", actual)
}

// ── LazyRegex — FindStringSubmatch / FindAllString ──

func Test_Cov5_LazyRegex_FindStringSubmatch(t *testing.T) {
	lr := regexnew.New.LazyLock(`(\d+)-(\d+)`)
	result := lr.FindStringSubmatch("abc 123-456 def")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LazyRegex FindStringSubmatch", actual)
}

func Test_Cov5_LazyRegex_FindAllString(t *testing.T) {
	lr := regexnew.New.LazyLock(`\d+`)
	result := lr.FindAllString("abc 123 def 456", -1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LazyRegex FindAllString", actual)
}

func Test_Cov5_LazyRegex_ReplaceAllString(t *testing.T) {
	lr := regexnew.New.LazyLock(`\d+`)
	result := lr.ReplaceAllString("abc 123 def", "NUM")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "abc NUM def"}
	expected.ShouldBeEqual(t, 0, "LazyRegex ReplaceAllString", actual)
}
