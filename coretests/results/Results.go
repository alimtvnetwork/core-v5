package results

import "fmt"

// Results is a two-value typed result for functions returning (T1, T2).
//
// Embeds Result[T1] for the primary value and adds Result2 for the second.
type Results[T1, T2 any] struct {
	Result[T1]
	Result2 T2
}

// String returns a human-readable summary including both values.
func (it Results[T1, T2]) String() string {
	if it.Panicked {
		return fmt.Sprintf(
			"Results{panicked: %v, panicValue: %v}",
			it.Panicked,
			it.PanicValue,
		)
	}

	if it.Error != nil {
		return fmt.Sprintf(
			"Results{value: %v, result2: %v, error: %s}",
			it.Value,
			it.Result2,
			it.Error.Error(),
		)
	}

	return fmt.Sprintf(
		"Results{value: %v, result2: %v}",
		it.Value,
		it.Result2,
	)
}

// IsResult2 checks whether Result2 matches the given expected value.
func (it Results[T1, T2]) IsResult2(expected any) bool {
	return fmt.Sprintf("%v", it.Result2) == fmt.Sprintf("%v", expected)
}

// Result2String returns Result2 formatted via %v.
func (it Results[T1, T2]) Result2String() string {
	return fmt.Sprintf("%v", it.Result2)
}
