package chmodhelper

import "os"

func IsDirectory(location string) bool {
	fileInfo, err := os.Stat(location)

	if err != nil {
		// Keep non-not-exist cases (e.g. permission denied) as "exists"
		// while avoiding nil fileInfo dereference.
		return !os.IsNotExist(err)
	}

	return fileInfo.IsDir()
}
