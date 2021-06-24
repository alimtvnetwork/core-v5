package chmodhelper

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type FilteredPathFileInfoMap struct {
	FilesToInfoMap           map[string]os.FileInfo
	MissingOrOtherPathIssues []string
	Error                    error
}

func InvalidFilteredPathFileInfoMap() *FilteredPathFileInfoMap {
	return &FilteredPathFileInfoMap{
		FilesToInfoMap:           map[string]os.FileInfo{},
		MissingOrOtherPathIssues: []string{},
		Error:                    nil,
	}
}

func (it *FilteredPathFileInfoMap) HasAnyValidFileInfo() bool {
	return len(it.FilesToInfoMap) > 0
}

func (it *FilteredPathFileInfoMap) IsEmptyValidFileInfos() bool {
	return len(it.FilesToInfoMap) == 0
}

func (it *FilteredPathFileInfoMap) LengthOfIssues() int {
	return len(it.MissingOrOtherPathIssues)
}

func (it *FilteredPathFileInfoMap) IsEmptyIssues() bool {
	return len(it.MissingOrOtherPathIssues) == 0
}

func (it *FilteredPathFileInfoMap) HasAnyIssues() bool {
	return it.Error != nil || it.HasAnyMissingPaths()
}

func (it *FilteredPathFileInfoMap) HasAnyMissingPaths() bool {
	return len(it.MissingOrOtherPathIssues) > 0
}

func (it *FilteredPathFileInfoMap) MissingPathsToString() string {
	return strings.Join(it.MissingOrOtherPathIssues, constants.CommaSpace)
}
