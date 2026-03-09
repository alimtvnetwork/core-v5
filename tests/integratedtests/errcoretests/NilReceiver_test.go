package errcoretests

import "testing"

// ==========================================
// ConcatMessageWithErr nil — CaseNilSafe pattern
// ==========================================

func Test_ConcatMessageWithErr_NilReceiver(t *testing.T) {
	for caseIndex, tc := range concatMessageNilSafeTestCases_v2 {
		tc.ShouldBeSafe(t, caseIndex)
	}
}
