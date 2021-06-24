package chmodhelper

import "os"

func IsPathExistsPlusFileInfo(location string) (bool, os.FileInfo) {
	fileInfo, err := os.Stat(location)
	isExist := err == nil || !os.IsNotExist(err)

	return isExist, fileInfo
}
