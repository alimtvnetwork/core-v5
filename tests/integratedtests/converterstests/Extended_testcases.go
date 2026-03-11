package converterstests

type stringToIntegerWithDefaultCase struct {
	name        string
	input       string
	defaultVal  int
	expectedVal int
	expectedOk  bool
}

var stringToIntegerWithDefaultCases = []stringToIntegerWithDefaultCase{
	{"valid integer", "42", 0, 42, true},
	{"empty string", "", 99, 99, false},
	{"non-numeric", "abc", 99, 99, false},
	{"zero", "0", 5, 0, true},
}

type stringToByteCase struct {
	name      string
	input     string
	expected  byte
	expectErr bool
}

var stringToByteCases = []stringToByteCase{
	{"zero", "0", 0, false},
	{"one", "1", 1, false},
	{"valid", "100", 100, false},
	{"max byte", "255", 255, false},
	{"over 255", "256", 0, true},
	{"negative", "-1", 0, true},
	{"empty", "", 0, true},
	{"non-numeric", "abc", 0, true},
}
