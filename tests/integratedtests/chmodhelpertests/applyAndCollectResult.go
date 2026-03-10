package chmodhelpertests

import (
	"fmt"
	"sort"
	"strings"

	"github.com/alimtvnetwork/core/tests/testwrappers/chmodhelpertestwrappers"
)

func applyAndCollectResult(
	wrapper *chmodhelpertestwrappers.RwxInstructionTestWrapper,
	applyFunc func(*chmodhelpertestwrappers.RwxInstructionTestWrapper) error,
) string {
	err := applyFunc(wrapper)
	if err != nil {
		return err.Error()
	}

	expected := wrapper.ExpectedAsRwxOwnerGroupOtherInstruction()
	expectedChmod := expected.String()

	var mismatches []string

	for _, createPath := range wrapper.CreatePaths {
		fileChmodMap := createPath.GetFilesChmodMap()

		for filePath, chmodValue := range fileChmodMap.Items() {
			if chmodValue != expectedChmod {
				mismatches = append(mismatches, fmt.Sprintf(
					"%s: expect %s got %s",
					filePath, expectedChmod, chmodValue,
				))
			}
		}
	}

	if len(mismatches) == 0 {
		return ""
	}

	sort.Strings(mismatches)

	return strings.Join(mismatches, "\n")
}
