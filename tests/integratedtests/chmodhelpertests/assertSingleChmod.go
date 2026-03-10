package chmodhelpertests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/errcore"
)

// assertSingleChmod , expectedChmodRwxFullString 10 chars "-rwxrwxrwx"
func assertSingleChmod(
	t *testing.T,
	testHeader string,
	createPath chmodhelper.DirFilesWithRwxPermission,
	expectedChmodRwxFullString string,
) {
	fileChmodMap := createPath.GetFilesChmodMap()
	caseIndex := 0

	for filePath, chmodValueString := range fileChmodMap.Items() {
		tc := coretestcases.CaseV1{
			Title:         testHeader,
			ExpectedInput: []string{expectedChmodRwxFullString},
		}

		if chmodValueString != expectedChmodRwxFullString {
			fmt.Println(
				errcore.Expecting(
					filePath,
					expectedChmodRwxFullString,
					chmodValueString,
				),
			)
		}

		tc.ShouldBeEqual(t, caseIndex, chmodValueString)
		caseIndex++
	}
}
