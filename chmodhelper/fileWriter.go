package chmodhelper

import (
	"errors"
	"os"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/internal/osconstsinternal"
	"github.com/alimtvnetwork/core/internal/pathinternal"
)

type fileWriter struct {
	Bytes  fileBytesWriter
	String fileStringWriter
	Any    anyItemWriter // writes any item using JSON
}

// AllLock
//
//	Writes contents to file system.
//
// parentDirPath:
//   - is a full path to the parent dir for checking
//     if parent dir exist if not then created
//
// writingFilePath:
//   - is a full path to the actual file where to write contents
func (it fileWriter) AllLock(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	isRemoveBeforeWrite,
	isApplyChmodMust,
	isApplyChmodOnMismatch bool, // only apply for file, dir will not be applied if already created
	isCreateDirOnRequired bool,
	parentDirPath string,
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.All(
		chmodDir,
		chmodFile,
		isRemoveBeforeWrite,
		isApplyChmodMust,
		isApplyChmodOnMismatch,
		isCreateDirOnRequired,
		parentDirPath,
		writingFilePath,
		contentsBytes,
	)
}

// All
//
//	Writes contents to file system.
//
// parentDirPath:
//   - is a full path to the parent dir for checking
//     if parent dir exist if not then created
//
// writingFilePath:
//   - is a full path to the actual file where to write contents
//
// Warning:
//   - Chmod will NOT be applied to dir if already created.
//     This may harm other files.
func (it fileWriter) All(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	isRemoveBeforeWrite,
	isApplyChmodMust,
	isApplyChmodOnMismatch bool, // only apply for file, dir will not be applied if already created
	isCreateDirOnRequired bool,
	parentDirPath string,
	writingFilePath string,
	contentsBytes []byte,
) error {
	// Create dir with permissive mode first so file write succeeds
	// even when chmodDir is restrictive (e.g., 0200 lacks execute).
	dirErr := internalDirCreator.If(
		isCreateDirOnRequired,
		dirDefaultChmod,
		parentDirPath,
	)

	if dirErr != nil {
		return dirErr
	}

	cleanUpErr := it.RemoveIf(
		isRemoveBeforeWrite,
		writingFilePath,
	)

	if cleanUpErr != nil {
		return cleanUpErr
	}

	err := os.WriteFile(
		writingFilePath,
		contentsBytes,
		chmodFile,
	)

	if err != nil {
		return errors.New(
			"writing failed " +
				"filePath : " + writingFilePath +
				", contents : " + corejson.BytesToString(contentsBytes) +
				", chmod file :" + chmodFile.String() + ", " +
				", chmod dir :" + chmodDir.String() + ", " +
				err.Error(),
		)
	}

	isNotApplyChmod := !isApplyChmodMust

	if isNotApplyChmod || osconstsinternal.IsWindows {
		// Apply dir chmod even when file chmod is not required
		return it.applyDirChmod(isCreateDirOnRequired, chmodDir, parentDirPath)
	}

	// unix, must chmod
	if isApplyChmodOnMismatch && ChmodVerify.IsEqual(writingFilePath, chmodFile) {
		return it.applyDirChmod(isCreateDirOnRequired, chmodDir, parentDirPath)
	}

	// not equal or apply anyway
	fileChmodErr := ChmodApply.Default(chmodFile, writingFilePath)
	if fileChmodErr != nil {
		return fileChmodErr
	}

	return it.applyDirChmod(isCreateDirOnRequired, chmodDir, parentDirPath)
}

// applyDirChmod applies the requested dir chmod after file operations are done.
func (it fileWriter) applyDirChmod(
	isCreateDirOnRequired bool,
	chmodDir os.FileMode,
	parentDirPath string,
) error {
	if !isCreateDirOnRequired {
		return nil
	}

	if chmodDir == dirDefaultChmod {
		return nil
	}

	chmodErr := os.Chmod(parentDirPath, chmodDir)
	if chmodErr != nil {
		return newError.chmodApplyFailed(
			chmodDir,
			parentDirPath,
			chmodErr,
		)
	}

	return nil
}

func (it fileWriter) Remove(removePath string) error {
	err := os.RemoveAll(removePath)

	if err != nil {
		return errors.New(
			"clean up or remove failed " +
				"filePath : " + removePath +
				", " +
				err.Error(),
		)
	}

	return nil
}

func (it fileWriter) RemoveIf(
	isRemove bool,
	removePath string,
) error {
	if !isRemove {
		return nil
	}

	if !pathinternal.IsPathExists(removePath) {
		return nil
	}

	return it.Remove(removePath)
}

func (it fileWriter) ParentDir(filePath string) string {
	return pathinternal.ParentDir(filePath)
}

func (it fileWriter) Chmod(
	isRemoveBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	filePath string,
	contentsBytes []byte,
) error {
	parentDir := it.ParentDir(filePath)

	return it.All(
		chmodDir,
		chmodFile,
		isRemoveBeforeWrite,
		true,
		true,
		true,
		parentDir,
		filePath,
		contentsBytes,
	)
}

func (it fileWriter) ChmodFile(
	isRemoveBeforeWrite bool,
	chmodFile os.FileMode,
	filePath string,
	contentsBytes []byte,
) error {
	parentDir := it.ParentDir(filePath)

	return it.All(
		dirDefaultChmod,
		chmodFile,
		isRemoveBeforeWrite,
		true,
		true,
		true,
		parentDir,
		filePath,
		contentsBytes,
	)
}
