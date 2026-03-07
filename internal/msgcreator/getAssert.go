package msgcreator

import (
	"fmt"
	"log/slog"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/convertinternal"
	"gitlab.com/auk-go/core/internal/msgformats"
)

type getAssert struct{}

// Quick
//
// Gives generic and consistent
// test message using msgformats.QuickIndexInputActualExpectedMessageFormat
func (it getAssert) Quick(
	when any,
	actual any,
	expected any,
	counter int,
) string {
	return fmt.Sprintf(
		msgformats.QuickIndexInputActualExpectedMessageFormat,
		counter,
		convertinternal.AnyTo.SmartString(when),
		convertinternal.AnyTo.SmartString(actual),
		convertinternal.AnyTo.SmartString(expected),
	)
}

// IsEqualMessage
//
// Gives generic and consistent
// test message using msgformats.IsEqualMessageFormat
func (it getAssert) IsEqualMessage(
	when any,
	actual any,
	expected any,
) string {
	return fmt.Sprintf(
		msgformats.IsEqualMessageFormat,
		convertinternal.AnyTo.SmartString(when),
		convertinternal.AnyTo.SmartString(actual),
		convertinternal.AnyTo.SmartString(expected),
	)
}

// IsNotEqualMessage
//
// Gives generic and consistent
// test message using msgformats.IsNotEqualMessageFormat
func (it getAssert) IsNotEqualMessage(
	when any,
	actual any,
	expected any,
) string {
	return fmt.Sprintf(
		msgformats.IsNotEqualMessageFormat,
		convertinternal.AnyTo.SmartString(when),
		convertinternal.AnyTo.SmartString(actual),
		convertinternal.AnyTo.SmartString(expected),
	)
}

// IsTrueMessage
//
// Gives generic and consistent
// test message using msgformats.IsTrueMessageFormat
func (it getAssert) IsTrueMessage(
	when any,
	actual any,
) string {
	return fmt.Sprintf(
		msgformats.IsTrueMessageFormat,
		convertinternal.AnyTo.SmartString(when),
		convertinternal.AnyTo.SmartString(actual),
	)
}

// IsFalseMessage
//
// Gives generic and consistent
// test message using msgformats.IsFalseMessageFormat
func (it getAssert) IsFalseMessage(
	when any,
	actual any,
) string {
	return fmt.Sprintf(
		msgformats.IsFalseMessageFormat,
		convertinternal.AnyTo.SmartString(when),
		convertinternal.AnyTo.SmartString(actual),
	)
}

// IsNilMessage
//
// Gives generic and consistent
// test message using msgformats.IsNilMessageFormat
func (it getAssert) IsNilMessage(
	when any,
	actual any,
) string {
	return fmt.Sprintf(
		msgformats.IsNilMessageFormat,
		convertinternal.AnyTo.SmartString(when),
		convertinternal.AnyTo.SmartString(actual),
	)
}

// IsNotNilMessage
//
// Gives generic and consistent
// test message using msgformats.IsNotNilMessageFormat
func (it getAssert) IsNotNilMessage(
	when any,
	actual any,
) string {
	return fmt.Sprintf(
		msgformats.IsNotNilMessageFormat,
		convertinternal.AnyTo.SmartString(when),
		convertinternal.AnyTo.SmartString(actual),
	)
}

// ShouldBeMessage
//
// Gives generic and consistent
// test message using msgformats.ShouldBeMessageFormat
func (it getAssert) ShouldBeMessage(
	title string,
	actual any,
	expected any,
) string {
	return fmt.Sprintf(
		msgformats.ShouldBeMessageFormat,
		title,
		convertinternal.AnyTo.SmartString(actual),
		convertinternal.AnyTo.SmartString(expected),
	)
}

// ShouldNotBeMessage
//
// Gives generic and consistent
// test message using msgformats.ShouldNotBeMessageFormat
func (it getAssert) ShouldNotBeMessage(
	title string,
	actual any,
	expected any,
) string {
	return fmt.Sprintf(
		msgformats.ShouldNotBeMessageFormat,
		title,
		convertinternal.AnyTo.SmartString(actual),
		convertinternal.AnyTo.SmartString(expected),
	)
}

func (it getAssert) SortedMessage(
	isPrint bool,
	message string,
	joiner string,
) string {
	whitespaceRemovedSplits := it.SortedArray(
		isPrint,
		true,
		message,
	)

	return strings.Join(whitespaceRemovedSplits, joiner)
}

func (it getAssert) SortedArray(
	isPrint bool,
	isSort bool,
	message string,
) []string {
	if isPrint {
		slog.Debug("sorted array", "message", message)
	}

	return SplitByEachWordTrimmedNoSpace(
		message,
		isSort,
	)
}

// SortedArrayNoPrint
//
// isPrint: false, isSort: true
func (it getAssert) SortedArrayNoPrint(
	message string,
) []string {
	return it.SortedArray(
		false,
		true, message,
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
//
// See also convertinternal.AnyTo.Strings
func (it getAssert) ToStrings(
	anyItem any,
) []string {
	return convertinternal.AnyTo.Strings(anyItem)
}

// ToStringsWithSpace
//
//	This function will display complex objects to simpler form
//	for the integration testing validation and expectations.
//	Usages a space prefix for each line.
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
//
// See also convertinternal.AnyTo.Strings
func (it getAssert) ToStringsWithSpace(
	spacePrefixCount int,
	anyItem any,
) []string {
	lines := convertinternal.AnyTo.Strings(anyItem)

	return it.StringsToWithSpaceLines(
		spacePrefixCount,
		lines...,
	)
}

func (it getAssert) ToStringsWithSpaceDefault(
	anyItem any,
) []string {
	return it.ToStringsWithSpace(2, anyItem)
}

func (it getAssert) ToStringWithSpace(
	spacePrefixCount int,
	anyItem any,
) string {
	lines := convertinternal.AnyTo.Strings(anyItem)

	withSpace := it.StringsToWithSpaceLines(
		spacePrefixCount,
		lines...,
	)

	return strings.Join(withSpace, constants.NewLineUnix)
}

func (it getAssert) StringsToWithSpaceLines(
	spaceCount int,
	lines ...string,
) []string {
	if len(lines) == 0 {
		return []string{}
	}

	newLines := make([]string, len(lines))
	prefix := strings.Repeat(
		" ",
		spaceCount,
	)

	for i, line := range lines {
		newLines[i] = fmt.Sprintf(
			"%s%s",
			prefix,
			line,
		)
	}

	return newLines
}

func (it getAssert) StringsToSpaceStringUsingFunc(
	spaceCount int,
	toStringFunc func(i int, spacePrefix, line string) string,
	lines ...string,
) []string {
	if len(lines) == 0 {
		return []string{}
	}

	newLines := make([]string, len(lines))
	prefix := strings.Repeat(
		" ",
		spaceCount,
	)

	for i, line := range lines {
		newLines[i] = toStringFunc(
			i,
			prefix,
			line,
		)
	}

	return newLines
}
