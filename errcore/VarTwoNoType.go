package errcore

func VarTwoNoType(
	firstName string,
	firstValue any,
	secondName string,
	secondValue any,
) string {
	return VarTwo(
		false,
		firstName, firstValue,
		secondName, secondValue)
}
