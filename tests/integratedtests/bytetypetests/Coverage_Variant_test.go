package bytetypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/bytetype"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

func Test_Cov_Variant_UnmarshalJSON_Error(t *testing.T) {
	v := new(bytetype.Variant)
	err := v.UnmarshalJSON([]byte("invalid"))
	if err == nil {
		t.Error("expected error")
	}
}

func Test_Cov_Variant_UnmarshallToValue(t *testing.T) {
	v := bytetype.Variant(1)
	jsonBytes, _ := corejson.Serialize.Raw(v)
	val, err := v.UnmarshallToValue(jsonBytes)
	if err != nil || val != 1 {
		t.Error("expected 1")
	}
}
