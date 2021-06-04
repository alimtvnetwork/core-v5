package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/msgtype"
)

func GetExistingChmod(filePath string) (os.FileMode, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return 0, msgtype.
			PathErrorMessage.
			Error(err.Error(), ", file:"+filePath)
	}

	return fileInfo.Mode(), err
}
