package coreinstruction

type BaseEnabler struct {
	IsEnabled bool `json:"IsEnabled"`
}

func (b *BaseEnabler) SetEnable() {
	b.IsEnabled = true
}

func (b *BaseEnabler) SetDisable() {
	b.IsEnabled = false
}

func (b *BaseEnabler) SetEnableVal(v bool) {
	b.IsEnabled = v
}
