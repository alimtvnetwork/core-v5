package corepayload

type PayloadCreateInstruction struct {
	Name, Identifier string
	TaskTypeName     string
	EntityType       string // for any type no need to entity type it will be collected by reflection.
	CategoryName     string
	HasManyRecords   bool
	Payloads         any // for any type no need to entity type it will be collected by reflection.
	Attributes       *Attributes
}
