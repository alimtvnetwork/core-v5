package pathinternal

import (
	"os"

	"gitlab.com/auk-go/core/errcore"
)

func RemoveDirIf(isRemoveAllDirBeforeCreate bool, dir string, funcName string) error {
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

func RemoveDirIfMust(isRemoveAllDirBeforeCreate bool, dir string, funcName string) {
	removeErr := RemoveDirIf(
		isRemoveAllDirBeforeCreate,
		dir,
		funcName,
	)

	if removeErr != nil {
		panic(removeErr)
	}
}

func RemoveDirMust(dir string, funcName string) {
	removeErr := RemoveDirIf(
		true,
		dir,
		funcName,
	)

	if removeErr != nil {
		panic(removeErr)
	}
}

func RemoveDirMustSimple(dir string) {
	removeErr := RemoveDirIf(
		true,
		dir,
		"",
	)

	if removeErr != nil {
		panic(removeErr)
	}
}
