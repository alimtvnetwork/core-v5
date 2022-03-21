package chmodhelper

import (
	"errors"
	"fmt"
	"os"

	"gitlab.com/evatix-go/core/errcore"
)

type fwChmodApplier struct {
	rw *SimpleFileReaderWriter
}

func (it fwChmodApplier) OnParent() error {
	return it.OnDir(it.rw.ParentDir)
}

func (it fwChmodApplier) OnDir(dir string) error {
	return it.Apply(
		it.rw.ChmodDir,
		dir)
}

func (it fwChmodApplier) OnFile() error {
	return it.Apply(
		it.rw.ChmodFile,
		it.rw.FilePath)
}

func (it fwChmodApplier) Apply(
	fileMode os.FileMode,
	location string,
) error {
	err := os.Chmod(
		location,
		fileMode)

	if err == nil {
		return nil
	}

	message := fmt.Sprintf(
		"applying chmod failed, path : %q, chmod: %q, is-file-exist: %v, err: %s",
		location,
		FileModeFriendlyString(fileMode),
		IsPathExists(location),
		err.Error())

	// has error
	return errors.New(message)
}

func (it fwChmodApplier) OnDiffFile(filePath string) error {
	return os.Chmod(filePath, it.rw.ChmodFile)
}

// OnAll
//
//  both file, parent dir
func (it fwChmodApplier) OnAll() error {
	err := it.OnParent()

	if err != nil {
		return err
	}

	return it.OnFile()
}

func (it fwChmodApplier) DirRecursive(
	isSkipOnInvalid bool,
	dir string,
) error {
	rwx := New.RwxWrapper.UsingFileMode(it.rw.ChmodDir)

	return rwx.ApplyRecursive(isSkipOnInvalid, dir)
}

func (it fwChmodApplier) OnParentRecursive() error {
	return it.DirRecursive(
		false,
		it.rw.ParentDir)
}

func (it fwChmodApplier) OnMismatch(
	isFile,
	isParentDir bool,
) error {
	if !isFile && !isParentDir {
		return nil
	}

	verifier := it.rw.ChmodVerifier()
	var fileErr, dirErr error

	if isFile && verifier.HasMismatchFile() {
		fileErr = it.OnFile()
	}

	if isParentDir && verifier.HasMismatchParentDir() {
		dirErr = it.OnParent()
	}

	return errcore.MergeErrors(fileErr, dirErr)
}
