package main

import (
	"fmt"

	"gitlab.com/auk-go/core/corecmp"
	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/coreversion"
	"gitlab.com/auk-go/core/enums/versionindexes"
)

func versionCompareTest(leftVersion, rightVersion string) corecomparator.Compare {
	fmt.Println("left, right = ", leftVersion, rightVersion)
	leftV := coreversion.New.Create(leftVersion)
	rightV := coreversion.New.Create(rightVersion)

	fmt.Println("   left, right = ", leftV, rightV)
	compareResult := leftV.Compare(&rightV)
	indexCompareResult := leftV.ComparisonValueIndexes(
		&rightV,
		versionindexes.AllVersionIndexes...,
	)
	leftVersionValues := leftV.AllVersionValues()
	rightVersionValues := rightV.AllVersionValues()

	fmt.Println("   (compareResult) left, right = ", compareResult)
	fmt.Println("   (indexCompareResult) left, right = ", indexCompareResult)
	fmt.Println("   (Values) left, right = ", leftVersionValues, rightVersionValues)

	sliceCompareResult := corecmp.VersionSliceInteger(
		leftVersionValues,
		rightVersionValues,
	)

	fmt.Println("   (sliceCompareResult) left, right = ", sliceCompareResult)

	return compareResult
}
