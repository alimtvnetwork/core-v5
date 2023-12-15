package codegen

import (
	"errors"

	"gitlab.com/auk-go/core/chmodhelper"
)

type newFinalCodeCreator struct{}

func (it newFinalCodeCreator) Default(
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

func (it newFinalCodeCreator) All(
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

func (it newFinalCodeCreator) Invalid(
	err error,
) *FinalCode {
	return &FinalCode{
		Error: err,
	}
}

func (it newFinalCodeCreator) InvalidMsg(
	msg string,
) *FinalCode {
	return &FinalCode{
		Error: errors.New(msg),
	}
}
