package chmodhelper

import "gitlab.com/evatix-go/core/msgtype"

func GetRwxLengthError(rwx string) error {
	if len(rwx) != SingleRwxLength {
		return msgtype.LengthShouldBeEqualToMessage.
			Error(
				"rwx length should be "+SingleRwxLengthString,
				len(rwx))
	}

	return nil
}
