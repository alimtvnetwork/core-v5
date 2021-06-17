package coreinstruction

import "gitlab.com/evatix-go/core/coredata/stringslice"

type FlatSpecification struct {
	Id       string    `json:"Id"`
	Display  string    `json:"Display"`
	Type     string    `json:"Type"`
	IsGlobal bool      `json:"IsGlobal"`
	Tags     *[]string `json:"Tags,omitempty"`
	IsValid  bool      `json:"IsValid,omitempty"`
	spec     *Specification
}

func InvalidFlatSpecification() *FlatSpecification {
	return &FlatSpecification{
		Id:       "",
		Display:  "",
		Type:     "",
		IsGlobal: false,
		Tags:     &[]string{},
		IsValid:  false,
	}
}

func NewFlatSpecificationUsingSpec(spec *Specification, isValid bool) *FlatSpecification {
	return &FlatSpecification{
		Id:       spec.Id,
		Display:  spec.Display,
		Type:     spec.Type,
		IsGlobal: spec.IsGlobal,
		Tags:     spec.Tags,
		IsValid:  isValid,
		spec:     spec,
	}
}

func (receiver *FlatSpecification) BaseIdentifier() BaseIdentifier {
	return receiver.Spec().BaseIdentifier
}

func (receiver *FlatSpecification) BaseTags() BaseTags {
	return receiver.Spec().BaseTags
}

func (receiver *FlatSpecification) BaseIsGlobal() BaseIsGlobal {
	return receiver.Spec().BaseIsGlobal
}

func (receiver *FlatSpecification) BaseDisplay() BaseDisplay {
	return receiver.Spec().BaseDisplay
}

func (receiver *FlatSpecification) BaseType() BaseType {
	return receiver.Spec().BaseType
}

func (receiver *FlatSpecification) Spec() *Specification {
	if receiver == nil {
		return nil
	}

	if receiver.spec != nil {
		return receiver.spec
	}

	receiver.spec = &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{Id: receiver.Id},
			BaseDisplay:    BaseDisplay{receiver.Display},
			BaseType:       BaseType{receiver.Type},
		},
		BaseTags: BaseTags{
			Tags: receiver.Tags,
		},
		BaseIsGlobal: BaseIsGlobal{IsGlobal: receiver.IsGlobal},
	}

	return receiver.spec
}

func (receiver *FlatSpecification) Clone() *FlatSpecification {
	if receiver == nil {
		return nil
	}

	return &FlatSpecification{
		Id:       receiver.Id,
		Display:  receiver.Display,
		Type:     receiver.Type,
		IsGlobal: receiver.IsGlobal,
		Tags:     stringslice.ClonePtr(receiver.Tags),
		IsValid:  receiver.IsValid,
	}
}
