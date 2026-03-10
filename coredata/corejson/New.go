package corejson

import (
	"encoding/json"

	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

func New(anyItem any) Result {
	jsonBytes, err := json.Marshal(anyItem)
	typeName := reflectinternal.TypeName(anyItem)

	if err != nil {
		nextErr := errcore.
			MarshallingFailedType.
			Error(
				err.Error(),
				typeName)

		return Result{
			Bytes:    jsonBytes,
			Error:    nextErr,
			TypeName: typeName,
		}
	}

	return Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}
