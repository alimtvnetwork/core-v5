package corevalidatortestwrappers

import "github.com/alimtvnetwork/core/corevalidator"

type SegmentValidatorWrapper struct {
	Header                string
	IsSkipOnContentsEmpty bool
	IsCaseSensitive       bool
	corevalidator.RangeSegmentsValidator
}
