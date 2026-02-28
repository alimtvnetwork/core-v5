package conditional

func NilCheck(
	canBeEmpty any,
	onNil any,
	onNonNil any,
) any {
	if canBeEmpty == nil {
		return onNil
	}

	return onNonNil
}
