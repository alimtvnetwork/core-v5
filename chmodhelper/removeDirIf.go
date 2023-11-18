package chmodhelper

import (
	"os"

	"gitlab.com/auk-go/core/errcore"
)

func removeDirIf(isRemoveAllDirBeforeCreate bool, dir string, funcName string) error {
	var removeErr error

	if isRemoveAllDirBeforeCreate && IsPathExists(dir) {
		removeErr = os.RemoveAll(dir)
	}

	if removeErr != nil {
		return errcore.PathMeaningfulError(
			errcore.PathCreateFailedType,
			funcName,
			removeErr,
			dir,
		)
	}

	return nil
}
