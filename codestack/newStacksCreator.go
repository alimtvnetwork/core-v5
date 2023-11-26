package codestack

import "gitlab.com/auk-go/core/constants"

type newStacksCreator struct{}

func NewStacks(
	isSkipInvalid,
	isBreakOnceInvalid bool,
	startSkipIndex, // should start from 1
	stackCount int,
) TraceCollection {
	traces := NewTraceCollection(stackCount + constants.Capacity2)

	return *traces.AddsUsingSkip(
		isSkipInvalid,
		isBreakOnceInvalid,
		startSkipIndex+defaultInternalSkip,
		stackCount,
	)
}

func NewStacksDefault(
	startSkipIndex,
	stackCount int,
) TraceCollection {
	return NewStacks(
		true,
		true,
		startSkipIndex+defaultInternalSkip,
		stackCount,
	)
}

func NewStacksDefaultCount(
	startSkipIndex int,
) TraceCollection {
	return NewStacks(
		true,
		true,
		startSkipIndex+defaultInternalSkip,
		DefaultStackCount,
	)
}

func NewStacksDefaultCountSkip1() TraceCollection {
	return NewStacks(
		true,
		true,
		Skip1+defaultInternalSkip,
		DefaultStackCount,
	)
}

func NewStacksDefaultCountSkipNone() TraceCollection {
	return NewStacks(
		true,
		true,
		defaultInternalSkip,
		DefaultStackCount,
	)
}

func NewStacksDefaultPtr(
	startSkipIndex,
	stackCount int,
) *TraceCollection {
	return NewStacksPtr(
		true,
		true,
		startSkipIndex+defaultInternalSkip,
		stackCount,
	)
}

func NewStacksPtr(
	isSkipInvalid,
	isBreakOnceInvalid bool,
	startSkipIndex,
	stackCount int,
) *TraceCollection {
	traces := NewTraceCollection(stackCount + constants.Capacity2)

	return traces.AddsUsingSkip(
		isSkipInvalid,
		isBreakOnceInvalid,
		startSkipIndex+defaultInternalSkip,
		stackCount,
	)
}
