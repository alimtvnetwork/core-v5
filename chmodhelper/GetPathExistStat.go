package chmodhelper

import (
	"os"
)

func GetPathExistStat(location string) *PathExistStat {
	fileInfo, err := os.Stat(location)
	isExist := err == nil || !os.IsNotExist(err)

	return &PathExistStat{
		FileInfo: fileInfo,
		IsExist:  isExist,
		Error:    err,
		Location: location,
	}
}
