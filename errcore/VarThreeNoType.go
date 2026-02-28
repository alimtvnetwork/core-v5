package errcore

func VarThreeNoType(
	var1 string,
	val1 any,
	var2 string,
	val2 any,
	var3 string,
	val3 any,
) string {
	return VarThree(
		false,
		var1, val1,
		var2, val2,
		var3, val3)
}
