package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/msgtype"
)

func GetExistingChmodRwxWrapperPtr(
	filePath string,
) (*RwxWrapper, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return nil, msgtype.PathErrorMessage.
			Error(err.Error(), ", file:"+filePath)
	}

	return NewUsingFileModePtr(fileInfo.Mode()), err
}
