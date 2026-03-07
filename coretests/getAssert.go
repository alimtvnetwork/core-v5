package coretests

import (
	"fmt"
	"log/slog"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/convertinternal"
	"gitlab.com/auk-go/core/internal/msgcreator"
	"gitlab.com/auk-go/core/internal/msgformats"
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
	expected any,
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
//  03. []any to []string
//  04. map[string]any (fmt - "%s : SmartJson(%s)") to []string
//  05. map[any]any (fmt - SmartJson("%s) : SmartJson(%s)") to []string
//  06. map[string]string (fmt - %s : %s)") to []string
//  07. map[string]int (fmt - %s : %d)") to []string
//  08. map[int]string (fmt - %d : %s)") to []string
//  09. int to []string
//  10. byte to []string
//  11. bool to []string
//  12. any to PrettyJSON
func (it getAssert) ToStrings(
	anyItem any,
) []string {
	return convertinternal.AnyTo.Strings(anyItem)
}

func (it getAssert) ToStringsWithSpace(
	spaceCount int,
	anyItem any,
) []string {
	return msgcreator.Assert.ToStringsWithSpace(
		spaceCount,
		anyItem,
	)
}

func (it getAssert) ErrorToLinesWithSpaces(
	spaceCount int,
	err error,
) []string {
	if err == nil {
		return []string{}
	}

	errStr := errcore.ToString(err)

	return it.ToStringsWithSpace(spaceCount, errStr)
}

func (it getAssert) ErrorToLinesWithSpacesDefault(
	err error,
) []string {
	return it.ErrorToLinesWithSpaces(2, err)
}

func (it getAssert) StringsToSpaceString(
	spaceCount int,
	lines ...string,
) []string {
	return msgcreator.Assert.StringsToWithSpaceLines(
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

// ToQuoteLines
//
// Converts from below lines to
//
//	line 1,
//	line 2,
//	line 3,
//
// Converts a strings lines to
//
//	{spaces} "line 1",
//	{spaces} "line 2",
//	{spaces} "line 3",
func (it getAssert) ToQuoteLines(
	spaceCount int,
	lines []string,
) []string {
	return errcore.LinesToDoubleQuoteLinesWithTabs(
		spaceCount,
		lines,
	)
}

// AnyToDoubleQuoteLines
//
// Converts from below lines or line to
//
//	line 1,
//	line 2,
//	line 3,
//
// Or, converts from below line to lines if string or converts it to line
//
//	"line 1,\nline 2,\nline 3"
//
// Converts a strings lines to
//
//	{spaces} "line 1",
//	{spaces} "line 2",
//	{spaces} "line 3",
func (it getAssert) AnyToDoubleQuoteLines(
	spaceCount int,
	anyItem any,
) []string {
	lines := convertinternal.AnyTo.Strings(anyItem)

	return it.ToQuoteLines(
		spaceCount,
		lines,
	)
}

// ConvertLinesToDoubleQuoteThenString
//
// Convert lines to double quote wrap and then adds a space prefix
func (it getAssert) ConvertLinesToDoubleQuoteThenString(
	spaceCount int,
	lines []string,
) string {
	finalLines := it.ToQuoteLines(
		spaceCount,
		lines,
	)

	return strings.Join(finalLines, constants.NewLineUnix)
}

// AnyToStringDoubleQuoteLine
//
// Convert Any to lines to double quote wrap
// and then adds a space prefix (using ConvertLinesToDoubleQuoteThenString)
func (it getAssert) AnyToStringDoubleQuoteLine(
	spaceCount int,
	anyItem any,
) string {
	lines := convertinternal.AnyTo.Strings(anyItem)

	return it.ConvertLinesToDoubleQuoteThenString(spaceCount, lines)
}

func (it getAssert) ToString(
	anyItem any,
) string {
	lines := convertinternal.AnyTo.Strings(anyItem)

	return strings.Join(lines, constants.NewLineUnix)
}

func LogOnFail(
	isPass bool,
	expected, actual any,
) {
	if isPass {
		return
	}

	logMessage := fmt.Sprintf(msgformats.LogFormat, expected, actual)
	slog.Warn("assertion failed", "message", logMessage)
}

func ToStringValues(anyItem any) string {
	if anyItem == nil {
		return constants.NilAngelBracket
	}

	return fmt.Sprintf(constants.SprintValueFormat, anyItem)
}

func ToStringNameValues(anyItem any) string {
	if anyItem == nil {
		return constants.NilAngelBracket
	}

	return fmt.Sprintf(
		constants.SprintFullPropertyNameValueFormat,
		anyItem,
	)
}
