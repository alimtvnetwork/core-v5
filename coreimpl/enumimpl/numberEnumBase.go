package enumimpl

import (
	"errors"
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/coreonce"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
)

type numberEnumBase struct {
	actualValueRanges    interface{}
	stringRanges         []string
	rangesCsvString      coreonce.StringOnce
	rangesInvalidMessage coreonce.StringOnce
	invalidError         coreonce.ErrorOnce
	typeName             string
}

// newNumberEnumBase
//
//  @actualRangesAnyType : []Byte, []int, []int8... not pointer
//
//  Lengths must match stringRanges and actualRangesAnyType
func newNumberEnumBase(
	typeName string,
	actualRangesAnyType interface{},
	stringRanges []string,
	min, max interface{},
) *numberEnumBase {
	if stringRanges == nil {
		errcore.MeaningfulErrorHandle(
			errcore.CannotBeNilType,
			"newNumberEnumBase",
			errors.New("StringRanges cannot be nil"))
	}

	rangesToCsvOnce := coreonce.NewStringOnce(func() string {
		return converters.StringsToCsvWithIndexes(
			stringRanges,
		)
	})

	invalidMessageOnce := coreonce.NewStringOnce(func() string {
		msg := errcore.EnumRangeNotMeet(
			min,
			max,
			rangesToCsvOnce.Value())

		return msg
	})

	return &numberEnumBase{
		actualValueRanges:    actualRangesAnyType,
		stringRanges:         stringRanges,
		rangesInvalidMessage: invalidMessageOnce,
		invalidError: coreonce.NewErrorOnce(func() error {
			return errors.New(invalidMessageOnce.Value())
		}),
		rangesCsvString: rangesToCsvOnce,
		typeName:        typeName,
	}
}

func (it *numberEnumBase) TypeName() string {
	return it.typeName
}

// NameWithValueOption
//
// Warning :
//
// Make sure non ptr is called +
// String should also be attached with non ptr.
func (it *numberEnumBase) NameWithValueOption(
	value interface{},
	isIncludeQuotation bool,
) string {
	if isIncludeQuotation {
		return fmt.Sprintf(
			constants.DoubleQuoteStringWithBracketWrapNumberFormat,
			value,
			value)
	}

	return fmt.Sprintf(
		constants.StringWithBracketWrapNumberFormat,
		value,
		value)
}

// NameWithValue
//
// Warning :
//
// Make sure non ptr is called +
// String should also be attached with non ptr.
func (it *numberEnumBase) NameWithValue(
	value interface{},
) string {
	return fmt.Sprintf(
		constants.StringWithBracketWrapNumberFormat,
		value,
		value)
}

func (it *numberEnumBase) ValueString(
	value interface{},
) string {
	return fmt.Sprintf(
		constants.SprintNumberFormat,
		value,
	)
}

// Format
//
//  Outputs name and
//  value by given format.
//
// sample-format :
//  - "Enum of {type-name} - {name} - {value}"
//
// sample-format-output :
//  - "Enum of EnumFullName - Invalid - 0"
//
// Key-Meaning :
//  - {type-name} : represents type-name string
//  - {name}      : represents name string
//  - {value}     : represents value string
func (it *numberEnumBase) Format(
	format string,
	value interface{},
) string {
	replacerMap := map[string]string{
		typeNameTemplateKey: it.TypeName(),
		nameKey:             it.ToEnumString(value),
		valueKey:            it.ValueString(value),
	}

	return utilstringinternal.ReplaceTemplateMap(
		true,
		format,
		replacerMap)
}

func (it *numberEnumBase) RangeNamesCsv() string {
	return it.rangesCsvString.Value()
}

func (it *numberEnumBase) RangesInvalidMessage() string {
	return it.rangesInvalidMessage.Value()
}

func (it *numberEnumBase) RangesInvalidErr() error {
	return it.invalidError.Value()
}

func (it *numberEnumBase) StringRangesPtr() *[]string {
	return &it.stringRanges
}

func (it *numberEnumBase) StringRanges() []string {
	return it.stringRanges
}

func (it *numberEnumBase) JsonString(input interface{}) string {
	return it.ToEnumString(input)
}

func (it *numberEnumBase) ToEnumString(
	input interface{},
) string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		input)
}
