package codestack

import "gitlab.com/auk-go/core/constants"

type stacksTo struct{}

func (it stacksTo) Bytes(stackSkipIndex int) []byte {
	return NewStacksDefaultCount(stackSkipIndex + defaultInternalSkip).
		StackTracesBytes()
}

func (it stacksTo) BytesDefault() []byte {
	return NewStacksDefaultCount(defaultInternalSkip).
		StackTracesBytes()
}

func (it stacksTo) String(
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

func (it stacksTo) StringUsingFmt(
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

func (it stacksTo) JsonString(
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

func (it stacksTo) JsonStringDefault() string {
	return it.JsonString(defaultInternalSkip)
}

func (it stacksTo) StringNoCount(
	startSkipIndex int,
) string {
	stacks := NewStacksDefaultCount(
		startSkipIndex + defaultInternalSkip,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}

func (it stacksTo) StringDefault() string {
	stacks := NewStacksDefaultCount(
		defaultInternalSkip,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}
