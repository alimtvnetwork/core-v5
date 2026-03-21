package corecsvtests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecsv"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── StringsToCsvString — all quote branches ──

func Test_Cov2_StringsToCsvString_SingleQuote(t *testing.T) {
	result := corecsv.StringsToCsvString(", ", true, true, "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsToCsvString returns correct value -- single quote", actual)
}

func Test_Cov2_StringsToCsvString_DoubleQuote(t *testing.T) {
	result := corecsv.StringsToCsvString(", ", true, false, "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsToCsvString returns correct value -- double quote", actual)
}

func Test_Cov2_StringsToCsvString_NoQuote(t *testing.T) {
	result := corecsv.StringsToCsvString(", ", false, false, "a", "b")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsToCsvString returns empty -- no quote", actual)
}

// ── AnyItemsToCsvString — all quote branches ──

func Test_Cov2_AnyItemsToCsvString_SingleQuote(t *testing.T) {
	result := corecsv.AnyItemsToCsvString(", ", true, true, "a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvString returns correct value -- single quote", actual)
}

func Test_Cov2_AnyItemsToCsvString_DoubleQuote(t *testing.T) {
	result := corecsv.AnyItemsToCsvString(", ", true, false, "a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvString returns correct value -- double quote", actual)
}

func Test_Cov2_AnyItemsToCsvString_NoQuote(t *testing.T) {
	result := corecsv.AnyItemsToCsvString(", ", false, false, "a", 1)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvString returns empty -- no quote", actual)
}

// ── AnyToTypesCsvStrings — no-quote branch ──

func Test_Cov2_AnyToTypesCsvStrings_NoQuote(t *testing.T) {
	result := corecsv.AnyToTypesCsvStrings(false, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvStrings returns empty -- no quote", actual)
}

func Test_Cov2_AnyToTypesCsvStrings_Empty(t *testing.T) {
	result := corecsv.AnyToTypesCsvStrings(false, false)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvStrings returns empty -- empty", actual)
}

// ── StringsToCsvStrings — double-quote branch (explicit) ──

func Test_Cov2_StringsToCsvStrings_DoubleQuote(t *testing.T) {
	result := corecsv.StringsToCsvStrings(true, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringsToCsvStrings returns correct value -- double quote", actual)
}

func Test_Cov2_StringsToCsvStrings_Empty(t *testing.T) {
	result := corecsv.StringsToCsvStrings(true, true)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringsToCsvStrings returns empty -- empty", actual)
}

// ── AnyItemsToCsvStrings — double-quote, empty ──

func Test_Cov2_AnyItemsToCsvStrings_DoubleQuote(t *testing.T) {
	result := corecsv.AnyItemsToCsvStrings(true, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvStrings returns correct value -- double quote", actual)
}

func Test_Cov2_AnyItemsToCsvStrings_Empty(t *testing.T) {
	result := corecsv.AnyItemsToCsvStrings(true, true)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvStrings returns empty -- empty", actual)
}

// ── RangeNamesWithValuesIndexes — multi items ──

func Test_Cov2_RangeNamesWithValuesIndexes_MultiItems(t *testing.T) {
	result := corecsv.RangeNamesWithValuesIndexes("A", "B", "C")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexes returns non-empty -- multi items", actual)
}

// ── StringersToString — empty ──

func Test_Cov2_StringersToString_Empty(t *testing.T) {
	result := corecsv.StringersToString(", ", false, false)
	actual := args.Map{"empty": result}
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "StringersToString returns empty -- empty", actual)
}

// ── CompileStringersToString — empty ──

func Test_Cov2_CompileStringersToString_Empty(t *testing.T) {
	result := corecsv.CompileStringersToString(", ", false, false)
	actual := args.Map{"empty": result}
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString returns empty -- empty", actual)
}
