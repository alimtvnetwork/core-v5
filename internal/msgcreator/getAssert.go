package msgcreator

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/msgformats"
)

type getAssert struct{}

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
	return fmt.Sprintf(
		msgformats.QuickIndexInputActualExpectedMessageFormat,
		counter,
		converters.AnyToSmartString(when),
		converters.AnyToSmartString(actual),
		converters.AnyToSmartString(expected),
	)
}

func (it getAssert) SortedMessage(
	isPrint bool,
	message,
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
		fmt.Println(message)
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

func (it getAssert) ToStrings(
	any interface{},
) []string {
	return convertinteranl.AnyTo.Strings(any)
}

func (it getAssert) ToStringsWithTab(
	spacePrefixCount int,
	any interface{},
) []string {
	lines := convertinteranl.AnyTo.Strings(any)

	return it.StringsToSpaceString(
		spacePrefixCount,
		lines...,
	)
}

func (it getAssert) StringsToSpaceString(
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
