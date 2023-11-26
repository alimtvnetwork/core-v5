package codestack

import "gitlab.com/auk-go/core/constants"

type stacksTo struct{}

func StackTracesBytes(stackSkipIndex int) []byte {
	return NewStacksDefaultCount(stackSkipIndex + defaultInternalSkip).
		StackTracesBytes()
}

func StackTracesBytesDefault() []byte {
	return NewStacksDefaultCount(defaultInternalSkip).
		StackTracesBytes()
}

func StacksCountString(
	startSkipIndex, count int,
) string {
	stacks := NewStacksDefault(
		startSkipIndex+defaultInternalSkip,
		count,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}

func StacksCountStringUsingFmt(
	formatter Formatter,
	startSkipIndex, count int,
) string {
	stacks := NewStacksDefault(
		startSkipIndex+defaultInternalSkip,
		count,
	)

	toString := stacks.JoinUsingFmt(
		formatter,
		constants.NewLineSpaceHyphenSpace,
	)
	stacks.Dispose()

	return toString
}

func StacksJsonString(
	startSkipIndex int,
) string {
	stacks := NewStacksDefaultCount(
		startSkipIndex + defaultInternalSkip,
	)

	json := stacks.JsonPtr()
	stacks.Dispose()
	json.HandleError()

	return json.JsonString()
}

func StacksJsonStringDefault() string {
	return StacksJsonString(defaultInternalSkip)
}

func StacksString(
	startSkipIndex int,
) string {
	stacks := NewStacksDefaultCount(
		startSkipIndex + defaultInternalSkip,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}

func StacksStringDefault() string {
	stacks := NewStacksDefaultCount(
		defaultInternalSkip,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}
