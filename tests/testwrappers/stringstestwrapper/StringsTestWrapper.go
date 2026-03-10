package stringstestwrapper

import (
	"github.com/alimtvnetwork/core/coretests"
)

// StringsTestWrapper wraps BaseTestCase with typed []string accessors
// for ArrangeInput and ExpectedInput. Use a type alias in package-local
// testWrapper.go files to keep test case definitions unchanged:
//
//	type testWrapper = stringstestwrapper.StringsTestWrapper
type StringsTestWrapper struct {
	coretests.BaseTestCase
}

// Arrange returns ArrangeInput cast to []string.
func (it StringsTestWrapper) Arrange() []string {
	return it.ArrangeInput.([]string)
}

// Expected returns ExpectedInput cast to []string.
func (it StringsTestWrapper) Expected() []string {
	return it.ExpectedInput.([]string)
}
