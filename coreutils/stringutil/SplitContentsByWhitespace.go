package stringutil

import (
	"sort"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/regexnew"
)

func SplitContentsByWhitespace(
	input string,
	isTrimEachLine,
	isNonEmptyWhitespace,
	isSort bool,
) []string {
	compiledStringSplits := regexnew.WhitespaceFinderRegex.
		Split(
			input,
			constants.TakeAllMinusOne)

	if isNonEmptyWhitespace && isTrimEachLine {
		compiledStringSplits = stringslice.NonWhitespaceTrimSlice(
			compiledStringSplits)
	} else if isNonEmptyWhitespace && !isTrimEachLine {
		compiledStringSplits = stringslice.NonWhitespaceSlice(
			compiledStringSplits)
	}

	if isSort {
		sort.Strings(compiledStringSplits)
	}

	return compiledStringSplits
}
