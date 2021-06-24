package chmodhelper

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

// RwxPartialApplyChmod
// rwxPartial can be any length in
// between 0-10 (rest will be fixed by wildcard)
//
// rwxPartial:
//  - "-rwx" will be "-rwx******"
//  - "-rwxr-x" will be "-rwxr-x***"
//  - "-rwxr-x" will be "-rwxr-x***"
func RwxPartialApplyChmod(
	rwxPartial string,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	rwxInstructionExecutor, err := RwxPartialToInstructionExecutor(
		rwxPartial,
		condition)

	if err != nil {
		return err
	}

	return rwxInstructionExecutor.
		ApplyOnPathsPtr(&locations)
}
