package convertinteranl

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type stringUtil struct{}

func (it stringUtil) PrependWithSpaces(
	joiner string,
	spaceCountLines int,
	existingLines []string,
	prependingLinesSpaceCount int,
	prependingLines ...string,
) string {
	toSlice := Util.Strings.PrependWithSpaces(
		spaceCountLines,
		existingLines,
		prependingLinesSpaceCount,
		prependingLines...,
	)

	return strings.Join(toSlice, joiner)
}

func (it stringUtil) PrependWithSpacesDefault(
	spaceCountLines int,
	existingLines []string,
	prependingLinesSpaceCount int,
	prependingLines ...string,
) string {
	toSlice := Util.Strings.PrependWithSpaces(
		spaceCountLines,
		existingLines,
		prependingLinesSpaceCount,
		prependingLines...,
	)

	return strings.Join(toSlice, constants.NewLineUnix)
}
