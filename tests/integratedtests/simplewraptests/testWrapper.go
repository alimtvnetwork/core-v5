package simplewraptests

import (
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/corevalidator"
)

type testWrapper struct {
	coretests.BaseTestCase
	IsExpectingError bool
	HasPanic         bool
	Validator        corevalidator.SliceValidator
}

func (it testWrapper) Arrange() []string {
	return it.ArrangeInput.([]string)
}

func (it testWrapper) Expected() []int {
	return it.ExpectedInput.([]int)
}
