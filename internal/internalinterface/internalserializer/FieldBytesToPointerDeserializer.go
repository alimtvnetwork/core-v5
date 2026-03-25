package internalserializer

type FieldBytesToPointerDeserializer interface {
	Deserialize(toPtr any) error
}
