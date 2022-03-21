package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/errcore"
)

func GetExistingChmodRwxWrapper(
	location string,
) (RwxWrapper, error) {
	fileInfo, err := os.Stat(location)

	if err != nil || fileInfo == nil {
		return RwxWrapper{}, errcore.
			PathErrorType.
			Error(err.Error(), ", file:"+location)
	}

	return New.RwxWrapper.UsingFileMode(fileInfo.Mode()), err
}
