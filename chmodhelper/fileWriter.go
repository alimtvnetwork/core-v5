package chmodhelper

import (
	"errors"
	"os"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/internal/osconstsinternal"
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
		isApplyChmodMust,
		isApplyChmodOnMismatch,
		isCreateDirOnRequired,
		parentDirPath,
		writingFilePath,
		contentsBytes)
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
	isApplyChmodMust,
	isApplyChmodOnMismatch bool, // only apply for file, dir will not be applied if already created
	isCreateDirOnRequired bool,
	parentDirPath string,
	writingFilePath string,
	contentsBytes []byte,
) error {
	dirErr := dirCreator{}.If(
		isCreateDirOnRequired,
		chmodDir,
		parentDirPath)

	if dirErr != nil {
		return dirErr
	}

	err := os.WriteFile(
		writingFilePath,
		contentsBytes,
		chmodFile)

	if err != nil {
		return errors.New(
			"file writing failed" +
				"filePath : " + writingFilePath +
				"contents : " + corejson.BytesToString(contentsBytes) +
				", chmod file :" + chmodFile.String() + ", " +
				", chmod dir :" + chmodDir.String() + ", " +
				err.Error())
	}

	isNotApplyChmod := !isApplyChmodMust

	if isNotApplyChmod || osconstsinternal.IsWindows {
		return nil
	}

	// unix, must chmod
	if isApplyChmodOnMismatch && ChmodVerify.IsEqual(writingFilePath, chmodFile) {
		return nil
	}

	// not equal or apply anyway
	return ChmodApply.Default(chmodFile, writingFilePath)
}
