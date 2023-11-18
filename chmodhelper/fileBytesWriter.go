package chmodhelper

import (
	"os"
	"path/filepath"
)

type fileBytesWriter struct{}

// WithDirChmodLock
//
// Create dir safely if required.
func (it fileBytesWriter) WithDirChmodLock(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WithDirChmod(
		chmodDir,
		chmodFile,
		writingFilePath,
		contentsBytes)
}

// WithDirChmod
//
// Create dir safely if required.
func (it fileBytesWriter) WithDirChmod(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	parentDir := filepath.Clean(
		filepath.Dir(writingFilePath))

	return fileWriter{}.All(
		chmodDir,
		chmodFile,
		true,
		true,
		true,
		parentDir,
		writingFilePath,
		contentsBytes)
}

// Chmod
//
// Create dir safely if required.
func (it fileBytesWriter) Chmod(
	chmodDir os.FileMode,
	chmodFile os.FileMode,
	writingFilePath string,
	contentsBytes []byte,
) error {
	parentDir := filepath.Clean(
		filepath.Dir(writingFilePath))

	return fileWriter{}.All(
		chmodDir,
		chmodFile,
		true,
		true,
		true,
		parentDir,
		writingFilePath,
		contentsBytes)
}

func (it fileBytesWriter) WithDirLock(
	writingFilePath string,
	contentsBytes []byte,
) error {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WithDir(
		writingFilePath,
		contentsBytes)
}

// WithDir
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it fileBytesWriter) WithDir(
	writingFilePath string,
	contentsBytes []byte,
) error {
	return it.WithDirChmod(
		dirDefaultChmod,
		fileDefaultChmod,
		writingFilePath,
		contentsBytes)
}

// Default
//
//	Applies default chmod (for dir - 0755, for file - 0644)
func (it fileBytesWriter) Default(
	writingFilePath string,
	contentsBytes []byte,
) error {
	return it.WithDirChmod(
		dirDefaultChmod,
		fileDefaultChmod,
		writingFilePath,
		contentsBytes)
}
