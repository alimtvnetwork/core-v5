package msgtype

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func ErrPrintWithTestIndex(
	caseIndex int,
	err error,
) {
	if err != nil {
		fmt.Print(
			"Case Index: ",
			caseIndex,
			constants.CommaSpace,
			constants.NewLineUnix,
			err)
	}
}
