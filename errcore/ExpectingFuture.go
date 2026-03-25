package errcore

// ExpectingFuture
//
// returns ExpectingRecord which will print
//
//	"%s - expecting (type:[%T]) : [\"%v\"], but received or actual (type:[%T]) : [\"%v\"]"
func ExpectingFuture(title string, wasExpecting any) *ExpectingRecord {
	return &ExpectingRecord{
		ExpectingTitle: title,
		WasExpecting:   wasExpecting,
	}
}
