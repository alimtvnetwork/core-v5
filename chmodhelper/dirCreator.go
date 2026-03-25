package chmodhelper

import (
	"os"
)

type dirCreator struct{}

// If
//
// if isCreate + is missing director then only create dir.
func (it dirCreator) If(
	isCreate bool,
	chmod os.FileMode,
	dirPath string,
) error {
	if !isCreate {
		return nil
	}

	return it.IfMissing(
		chmod,
		dirPath,
	)
}

func (it dirCreator) IfMissingLock(
	applyChmod os.FileMode,
	dirPath string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.IfMissing(
		applyChmod,
		dirPath,
	)
}

// IfMissing
//
// Only create dir if missing.
func (it dirCreator) IfMissing(
	applyChmod os.FileMode,
	dirPath string,
) error {
	if IsPathExists(dirPath) {
		return nil
	}

	// Use permissive mode for MkdirAll to allow nested dir creation,
	// then apply the requested chmod afterward.
	err := os.MkdirAll(
		dirPath,
		dirDefaultChmod,
	)

	if err != nil {
		// has err
		return newError.pathErrorWithDirValidate(
			"dir creation failed",
			applyChmod,
			dirPath,
			err,
		)
	}

	// Apply the requested chmod if it differs from the default
	if applyChmod != dirDefaultChmod {
		chmodErr := os.Chmod(dirPath, applyChmod)
		if chmodErr != nil {
			return newError.chmodApplyFailed(
				applyChmod,
				dirPath,
				chmodErr,
			)
		}
	}

	return nil
}

// ByChecking
//
// Check only if the dir is missing and apply chmod.
func (it dirCreator) ByChecking(
	applyChmod os.FileMode,
	dirPath string,
) error {
	isExists := IsPathExists(dirPath)
	isDir := IsDirectory(dirPath)

	if isExists && isDir {
		return os.Chmod(dirPath, applyChmod)
	}

	if isExists && !isDir {
		return newError.notDirError(dirPath)
	}

	// Use permissive mode for MkdirAll to allow nested dir creation,
	// then apply the requested chmod afterward.
	err := os.MkdirAll(
		dirPath,
		dirDefaultChmod,
	)

	if err != nil {
		return newError.pathErrorWithDirValidate(
			"dir creation failed",
			applyChmod,
			dirPath,
			err,
		)
	}

	chmodErr := os.Chmod(
		dirPath,
		applyChmod,
	)

	if chmodErr != nil {
		return newError.chmodApplyFailed(
			applyChmod,
			dirPath,
			chmodErr,
		)
	}

	return nil
}

func (it dirCreator) DefaultLock(
	applyChmod os.FileMode,
	dirPath string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Default(
		applyChmod,
		dirPath,
	)
}

// Default
//
// Direct try to create without checking if directory exists.
func (it dirCreator) Default(
	applyChmod os.FileMode,
	dirPath string,
) error {
	err := os.MkdirAll(
		dirPath,
		applyChmod,
	)

	if err == nil {
		return nil
	}

	// has err
	return newError.dirError(dirPath, err)
}

func (it dirCreator) DirectLock(
	dirPath string,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Direct(
		dirPath,
	)
}

// Direct
//
// Dir default chmod 0755
func (it dirCreator) Direct(
	dirPath string,
) error {
	err := os.MkdirAll(
		dirPath,
		dirDefaultChmod,
	)

	if err == nil {
		return nil
	}

	return newError.dirError(dirPath, err)
}
