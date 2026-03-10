package chmodhelper

import "github.com/alimtvnetwork/core/errcore"

func GetRwxLengthError(rwx string) error {
	if len(rwx) != SingleRwxLength {
		return errcore.LengthShouldBeEqualToType.
			Error(
				"rwx length should be "+SingleRwxLengthString,
				len(rwx))
	}

	return nil
}
