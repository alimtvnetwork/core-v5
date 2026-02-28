package errcore

func VarThreeNoType(
	firstName string,
	firstValue any,
	secondName string,
	secondValue any,
	thirdName string,
	thirdValue any,
) string {
	return VarThree(
		false,
		firstName, firstValue,
		secondName, secondValue,
		thirdName, thirdValue)
}
