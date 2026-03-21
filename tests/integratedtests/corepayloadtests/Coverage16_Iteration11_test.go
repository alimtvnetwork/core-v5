package corepayloadtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

type stringer16 struct{ val string }

func (s stringer16) String() string { return s.val }

// ==========================================================================
// newPayloadWrapperCreator — UsingBytesCreateInstruction
// ==========================================================================

func Test_I11_NewPW_UsingBytesCreateInstruction(t *testing.T) {
	b, _ := corejson.Serialize.Raw("hello")
	pw := corepayload.New.PayloadWrapper.UsingBytesCreateInstruction(
		&corepayload.BytesCreateInstruction{
			Name: "n", Identifier: "id", TaskTypeName: "task",
			EntityType: "entity", CategoryName: "cat",
			HasManyRecords: false, Payloads: b,
		})
	actual := args.Map{"name": pw.Name, "id": pw.Identifier, "entity": pw.EntityType}
	expected := args.Map{"name": "n", "id": "id", "entity": "entity"}
	expected.ShouldBeEqual(t, 0, "UsingBytesCreateInstruction returns correct value -- with args", actual)
}

func Test_I11_NewPW_UsingBytesCreateInstructionTypeStringer(t *testing.T) {
	b, _ := corejson.Serialize.Raw("hello")
	pw := corepayload.New.PayloadWrapper.UsingBytesCreateInstructionTypeStringer(
		&corepayload.BytesCreateInstructionStringer{
			Name: "n", Identifier: "id",
			TaskTypeName: stringer16{"task"},
			CategoryName: stringer16{"cat"},
			EntityType:   "entity", Payloads: b,
		})
	actual := args.Map{"name": pw.Name, "task": pw.TaskTypeName, "cat": pw.CategoryName}
	expected := args.Map{"name": "n", "task": "task", "cat": "cat"}
	expected.ShouldBeEqual(t, 0, "UsingBytesCreateInstructionTypeStringer returns correct value -- with args", actual)
}

// ==========================================================================
// UsingCreateInstructionTypeStringer
// ==========================================================================

func Test_I11_NewPW_UsingCreateInstructionTypeStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.UsingCreateInstructionTypeStringer(
		&corepayload.PayloadCreateInstructionTypeStringer{
			Name: "n", Identifier: "id",
			TaskTypeNameStringer: stringer16{"task"},
			CategoryNameStringer: stringer16{"cat"},
			Payloads:             testUser{Name: "Alice"},
		})
	actual := args.Map{"noErr": err == nil, "name": pw.Name, "task": pw.TaskTypeName}
	expected := args.Map{"noErr": true, "name": "n", "task": "task"}
	expected.ShouldBeEqual(t, 0, "UsingCreateInstructionTypeStringer returns correct value -- with args", actual)
}

// ==========================================================================
// UsingCreateInstruction — string payload branch
// ==========================================================================

func Test_I11_NewPW_UsingCreateInstruction_StringPayload(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.UsingCreateInstruction(
		&corepayload.PayloadCreateInstruction{
			Name: "n", Identifier: "id", TaskTypeName: "task",
			EntityType: "entity", CategoryName: "cat",
			Payloads: `{"Name":"Bob"}`,
		})
	actual := args.Map{"noErr": err == nil, "entity": pw.EntityType}
	expected := args.Map{"noErr": true, "entity": "entity"}
	expected.ShouldBeEqual(t, 0, "UsingCreateInstruction returns correct value -- string payload", actual)
}

// ==========================================================================
// CreateUsingTypeStringer
// ==========================================================================

func Test_I11_NewPW_CreateUsingTypeStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.CreateUsingTypeStringer(
		"n", "id", stringer16{"task"}, stringer16{"cat"},
		testUser{Name: "X"})
	actual := args.Map{"noErr": err == nil, "task": pw.TaskTypeName, "cat": pw.CategoryName}
	expected := args.Map{"noErr": true, "task": "task", "cat": "cat"}
	expected.ShouldBeEqual(t, 0, "CreateUsingTypeStringer returns correct value -- with args", actual)
}

// ==========================================================================
// NameIdCategoryStringer
// ==========================================================================

func Test_I11_NewPW_NameIdCategoryStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdCategoryStringer(
		"n", "id", stringer16{"cat"}, testUser{Name: "Y"})
	actual := args.Map{"noErr": err == nil, "cat": pw.CategoryName}
	expected := args.Map{"noErr": true, "cat": "cat"}
	expected.ShouldBeEqual(t, 0, "NameIdCategoryStringer returns correct value -- with args", actual)
}

