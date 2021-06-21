package coreinstruction

import (
	"regexp"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

type BaseTags struct {
	tagsHashset *corestr.Hashset
	Tags        *[]string `json:"Tags,omitempty"`
}

func NewTagsPtr(tags *[]string) *BaseTags {
	if tags == nil || len(*tags) == 0 {
		return NewTags(nil)
	}

	return NewTags(*tags)
}

func NewTags(tags []string) *BaseTags {
	if len(tags) == 0 {
		return &BaseTags{
			Tags: &[]string{},
		}
	}

	return &BaseTags{
		Tags: &tags,
	}
}

func (receiver BaseTags) TagsLength() int {
	if receiver.Tags == nil {
		return constants.Zero
	}

	return len(*receiver.Tags)
}

func (receiver BaseTags) IsTagsEmpty() bool {
	return receiver.TagsLength() == 0
}

func (receiver BaseTags) TagsHashset() *corestr.Hashset {
	if receiver.tagsHashset != nil {
		return receiver.tagsHashset
	}

	receiver.tagsHashset = corestr.NewHashsetUsingStrings(
		receiver.Tags)

	return receiver.tagsHashset
}

func (receiver BaseTags) HasAllTags(tags ...string) bool {
	if len(tags) == 0 {
		return true
	}

	hashset := receiver.TagsHashset()

	return hashset.HasAll(tags...)
}

func (receiver BaseTags) HasAnyTags(tags ...string) bool {
	if len(tags) == 0 {
		return true
	}

	hashset := receiver.TagsHashset()

	return hashset.HasAny(tags...)
}

func (receiver BaseTags) IsAnyTagMatchesRegex(regexp2 *regexp.Regexp) bool {
	if receiver.IsTagsEmpty() {
		return false
	}

	for _, s := range *receiver.Tags {
		if regexp2.MatchString(s) {
			return true
		}
	}

	return false
}
