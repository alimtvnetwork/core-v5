package coreinstruction

type Specification struct {
	BaseIdDisplayType
	BaseTags
	BaseIsGlobal
	flatSpec *FlatSpecification
}

func NewSpecificationSimple(
	id,
	display,
	typeName string,
) *Specification {
	return &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{Id: id},
			BaseDisplay:    BaseDisplay{display},
			BaseType:       BaseType{typeName},
		},
		BaseTags:     *NewTags(nil),
		BaseIsGlobal: BaseIsGlobal{false},
	}
}

func NewSpecificationSimpleGlobal(
	id,
	display,
	typeName string,
) *Specification {
	return &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{Id: id},
			BaseDisplay:    BaseDisplay{display},
			BaseType:       BaseType{typeName},
		},
		BaseTags:     *NewTags(nil),
		BaseIsGlobal: BaseIsGlobal{true},
	}
}

func NewSpecification(
	id,
	display,
	typeName string,
	tags []string,
	isGlobal bool,
) *Specification {
	return &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{Id: id},
			BaseDisplay:    BaseDisplay{display},
			BaseType:       BaseType{typeName},
		},
		BaseTags:     *NewTags(tags),
		BaseIsGlobal: BaseIsGlobal{isGlobal},
	}
}

func (r *Specification) Clone() *Specification {
	if r == nil {
		return nil
	}

	clonedTags := make([]string, len(r.Tags))
	copy(clonedTags, r.Tags)

	return &Specification{
		BaseIdDisplayType: BaseIdDisplayType{
			BaseIdentifier: BaseIdentifier{r.Id},
			BaseDisplay:    BaseDisplay{r.Display},
			BaseType:       BaseType{r.Type},
		},
		BaseTags: BaseTags{
			Tags: clonedTags,
		},
		BaseIsGlobal: BaseIsGlobal{r.IsGlobal},
	}
}

func (r *Specification) FlatSpecification() *FlatSpecification {
	if r == nil {
		return nil
	}

	if r.flatSpec != nil {
		return r.flatSpec
	}

	r.flatSpec = NewFlatSpecificationUsingSpec(
		r,
		true)

	return r.flatSpec
}
