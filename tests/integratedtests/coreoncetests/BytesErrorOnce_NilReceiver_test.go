package coreoncetests

import (
	"testing"
)

// =============================================================================
// BytesErrorOnce — Nil Receiver (CaseNilSafe pattern)
// =============================================================================

func Test_BytesErrorOnce_NilReceiver(t *testing.T) {
	for caseIndex, tc := range bytesErrorOnceNilSafeTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}
