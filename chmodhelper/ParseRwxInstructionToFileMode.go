package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/msgtype"
)

func ParseRwxOwnerGroupOtherToFileModeMust(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) os.FileMode {
	fileMode, err := ParseRwxOwnerGroupOtherToFileMode(
		rwxOwnerGroupOther)

	if err != nil {
		panic(err)
	}

	return fileMode
}

func ParseRwxOwnerGroupOtherToFileMode(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) (os.FileMode, error) {
	varWrapper, err := ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		rwxOwnerGroupOther)

	if err != nil {
		return 0, msgtype.MeaningFulErrorWithData(
			msgtype.FailedToConvert,
			"ParseRwxOwnerGroupOtherToFileMode",
			err,
			rwxOwnerGroupOther)
	}

	return varWrapper.ToCompileFixedPtr().ToFileMode(), nil
}
