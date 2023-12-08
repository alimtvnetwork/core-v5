package codegen

import "gitlab.com/auk-go/core/chmodhelper"

type CodeOutput struct {
	UnitTest string
	TestCase string
	Error    error
	chmodhelper.SimpleFileReaderWriter
}

func (it *CodeOutput) IsValid() bool {
	return it != nil &&
		it.Error == nil &&
		it.UnitTest != "" &&
		it.TestCase != ""
}

func (it *CodeOutput) IsInvalid() bool {
	return !it.IsValid()
}

func (it *CodeOutput) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *CodeOutput) IsEmptyError() bool {
	return it == nil && it.Error == nil
}

func (it *CodeOutput) ErrorString() string {
	if it.IsEmptyError() {
		return ""
	}

	return it.Error.Error()
}

func (it *CodeOutput) HasUnitTest() bool {
	return it != nil &&
		it.UnitTest != ""
}

func (it *CodeOutput) HasTestCase() bool {
	return it != nil &&
		it.TestCase != ""
}
