package codegen

import (
	"fmt"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/iserror"
)

type FinalCode struct {
	UnitTest             *GoCode
	TestCase             *GoCode
	StructName, FuncName string
	Error                error
	FileWriter           *chmodhelper.SimpleFileReaderWriter
}

func (it *FinalCode) IsValid() bool {
	return it != nil &&
		it.Error == nil &&
		it.UnitTest.IsCodeDefined() ||
		it.TestCase.IsCodeDefined()
}

func (it *FinalCode) IsInvalid() bool {
	return !it.IsValid()
}

func (it *FinalCode) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *FinalCode) IsEmptyError() bool {
	return it == nil && it.Error == nil
}

func (it *FinalCode) ErrorString() string {
	if it.IsEmptyError() {
		return ""
	}

	return it.Error.Error()
}

func (it *FinalCode) HasUnitTest() bool {
	return it != nil &&
		it.UnitTest.IsCodeDefined()
}

func (it *FinalCode) HasTestCase() bool {
	return it != nil &&
		it.TestCase.IsCodeDefined()
}

func (it *FinalCode) Write() errcore.RawErrCollection {
	var rawErrCollection errcore.RawErrCollection

	if it == nil {
		rawErrCollection.AddMsg("code output is nil")

		return rawErrCollection
	}

	if it.IsInvalid() {
		rawErrCollection.Add(it.Error)

		return rawErrCollection
	}

	if it.HasUnitTest() {
		rawErrCollection.Add(it.WriteUnitTestFile())
	}

	if it.HasTestCase() {
		rawErrCollection.Add(it.WriteTestCaseFile())
	}

	return rawErrCollection
}

func (it *FinalCode) WriteUnitTestFile() error {
	filePath := it.unitTestFileName()
	code, err := it.UnitTest.CompileFullCode()

	if iserror.Defined(err) {
		return err
	}

	return it.FileWriter.WriteRelativePath(
		it.FileWriter.IsRemoveBeforeWrite,
		filePath,
		[]byte(code),
	)
}

func (it *FinalCode) WriteTestCaseFile() error {
	filePath := it.testCaseFileName()
	code, err := it.TestCase.CompileFullCode()

	if iserror.Defined(err) {
		return err
	}

	return it.FileWriter.WriteRelativePath(
		it.FileWriter.IsRemoveBeforeWrite,
		filePath,
		[]byte(code),
	)
}

func (it *FinalCode) unitTestFileName() string {
	if len(it.StructName) == 0 {
		return fmt.Sprintf(
			"%s_test.go",
			it.FuncName,
		)
	}

	return fmt.Sprintf(
		"%s_%s_test.go",
		it.StructName,
		it.FuncName,
	)
}

func (it *FinalCode) testCaseFileName() string {
	if len(it.StructName) == 0 {
		return fmt.Sprintf(
			"%s_testcases.go",
			it.FuncName,
		)
	}

	return fmt.Sprintf(
		"%s_%s_testcases.go",
		it.StructName,
		it.FuncName,
	)
}
