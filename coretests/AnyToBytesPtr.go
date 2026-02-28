package coretests

func AnyToBytesPtr(anyItem any) *[]byte {
	toBytes := AnyToBytes(anyItem)

	return &toBytes
}
