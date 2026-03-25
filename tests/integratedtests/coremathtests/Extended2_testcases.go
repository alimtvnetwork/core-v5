package coremathtests

type byteCase struct {
	name     string
	left     byte
	right    byte
	expected byte
}

var maxByteCases = []byteCase{
	{"left larger", 10, 5, 10},
	{"right larger", 3, 8, 8},
	{"equal", 7, 7, 7},
	{"zero and max", 0, 255, 255},
}

var minByteCases = []byteCase{
	{"left smaller", 3, 8, 3},
	{"right smaller", 10, 5, 5},
	{"equal", 7, 7, 7},
	{"zero and max", 0, 255, 0},
}
