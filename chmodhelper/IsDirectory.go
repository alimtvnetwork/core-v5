package chmodhelper

import "os"

func IsDirectory(location string) bool {
	fileInfo, err := os.Stat(location)

	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}
