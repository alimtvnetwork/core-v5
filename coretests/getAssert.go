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

// ToStrings
//
//	This function will display complex objects to simpler form
//	for the integration testing validation and expectations.
//
// # Steps:
//  01. string to []string
//  02. []string to as is.
//  03. []interface{} to []string
//  04. map[string]interface{} (fmt - "%s : SmartJson(%s)") to []string
//  05. map[interface{}]interface{} (fmt - SmartJson("%s) : SmartJson(%s)") to []string
//  06. map[string]string (fmt - %s : %s)") to []string
//  07. map[string]int (fmt - %s : %d)") to []string
//  08. map[int]string (fmt - %d : %s)") to []string
//  09. int to []string
//  10. byte to []string
//  11. bool to []string
//  12. any to PrettyJSON
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
	converterFunc ToLineConverterFunc,
	lines ...string,
) []string {
	return msgcreator.Assert.StringsToSpaceStringUsingFunc(
		spaceCount,
		converterFunc,
		lines...,
	)
}
