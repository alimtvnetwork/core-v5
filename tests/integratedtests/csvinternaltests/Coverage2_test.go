package csvinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/csvinternal"
)

// ── StringsToCsvStrings all branches ──

func Test_Cov2_StringsToCsvStrings_SingleQuote(t *testing.T) {
	result := csvinternal.StringsToCsvStrings(true, true, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SingleQuote", actual)
}

func Test_Cov2_StringsToCsvStrings_NoQuote(t *testing.T) {
	result := csvinternal.StringsToCsvStrings(false, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NoQuote", actual)
}

func Test_Cov2_StringsToCsvStrings_Empty(t *testing.T) {
	result := csvinternal.StringsToCsvStrings(false, false)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty", actual)
}

func Test_Cov2_StringsToCsvStringsDefault(t *testing.T) {
	result := csvinternal.StringsToCsvStringsDefault("a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Default", actual)
}

func Test_Cov2_StringsToStringDefault(t *testing.T) {
	result := csvinternal.StringsToStringDefault("a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsToStringDefault", actual)
}

func Test_Cov2_StringsToStringDefaultNoQuotations(t *testing.T) {
	result := csvinternal.StringsToStringDefaultNoQuotations("a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NoQuotations", actual)
}

// ── AnyItemsToCsvStrings all branches ──

func Test_Cov2_AnyItemsToCsvStrings_SingleQuote(t *testing.T) {
	result := csvinternal.AnyItemsToCsvStrings(true, true, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Any_SingleQuote", actual)
}

func Test_Cov2_AnyItemsToCsvStrings_DoubleQuote(t *testing.T) {
	result := csvinternal.AnyItemsToCsvStrings(true, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Any_DoubleQuote", actual)
}

func Test_Cov2_AnyItemsToCsvStrings_Empty(t *testing.T) {
	result := csvinternal.AnyItemsToCsvStrings(false, false)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Any_Empty", actual)
}

func Test_Cov2_AnyItemsToStringDefault(t *testing.T) {
	result := csvinternal.AnyItemsToStringDefault("a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToStringDefault", actual)
}

// ── Stringers all branches ──

type cov2Stringer struct{ val string }

func (s cov2Stringer) String() string { return s.val }

func Test_Cov2_StringersToCsvStrings_AllBranches(t *testing.T) {
	s := cov2Stringer{val: "x"}
	single := csvinternal.StringersToCsvStrings(true, true, s)
	double := csvinternal.StringersToCsvStrings(true, false, s)
	noQuote := csvinternal.StringersToCsvStrings(false, false, s)
	empty := csvinternal.StringersToCsvStrings(false, false)
	actual := args.Map{
		"singleLen": len(single), "doubleLen": len(double),
		"noQuoteLen": len(noQuote), "emptyLen": len(empty),
	}
	expected := args.Map{
		"singleLen": 1, "doubleLen": 1,
		"noQuoteLen": 1, "emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "StringersToCsvStrings", actual)
}

func Test_Cov2_StringersToStringDefault(t *testing.T) {
	s := cov2Stringer{val: "x"}
	result := csvinternal.StringersToStringDefault(s)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToStringDefault", actual)
}

// ── CompileStringers all branches ──

func Test_Cov2_CompileStringersToCsvStrings_AllBranches(t *testing.T) {
	f := func() string { return "x" }
	single := csvinternal.CompileStringersToCsvStrings(true, true, f)
	double := csvinternal.CompileStringersToCsvStrings(true, false, f)
	noQuote := csvinternal.CompileStringersToCsvStrings(false, false, f)
	empty := csvinternal.CompileStringersToCsvStrings(false, false)
	actual := args.Map{
		"singleLen": len(single), "doubleLen": len(double),
		"noQuoteLen": len(noQuote), "emptyLen": len(empty),
	}
	expected := args.Map{
		"singleLen": 1, "doubleLen": 1,
		"noQuoteLen": 1, "emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "CompileStringersCsvStrings", actual)
}

func Test_Cov2_CompileStringersToStringDefault(t *testing.T) {
	f := func() string { return "x" }
	result := csvinternal.CompileStringersToStringDefault(f)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToStringDefault", actual)
}

// ── RangeNames ──

func Test_Cov2_RangeNamesWithValuesIndexesCsvString(t *testing.T) {
	result := csvinternal.RangeNamesWithValuesIndexesCsvString("A", "B")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsvString", actual)
}

// ── AnyItemsToCsvString single quote ──

func Test_Cov2_AnyItemsToCsvString_SingleQuote(t *testing.T) {
	result := csvinternal.AnyItemsToCsvString(", ", true, true, "a")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItems_CsvString_SingleQuote", actual)
}

// ── StringsToCsvString single quote ──

func Test_Cov2_StringsToCsvString_SingleQuote(t *testing.T) {
	result := csvinternal.StringsToCsvString(", ", true, true, "a")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Strings_CsvString_SingleQuote", actual)
}

