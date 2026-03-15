package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Cover receiver method RangeNamesCsv (vs package-level func) ──

func Test_Cov7_Compare_ReceiverRangeNamesCsv(t *testing.T) {
	csv := corecomparator.Equal.RangeNamesCsv()
	actual := args.Map{"notEmpty": csv != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Compare.RangeNamesCsv receiver -- not empty", actual)
}

// ── IsCompareEqualLogically: branch where expectedCompare is Equal but it is not ──

func Test_Cov7_IsCompareEqualLogically_ExpectedEqual_ItNotEqual(t *testing.T) {
	actual := args.Map{
		"result": corecomparator.LeftGreater.IsCompareEqualLogically(corecomparator.Equal),
	}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsCompareEqualLogically expected=Equal it=LeftGreater -- false", actual)
}

// ── OnlySupportedDirectErr error path ──

func Test_Cov7_OnlySupportedDirectErr_NotMatching(t *testing.T) {
	err := corecomparator.Inconclusive.OnlySupportedDirectErr(corecomparator.Equal, corecomparator.LeftGreater)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedDirectErr no match -- error", actual)
}

// ── IsLeftLessEqualLogically on non-less values ──

func Test_Cov7_IsLeftLessEqualLogically_Greater(t *testing.T) {
	actual := args.Map{
		"greater": corecomparator.LeftGreater.IsLeftLessEqualLogically(),
		"notEq":   corecomparator.NotEqual.IsLeftLessEqualLogically(),
	}
	expected := args.Map{"greater": false, "notEq": false}
	expected.ShouldBeEqual(t, 0, "IsLeftLessEqualLogically non-less values -- false", actual)
}

// ── IsLeftGreaterEqualLogically on non-greater values ──

func Test_Cov7_IsLeftGreaterEqualLogically_Less(t *testing.T) {
	actual := args.Map{
		"less":  corecomparator.LeftLess.IsLeftGreaterEqualLogically(),
		"notEq": corecomparator.NotEqual.IsLeftGreaterEqualLogically(),
	}
	expected := args.Map{"less": false, "notEq": false}
	expected.ShouldBeEqual(t, 0, "IsLeftGreaterEqualLogically non-greater values -- false", actual)
}

// ── MinLength equal values ──

func Test_Cov7_MinLength_Equal(t *testing.T) {
	actual := args.Map{"result": corecomparator.MinLength(3, 3)}
	expected := args.Map{"result": 3}
	expected.ShouldBeEqual(t, 0, "MinLength equal -- returns same", actual)
}
