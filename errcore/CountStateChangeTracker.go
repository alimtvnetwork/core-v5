package errcore

type CountStateChangeTracker struct {
	lengthGetter
	initLength int
}

func NewCountStateChangeTracker(
	lengthGetter lengthGetter,
) CountStateChangeTracker {
	return CountStateChangeTracker{
		lengthGetter: lengthGetter,
		initLength:   lengthGetter.Length(),
	}
}

func (it CountStateChangeTracker) IsSameStateUsingCount(
	currentCount int,
) bool {
	return currentCount == it.initLength
}

// IsSameState returns true if the current length matches the initial length.
// IsValid and IsSuccess are aliases with identical behavior.
func (it CountStateChangeTracker) IsSameState() bool {
	return it.lengthGetter.Length() == it.initLength
}

// IsValid is an alias for IsSameState.
func (it CountStateChangeTracker) IsValid() bool {
	return it.IsSameState()
}

// IsSuccess is an alias for IsSameState.
func (it CountStateChangeTracker) IsSuccess() bool {
	return it.IsSameState()
}

// HasChanges returns true if the current length differs from the initial length.
// IsFailed is an alias with identical behavior.
func (it CountStateChangeTracker) HasChanges() bool {
	return it.lengthGetter.Length() != it.initLength
}

// IsFailed is an alias for HasChanges.
func (it CountStateChangeTracker) IsFailed() bool {
	return it.HasChanges()
}
