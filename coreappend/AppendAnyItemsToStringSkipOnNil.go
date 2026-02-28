package coreappend

import "strings"

func AppendAnyItemsToStringSkipOnNil(
	joiner string,
	appendItem any,
	anyItems ...any,
) string {
	slice := PrependAppendAnyItemsToStringsSkipOnNil(
		nil,
		appendItem,
		anyItems...)

	return strings.Join(slice, joiner)
}
