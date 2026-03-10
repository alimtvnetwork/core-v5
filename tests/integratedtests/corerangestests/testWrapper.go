package corerangestests

import (
	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/corevalidator"
)

type testWrapper struct {
	coretests.BaseTestCase
	IsExpectingError bool
	HasPanic         bool
	Validator        corevalidator.SliceValidator
}

func (it testWrapper) Arrange() []corerange.MinMaxInt {
	return it.ArrangeInput.([]corerange.MinMaxInt)
}

func (it testWrapper) Expected() []int {
	return it.ExpectedInput.([]int)
}
