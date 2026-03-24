package pathinternal

import (
	"fmt"
	"os"
)

func RemoveDirIf(isRemoveAllDirBeforeCreate bool, dir string, funcName string) error {
	var removeErr error

	if isRemoveAllDirBeforeCreate && IsPathExists(dir) {
		// Ensure directory is writable before removal to avoid permission denied errors
		_ = os.Chmod(dir, 0777)
		removeErr = os.RemoveAll(dir)
	}

	if removeErr != nil {
		return pathMeaningfulError(
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

func pathMeaningfulError(
	funcName string,
	err error,
	location string,
) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error() +
		", location: [" + location + "]"

	return fmt.Errorf(
		"%s - %s %s, location: [%s]",
		funcName,
		errMsg,
		err.Error(),
		location,
	)
}
