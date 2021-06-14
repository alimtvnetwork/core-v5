package coreinstruction

type Specification struct {
	BaseIdDisplayType
	BaseTags
	BaseIsGlobal
}

func (r Specification) ClonePtr() *Specification {
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
