package versionindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/enums/versionindexes"
)

func Test_QW_Index_JsonParseSelfInject_NilResult(t *testing.T) {
	idx := versionindexes.Major
	err := idx.JsonParseSelfInject(nil)
	if err == nil {
		t.Fatal("expected error for nil json result")
	}
}

func Test_QW_Index_JsonParseSelfInject_ErrorResult(t *testing.T) {
	idx := versionindexes.Major
	bad := corejson.NewResult.UsingString(`invalid`)
	err := idx.JsonParseSelfInject(bad)
	if err == nil {
		t.Fatal("expected error for invalid json")
	}
}
