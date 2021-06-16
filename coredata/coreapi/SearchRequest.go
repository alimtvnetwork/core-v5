package coreapi

type SearchRequest struct {
	SearchTerm                                                     string
	IsNaturalSearch, IsContains, IsStartsWith, IsEndsWith, IsRegex bool
}

func (receiver *SearchRequest) Clone() *SearchRequest {
	if receiver == nil {
		return nil
	}

	return &SearchRequest{
		SearchTerm:      receiver.SearchTerm,
		IsNaturalSearch: receiver.IsNaturalSearch,
		IsContains:      receiver.IsContains,
		IsStartsWith:    receiver.IsStartsWith,
		IsEndsWith:      receiver.IsEndsWith,
		IsRegex:         receiver.IsRegex,
	}
}
