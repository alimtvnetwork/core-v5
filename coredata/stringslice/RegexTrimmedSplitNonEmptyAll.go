package stringslice

import (
	"regexp"

	"github.com/alimtvnetwork/core/constants"
)

func RegexTrimmedSplitNonEmptyAll(
	regexp *regexp.Regexp,
	content string,
) []string {
	items := regexp.Split(
		content,
		constants.TakeAllMinusOne)

	return TrimmedEachWords(items)
}
