package pathinternal

import (
	"fmt"
	"os"
	"path/filepath"
)

func RemoveDirIf(isRemoveAllDirBeforeCreate bool, dir string, funcName string) error {
	var removeErr error

	if isRemoveAllDirBeforeCreate && IsPathExists(dir) {
		// Recursively chmod to 0777 before removal to handle nested restrictive permissions
		chmodRecursive(dir)
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

// chmodRecursive walks the directory tree and sets 0777 on all entries
// so that os.RemoveAll can succeed even on restrictively-permissioned trees.
func chmodRecursive(dir string) {
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// If we can't stat, try to chmod the path anyway
			_ = os.Chmod(path, 0777)
			return nil
		}
		_ = os.Chmod(path, 0777)
		return nil
	})
}
