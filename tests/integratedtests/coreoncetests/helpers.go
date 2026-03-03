package coreoncetests

// callPanics returns true if fn panics, false otherwise.
func callPanics(fn func()) bool {
	didPanic := true

	func() {
		defer func() { recover() }()
		fn()
		didPanic = false
	}()

	return didPanic
}
