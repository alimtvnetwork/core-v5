package coreappend

import "strings"

func PrependAnyItemsToStringSkipOnNil(
	joiner string,
	prependItem any,
	anyItems ...any,
) string {
	slice := PrependAppendAnyItemsToStringsSkipOnNil(
		prependItem,
		nil,
		anyItems...)

	return strings.Join(slice, joiner)
}
