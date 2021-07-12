package chmodhelper

import "os"

func GetPathExistStatExpand(location string) (fileInfo os.FileInfo, isExist bool, err error) {
	fileInfo, err = os.Stat(location)
	isExist = err == nil || !os.IsNotExist(err)

	return fileInfo, isExist, err
}
