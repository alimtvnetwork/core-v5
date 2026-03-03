package coretestcases



// IsFailedToMatch returns the inverse of IsMatching.
//
// Use when validating that a mismatch is expected.
func (it *GenericGherkins[TInput, TExpect]) IsFailedToMatch() bool {
	return !it.IsMatching
}

// HasExtraArgs returns true if ExtraArgs is defined and non-empty.
func (it *GenericGherkins[TInput, TExpect]) HasExtraArgs() bool {
	return it != nil && len(it.ExtraArgs) > 0
}

// GetExtra returns the value for a key in ExtraArgs, or nil if not found.
func (it *GenericGherkins[TInput, TExpect]) GetExtra(key string) any {
	if it == nil || it.ExtraArgs == nil {
		return nil
	}

	return it.ExtraArgs[key]
}

// GetExtraAsString returns the string value for a key in ExtraArgs.
func (it *GenericGherkins[TInput, TExpect]) GetExtraAsString(key string) (string, bool) {
	if it == nil || it.ExtraArgs == nil {
		return "", false
	}

	return it.ExtraArgs.GetAsString(key)
}

// GetExtraAsBool returns the bool value for a key in ExtraArgs.
func (it *GenericGherkins[TInput, TExpect]) GetExtraAsBool(key string) (value bool, isValid bool) {
	if it == nil || it.ExtraArgs == nil {
		return false, false
	}

	val, has := it.ExtraArgs.Get(key)
	isNotFound := !has

	if isNotFound {
		return false, false
	}

	b, ok := val.(bool)

	return b, ok
}
