package chmodhelper

import (
	"errors"
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

// VerifyChmod - expectedHyphenedRwx should be 10 chars example "-rwxrwxrwx"
func VerifyChmod(location string, expectedHyphenedRwx string) error {
	if len(expectedHyphenedRwx) != HyphenedRwxLength {
		return msgtype.MeaningFulError(
			msgtype.LengthShouldBeEqualToMessage,
			"VerifyChmod"+constants.HypenAngelRight+location,
			hyphenedRwxLengthErr)
	}

	fileInfo, err := os.Stat(location)

	if os.IsNotExist(err) || fileInfo == nil {
		return msgtype.MeaningFulError(
			msgtype.PathInvalidErrorMessage,
			"VerifyChmod"+constants.HypenAngelRight+location,
			err)
	}

	existingFileMode := fileInfo.Mode().String()
	if existingFileMode == expectedHyphenedRwx {
		return nil
	}

	expectationFailedMessage := msgtype.Expecting(
		chmod,
		expectedHyphenedRwx,
		existingFileMode)

	return msgtype.MeaningFulError(
		msgtype.PathChmodMismatchErrorMessage,
		"VerifyChmod"+constants.HypenAngelRight+location,
		errors.New(expectationFailedMessage))
}
