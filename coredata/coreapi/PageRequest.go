package coreapi

type PageRequest struct {
	PageSize  int
	PageIndex int
}

func (p *PageRequest) Clone() *PageRequest {
	if p == nil {
		return nil
	}

	return &PageRequest{
		PageSize:  p.PageSize,
		PageIndex: p.PageIndex,
	}
}
