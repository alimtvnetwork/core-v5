package coreinterface

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corejson"
)

type CoreDefiner interface {
	corejson.Jsoner
	AllSerializer

	fmt.Stringer
}