// ==========================================================================
// RecordsTypeStringer, RecordTypeStringer
// ==========================================================================

func Test_I11_NewPW_RecordsTypeStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.RecordsTypeStringer(
		"n", "id", stringer16{"task"}, stringer16{"cat"},
		[]testUser{{Name: "A"}, {Name: "B"}})
	actual := args.Map{"noErr": err == nil, "many": pw.HasManyRecords}
	expected := args.Map{"noErr": true, "many": true}
	expected.ShouldBeEqual(t, 0, "RecordsTypeStringer returns correct value -- with args", actual)
}

func Test_I11_NewPW_RecordTypeStringer(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.RecordTypeStringer(
		"n", "id", stringer16{"task"}, stringer16{"cat"},
		testUser{Name: "Z"})
	actual := args.Map{"noErr": err == nil, "task": pw.TaskTypeName}
	expected := args.Map{"noErr": true, "task": "task"}
	expected.ShouldBeEqual(t, 0, "RecordTypeStringer returns correct value -- with args", actual)
}

// ==========================================================================
// NameIdTaskStringerRecord, NameTaskNameRecord
// ==========================================================================

func Test_I11_NewPW_NameIdTaskStringerRecord(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameIdTaskStringerRecord(
		"n", "id", stringer16{"task"}, testUser{Name: "Q"})
	actual := args.Map{"noErr": err == nil, "task": pw.TaskTypeName}
	expected := args.Map{"noErr": true, "task": "task"}
	expected.ShouldBeEqual(t, 0, "NameIdTaskStringerRecord returns error -- with args", actual)
}

func Test_I11_NewPW_NameTaskNameRecord(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.NameTaskNameRecord(
		"id", "task", testUser{Name: "R"})
	actual := args.Map{"noErr": err == nil, "task": pw.TaskTypeName}
	expected := args.Map{"noErr": true, "task": "task"}
	expected.ShouldBeEqual(t, 0, "NameTaskNameRecord returns correct value -- with args", actual)
}

// ==========================================================================
// ManyRecords
// ==========================================================================

func Test_I11_NewPW_ManyRecords(t *testing.T) {
	pw, err := corepayload.New.PayloadWrapper.ManyRecords(
		"n", "id", "task", "cat",
		[]testUser{{Name: "A"}})
	actual := args.Map{"noErr": err == nil, "name": pw.Name}
	expected := args.Map{"noErr": true, "name": "n"}
	expected.ShouldBeEqual(t, 0, "ManyRecords returns correct value -- with args", actual)
}

// ==========================================================================
// PayloadsCollection — DeserializeMust, DeserializeToMany, DeserializeUsingJsonResult
// ==========================================================================

func Test_I11_NewPC_DeserializeMust(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	b, _ := corejson.Serialize.Raw(pc)
	result := corepayload.New.PayloadsCollection.DeserializeMust(b)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewPC.DeserializeMust returns correct value -- with args", actual)
}

func Test_I11_NewPC_DeserializeToMany_Valid(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	b, _ := corejson.Serialize.Raw([]*corepayload.PayloadsCollection{pc})
	result, err := corepayload.New.PayloadsCollection.DeserializeToMany(b)
	actual := args.Map{"noErr": err == nil, "len": len(result)}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "NewPC.DeserializeToMany returns non-empty -- valid", actual)
}

func Test_I11_NewPC_DeserializeUsingJsonResult_Valid(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	b, _ := corejson.Serialize.Raw(pc)
	jr := corejson.NewResult.UsingTypeBytesPtr("test", b)
	result, err := corepayload.New.PayloadsCollection.DeserializeUsingJsonResult(jr)
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "NewPC.DeserializeUsingJsonResult returns non-empty -- valid", actual)
}

// ==========================================================================
// PayloadsCollectionGetters — Dynamic accessors
// ==========================================================================

func makePC3(t *testing.T) *corepayload.PayloadsCollection {
	t.Helper()
	pw1, _ := corepayload.New.PayloadWrapper.NameIdCategory("n1", "1", "cat", "a")
	pw2, _ := corepayload.New.PayloadWrapper.NameIdCategory("n2", "2", "cat", "b")
	pw3, _ := corepayload.New.PayloadWrapper.NameIdCategory("n3", "3", "cat", "c")
	return corepayload.New.PayloadsCollection.UsingWrappers(pw1, pw2, pw3)
}

