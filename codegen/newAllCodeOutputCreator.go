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
) *FinalCode {
	return &FinalCode{
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
) *FinalCode {
	return &FinalCode{
		UnitTest:   unit,
		TestCase:   testCase,
		StructName: structName,
		FuncName:   funcName,
		FileWriter: fileWriter,
	}
}

func (it newAllCodeOutputCreator) Invalid(
	err error,
) *FinalCode {
	return &FinalCode{
		Error: err,
	}
}

func (it newAllCodeOutputCreator) InvalidMsg(
	msg string,
) *FinalCode {
	return &FinalCode{
		Error: errors.New(msg),
	}
}
