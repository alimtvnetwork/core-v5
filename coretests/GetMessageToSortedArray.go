package coretests

import (
	"fmt"
	"sort"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/testconsts"
)

func GetMessageToSortedArray(
	isPrint bool,
	isSort bool,
	message string,
) []string {
	if isPrint {
		fmt.Println(message)
	}

	if message == "" {
		return []string{}
	}

	items := testconsts.WhitespaceOrPipeFinderRegex.Split(
		message,
		constants.TakeAllMinusOne)

	if isSort {
		sort.Strings(items)
	}

	return items
}
