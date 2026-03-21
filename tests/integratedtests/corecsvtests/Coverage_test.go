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
	expected.ShouldBeEqual(t, 0, "DefaultCsv returns correct value -- with args", actual)
}

func Test_Cov_DefaultCsvStrings(t *testing.T) {
	result := corecsv.DefaultCsvStrings("a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultCsvStrings returns correct value -- with args", actual)
}

func Test_Cov_DefaultCsvUsingJoiner(t *testing.T) {
	result := corecsv.DefaultCsvUsingJoiner(" | ", "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DefaultCsvUsingJoiner returns correct value -- with args", actual)
}

// ── DefaultAnyCsv / DefaultAnyCsvStrings ──

func Test_Cov_DefaultAnyCsv(t *testing.T) {
	result := corecsv.DefaultAnyCsv("a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsv returns correct value -- with args", actual)
}

func Test_Cov_DefaultAnyCsvStrings(t *testing.T) {
	result := corecsv.DefaultAnyCsvStrings("a", 1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsvStrings returns correct value -- with args", actual)
}

func Test_Cov_DefaultAnyCsvUsingJoiner(t *testing.T) {
	result := corecsv.DefaultAnyCsvUsingJoiner(" | ", "a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsvUsingJoiner returns correct value -- with args", actual)
}

// ── StringsToCsvStrings all quote branches ──

func Test_Cov_StringsToCsvStrings_SingleQuote(t *testing.T) {
	result := corecsv.StringsToCsvStrings(true, true, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SingleQuote returns correct value -- with args", actual)
}

func Test_Cov_StringsToCsvStrings_NoQuote(t *testing.T) {
	result := corecsv.StringsToCsvStrings(false, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NoQuote returns correct value -- with args", actual)
}

func Test_Cov_StringsToCsvStringsDefault(t *testing.T) {
	result := corecsv.StringsToCsvStringsDefault("a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsToCsvStringsDefault returns correct value -- with args", actual)
}

func Test_Cov_StringsToStringDefault(t *testing.T) {
	result := corecsv.StringsToStringDefault("a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsToStringDefault returns correct value -- with args", actual)
}

// ── AnyItemsToCsvStrings all quote branches ──

func Test_Cov_AnyItemsToCsvStrings_SingleQuote(t *testing.T) {
	result := corecsv.AnyItemsToCsvStrings(true, true, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItems_SingleQuote returns correct value -- with args", actual)
}

func Test_Cov_AnyItemsToCsvStrings_NoQuote(t *testing.T) {
	result := corecsv.AnyItemsToCsvStrings(false, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItems_NoQuote returns correct value -- with args", actual)
}

func Test_Cov_AnyItemsToStringDefault(t *testing.T) {
	result := corecsv.AnyItemsToStringDefault("a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToStringDefault returns correct value -- with args", actual)
}

// ── AnyToTypesCsvStrings all branches ──

func Test_Cov_AnyToTypesCsvStrings_SingleQuote(t *testing.T) {
	result := corecsv.AnyToTypesCsvStrings(true, true, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Types_SingleQuote returns correct value -- with args", actual)
}

func Test_Cov_AnyToTypesCsvStrings_DoubleQuote(t *testing.T) {
	result := corecsv.AnyToTypesCsvStrings(true, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Types_DoubleQuote returns correct value -- with args", actual)
}

func Test_Cov_AnyToTypesCsvDefault(t *testing.T) {
	result := corecsv.AnyToTypesCsvDefault("a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvDefault returns correct value -- with args", actual)
}

// ── AnyToValuesTypeStrings ──

func Test_Cov_AnyToValuesTypeStrings(t *testing.T) {
	result := corecsv.AnyToValuesTypeStrings("a", 1)
	empty := corecsv.AnyToValuesTypeStrings()
	actual := args.Map{"len": len(result), "emptyLen": len(empty)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeStrings returns non-empty -- with args", actual)
}

func Test_Cov_AnyToValuesTypeString(t *testing.T) {
	result := corecsv.AnyToValuesTypeString("a")
	empty := corecsv.AnyToValuesTypeString()
	actual := args.Map{"notEmpty": result != "", "empty": empty}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeString returns non-empty -- with args", actual)
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
	expected.ShouldBeEqual(t, 0, "StringersToCsvStrings returns correct value -- with args", actual)
}

func Test_Cov_StringersToString(t *testing.T) {
	s := covTestStringer{val: "x"}
	result := corecsv.StringersToString(", ", false, false, s)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToString returns correct value -- with args", actual)
}

func Test_Cov_StringersToStringDefault(t *testing.T) {
	s := covTestStringer{val: "x"}
	result := corecsv.StringersToStringDefault(s)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToStringDefault returns correct value -- with args", actual)
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
	expected.ShouldBeEqual(t, 0, "CompileStringersToCsvStrings returns correct value -- with args", actual)
}

func Test_Cov_CompileStringersToString(t *testing.T) {
	f := func() string { return "x" }
	result := corecsv.CompileStringersToString(", ", false, false, f)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString returns correct value -- with args", actual)
}

func Test_Cov_CompileStringersToStringDefault(t *testing.T) {
	f := func() string { return "x" }
	result := corecsv.CompileStringersToStringDefault(f)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToStringDefault returns correct value -- with args", actual)
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
	expected.ShouldBeEqual(t, 0, "StringFunctionsToString returns correct value -- with args", actual)
}

// ── RangeNames ──

func Test_Cov_RangeNamesWithValuesIndexesCsvString(t *testing.T) {
	result := corecsv.RangeNamesWithValuesIndexesCsvString("A", "B")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexesCsvString returns non-empty -- with args", actual)
}

func Test_Cov_RangeNamesWithValuesIndexesString(t *testing.T) {
	result := corecsv.RangeNamesWithValuesIndexesString(" | ", "A", "B")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexesString returns non-empty -- with args", actual)
}

// ── Empty-items branch coverage ──

func Test_Cov_AnyItemsToCsvString_EmptyItems(t *testing.T) {
	result := corecsv.AnyItemsToCsvString(", ", true, false)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvString returns empty -- no items", actual)
}

func Test_Cov_StringsToCsvString_EmptyItems(t *testing.T) {
	result := corecsv.StringsToCsvString(", ", true, false)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "StringsToCsvString returns empty -- no items", actual)
}

func Test_Cov_RangeNamesWithValuesIndexes_EmptyItems(t *testing.T) {
	result := corecsv.RangeNamesWithValuesIndexes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexes returns empty -- no items", actual)
}

func Test_Cov_AnyToValuesTypeStrings_EmptyStringItem(t *testing.T) {
	// Covers the finalString == "" branch inside AnyToValuesTypeStrings
	result := corecsv.AnyToValuesTypeStrings(nil)
	actual := args.Map{"len": len(result), "notPanicked": true}
	expected := args.Map{"len": 1, "notPanicked": true}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeStrings covers empty-string item -- nil input", actual)
}
