package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/tests/testwrappers/chmodhelpertestwrappers"
)

func applyPathInstructions(
	testCase *chmodhelpertestwrappers.RwxInstructionTestWrapper,
) error {
	executors, err := chmodhelper.ParseRwxInstructionsToExecutors(
		testCase.RwxInstructions,
	)

	errcore.SimpleHandleErr(err, "applyPathInstructions")

	for _, createPath := range testCase.CreatePaths {
		applyErr := executors.ApplyOnPaths(createPath.GetPaths())

		if applyErr != nil {
			return applyErr
		}
	}

	return nil
}
