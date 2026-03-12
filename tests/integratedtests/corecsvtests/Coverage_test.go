package corecsvtests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecsv"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── DefaultCsv / DefaultCsvStrings ──

func Test_Cov_DefaultCsv(t *testing.T) {
	result := corecsv.DefaultCsv("a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DefaultCsv", actual)
}

func Test_Cov_DefaultCsvStrings(t *testing.T) {
	result := corecsv.DefaultCsvStrings("a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultCsvStrings", actual)
}

func Test_Cov_DefaultCsvUsingJoiner(t *testing.T) {
	result := corecsv.DefaultCsvUsingJoiner(" | ", "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DefaultCsvUsingJoiner", actual)
}

// ── DefaultAnyCsv / DefaultAnyCsvStrings ──

func Test_Cov_DefaultAnyCsv(t *testing.T) {
	result := corecsv.DefaultAnyCsv("a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsv", actual)
}

func Test_Cov_DefaultAnyCsvStrings(t *testing.T) {
	result := corecsv.DefaultAnyCsvStrings("a", 1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsvStrings", actual)
}

func Test_Cov_DefaultAnyCsvUsingJoiner(t *testing.T) {
	result := corecsv.DefaultAnyCsvUsingJoiner(" | ", "a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsvUsingJoiner", actual)
}

// ── StringsToCsvStrings all quote branches ──

func Test_Cov_StringsToCsvStrings_SingleQuote(t *testing.T) {
	result := corecsv.StringsToCsvStrings(true, true, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SingleQuote", actual)
}

func Test_Cov_StringsToCsvStrings_NoQuote(t *testing.T) {
	result := corecsv.StringsToCsvStrings(false, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NoQuote", actual)
}

func Test_Cov_StringsToCsvStringsDefault(t *testing.T) {
	result := corecsv.StringsToCsvStringsDefault("a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsToCsvStringsDefault", actual)
}

func Test_Cov_StringsToStringDefault(t *testing.T) {
	result := corecsv.StringsToStringDefault("a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsToStringDefault", actual)
}

// ── AnyItemsToCsvStrings all quote branches ──

func Test_Cov_AnyItemsToCsvStrings_SingleQuote(t *testing.T) {
	result := corecsv.AnyItemsToCsvStrings(true, true, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItems_SingleQuote", actual)
}

func Test_Cov_AnyItemsToCsvStrings_NoQuote(t *testing.T) {
	result := corecsv.AnyItemsToCsvStrings(false, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItems_NoQuote", actual)
}

func Test_Cov_AnyItemsToStringDefault(t *testing.T) {
	result := corecsv.AnyItemsToStringDefault("a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToStringDefault", actual)
}

// ── AnyToTypesCsvStrings all branches ──

func Test_Cov_AnyToTypesCsvStrings_SingleQuote(t *testing.T) {
	result := corecsv.AnyToTypesCsvStrings(true, true, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Types_SingleQuote", actual)
}

func Test_Cov_AnyToTypesCsvStrings_DoubleQuote(t *testing.T) {
	result := corecsv.AnyToTypesCsvStrings(true, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Types_DoubleQuote", actual)
}

func Test_Cov_AnyToTypesCsvDefault(t *testing.T) {
	result := corecsv.AnyToTypesCsvDefault("a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvDefault", actual)
}

// ── AnyToValuesTypeStrings ──

func Test_Cov_AnyToValuesTypeStrings(t *testing.T) {
	result := corecsv.AnyToValuesTypeStrings("a", 1)
	empty := corecsv.AnyToValuesTypeStrings()
	actual := args.Map{"len": len(result), "emptyLen": len(empty)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeStrings", actual)
}

func Test_Cov_AnyToValuesTypeString(t *testing.T) {
	result := corecsv.AnyToValuesTypeString("a")
	empty := corecsv.AnyToValuesTypeString()
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeString", actual)
}

// ── Stringers ──

type covTestStringer struct{ val string }

func (s covTestStringer) String() string { return s.val }

func Test_Cov_StringersToCsvStrings_AllBranches(t *testing.T) {
	s := covTestStringer{val: "x"}
	single := corecsv.StringersToCsvStrings(true, true, s)
	double := corecsv.StringersToCsvStrings(true, false, s)
	noQuote := corecsv.StringersToCsvStrings(false, false, s)
	empty := corecsv.StringersToCsvStrings(false, false)
	actual := args.Map{
		"singleLen":  len(single),
		"doubleLen":  len(double),
		"noQuoteLen": len(noQuote),
		"emptyLen":   len(empty),
	}
	expected := args.Map{
		"singleLen":  1,
		"doubleLen":  1,
		"noQuoteLen": 1,
		"emptyLen":   0,
	}
	expected.ShouldBeEqual(t, 0, "StringersToCsvStrings", actual)
}

func Test_Cov_StringersToString(t *testing.T) {
	s := covTestStringer{val: "x"}
	result := corecsv.StringersToString(", ", false, false, s)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToString", actual)
}

func Test_Cov_StringersToStringDefault(t *testing.T) {
	s := covTestStringer{val: "x"}
	result := corecsv.StringersToStringDefault(s)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToStringDefault", actual)
}

// ── CompileStringers ──

func Test_Cov_CompileStringersToCsvStrings_AllBranches(t *testing.T) {
	f := func() string { return "x" }
	single := corecsv.CompileStringersToCsvStrings(true, true, f)
	double := corecsv.CompileStringersToCsvStrings(true, false, f)
	noQuote := corecsv.CompileStringersToCsvStrings(false, false, f)
	empty := corecsv.CompileStringersToCsvStrings(false, false)
	actual := args.Map{
		"singleLen":  len(single),
		"doubleLen":  len(double),
		"noQuoteLen": len(noQuote),
		"emptyLen":   len(empty),
	}
	expected := args.Map{
		"singleLen":  1,
		"doubleLen":  1,
		"noQuoteLen": 1,
		"emptyLen":   0,
	}
	expected.ShouldBeEqual(t, 0, "CompileStringersToCsvStrings", actual)
}

func Test_Cov_CompileStringersToString(t *testing.T) {
	f := func() string { return "x" }
	result := corecsv.CompileStringersToString(", ", false, false, f)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString", actual)
}

func Test_Cov_CompileStringersToStringDefault(t *testing.T) {
	f := func() string { return "x" }
	result := corecsv.CompileStringersToStringDefault(f)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToStringDefault", actual)
}

// ── StringFunctionsToString all branches ──

func Test_Cov_StringFunctionsToString_AllBranches(t *testing.T) {
	f := func() string { return "x" }
	single := corecsv.StringFunctionsToString(true, true, f)
	double := corecsv.StringFunctionsToString(true, false, f)
	noQuote := corecsv.StringFunctionsToString(false, false, f)
	empty := corecsv.StringFunctionsToString(false, false)
	actual := args.Map{
		"singleLen":  len(single),
		"doubleLen":  len(double),
		"noQuoteLen": len(noQuote),
		"emptyLen":   len(empty),
	}
	expected := args.Map{
		"singleLen":  1,
		"doubleLen":  1,
		"noQuoteLen": 1,
		"emptyLen":   0,
	}
	expected.ShouldBeEqual(t, 0, "StringFunctionsToString", actual)
}

// ── RangeNames ──

func Test_Cov_RangeNamesWithValuesIndexesCsvString(t *testing.T) {
	result := corecsv.RangeNamesWithValuesIndexesCsvString("A", "B")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexesCsvString", actual)
}

func Test_Cov_RangeNamesWithValuesIndexesString(t *testing.T) {
	result := corecsv.RangeNamesWithValuesIndexesString(" | ", "A", "B")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexesString", actual)
}
