package corevalidatortestwrappers

import "gitlab.com/evatix-go/core/corevalidator"

type SegmentValidatorWrapper struct {
	Header                string
	IsSkipOnContentsEmpty bool
	IsCaseSensitive       bool
	corevalidator.SimpleSliceRangeValidator
}
