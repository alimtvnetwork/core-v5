package coretests

import "strings"

func GetMessageToSorted(
	isPrint bool,
	message,
	joiner string,
) string {
	whitespaceRemovedSplits := GetMessageToSortedArray(
		isPrint,
		true,
		message)

	return strings.Join(whitespaceRemovedSplits, joiner)
}
