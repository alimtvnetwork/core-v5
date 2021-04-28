package chmodhelper

import "os"

func GetExistingChmod(filePath string) (os.FileMode, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return 0, err
	}

	return fileInfo.Mode(), err
}
