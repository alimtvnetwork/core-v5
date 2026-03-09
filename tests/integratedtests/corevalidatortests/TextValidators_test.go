package corevalidatortests

import (
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

// ==========================================
// TextValidators — collection operations
// ==========================================

func Test_TextValidators_NewEmpty(t *testing.T) {
	tc := textValidatorsNewEmptyTestCase
	v := corevalidator.NewTextValidators(5)

	actual := args.Map{
		"isEmpty": v.IsEmpty(),
		"length":  v.Length(),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_Add(t *testing.T) {
	tc := textValidatorsAddTestCase
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	v.Add(corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal})

	actual := args.Map{"length": v.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_Adds(t *testing.T) {
	tc := textValidatorsAddsTestCase
	v := corevalidator.NewTextValidators(2)
	v.Adds(
		corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
		corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal},
	)

	actual := args.Map{"length": v.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_Adds_Empty(t *testing.T) {
	tc := textValidatorsAddsEmptyTestCase
	v := corevalidator.NewTextValidators(2)
	v.Adds()

	actual := args.Map{"length": v.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_AddSimple(t *testing.T) {
	tc := textValidatorsAddSimpleTestCase
	v := corevalidator.NewTextValidators(1)
	v.AddSimple("test", stringcompareas.Contains)

	actual := args.Map{"length": v.Length()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_HasIndex(t *testing.T) {
	tc := textValidatorsHasIndexTestCase
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})

	actual := args.Map{
		"hasIndex0": v.HasIndex(0),
		"hasIndex1": v.HasIndex(1),
	}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_LastIndex(t *testing.T) {
	tc := textValidatorsLastIndexTestCase
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	v.Add(corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal})

	actual := args.Map{"lastIndex": v.LastIndex()}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidators.IsMatch
// ==========================================

func Test_TextValidators_IsMatch_EmptyValidators(t *testing.T) {
	tc := textValidatorsIsMatchEmptyTestCase
	v := corevalidator.NewTextValidators(0)

	actual := args.Map{"isMatch": v.IsMatch("anything", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_IsMatch_AllPass(t *testing.T) {
	tc := textValidatorsIsMatchAllPassTestCase
	v := corevalidator.NewTextValidators(2)
	v.AddSimple("hello", stringcompareas.Contains)
	v.AddSimple("world", stringcompareas.Contains)

	actual := args.Map{"isMatch": v.IsMatch("hello world", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_IsMatch_OneFails(t *testing.T) {
	tc := textValidatorsIsMatchOneFailsTestCase
	v := corevalidator.NewTextValidators(2)
	v.AddSimple("hello", stringcompareas.Contains)
	v.AddSimple("xyz", stringcompareas.Contains)

	actual := args.Map{"isMatch": v.IsMatch("hello world", true)}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidators.IsMatchMany
// ==========================================

func Test_TextValidators_IsMatchMany_EmptyValidators(t *testing.T) {
	tc := textValidatorsIsMatchManyEmptyTestCase
	v := corevalidator.NewTextValidators(0)

	actual := args.Map{"isMatch": v.IsMatchMany(false, true, "a", "b")}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidators.VerifyFirstError
// ==========================================

func Test_TextValidators_VerifyFirstError_AllPass(t *testing.T) {
	tc := textValidatorsVerifyFirstAllPassTestCase
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})

	actual := args.Map{"hasError": v.VerifyFirstError(0, "hello", true) != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_VerifyFirstError_Fails(t *testing.T) {
	tc := textValidatorsVerifyFirstFailTestCase
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})

	actual := args.Map{"hasError": v.VerifyFirstError(0, "world", true) != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_VerifyFirstError_Empty(t *testing.T) {
	tc := textValidatorsVerifyFirstEmptyTestCase
	v := corevalidator.NewTextValidators(0)

	actual := args.Map{"hasError": v.VerifyFirstError(0, "anything", true) != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidators.AllVerifyError
// ==========================================

func Test_TextValidators_AllVerifyError_AllPass(t *testing.T) {
	tc := textValidatorsAllVerifyPassTestCase
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})

	actual := args.Map{"hasError": v.AllVerifyError(0, "hello", true) != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_AllVerifyError_MultipleFail(t *testing.T) {
	tc := textValidatorsAllVerifyFailTestCase
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{
		Search:    "x",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	v.Add(corevalidator.TextValidator{
		Search:    "y",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})

	actual := args.Map{"hasError": v.AllVerifyError(0, "z", true) != nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidators.Dispose
// ==========================================

func Test_TextValidators_Dispose(t *testing.T) {
	tc := textValidatorsDisposeTestCase
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	v.Dispose()

	actual := args.Map{"isNil": v.Items == nil}

	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver tests migrated to TextValidators_NilReceiver_testcases.go)
