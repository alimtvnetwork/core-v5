package chmodhelper

import (
	"os"
)

// IsChmod expectedHyphenedRwx should be 10 chars example "-rwxrwxrwx"
func IsChmod(location string, expectedHyphenedRwx string) bool {
	if len(expectedHyphenedRwx) != HyphenedRwxLength {
		return false
	}

	if location == "" {
		return false
	}

	fileInfo, err := os.Stat(location)

	if os.IsNotExist(err) || fileInfo == nil {
		return false
	}

	existingFileMode := fileInfo.Mode().String()[1:]

	return existingFileMode == expectedHyphenedRwx[1:]
}
