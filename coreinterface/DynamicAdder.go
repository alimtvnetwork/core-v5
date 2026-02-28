package coreinterface

type DynamicAdder interface {
	AddDynamic(addingItem any) (isSuccess bool)
	AddDynamicMany(addingItem ...any) (isSuccess bool)
}
