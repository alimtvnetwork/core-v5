package corestr

import "gitlab.com/auk-go/core/coredata/coregeneric"

// LeftRightFromSplit splits a string by separator into a LeftRight.
//
// Delegates to coregeneric.PairFromSplit.
//
// Usage:
//
//	lr := corestr.LeftRightFromSplit("key=value", "=")
//	lr.Left  // "key"
//	lr.Right // "value"
func LeftRightFromSplit(input, sep string) *LeftRight {
	pair := coregeneric.PairFromSplit(input, sep)

	return &LeftRight{Pair: *pair}
}

// LeftRightFromSplitTrimmed splits and trims whitespace from both parts.
//
// Delegates to coregeneric.PairFromSplitTrimmed.
func LeftRightFromSplitTrimmed(input, sep string) *LeftRight {
	pair := coregeneric.PairFromSplitTrimmed(input, sep)

	return &LeftRight{Pair: *pair}
}

// LeftRightFromSplitFull splits at the first separator only.
// Right gets everything after the first separator.
//
// Delegates to coregeneric.PairFromSplitFull.
//
// Usage:
//
//	lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
//	lr.Left  // "a"
//	lr.Right // "b:c:d"
func LeftRightFromSplitFull(input, sep string) *LeftRight {
	pair := coregeneric.PairFromSplitFull(input, sep)

	return &LeftRight{Pair: *pair}
}

// LeftRightFromSplitFullTrimmed is LeftRightFromSplitFull with trimmed output.
//
// Delegates to coregeneric.PairFromSplitFullTrimmed.
func LeftRightFromSplitFullTrimmed(input, sep string) *LeftRight {
	pair := coregeneric.PairFromSplitFullTrimmed(input, sep)

	return &LeftRight{Pair: *pair}
}
