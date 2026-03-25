package coreinterface

import (
	"fmt"

	"github.com/alimtvnetwork/core/coredata/corejson"
)

type CoreDefiner interface {
	corejson.Jsoner
	AllSerializer

	fmt.Stringer
}
