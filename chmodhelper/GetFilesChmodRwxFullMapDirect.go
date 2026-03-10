package chmodhelper

import "github.com/alimtvnetwork/core/coredata/corestr"

// GetFilesChmodRwxFullMapDirect returns filePath -> "-rwxrwxrwx"
func GetFilesChmodRwxFullMapDirect(
	requestedPaths ...string,
) (*corestr.Hashmap, error) {
	return GetFilesChmodRwxFullMap(requestedPaths)
}
