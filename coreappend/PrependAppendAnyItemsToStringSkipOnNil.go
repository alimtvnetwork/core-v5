package coreappend

import "strings"

func PrependAppendAnyItemsToStringSkipOnNil(
	joiner string,
	prependItem, appendItem any,
	anyItems ...any,
) string {
	slice := PrependAppendAnyItemsToStringsSkipOnNil(
		prependItem,
		appendItem,
		anyItems...)

	return strings.Join(slice, joiner)
}
