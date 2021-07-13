package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/msgtype"
)

func ParseRwxOwnerGroupOtherToFileMode(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) (os.FileMode, error) {
	varWrapper, err := ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		rwxOwnerGroupOther)

	if err != nil {
		return 0, msgtype.MeaningfulErrorWithData(
			msgtype.FailedToConvert,
			"ParseRwxOwnerGroupOtherToFileMode",
			err,
			rwxOwnerGroupOther)
	}

	return varWrapper.ToCompileFixedPtr().ToFileMode(), nil
}
