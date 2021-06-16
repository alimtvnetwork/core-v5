package coreapi

func InvalidSimpleGenericRequest(attribute *RequestAttribute) *SimpleGenericRequest {
	if attribute == nil {
		return &SimpleGenericRequest{
			Attribute: InvalidRequestAttribute(),
			Request:   nil,
		}
	}

	return &SimpleGenericRequest{
		Attribute: attribute,
		Request:   nil,
	}
}
