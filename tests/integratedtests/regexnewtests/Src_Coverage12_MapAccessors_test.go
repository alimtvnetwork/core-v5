package regexnewtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

// ══════════════════════════════════════════════════════════════════════════════
// Tests for newLazyRegexCreator exported accessors that delegate to
// unexported lazyRegexMap methods. Closes the integrated coverage gap.
// ══════════════════════════════════════════════════════════════════════════════

// ── Count / CountLock ──

func Test_Src12_LazyRegexCreator_Count(t *testing.T) {
	// Arrange — ensure at least one entry exists
	_ = regexnew.New.Lazy(`^src12-count$`)

	// Act
	count := regexnew.New.LazyRegex.Count()

	// Assert
	actual := args.Map{"positive": count > 0}
	expected := args.Map{"positive": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.Count returns positive -- after New.Lazy", actual)
}

func Test_Src12_LazyRegexCreator_CountLock(t *testing.T) {
	// Arrange
	_ = regexnew.New.Lazy(`^src12-countlock$`)

	// Act
	count := regexnew.New.LazyRegex.CountLock()

	// Assert
	actual := args.Map{"positive": count > 0}
	expected := args.Map{"positive": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.CountLock returns positive -- after New.Lazy", actual)
}

// ── IsEmpty / IsEmptyLock ──

func Test_Src12_LazyRegexCreator_IsEmpty(t *testing.T) {
	// Arrange — the global map is non-empty from init vars
	// Act
	result := regexnew.New.LazyRegex.IsEmpty()

	// Assert
	actual := args.Map{"isEmpty": result}
	expected := args.Map{"isEmpty": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsEmpty returns false -- has items", actual)
}

func Test_Src12_LazyRegexCreator_IsEmptyLock(t *testing.T) {
	// Arrange & Act
	result := regexnew.New.LazyRegex.IsEmptyLock()

	// Assert
	actual := args.Map{"isEmpty": result}
	expected := args.Map{"isEmpty": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.IsEmptyLock returns false -- has items", actual)
}

// ── HasAnyItem / HasAnyItemLock ──

func Test_Src12_LazyRegexCreator_HasAnyItem(t *testing.T) {
	// Arrange & Act
	result := regexnew.New.LazyRegex.HasAnyItem()

	// Assert
	actual := args.Map{"hasAnyItem": result}
	expected := args.Map{"hasAnyItem": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.HasAnyItem returns true -- has items", actual)
}

func Test_Src12_LazyRegexCreator_HasAnyItemLock(t *testing.T) {
	// Arrange & Act
	result := regexnew.New.LazyRegex.HasAnyItemLock()

	// Assert
	actual := args.Map{"hasAnyItem": result}
	expected := args.Map{"hasAnyItem": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.HasAnyItemLock returns true -- has items", actual)
}

// ── Has / HasLock ──

func Test_Src12_LazyRegexCreator_Has_Existing(t *testing.T) {
	// Arrange
	pattern := `^src12-has-existing$`
	_ = regexnew.New.Lazy(pattern)

	// Act
	result := regexnew.New.LazyRegex.Has(pattern)

	// Assert
	actual := args.Map{"has": result}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.Has returns true -- existing pattern", actual)
}

func Test_Src12_LazyRegexCreator_Has_Missing(t *testing.T) {
	// Arrange & Act
	result := regexnew.New.LazyRegex.Has(`^never-registered-src12$`)

	// Assert
	actual := args.Map{"has": result}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.Has returns false -- missing pattern", actual)
}

func Test_Src12_LazyRegexCreator_HasLock_Existing(t *testing.T) {
	// Arrange
	pattern := `^src12-haslock-existing$`
	_ = regexnew.New.Lazy(pattern)

	// Act
	result := regexnew.New.LazyRegex.HasLock(pattern)

	// Assert
	actual := args.Map{"has": result}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "LazyRegex.HasLock returns true -- existing pattern", actual)
}

func Test_Src12_LazyRegexCreator_HasLock_Missing(t *testing.T) {
	// Arrange & Act
	result := regexnew.New.LazyRegex.HasLock(`^never-registered-src12-lock$`)

	// Assert
	actual := args.Map{"has": result}
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "LazyRegex.HasLock returns false -- missing pattern", actual)
}

// ── NewLockIf (refactored to use CreateOrExistingLockIf) ──

func Test_Src12_LazyRegexCreator_NewLockIf_Locked(t *testing.T) {
	// Arrange & Act
	lr := regexnew.New.LazyRegex.NewLockIf(true, `^src12-lockif-locked$`)

	// Assert
	actual := args.Map{
		"notNil":    lr != nil,
		"isDefined": lr.IsDefined(),
	}
	expected := args.Map{
		"notNil":    true,
		"isDefined": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.NewLockIf returns defined -- locked", actual)
}

func Test_Src12_LazyRegexCreator_NewLockIf_Unlocked(t *testing.T) {
	// Arrange & Act
	lr := regexnew.New.LazyRegex.NewLockIf(false, `^src12-lockif-unlocked$`)

	// Assert
	actual := args.Map{
		"notNil":    lr != nil,
		"isDefined": lr.IsDefined(),
	}
	expected := args.Map{
		"notNil":    true,
		"isDefined": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.NewLockIf returns defined -- unlocked", actual)
}

// ── NewWithCompiler ──

func Test_Src12_LazyRegexCreator_NewWithCompiler_Valid(t *testing.T) {
	// Arrange
	compiler := func(pattern string) (*regexp.Regexp, error) {
		return regexp.Compile(pattern)
	}

	// Act
	lr := regexnew.New.LazyRegex.NewWithCompiler(`^src12-custom$`, compiler)

	// Assert
	actual := args.Map{
		"notNil":       lr != nil,
		"isDefined":    lr.IsDefined(),
		"isApplicable": lr.IsApplicable(),
	}
	expected := args.Map{
		"notNil":       true,
		"isDefined":    true,
		"isApplicable": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.NewWithCompiler returns applicable -- valid custom compiler", actual)
}

func Test_Src12_LazyRegexCreator_NewWithCompiler_Invalid(t *testing.T) {
	// Arrange
	compiler := func(pattern string) (*regexp.Regexp, error) {
		return regexp.Compile(pattern)
	}

	// Act
	lr := regexnew.New.LazyRegex.NewWithCompiler(`[invalid`, compiler)

	// Assert
	actual := args.Map{
		"notNil":       lr != nil,
		"isApplicable": lr.IsApplicable(),
		"hasError":     lr.HasError(),
	}
	expected := args.Map{
		"notNil":       true,
		"isApplicable": false,
		"hasError":     true,
	}
	expected.ShouldBeEqual(t, 0, "LazyRegex.NewWithCompiler returns not applicable -- invalid pattern", actual)
}
