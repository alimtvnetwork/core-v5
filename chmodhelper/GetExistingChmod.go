package chmodhelper

import (
	"os"

	"github.com/alimtvnetwork/core/errcore"
)

func GetExistingChmod(filePath string) (os.FileMode, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil || fileInfo == nil {
		return 0, errcore.
			PathErrorType.
			Error(err.Error(), ", file:"+filePath)
	}

	return fileInfo.Mode(), err
}
