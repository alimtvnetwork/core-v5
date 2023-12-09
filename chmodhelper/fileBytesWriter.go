package chmodhelper

import (
	"os"

	"gitlab.com/auk-go/core/internal/pathinternal"
)

type fileBytesWriter struct{}

// WithDirChmodLock
//
// Create dir safely if required.
func (it fileBytesWriter) WithDirChmodLock(
	isCleanBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WithDirChmod(
		isCleanBeforeWrite,
		chmodDir,
		chmodFile,
		writingFilePath,
		contentsBytes,
	)
}

// WithDirChmod
//
// Create dir safely if required.
func (it fileBytesWriter) WithDirChmod(
	isCleanBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	parentDir := pathinternal.ParentDir(writingFilePath)

	return fileWriter{}.All(
		chmodDir,
		chmodFile,
		isCleanBeforeWrite,
		true,
		true,
		true,
		parentDir,
		writingFilePath,
		contentsBytes,
	)
}

// Chmod
//
// Create dir safely if required.
func (it fileBytesWriter) Chmod(
	isCleanBeforeWrite bool,
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	parentDir := pathinternal.ParentDir(writingFilePath)

	return fileWriter{}.All(
		chmodDir,
		chmodFile,
		isCleanBeforeWrite,
		true,
		true,
		true,
		parentDir,
		writingFilePath,
		contentsBytes,
	)
}

func (it fileBytesWriter) WithDirLock(
	isCleanBeforeWrite bool,
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WithDir(
		isCleanBeforeWrite,
		writingFilePath,
		contentsBytes,
	)
}

// WithDir
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it fileBytesWriter) WithDir(
	isCleanBeforeWrite bool,
	writingFilePath string,
	contentsBytes []byte,
) error {
	return it.WithDirChmod(
		isCleanBeforeWrite,
		dirDefaultChmod,
		fileDefaultChmod,
		writingFilePath,
		contentsBytes,
	)
}

// Default
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it fileBytesWriter) Default(
	isCleanBeforeWrite bool,
	writingFilePath string,
	contentsBytes []byte,
) error {
	return it.WithDirChmod(
		isCleanBeforeWrite,
		dirDefaultChmod,
		fileDefaultChmod,
		writingFilePath,
		contentsBytes,
	)
}
