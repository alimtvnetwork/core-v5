package serializerinf

type FieldBytesToPointerDeserializer interface {
	Deserialize(toPtr any) error
}
