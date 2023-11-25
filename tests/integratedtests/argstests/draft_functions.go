package argstests

import "errors"

func someFunctionV1(arg1, arg2 string) string {
	return arg1 + arg2 + " - final output"
}

func someFunctionV2(arg1, arg2 string) (string, error) {
	return arg1 + arg2 + " - final output", errors.New("some err")
}

func someFunctionV3(arg1, arg2 string) (int, string, error) {
	return 5, arg1 + arg2 + " - final output", errors.New("some err")
}
