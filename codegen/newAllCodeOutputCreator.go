package codegen

import (
	"errors"

	"gitlab.com/auk-go/core/chmodhelper"
)

type newAllCodeOutputCreator struct{}

func (it newAllCodeOutputCreator) Default(
	structName, funcName string,
	unit, testCase *GoCode,
	fileWriter *chmodhelper.SimpleFileReaderWriter,
) *AllCodeOutput {
	return &AllCodeOutput{
		UnitTest:   unit,
		TestCase:   testCase,
		StructName: structName,
		FuncName:   funcName,
		FileWriter: fileWriter,
	}
}

func (it newAllCodeOutputCreator) All(
	structName, funcName string,
	unit, testCase *GoCode,
	fileWriter *chmodhelper.SimpleFileReaderWriter,
) *AllCodeOutput {
	return &AllCodeOutput{
		UnitTest:   unit,
		TestCase:   testCase,
		StructName: structName,
		FuncName:   funcName,
		FileWriter: fileWriter,
	}
}

func (it newAllCodeOutputCreator) Invalid(
	err error,
) *AllCodeOutput {
	return &AllCodeOutput{
		Error: err,
	}
}

func (it newAllCodeOutputCreator) InvalidMsg(
	msg string,
) *AllCodeOutput {
	return &AllCodeOutput{
		Error: errors.New(msg),
	}
}
