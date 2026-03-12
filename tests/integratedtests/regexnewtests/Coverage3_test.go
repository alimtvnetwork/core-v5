package regexnewtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// =============================================================================
// IsMatchLock — additional branches
// =============================================================================

func Test_Cov3_IsMatchLock_Valid(t *testing.T) {
	isMatch := regexnew.IsMatchLock(`^\d+$`, "123")
	actual := args.Map{"isMatch": isMatch}
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns true -- valid match", actual)
}

func Test_Cov3_IsMatchLock_NoMatch(t *testing.T) {
	isMatch := regexnew.IsMatchLock(`^\d+$`, "abc")
	actual := args.Map{"isMatch": isMatch}
	expected := args.Map{"isMatch": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns false -- no match", actual)
}

func Test_Cov3_IsMatchLock_InvalidPattern(t *testing.T) {
	isMatch := regexnew.IsMatchLock(`[invalid`, "abc")
	actual := args.Map{"isMatch": isMatch}
	expected := args.Map{"isMatch": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns false -- invalid pattern", actual)
}

// =============================================================================
// Create — additional branches
// =============================================================================

func Test_Cov3_Create_Valid(t *testing.T) {
	r, err := regexnew.Create(`^\d+$`)
	actual := args.Map{"notNil": r != nil, "hasErr": err != nil}
	expected := args.Map{"notNil": true, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "Create returns regex -- valid pattern", actual)
}

func Test_Cov3_Create_Invalid(t *testing.T) {
	_, err := regexnew.Create(`[invalid`)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Create returns error -- invalid pattern", actual)
}

func Test_Cov3_CreateLock_Valid(t *testing.T) {
	r, err := regexnew.CreateLock(`^\d+$`)
	actual := args.Map{"notNil": r != nil, "hasErr": err != nil}
	expected := args.Map{"notNil": true, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "CreateLock returns regex -- valid", actual)
}

func Test_Cov3_CreateMust_Valid(t *testing.T) {
	r := regexnew.CreateMust(`^\d+$`)
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CreateMust returns regex -- valid", actual)
}

// =============================================================================
// LazyRegex — additional methods
// =============================================================================

func Test_Cov3_LazyRegex_Compile(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	compiled, err := lazy.Compile()
	actual := args.Map{"notNil": compiled != nil, "hasErr": err != nil}
	expected := args.Map{"notNil": true, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex Compile returns regex -- valid", actual)
}

func Test_Cov3_LazyRegex_IsMatch(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	actual := args.Map{
		"matchDigits": lazy.IsMatch("123"),
		"failAlpha":   lazy.IsMatch("abc"),
	}
	expected := args.Map{"matchDigits": true, "failAlpha": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsMatch returns expected -- digits vs alpha", actual)
}

func Test_Cov3_LazyRegex_IsMatchBytes(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	actual := args.Map{
		"matchDigits": lazy.IsMatchBytes([]byte("123")),
		"failAlpha":   lazy.IsMatchBytes([]byte("abc")),
	}
	expected := args.Map{"matchDigits": true, "failAlpha": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsMatchBytes returns expected -- digits vs alpha", actual)
}

func Test_Cov3_LazyRegex_IsFailedMatch(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	actual := args.Map{
		"failAlpha":   lazy.IsFailedMatch("abc"),
		"failDigits":  lazy.IsFailedMatch("123"),
	}
	expected := args.Map{"failAlpha": true, "failDigits": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsFailedMatch returns expected -- alpha fails", actual)
}

func Test_Cov3_LazyRegex_IsFailedMatchBytes(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	actual := args.Map{
		"failAlpha": lazy.IsFailedMatchBytes([]byte("abc")),
	}
	expected := args.Map{"failAlpha": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsFailedMatchBytes returns true -- alpha", actual)
}

func Test_Cov3_LazyRegex_OnRequiredCompiled(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	err := lazy.OnRequiredCompiled()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex OnRequiredCompiled returns nil -- valid", actual)
}

func Test_Cov3_LazyRegex_OnRequiredCompiledMust(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	lazy.OnRequiredCompiledMust() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex OnRequiredCompiledMust no panic -- valid", actual)
}

func Test_Cov3_LazyRegex_StringAndPattern(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	actual := args.Map{
		"string":  lazy.String(),
		"pattern": lazy.Pattern(),
		"isNull":  lazy.IsNull(),
		"isDef":   lazy.IsDefined(),
		"isUndef": lazy.IsUndefined(),
		"isComp":  lazy.IsCompiled(),
		"isAppl":  lazy.IsApplicable(),
	}
	expected := args.Map{
		"string": `^\d+$`, "pattern": `^\d+$`,
		"isNull": false, "isDef": true, "isUndef": false,
		"isComp": true, "isAppl": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex String/Pattern returns expected -- valid", actual)
}

func Test_Cov3_LazyRegex_MatchError_Valid(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	actual := args.Map{
		"matchErr":   lazy.MatchError("123") == nil,
		"noMatchErr": lazy.MatchError("abc") != nil,
	}
	expected := args.Map{"matchErr": true, "noMatchErr": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex MatchError returns expected -- match vs no match", actual)
}

func Test_Cov3_LazyRegex_FirstMatchLine(t *testing.T) {
	lazy := regexnew.New.Lazy(`\d+`)
	match, isInvalid := lazy.FirstMatchLine("abc123def")
	actual := args.Map{"match": match, "isInvalid": isInvalid}
	expected := args.Map{"match": "123", "isInvalid": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex FirstMatchLine returns 123 -- valid", actual)
}

func Test_Cov3_LazyRegex_FirstMatchLine_NoMatch(t *testing.T) {
	lazy := regexnew.New.Lazy(`\d+`)
	match, isInvalid := lazy.FirstMatchLine("abc")
	actual := args.Map{"match": match, "isInvalid": isInvalid}
	expected := args.Map{"match": "", "isInvalid": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex FirstMatchLine returns empty -- no match", actual)
}

func Test_Cov3_LazyRegex_HasError(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	actual := args.Map{"hasError": lazy.HasError(), "hasIssues": lazy.HasAnyIssues(), "isInvalid": lazy.IsInvalid()}
	expected := args.Map{"hasError": false, "hasIssues": false, "isInvalid": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex HasError returns false -- valid pattern", actual)
}

func Test_Cov3_LazyRegex_CompiledError(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	actual := args.Map{"hasErr": lazy.CompiledError() != nil, "errNil": lazy.Error() == nil}
	expected := args.Map{"hasErr": false, "errNil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex CompiledError returns nil -- valid", actual)
}

func Test_Cov3_LazyRegex_MustBeSafe(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	lazy.MustBeSafe() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex MustBeSafe no panic -- valid", actual)
}

func Test_Cov3_LazyRegex_FullString(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	actual := args.Map{"notEmpty": lazy.FullString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex FullString returns non-empty -- valid", actual)
}

func Test_Cov3_LazyRegex_CompileMust(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	r := lazy.CompileMust()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex CompileMust returns regex -- valid", actual)
}
