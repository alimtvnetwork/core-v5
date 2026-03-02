package coredynamic

import (
	"errors"
	"reflect"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/strutilinternal"
)

type SimpleResult struct {
	Dynamic
	Result  any
	Message string
	err     error
}

func InvalidSimpleResultNoMessage() *SimpleResult {
	return &SimpleResult{
		Result:  nil,
		Dynamic: NewDynamic(nil, false),
		Message: constants.EmptyString,
	}
}

func InvalidSimpleResult(
	invalidMessage string,
) *SimpleResult {
	return &SimpleResult{
		Result:  nil,
		Dynamic: NewDynamic(nil, false),
		Message: invalidMessage,
	}
}

func NewSimpleResultValid(
	result any,
) *SimpleResult {
	return &SimpleResult{
		Result:  result,
		Dynamic: NewDynamic(result, true),
		Message: constants.EmptyString,
	}
}

func NewSimpleResult(
	result any,
	isValid bool,
	invalidMessage string,
) *SimpleResult {
	return &SimpleResult{
		Result:  result,
		Dynamic: NewDynamic(result, isValid),
		Message: invalidMessage,
	}
}

func (it *SimpleResult) GetErrorOnTypeMismatch(
	typeMatch reflect.Type,
	isIncludeInvalidMessage bool,
) error {
	if it.IsReflectTypeOf(typeMatch) {
		return nil
	}

	typeMismatchMessage := errcore.CombineWithMsgTypeNoStack(
		errcore.TypeMismatchType,
		"Current type - ["+it.ReflectTypeName()+"], expected type",
		typeMatch,
	) + constants.NewLineUnix

	isExcludeMessage := !isIncludeInvalidMessage

	if isExcludeMessage {
		return errors.New(typeMismatchMessage)
	}

	return errors.New(typeMismatchMessage + it.Message)
}

func (it *SimpleResult) InvalidError() error {
	if it.err != nil {
		return it.err
	}

	if strutilinternal.IsEmptyOrWhitespace(it.Message) {
		return nil
	}

	if it.err == nil {
		it.err = errors.New(it.Message)
	}

	return it.err
}

func (it *SimpleResult) Clone() SimpleResult {
	return SimpleResult{
		Dynamic: it.Dynamic.Clone(),
		Result:  it.Result,
		Message: it.Message,
	}
}

func (it *SimpleResult) ClonePtr() *SimpleResult {
	if it == nil {
		return nil
	}

	return &SimpleResult{
		Dynamic: it.Dynamic.Clone(),
		Result:  it.Result,
		Message: it.Message,
	}
}
