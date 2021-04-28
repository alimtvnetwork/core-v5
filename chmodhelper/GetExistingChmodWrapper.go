package chmodhelper

import "os"

func GetExistingChmodWrapper(filePath string) (Wrapper, error) {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return Wrapper{}, err
	}

	return NewUsingFileMode(fileInfo.Mode()), err
}
