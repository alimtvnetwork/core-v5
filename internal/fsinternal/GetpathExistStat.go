package fsinternal

import "os"

func GetPathExistStat(location string) (fileInfo os.FileInfo, isExist bool, err error) {
	fileInfo, err = os.Stat(location)
	isExist = err == nil || !os.IsNotExist(err)

	return fileInfo, isExist, err
}
