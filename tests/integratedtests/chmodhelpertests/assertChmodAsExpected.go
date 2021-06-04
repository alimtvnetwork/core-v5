package chmodhelpertests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func assertChmodAsExpected(
	t *testing.T,
	testCase *chmodhelpertestwrappers.RwxInstructionTestWrapper,
	testHeader string,
) {
	expected := testCase.ExpectedAsRwxOwnerGroupOtherInstruction()
	expectedChmod := expected.String()

	for _, createPath := range testCase.CreatePaths {
		fileChmodMap := createPath.GetFilesChmodMap()
		for filePath, chmodValueString := range *fileChmodMap.Items() {
			Convey(testHeader, t, func() {
				isEqual := chmodValueString == expectedChmod

				if !isEqual {
					fmt.Println(
						msgtype.Expecting(
							filePath,
							expectedChmod,
							chmodValueString))

				}

				So(isEqual, ShouldBeTrue)
			})
		}
	}
}
