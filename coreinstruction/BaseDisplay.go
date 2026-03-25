package coreinstruction

import (
	"regexp"
	"strings"
)

type BaseDisplay struct {
	Display string `json:"Display"`
}

func (receiver BaseDisplay) IsDisplay(display string) bool {
	return receiver.Display == display
}

func (receiver BaseDisplay) IsDisplayCaseInsensitive(display string) bool {
	return strings.EqualFold(receiver.Display, display)
}

func (receiver BaseDisplay) IsDisplayContains(displayContains string) bool {
	return strings.Contains(receiver.Display, displayContains)
}

func (receiver BaseDisplay) IsDisplayRegexMatches(regex *regexp.Regexp) bool {
	return regex.MatchString(receiver.Display)
}
