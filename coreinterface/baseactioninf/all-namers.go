package baseactioninf

import "github.com/alimtvnetwork/core/coreinterface/enuminf"

type CategoryTypeNamer interface {
	TypeName() enuminf.BasicEnumer
	Category() enuminf.BasicEnumer
}
