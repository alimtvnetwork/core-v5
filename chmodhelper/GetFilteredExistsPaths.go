package chmodhelper

import (
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/internal/fsinternal"
)

func GetFilteredExistsPaths(locations []string) (
	foundFiles, missingOrPathsWithIssues []string,
) {
	if len(locations) == 0 {
		return []string{}, []string{}
	}

	results := make(
		[]string,
		constants.Zero,
		len(locations)+constants.Capacity2)

	for _, location := range locations {
		if fsinternal.IsPathExists(location) {
			results = append(results, location)
		} else {
			missingOrPathsWithIssues = append(
				missingOrPathsWithIssues,
				location)
		}
	}

	return results, missingOrPathsWithIssues
}
