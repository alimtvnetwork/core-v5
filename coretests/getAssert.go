package coretests

import (
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/msgcreator"
)

type getAssert struct {
	SimpleTestCaseWrapper getAssertSimpleTestCaseWrapper
}

// Quick
//
// Gives generic and consistent
// test message using msgformats.QuickIndexInputActualExpectedMessageFormat
func (it getAssert) Quick(
	when,
	actual,
	expected interface{},
	counter int,
) string {
	return msgcreator.Assert.Quick(
		when,
		actual,
		expected,
		counter,
	)
}

func (it getAssert) SortedMessage(
	isPrint bool,
	message,
	joiner string,
) string {
	return msgcreator.Assert.SortedMessage(
		isPrint,
		message,
		joiner,
	)
}

func (it getAssert) SortedArray(
	isPrint bool,
	isSort bool,
	message string,
) []string {
	return msgcreator.Assert.SortedArray(
		isPrint,
		isSort,
		message,
	)
}

// SortedArrayNoPrint
//
// isPrint: false, isSort: true
func (it getAssert) SortedArrayNoPrint(
	message string,
) []string {
	return msgcreator.Assert.SortedArrayNoPrint(
		message,
	)
}

func (it getAssert) ToStrings(
	any interface{},
) []string {
	return convertinteranl.AnyTo.Strings(any)
}

func (it getAssert) ToStringsWithTab(
	spacePrefixCount int,
	any interface{},
) []string {
	return msgcreator.Assert.ToStringsWithTab(
		spacePrefixCount,
		any,
	)
}

func (it getAssert) StringsToSpaceString(
	spaceCount int,
	lines ...string,
) []string {
	return msgcreator.Assert.StringsToSpaceString(
		spaceCount,
		lines...,
	)
}

func (it getAssert) StringsToSpaceStringUsingFunc(
	spaceCount int,
	toStringFunc func(i int, spacePrefix, line string) string,
	lines ...string,
) []string {
	return msgcreator.Assert.StringsToSpaceStringUsingFunc(
		spaceCount,
		toStringFunc,
		lines...,
	)
}
