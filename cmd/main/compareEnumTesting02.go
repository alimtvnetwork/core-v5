package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/corecomparator"
)

func compareEnumTesting02() {
	eq := corecomparator.Equal

	err := eq.
		OnlySupportedErr(
			"dining doesn't support more",
			corecomparator.Inconclusive,
			corecomparator.NotEqual)

	fmt.Println(err)
}
