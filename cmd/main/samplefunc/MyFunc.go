package samplefunc

import (
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
)

type AlimStruct struct {
	First     string
	LeftRight args.LeftRight
	Draft     *coretests.DraftType
}

func MyFunc(
	x int,
	arg1, arg2 string,
	alim *AlimStruct,
	alim2 *[]*AlimStruct,
) (r1 string, r2 int, r3 *[]**AlimStruct) {
	toAlim := &AlimStruct{
		First:     "someName - " + alim.First + (*alim2)[0].First,
		LeftRight: alim.LeftRight,
		Draft:     (*alim2)[0].Draft,
	}

	return arg1 + " " + arg2 + "-> Processed", x + 1, &[]**AlimStruct{&toAlim}
}
