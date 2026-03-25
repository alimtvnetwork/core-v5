package stringutil

import (
	"sort"
	"strings"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coredata/stringslice"
)

func SplitContentsByWhitespaceConditions(
	input string,
	isTrimEachLine,
	isNonEmptyWhitespace,
	isSort bool,
	isUnique bool,
	isLowerCase bool,
) []string {
	if isLowerCase || isUnique {
		input = strings.ToLower(input)
	}

	compiledStringSplits := strings.Fields(input)

	if isNonEmptyWhitespace && isTrimEachLine {
		compiledStringSplits = stringslice.TrimmedEachWords(
			compiledStringSplits,
		)
	} else if isNonEmptyWhitespace && !isTrimEachLine {
		compiledStringSplits = stringslice.NonWhitespace(
			compiledStringSplits,
		)
	}

	if isUnique {
		hashset := corestr.New.Hashset.Strings(compiledStringSplits)
		compiledStringSplits = hashset.List()
	}

	if isSort {
		sort.Strings(compiledStringSplits)
	}

	return compiledStringSplits
}
