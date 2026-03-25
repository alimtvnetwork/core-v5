package chmodhelper

import (
	"os"

	"github.com/alimtvnetwork/core/errcore"
)

func GetExistingChmodRwxWrapperPtr(
	location string,
) (*RwxWrapper, error) {
	fileInfo, err := os.Stat(location)

	if err != nil {
		return nil, errcore.PathErrorType.
			Error(err.Error(), ", file:"+location)
	}

	return New.RwxWrapper.UsingFileModePtr(fileInfo.Mode()), err
}
