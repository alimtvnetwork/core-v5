package errcore

import (
	"errors"
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func ExpectingErrorSimpleNoType(
	title,
	wasExpecting,
	actual any,
) error {
	msg := ExpectingSimpleNoType(
		title,
		wasExpecting,
		actual)

	return errors.New(msg)
}

func ExpectingErrorSimpleNoTypeNewLineEnds(
	title,
	wasExpecting,
	actual any,
) error {
	msg := ExpectingSimpleNoType(
		title,
		wasExpecting,
		actual) +
		constants.NewLineUnix

	return errors.New(msg)
}

func WasExpectingErrorF(
	wasExpecting,
	actual any,
	titleFormat string,
	v ...any,
) error {
	title := fmt.Sprintf(
		titleFormat,
		v...)

	msg := ExpectingSimpleNoType(
		title,
		wasExpecting,
		actual)

	return errors.New(msg)
}