func Test_I11_PC_FirstDynamic(t *testing.T) {
	pc := makePC3(t)
	actual := args.Map{"notNil": pc.FirstDynamic() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstDynamic returns correct value -- with args", actual)
}

func Test_I11_PC_FirstDynamic_Nil(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	actual := args.Map{"nil": pc.FirstDynamic() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "FirstDynamic returns nil -- nil", actual)
}

func Test_I11_PC_LastDynamic(t *testing.T) {
	pc := makePC3(t)
	actual := args.Map{"notNil": pc.LastDynamic() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LastDynamic returns correct value -- with args", actual)
}

func Test_I11_PC_LastDynamic_Nil(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	actual := args.Map{"nil": pc.LastDynamic() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "LastDynamic returns nil -- nil", actual)
}

func Test_I11_PC_FirstOrDefaultDynamic(t *testing.T) {
	pc := makePC3(t)
	actual := args.Map{"notNil": pc.FirstOrDefaultDynamic() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultDynamic returns correct value -- with args", actual)
}

func Test_I11_PC_LastOrDefaultDynamic(t *testing.T) {
	pc := makePC3(t)
	actual := args.Map{"notNil": pc.LastOrDefaultDynamic() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LastOrDefaultDynamic returns correct value -- with args", actual)
}

// ==========================================================================
// PayloadsCollectionGetters — Slice operations
// ==========================================================================

func Test_I11_PC_SkipDynamic(t *testing.T) {
	pc := makePC3(t)
	result := pc.SkipDynamic(1)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SkipDynamic returns correct value -- with args", actual)
}

func Test_I11_PC_SkipCollection(t *testing.T) {
	pc := makePC3(t)
	result := pc.SkipCollection(1)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SkipCollection returns correct value -- with args", actual)
}

func Test_I11_PC_TakeDynamic(t *testing.T) {
	pc := makePC3(t)
	result := pc.TakeDynamic(2)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TakeDynamic returns correct value -- with args", actual)
}

func Test_I11_PC_TakeCollection(t *testing.T) {
	pc := makePC3(t)
	result := pc.TakeCollection(2)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TakeCollection returns correct value -- with args", actual)
}

func Test_I11_PC_LimitCollection(t *testing.T) {
	pc := makePC3(t)
	result := pc.LimitCollection(1)
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LimitCollection returns correct value -- with args", actual)
}

func Test_I11_PC_LimitDynamic(t *testing.T) {
	pc := makePC3(t)
	result := pc.LimitDynamic(2)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LimitDynamic returns correct value -- with args", actual)
}

func Test_I11_PC_Limit(t *testing.T) {
	pc := makePC3(t)
	result := pc.Limit(2)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Limit returns correct value -- with args", actual)
}

// ==========================================================================
// PayloadsCollectionGetters — IsEqualItems
// ==========================================================================

func Test_I11_PC_IsEqualItems_Same(t *testing.T) {
	pc := makePC3(t)
	actual := args.Map{"val": pc.IsEqualItems(pc.Items[0], pc.Items[1], pc.Items[2])}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEqualItems returns correct value -- same", actual)
}

func Test_I11_PC_IsEqualItems_DiffLen(t *testing.T) {
	pc := makePC3(t)
	actual := args.Map{"val": pc.IsEqualItems(pc.Items[0])}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsEqualItems returns correct value -- diff len", actual)
}

func Test_I11_PC_IsEqualItems_NilPC(t *testing.T) {
	var pc *corepayload.PayloadsCollection
	actual := args.Map{"val": pc.IsEqualItems(nil)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEqualItems returns nil -- nil pc nil items", actual)
}

// ==========================================================================
// PayloadsCollectionPaging — GetPagedCollection, GetSinglePageCollection
// ==========================================================================

func Test_I11_PC_GetPagedCollection(t *testing.T) {
	pc := makePC3(t)
	pages := pc.GetPagedCollection(2)
	actual := args.Map{"pages": len(pages)}
	expected := args.Map{"pages": 2}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns correct value -- with args", actual)
}

func Test_I11_PC_GetPagedCollection_SmallEnough(t *testing.T) {
	pc := makePC3(t)
	pages := pc.GetPagedCollection(10)
	actual := args.Map{"pages": len(pages)}
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns correct value -- small", actual)
}

func Test_I11_PC_GetSinglePageCollection(t *testing.T) {
	pc := makePC3(t)
	page := pc.GetSinglePageCollection(2, 2)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection returns correct value -- page 2", actual)
}

func Test_I11_PC_GetSinglePageCollection_SmallEnough(t *testing.T) {
	pc := makePC3(t)
	page := pc.GetSinglePageCollection(10, 1)
	actual := args.Map{"len": page.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection returns correct value -- small", actual)
}

// ==========================================================================
// PayloadCreateInstructionTypeStringer — PayloadCreateInstruction()
// ==========================================================================

func Test_I11_PayloadCreateInstructionTypeStringer(t *testing.T) {
	pci := corepayload.PayloadCreateInstructionTypeStringer{
		Name: "n", Identifier: "id",
		TaskTypeNameStringer: stringer16{"task"},
		CategoryNameStringer: stringer16{"cat"},
		Payloads:             "data",
	}
	pi := pci.PayloadCreateInstruction()
	actual := args.Map{"name": pi.Name, "task": pi.TaskTypeName, "cat": pi.CategoryName}
	expected := args.Map{"name": "n", "task": "task", "cat": "cat"}
	expected.ShouldBeEqual(t, 0, "PayloadCreateInstructionTypeStringer returns correct value -- with args", actual)
}

// ==========================================================================
// PayloadTypeExpander — struct coverage
// ==========================================================================

func Test_I11_PayloadTypeExpander(t *testing.T) {
	pte := corepayload.PayloadTypeExpander{
		CategoryStringer: stringer16{"cat"},
		TaskTypeStringer: stringer16{"task"},
	}
	actual := args.Map{
		"cat":  pte.CategoryStringer.String(),
		"task": pte.TaskTypeStringer.String(),
	}
	expected := args.Map{"cat": "cat", "task": "task"}
	expected.ShouldBeEqual(t, 0, "PayloadTypeExpander returns correct value -- with args", actual)
}

// ==========================================================================
// BytesCreateInstructionStringer — struct fields
// ==========================================================================

func Test_I11_BytesCreateInstructionStringer(t *testing.T) {
	bci := corepayload.BytesCreateInstructionStringer{
		Name: "n", Identifier: "id",
		TaskTypeName: stringer16{"task"},
		EntityType:   "entity",
		CategoryName: stringer16{"cat"},
		Payloads:     []byte("data"),
	}
	actual := args.Map{"name": bci.Name, "task": bci.TaskTypeName.String(), "cat": fmt.Sprintf("%v", bci.CategoryName)}
	expected := args.Map{"name": "n", "task": "task", "cat": "cat"}
	expected.ShouldBeEqual(t, 0, "BytesCreateInstructionStringer returns correct value -- with args", actual)
}

// ==========================================================================
// CastOrDeserializeFrom — valid path
// ==========================================================================

func Test_I11_NewPW_CastOrDeserializeFrom_Valid(t *testing.T) {
	pw, _ := corepayload.New.PayloadWrapper.NameIdCategory("n", "id", "cat", "data")
	pw2, err := corepayload.New.PayloadWrapper.CastOrDeserializeFrom(pw)
	actual := args.Map{"noErr": err == nil, "name": pw2.Name}
	expected := args.Map{"noErr": true, "name": "n"}
	expected.ShouldBeEqual(t, 0, "CastOrDeserializeFrom returns non-empty -- valid", actual)
}

// ==========================================================================
// DeserializeToMany — valid path for PayloadWrapper
// ==========================================================================

func Test_I11_NewPW_DeserializeToMany_Valid(t *testing.T) {
	pw, _ := corepayload.New.PayloadWrapper.NameIdCategory("n", "id", "cat", "data")
	b, _ := corejson.Serialize.Raw([]*corepayload.PayloadWrapper{pw})
	result, err := corepayload.New.PayloadWrapper.DeserializeToMany(b)
	actual := args.Map{"noErr": err == nil, "len": len(result)}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "NewPW.DeserializeToMany returns non-empty -- valid", actual)
}

// ==========================================================================
// DeserializeToCollection
// ==========================================================================

func Test_I11_NewPW_DeserializeToCollection(t *testing.T) {
	pc := corepayload.New.PayloadsCollection.Empty()
	b, _ := corejson.Serialize.Raw(pc)
	result, err := corepayload.New.PayloadWrapper.DeserializeToCollection(b)
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "DeserializeToCollection returns correct value -- with args", actual)
}

// ==========================================================================
// emptyCreator — remaining methods
// ==========================================================================

func Test_I11_Empty_PayloadWrapper(t *testing.T) {
	pw := corepayload.Empty.PayloadWrapper()
	actual := args.Map{"notNil": pw != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Empty.PayloadWrapper returns empty -- with args", actual)
}

func Test_I11_Empty_PayloadsCollection(t *testing.T) {
	pc := corepayload.Empty.PayloadsCollection()
	actual := args.Map{"len": pc.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty.PayloadsCollection returns empty -- with args", actual)
}
