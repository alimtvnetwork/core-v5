package coretestcases

import "testing"

// ShouldBeEqualFirst asserts using ShouldBeEqual with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeEqualFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldBeEqual(
		t,
		0,
		actualElements...,
	)
}

// ShouldBeTrimEqualFirst asserts using ShouldBeTrimEqual with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeTrimEqualFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldBeTrimEqual(
		t,
		0,
		actualElements...,
	)
}

// ShouldBeSortedEqualFirst asserts using ShouldBeSortedEqual with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeSortedEqualFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldBeSortedEqual(
		t,
		0,
		actualElements...,
	)
}

// ShouldContainsFirst asserts using ShouldContains with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldContainsFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldContains(
		t,
		0,
		actualElements...,
	)
}

// ShouldStartsWithFirst asserts using ShouldStartsWith with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldStartsWithFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldStartsWith(
		t,
		0,
		actualElements...,
	)
}

// ShouldEndsWithFirst asserts using ShouldEndsWith with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldEndsWithFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldEndsWith(
		t,
		0,
		actualElements...,
	)
}

// ShouldBeNotEqualFirst asserts using ShouldBeNotEqual with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeNotEqualFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldBeNotEqual(
		t,
		0,
		actualElements...,
	)
}

// ShouldBeRegexFirst asserts using ShouldBeRegex with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeRegexFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldBeRegex(
		t,
		0,
		actualElements...,
	)
}
