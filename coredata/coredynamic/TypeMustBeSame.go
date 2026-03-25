package coredynamic

func TypeMustBeSame(
	left, right any,
) {
	err := TypeNotEqualErr(left, right)

	if err != nil {
		panic(err)
	}
}
