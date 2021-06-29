package chmodhelper

import "os"

type LocationFileInfoRwxWrapper struct {
	Location   string
	FileInfo   os.FileInfo
	RwxWrapper *RwxWrapper
}
