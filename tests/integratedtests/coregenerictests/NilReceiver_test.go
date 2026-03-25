package coregenerictests

import "testing"

// ==========================================
// LinkedList — CaseNilSafe pattern
// ==========================================

func Test_LinkedList_NilReceiver(t *testing.T) {
	for caseIndex, tc := range linkedListNilSafeTestCases {
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// Hashset — CaseNilSafe pattern
// ==========================================

func Test_Hashset_NilReceiver(t *testing.T) {
	for caseIndex, tc := range hashsetNilSafeTestCases {
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// Hashmap — CaseNilSafe pattern
// ==========================================

func Test_Hashmap_NilReceiver(t *testing.T) {
	for caseIndex, tc := range hashmapNilSafeTestCases {
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// Pair — CaseNilSafe pattern
// ==========================================

func Test_Pair_NilReceiver(t *testing.T) {
	for caseIndex, tc := range pairNilSafeTestCases {
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// Triple — CaseNilSafe pattern
// ==========================================

func Test_Triple_NilReceiver(t *testing.T) {
	for caseIndex, tc := range tripleNilSafeTestCases {
		tc.ShouldBeSafe(t, caseIndex)
	}
}
