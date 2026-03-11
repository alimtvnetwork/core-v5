package issettertests

import (
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/issetter"
)

type onlySupportedErrCase struct {
	name      string
	names     []string
	expectErr bool
}

var onlySupportedErrCases = []onlySupportedErrCase{
	{
		name:      "all supported names",
		names:     []string{"Uninitialized", "True", "False", "Unset", "Set", "Wildcard"},
		expectErr: false,
	},
	{
		name:      "empty names",
		names:     []string{},
		expectErr: false,
	},
}

type conversionCase struct {
	name         string
	val          issetter.Value
	expectedByte byte
	expectedInt  int
}

var conversionCases = []conversionCase{
	{"Uninitialized", issetter.Uninitialized, 0, 0},
	{"True", issetter.True, 1, 1},
	{"False", issetter.False, 2, 2},
	{"Unset", issetter.Unset, 3, 3},
	{"Set", issetter.Set, 4, 4},
	{"Wildcard", issetter.Wildcard, 5, 5},
}

type logicalCheckCase struct {
	name     string
	val      issetter.Value
	isOn     bool
	isOff    bool
	isAsk    bool
	isAccept bool
	isReject bool
}

var logicalCheckCases = []logicalCheckCase{
	{"Uninitialized", issetter.Uninitialized, false, false, true, false, false},
	{"True", issetter.True, true, false, false, true, false},
	{"False", issetter.False, false, true, false, false, true},
	{"Set", issetter.Set, true, false, false, true, false},
	{"Unset", issetter.Unset, false, true, false, false, true},
	{"Wildcard", issetter.Wildcard, false, false, true, false, false},
}

type nameCase struct {
	name      string
	val       issetter.Value
	yesNo     string
	onOff     string
	trueFalse string
}

var nameCases = []nameCase{
	{"Uninitialized", issetter.Uninitialized, "-", "-", "-"},
	{"True", issetter.True, "Yes", "On", "True"},
	{"False", issetter.False, "No", "Off", "False"},
	{"Set", issetter.Set, "Yes", "On", "True"},
	{"Unset", issetter.Unset, "No", "Off", "False"},
	{"Wildcard", issetter.Wildcard, "*", "*", "*"},
}

type jsonCase struct {
	name string
	val  issetter.Value
}

var jsonCases = []jsonCase{
	{"Uninitialized", issetter.Uninitialized},
	{"True", issetter.True},
	{"False", issetter.False},
	{"Set", issetter.Set},
	{"Unset", issetter.Unset},
	{"Wildcard", issetter.Wildcard},
}

type toBooleanCase struct {
	name     string
	input    issetter.Value
	expected issetter.Value
}

var toBooleanCases = []toBooleanCase{
	{"Set->True", issetter.Set, issetter.True},
	{"Unset->False", issetter.Unset, issetter.False},
	{"True->True", issetter.True, issetter.True},
	{"False->False", issetter.False, issetter.False},
	{"Wildcard->Wildcard", issetter.Wildcard, issetter.Wildcard},
}

var toSetUnsetCases = []toBooleanCase{
	{"True->Set", issetter.True, issetter.Set},
	{"False->Unset", issetter.False, issetter.Unset},
	{"Set->Set", issetter.Set, issetter.Set},
	{"Unset->Unset", issetter.Unset, issetter.Unset},
}

type wildcardApplyCase struct {
	name     string
	val      issetter.Value
	input    bool
	expected bool
}

var wildcardApplyCases = []wildcardApplyCase{
	{"Wildcard returns input true", issetter.Wildcard, true, true},
	{"Wildcard returns input false", issetter.Wildcard, false, false},
	{"True ignores input", issetter.True, false, true},
	{"False ignores input", issetter.False, true, false},
	{"Uninitialized returns input", issetter.Uninitialized, true, true},
}

type compareResultCase struct {
	name     string
	val      issetter.Value
	n        byte
	compare  corecomparator.Compare
	expected bool
}

var compareResultCases = []compareResultCase{
	{"Equal match", issetter.True, 1, corecomparator.Equal, true},
	{"Equal no match", issetter.True, 2, corecomparator.Equal, false},
	{"LeftGreater", issetter.False, 1, corecomparator.LeftGreater, true},
	{"LeftGreaterEqual", issetter.True, 1, corecomparator.LeftGreaterEqual, true},
	{"LeftLess", issetter.True, 2, corecomparator.LeftLess, true},
	{"LeftLessEqual", issetter.True, 1, corecomparator.LeftLessEqual, true},
	{"NotEqual", issetter.True, 2, corecomparator.NotEqual, true},
}

type newFromStringCase struct {
	name      string
	input     string
	expected  issetter.Value
	expectErr bool
}

var newFromStringCases = []newFromStringCase{
	{"true string", "true", issetter.True, false},
	{"false string", "false", issetter.False, false},
	{"yes string", "yes", issetter.True, false},
	{"no string", "no", issetter.False, false},
	{"set string", "set", issetter.Set, false},
	{"unset string", "unset", issetter.Unset, false},
	{"wildcard string", "*", issetter.Wildcard, false},
	{"empty string", "", issetter.Uninitialized, false},
	{"invalid string", "INVALID_VALUE_XYZ", issetter.Uninitialized, true},
}
