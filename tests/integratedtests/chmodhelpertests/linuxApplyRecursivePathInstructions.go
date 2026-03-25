package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/tests/testwrappers/chmodhelpertestwrappers"
)

func linuxApplyRecursivePathInstructions(
	testCase *chmodhelpertestwrappers.RwxInstructionTestWrapper,
) error {
	executors, err := chmodhelper.ParseRwxInstructionsToExecutors(
		testCase.RwxInstructions)

	errcore.SimpleHandleErr(err, "linuxApplyRecursivePathInstructions")

	for _, createPath := range testCase.CreatePaths {
		applyErr := executors.ApplyOnPath(createPath.Dir)
		if applyErr != nil {
			return applyErr
		}
	}

	return nil
}
