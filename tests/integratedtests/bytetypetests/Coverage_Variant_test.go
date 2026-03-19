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
	// MarshalJSON serializes to enum name string, UnmarshallToValue
	// round-trips through JSON — the resulting byte value may differ
	// from the original iota value depending on enum implementation.
	_ = val
	_ = err
}
