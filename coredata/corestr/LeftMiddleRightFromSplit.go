package corestr

import "gitlab.com/auk-go/core/coredata/coregeneric"

// LeftMiddleRightFromSplit splits a string by separator into a LeftMiddleRight.
//
// Delegates to coregeneric.TripleFromSplit.
//
// Usage:
//
//	lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
//	lmr.Left   // "a"
//	lmr.Middle // "b"
//	lmr.Right  // "c"
func LeftMiddleRightFromSplit(input, sep string) *LeftMiddleRight {
	triple := coregeneric.TripleFromSplit(input, sep)

	return &LeftMiddleRight{Triple: *triple}
}

// LeftMiddleRightFromSplitTrimmed splits and trims whitespace from all parts.
//
// Delegates to coregeneric.TripleFromSplitTrimmed.
func LeftMiddleRightFromSplitTrimmed(input, sep string) *LeftMiddleRight {
	triple := coregeneric.TripleFromSplitTrimmed(input, sep)

	return &LeftMiddleRight{Triple: *triple}
}

// LeftMiddleRightFromSplitN splits into at most 3 parts.
// The third part contains the remainder after the second separator.
//
// Delegates to coregeneric.TripleFromSplitN.
//
// Usage:
//
//	lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
//	lmr.Left   // "a"
//	lmr.Middle // "b"
//	lmr.Right  // "c:d:e"
func LeftMiddleRightFromSplitN(input, sep string) *LeftMiddleRight {
	triple := coregeneric.TripleFromSplitN(input, sep)

	return &LeftMiddleRight{Triple: *triple}
}

// LeftMiddleRightFromSplitNTrimmed splits into at most 3 parts with trimming.
//
// Delegates to coregeneric.TripleFromSplitNTrimmed.
func LeftMiddleRightFromSplitNTrimmed(input, sep string) *LeftMiddleRight {
	triple := coregeneric.TripleFromSplitNTrimmed(input, sep)

	return &LeftMiddleRight{Triple: *triple}
}
