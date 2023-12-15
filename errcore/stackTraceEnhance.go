package errcore

import (
	"errors"
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type stackTraceEnhance struct{}

func (it stackTraceEnhance) Error(err error) error {
	return it.ErrorSkip(1, err)
}

func (it stackTraceEnhance) ErrorSkip(skip int, err error) error {
	if err == nil {
		return nil
	}

	msg := err.Error()

	return errors.New(it.MsgSkip(1+skip, msg))
}

func (it stackTraceEnhance) MsgToErrSkip(skip int, msg string) error {
	if len(msg) == 0 {
		return nil
	}

	return errors.New(it.MsgSkip(1+skip, msg))
}

func (it stackTraceEnhance) FmtSkip(skip int, format string, v ...interface{}) error {
	if len(format) == 0 {
		return nil
	}

	msg := fmt.Sprintf(format, v...)

	return errors.New(it.MsgSkip(1+skip, msg))
}

func (it stackTraceEnhance) Msg(msg string) string {
	if len(msg) == 0 {
		return ""
	}

	return it.MsgSkip(1, msg)
}

func (it stackTraceEnhance) MsgSkip(skip int, msg string) string {
	if len(msg) == 0 {
		return ""
	}

	if strings.Contains(msg, constants.StackTrace) {
		return msg
	}

	toTrace := it.trace(skip)

	if len(toTrace) == 0 {
		return fmt.Sprintf(
			"%s - %s",
			it.methodName(skip+1),
			msg,
		)
	}

	fullMessage := fmt.Sprintf(
		"%s - %s\n %s \n  - %s",
		it.methodName(skip+1),
		msg,
		constants.StackTrace,
		it.trace(skip+1),
	)

	return fullMessage
}

func (it stackTraceEnhance) methodName(skip int) string {
	return reflectinternal.CodeStack.MethodName(1 + skip)
}

func (it stackTraceEnhance) trace(skip int) string {
	return reflectinternal.CodeStack.SingleStack(2 + skip)
}

func (it stackTraceEnhance) MsgErrorSkip(skip int, msg string, err error) string {
	if err == nil {
		return ""
	}

	compiledMsg := fmt.Sprintf(
		"%s %s",
		msg,
		err,
	)

	if strings.Contains(msg, constants.StackTrace) {
		return compiledMsg
	}

	toTrace := it.trace(skip)

	if len(toTrace) == 0 {
		return fmt.Sprintf(
			"%s - %s",
			it.methodName(skip+1),
			msg,
		)
	}

	fullMessage := fmt.Sprintf(
		"%s - %s\n %s \n  - %s",
		it.methodName(skip+1),
		msg,
		constants.StackTrace,
		it.trace(skip+1),
	)

	return fullMessage
}

func (it stackTraceEnhance) MsgErrorToErrSkip(skip int, msg string, err error) error {
	if err == nil {
		return nil
	}

	return errors.New(it.MsgErrorSkip(1+skip, msg, err))
}
