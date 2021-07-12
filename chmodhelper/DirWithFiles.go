package chmodhelper

import "os"

type DirWithFiles struct {
	Dir   string
	Files []string
}

func (it *DirWithFiles) CreatePaths(
	isRemoveBeforeCreate bool,
	fileMode os.FileMode,
) error {
	return CreateDirWithFiles(
		isRemoveBeforeCreate,
		fileMode,
		it,
	)
}
