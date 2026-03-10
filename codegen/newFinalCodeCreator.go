package codegen

import (
	"errors"

	"github.com/alimtvnetwork/core/chmodhelper"
)

type newFinalCodeCreator struct{}

func (it newFinalCodeCreator) Default(
	structName, funcName string,
	unit, testCase *GoCode,
	fileWriter *chmodhelper.SimpleFileReaderWriter,
	options Options,
) *FinalCode {
	return &FinalCode{
		UnitTest:   unit,
		TestCase:   testCase,
		StructName: structName,
		FuncName:   funcName,
		FileWriter: fileWriter,
		Options:    options,
	}
}

func (it newFinalCodeCreator) UsingGeneratorFunc(
	generateFunc BaseGenerator,
	unit, testCase *GoCode,
	options Options,
) *FinalCode {
	return &FinalCode{
		UnitTest:   unit,
		TestCase:   testCase,
		StructName: generateFunc.StructName(),
		FuncName:   generateFunc.FuncName(),
		FileWriter: generateFunc.FileWriter(),
		Options:    options,
	}
}

func (it newFinalCodeCreator) All(
	structName, funcName string,
	unit, testCase *GoCode,
	fileWriter *chmodhelper.SimpleFileReaderWriter,
	options Options,
) *FinalCode {
	return &FinalCode{
		UnitTest:   unit,
		TestCase:   testCase,
		StructName: structName,
		FuncName:   funcName,
		FileWriter: fileWriter,
		Options:    options,
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
