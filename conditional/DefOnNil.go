package conditional

func DefOnNil(
	canBeEmpty any,
	onNonNil any,
) any {
	if canBeEmpty == nil {
		return onNonNil
	}

	return canBeEmpty
}
