package coreapi

func InvalidRequestAttribute() *RequestAttribute {
	return &RequestAttribute{
		IsValid: false,
	}
}
