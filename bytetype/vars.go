package bytetype

import (
	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

var (
	BasicEnumImpl = enumimpl.New.BasicByte.CreateUsingMap(
		reflectinternal.TypeName(Variant(0)),
		map[byte]string{
			Zero.Value():  "Zero",
			Min.Value():   "Min",
			One.Value():   "One",
			Two.Value():   "Two",
			Three.Value(): "Three",
			Max.Value():   "Max",
		})
)
