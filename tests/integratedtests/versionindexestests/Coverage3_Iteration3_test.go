package versionindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/enums/versionindexes"
)

// Test_Cov3_JsonParseSelfInject_HasError covers
// enums/versionindexes/Index.go L197-199: jsonResult.HasError() branch.
func Test_Cov3_JsonParseSelfInject_HasError(t *testing.T) {
	// Arrange
	idx := versionindexes.Major
	errResult := corejson.NewPtr("invalid-data")
	errResult.SetError("simulated parse error")

	// Act
	err := idx.JsonParseSelfInject(errResult)

	// Assert
	coretests.ShouldNotBeNil(t, err)
}
