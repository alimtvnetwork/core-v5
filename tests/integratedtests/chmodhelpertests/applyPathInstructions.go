package chmodhelpertests

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

func applyPathInstructions(
	testCase *chmodhelpertestwrappers.RwxInstructionTestWrapper,
) error {
	executors, err := chmodhelper.ParseRwxInstructionsToExecutors(
		&testCase.RwxInstructions)

	msgtype.SimpleHandleErr(err, "applyPathInstructions")

	for _, createPath := range testCase.CreatePaths {
		err2 := executors.ApplyOnPaths(createPath.GetPaths())

		if err2 != nil {
			return err2
		}
	}

	return nil
}

func unixApplyRecursivePathInstructions(
	testCase *chmodhelpertestwrappers.RwxInstructionTestWrapper,
) error {
	executors, err := chmodhelper.ParseRwxInstructionsToExecutors(
		&testCase.RwxInstructions)

	msgtype.SimpleHandleErr(err, "applyPathInstructions")

	for _, createPath := range testCase.CreatePaths {
		err2 := executors.ApplyOnPath(createPath.Dir)
		if err2 != nil {
			return err2
		}
	}

	return nil
}
