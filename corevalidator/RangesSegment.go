package corevalidator

import (
	"gitlab.com/evatix-go/core/coredata/corerange"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
)

type RangesSegment struct {
	corerange.RangeInt
	ExpectedLines []string
	CompareAs     stringcompareas.Variant
	ValidatorCoreCondition
}
