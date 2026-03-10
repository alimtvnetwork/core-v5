package corevalidator

import (
	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

type RangesSegment struct {
	corerange.RangeInt
	ExpectedLines []string
	CompareAs     stringcompareas.Variant
	Condition
}
