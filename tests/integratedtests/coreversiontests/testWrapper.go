package coreversiontests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

type testWrapper coretestcases.CaseV1

func (it testWrapper) Arrange() []string {
	return it.ArrangeInput.([]string)
}

func (it testWrapper) Expected() []string {
	return it.ExpectedInput.([]string)
}

func (it testWrapper) AsCaseV1() coretestcases.CaseV1 {
	return coretestcases.CaseV1(it)
}
