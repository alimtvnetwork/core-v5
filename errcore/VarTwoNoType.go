package errcore

func VarTwoNoType(
	var1 string,
	val1 any,
	var2 string,
	val2 any,
) string {
	return VarTwo(
		false,
		var1, val1,
		var2, val2)
}
