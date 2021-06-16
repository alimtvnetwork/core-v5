package coreinstruction

type Specification struct {
	BaseIdDisplayType
	BaseTags
	BaseIsGlobal
	flatSpec *FlatSpecification
}

func (r *Specification) Clone() *Specification {
	return &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{r.Id},
			BaseDisplay:    BaseDisplay{r.Display},
			BaseType:       BaseType{r.Type},
		},
		BaseTags: BaseTags{
			Tags: r.Tags,
		},
		BaseIsGlobal: BaseIsGlobal{r.IsGlobal},
	}
}

func (r *Specification) FlatSpecification() *FlatSpecification {
	if r.flatSpec != nil {
		return r.flatSpec
	}

	r.flatSpec = NewFlatSpecificationUsingSpec(
		r,
		true)

	return r.flatSpec
}
