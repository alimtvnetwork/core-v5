package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/msgtype"
)

func GetExistingChmodWrapperPtr(
	filePath string,
) (*Wrapper, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return nil, msgtype.PathErrorMessage.
			Error(err.Error(), ", file:"+filePath)
	}

	return NewUsingFileModePtr(fileInfo.Mode()), err
}
