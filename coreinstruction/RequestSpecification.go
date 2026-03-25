package coreinstruction

type RequestSpecification struct {
	BaseIdentifier
	BaseTypeDotFilter
	BaseTags
	BaseIsGlobal
	BaseIsContinueOnError
	BaseIsRunAll
}

func (r RequestSpecification) Clone() *RequestSpecification {
	clonedTags := make([]string, len(r.Tags))
	copy(clonedTags, r.Tags)

	return &RequestSpecification{
		BaseIdentifier: BaseIdentifier{r.Id},
		BaseTypeDotFilter: BaseTypeDotFilter{
			splitDotFilters: nil,
			TypeDotFilter:   r.TypeDotFilter,
		},
		BaseTags: BaseTags{
			Tags: clonedTags,
		},
		BaseIsGlobal:          BaseIsGlobal{r.IsGlobal},
		BaseIsContinueOnError: BaseIsContinueOnError{r.IsContinueOnError},
		BaseIsRunAll:          BaseIsRunAll{r.IsRunAll},
	}
}
