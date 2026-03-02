package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/issetter"
)

// BaseTestCase
//
//   - Title : Test case header
//   - ArrangeInput : Preparing input
//   - ActualInput : Input for the act method
//   - ExpectedInput : Set expectations for the unit test (what we are going receive from invoking something)
//   - Will verify type using VerifyTypeOf
//
// Getters and formatting are in BaseTestCaseGetters.go.
// Type validation logic is in BaseTestCaseValidation.go.
// Assertion methods are in BaseTestCaseAssertions.go.
type BaseTestCase struct {
	Title           string         `json:",omitempty"` // consider as header
	ArrangeInput    any            `json:",omitempty"` // preparing input, initial input
	ActualInput     any            `json:",omitempty"` // (dynamically set) : must be set after running Act, using SetActual
	ExpectedInput   any            `json:",omitempty"` // expectation set from the test
	Additional      any            `json:",omitempty"` // additional input to do
	CustomFormat    string         `json:",omitempty"` // custom format for the test case
	VerifyTypeOf    *VerifyTypeOf  `json:",omitempty"` // Setting this creates the verify auto, verifies ArrangeInput, ActualInput, ExpectedInput type
	Parameters      *args.Holder   `json:",omitempty"` // If Act function / or any function requires more parameters it can be defined in the Holder.
	IsEnable        issetter.Value `json:",omitempty"` // Only false makes it disabled.
	HasError        bool           `json:",omitempty"`
	HasPanic        bool           `json:",omitempty"`
	IsValidateError bool           `json:",omitempty"`
}

func (it *BaseTestCase) CaseTitle() string {
	return it.Title
}

func (it *BaseTestCase) ArrangeTypeName() string {
	return reflectinternal.TypeName(it.ArrangeInput)
}

func (it *BaseTestCase) IsTypeInvalidOrSkipVerify() bool {
	return it == nil ||
		it.VerifyTypeOf == nil ||
		it.VerifyTypeOf.IsInvalidOrSkipVerify()
}

func (it *BaseTestCase) HasParameters() bool {
	return it != nil &&
		it.Parameters != nil
}

func (it *BaseTestCase) IsInvalidParameters() bool {
	return it == nil || it.Parameters == nil
}

func (it *BaseTestCase) FirstParam() any {
	if it.IsInvalidParameters() {
		return nil
	}

	return it.Parameters.First
}

func (it *BaseTestCase) SecondParam() any {
	if it.IsInvalidParameters() {
		return nil
	}

	return it.Parameters.Second
}

func (it *BaseTestCase) ThirdParam() any {
	if it.IsInvalidParameters() {
		return nil
	}

	return it.Parameters.Third
}

func (it *BaseTestCase) FourthParam() any {
	if it.IsInvalidParameters() {
		return nil
	}

	return it.Parameters.Fourth
}

func (it *BaseTestCase) FifthParam() any {
	if it.IsInvalidParameters() {
		return nil
	}

	return it.Parameters.Fifth
}

func (it *BaseTestCase) HashmapParam() (hasMapItem bool, hashMap map[string]any) {
	if it.IsInvalidParameters() {
		return false, map[string]any{}
	}

	hashMap = it.Parameters.Hashmap

	return len(hashMap) > 0, hashMap
}

func (it *BaseTestCase) HasValidHashmapParam() bool {
	if it.IsInvalidParameters() {
		return false
	}

	hashMap := it.Parameters.Hashmap

	return len(hashMap) > 0
}

func (it *BaseTestCase) IsVerifyType() bool {
	return it != nil && !it.IsTypeInvalidOrSkipVerify()
}

// ArrangeString
//
//	returns ArrangeInput in string
//	format using constants.SprintValueFormat
func (it *BaseTestCase) ArrangeString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ArrangeInput,
	)
}

func (it *BaseTestCase) AsSimpleTestCaseWrapper() SimpleTestCaseWrapper {
	return it
}

func (it *BaseTestCase) AsBaseTestCaseWrapper() BaseTestCaseWrapper {
	return it
}
