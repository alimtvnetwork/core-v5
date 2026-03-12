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
	isMatch, err := regexnew.IsMatchLock(`^\d+$`, "123")
	actual := args.Map{"isMatch": isMatch, "hasErr": err != nil}
	expected := args.Map{"isMatch": true, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns true -- valid match", actual)
}

func Test_Cov3_IsMatchLock_NoMatch(t *testing.T) {
	isMatch, err := regexnew.IsMatchLock(`^\d+$`, "abc")
	actual := args.Map{"isMatch": isMatch, "hasErr": err != nil}
	expected := args.Map{"isMatch": false, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns false -- no match", actual)
}

func Test_Cov3_IsMatchLock_InvalidPattern(t *testing.T) {
	_, err := regexnew.IsMatchLock(`[invalid`, "abc")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "IsMatchLock returns error -- invalid pattern", actual)
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

func Test_Cov3_LazyRegex_Compiled(t *testing.T) {
	lazy := regexnew.New.Lazy(`^\d+$`)
	compiled, err := lazy.Compiled()
	actual := args.Map{"notNil": compiled != nil, "hasErr": err != nil}
	expected := args.Map{"notNil": true, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex Compiled returns regex -- valid", actual)
}

func Test_Cov3_LazyRegex_CompiledLock(t *testing.T) {
	lazy := regexnew.New.LazyLock(`^\d+$`)
	compiled, err := lazy.CompiledLock()
	actual := args.Map{"notNil": compiled != nil, "hasErr": err != nil}
	expected := args.Map{"notNil": true, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex CompiledLock returns regex -- valid lock", actual)
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

func Test_Cov3_LazyRegex_IsMatchLock(t *testing.T) {
	lazy := regexnew.New.LazyLock(`^\d+$`)
	actual := args.Map{
		"matchDigits": lazy.IsMatchLock("123"),
		"failAlpha":   lazy.IsMatchLock("abc"),
	}
	expected := args.Map{"matchDigits": true, "failAlpha": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex IsMatchLock returns expected -- digits vs alpha", actual)
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

func Test_Cov3_LazyRegex_FindAllMatches(t *testing.T) {
	lazy := regexnew.New.Lazy(`\d+`)
	matches := lazy.FindAllMatches("abc123def456")
	actual := args.Map{"len": len(matches)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LazyRegex FindAllMatches returns 2 -- two numbers", actual)
}

func Test_Cov3_LazyRegex_ReplaceAll(t *testing.T) {
	lazy := regexnew.New.Lazy(`\d+`)
	result := lazy.ReplaceAll("abc123def456", "X")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "abcXdefX"}
	expected.ShouldBeEqual(t, 0, "LazyRegex ReplaceAll returns replaced -- digits to X", actual)
}
