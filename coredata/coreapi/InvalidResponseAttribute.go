package coreapi

func InvalidResponseAttribute(message string) *ResponseAttribute {
	return &ResponseAttribute{
		IsValid: false,
		Message: message,
	}
}
