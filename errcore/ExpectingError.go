package errcore

import (
	"errors"
	"gitlab.com/auk-go/core/constants"
)

func ExpectingErrorSimpleNoType(
	title,
	wasExpecting,
	actual interface{},
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
	actual interface{},
) error {
	msg := ExpectingSimpleNoType(
		title,
		wasExpecting,
		actual) +
		constants.NewLineUnix

	return errors.New(msg)
}
