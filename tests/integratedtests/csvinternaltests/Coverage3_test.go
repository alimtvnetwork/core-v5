package csvinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/csvinternal"
)

// ── StringsToCsvStrings double quote branch ──

func Test_Cov3_StringsToCsvStrings_DoubleQuote(t *testing.T) {
	result := csvinternal.StringsToCsvStrings(true, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Strings_DoubleQuote returns correct value -- with args", actual)
}

// ── AnyItemsToCsvStrings no quote ──

func Test_Cov3_AnyItemsToCsvStrings_NoQuote(t *testing.T) {
	result := csvinternal.AnyItemsToCsvStrings(false, false, "a")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Any_NoQuote returns correct value -- with args", actual)
}

// ── AnyItemsToCsvString all branches ──

func Test_Cov3_AnyItemsToCsvString_DoubleQuote(t *testing.T) {
	result := csvinternal.AnyItemsToCsvString(", ", true, false, "a")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItems_CsvString_DoubleQuote returns correct value -- with args", actual)
}

func Test_Cov3_AnyItemsToCsvString_NoQuote(t *testing.T) {
	result := csvinternal.AnyItemsToCsvString(", ", false, false, "a")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItems_CsvString_NoQuote returns correct value -- with args", actual)
}

func Test_Cov3_AnyItemsToCsvString_Empty(t *testing.T) {
	result := csvinternal.AnyItemsToCsvString(", ", false, false)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyItems_CsvString_Empty returns empty -- with args", actual)
}

// ── StringsToCsvString all branches ──

func Test_Cov3_StringsToCsvString_DoubleQuote(t *testing.T) {
	result := csvinternal.StringsToCsvString(", ", true, false, "a")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Strings_CsvString_DoubleQuote returns correct value -- with args", actual)
}

func Test_Cov3_StringsToCsvString_NoQuote(t *testing.T) {
	result := csvinternal.StringsToCsvString(", ", false, false, "a")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Strings_CsvString_NoQuote returns correct value -- with args", actual)
}

func Test_Cov3_StringsToCsvString_Empty(t *testing.T) {
	result := csvinternal.StringsToCsvString(", ", false, false)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Strings_CsvString_Empty returns empty -- with args", actual)
}

// ── StringersToString all branches ──

type cov3Stringer struct{ v string }

func (s cov3Stringer) String() string { return s.v }

func Test_Cov3_StringersToString_SingleQuote(t *testing.T) {
	s := cov3Stringer{v: "x"}
	result := csvinternal.StringersToString(", ", true, true, s)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToString_SingleQuote returns correct value -- with args", actual)
}

func Test_Cov3_StringersToString_DoubleQuote(t *testing.T) {
	s := cov3Stringer{v: "x"}
	result := csvinternal.StringersToString(", ", true, false, s)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToString_DoubleQuote returns correct value -- with args", actual)
}

func Test_Cov3_StringersToString_NoQuote(t *testing.T) {
	s := cov3Stringer{v: "x"}
	result := csvinternal.StringersToString(", ", false, false, s)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToString_NoQuote returns correct value -- with args", actual)
}

func Test_Cov3_StringersToString_Empty(t *testing.T) {
	result := csvinternal.StringersToString(", ", false, false)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "StringersToString_Empty returns empty -- with args", actual)
}

// ── CompileStringersToString all branches ──

func Test_Cov3_CompileStringersToString_SingleQuote(t *testing.T) {
	f := func() string { return "x" }
	result := csvinternal.CompileStringersToString(", ", true, true, f)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString_SingleQuote returns correct value -- with args", actual)
}

func Test_Cov3_CompileStringersToString_DoubleQuote(t *testing.T) {
	f := func() string { return "x" }
	result := csvinternal.CompileStringersToString(", ", true, false, f)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString_DoubleQuote returns correct value -- with args", actual)
}

func Test_Cov3_CompileStringersToString_NoQuote(t *testing.T) {
	f := func() string { return "x" }
	result := csvinternal.CompileStringersToString(", ", false, false, f)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString_NoQuote returns correct value -- with args", actual)
}

func Test_Cov3_CompileStringersToString_Empty(t *testing.T) {
	result := csvinternal.CompileStringersToString(", ", false, false)
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString_Empty returns empty -- with args", actual)
}

// ── RangeNamesWithValuesIndexes ──

func Test_Cov3_RangeNamesWithValuesIndexes(t *testing.T) {
	result := csvinternal.RangeNamesWithValuesIndexes("A", "B")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexes returns non-empty -- with args", actual)
}

func Test_Cov3_RangeNamesWithValuesIndexes_Empty(t *testing.T) {
	result := csvinternal.RangeNamesWithValuesIndexes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexes returns empty -- empty", actual)
}

func Test_Cov3_RangeNamesWithValuesIndexesCsvString_Empty(t *testing.T) {
	result := csvinternal.RangeNamesWithValuesIndexesCsvString()
	actual := args.Map{"empty": result == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsvString returns empty -- empty", actual)
}
