package codegen

import (
	"errors"

	"gitlab.com/auk-go/core/chmodhelper"
)

type newCodeOutputCreator struct{}

func (it newCodeOutputCreator) Default(
	unit, testCase string,
	fileWriter *chmodhelper.SimpleFileReaderWriter,
) *CodeOutput {
	return &CodeOutput{
		UnitTest:   unit,
		TestCase:   testCase,
		FileWriter: fileWriter,
	}
}

func (it newCodeOutputCreator) All(
	unit, testCase string,
	fileWriter *chmodhelper.SimpleFileReaderWriter,
) *CodeOutput {
	return &CodeOutput{
		UnitTest:   unit,
		TestCase:   testCase,
		FileWriter: fileWriter,
	}
}

func (it newCodeOutputCreator) Invalid(
	err error,
) *CodeOutput {
	return &CodeOutput{
		Error: err,
	}
}

func (it newCodeOutputCreator) InvalidMsg(
	msg string,
) *CodeOutput {
	return &CodeOutput{
		Error: errors.New(msg),
	}
}
