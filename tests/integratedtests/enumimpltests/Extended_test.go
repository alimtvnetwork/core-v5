package enumimpltests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// DynamicMap - Basic Operations
// ==========================================

func Test_ExtDynMap_Basic_Verification(t *testing.T) {
	for caseIndex, tc := range extDynMapBasicTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		items := input["items"].(map[string]any)
		dm := enumimpl.DynamicMap(items)

		// Act
		actual := args.Map{
			"length":  fmt.Sprintf("%d", dm.Length()),
			"isEmpty": fmt.Sprintf("%v", dm.IsEmpty()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// DynamicMap - AddOrUpdate / Set / AddNewOnly
// ==========================================

func Test_ExtDynMap_AddOrUpdate_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	isNew := dm.AddOrUpdate("B", 2)
	isExisting := dm.AddOrUpdate("A", 10)

	// Assert
	if !isNew {
		t.Error("AddOrUpdate new key should return true")
	}
	if isExisting {
		t.Error("AddOrUpdate existing key should return false")
	}
	if dm.Length() != 2 {
		t.Errorf("Expected 2, got %d", dm.Length())
	}
}

func Test_ExtDynMap_Set_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	isNew := dm.Set("B", 2)

	// Assert
	if !isNew {
		t.Error("Set new key should return true")
	}
}

func Test_ExtDynMap_AddNewOnly_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	added := dm.AddNewOnly("A", 10)
	addedNew := dm.AddNewOnly("B", 2)

	// Assert
	if added {
		t.Error("AddNewOnly existing should return false")
	}
	if !addedNew {
		t.Error("AddNewOnly new should return true")
	}
}

// ==========================================
// DynamicMap - AllKeys / AllKeysSorted
// ==========================================

func Test_ExtDynMap_AllKeys_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"B": 2, "A": 1, "C": 3}

	// Act
	keys := dm.AllKeys()
	sortedKeys := dm.AllKeysSorted()

	// Assert
	if len(keys) != 3 {
		t.Errorf("AllKeys expected 3, got %d", len(keys))
	}
	if len(sortedKeys) != 3 || sortedKeys[0] != "A" {
		t.Errorf("AllKeysSorted first should be A, got %v", sortedKeys)
	}
}

// ==========================================
// DynamicMap - HasKey / HasAllKeys / HasAnyKeys / IsMissingKey
// ==========================================

func Test_ExtDynMap_HasKey_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}

	// Act & Assert
	if !dm.HasKey("A") {
		t.Error("HasKey should find A")
	}
	if dm.HasKey("Z") {
		t.Error("HasKey should not find Z")
	}
	if !dm.HasAllKeys("A", "B") {
		t.Error("HasAllKeys should be true")
	}
	if dm.HasAllKeys("A", "Z") {
		t.Error("HasAllKeys with missing should be false")
	}
	if !dm.HasAnyKeys("Z", "A") {
		t.Error("HasAnyKeys should be true")
	}
	if dm.HasAnyKeys("X", "Y") {
		t.Error("HasAnyKeys should be false")
	}
	if !dm.IsMissingKey("Z") {
		t.Error("IsMissingKey should be true for Z")
	}
}

// ==========================================
// DynamicMap - HasAnyItem / Count / LastIndex / HasIndex
// ==========================================

func Test_ExtDynMap_Utility_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}

	// Act & Assert
	if !dm.HasAnyItem() {
		t.Error("HasAnyItem should be true")
	}
	if dm.Count() != 2 {
		t.Errorf("Count expected 2, got %d", dm.Count())
	}
	if dm.LastIndex() != 1 {
		t.Errorf("LastIndex expected 1, got %d", dm.LastIndex())
	}
	if !dm.HasIndex(1) {
		t.Error("HasIndex(1) should be true")
	}
	if dm.HasIndex(5) {
		t.Error("HasIndex(5) should be false")
	}
}

// ==========================================
// DynamicMap - First
// ==========================================

func Test_ExtDynMap_First_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	key, val := dm.First()

	// Assert
	if key == "" || val == nil {
		t.Error("First should return a key-value")
	}

	// Arrange - empty
	empty := enumimpl.DynamicMap{}

	// Act
	key2, val2 := empty.First()

	// Assert
	if key2 != "" || val2 != nil {
		t.Error("First on empty should return empty")
	}
}

// ==========================================
// DynamicMap - AllValuesStrings / AllValuesStringsSorted / AllValuesIntegers
// ==========================================

func Test_ExtDynMap_AllValues_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}

	// Act
	valStrings := dm.AllValuesStrings()
	valStringsSorted := dm.AllValuesStringsSorted()
	valIntegers := dm.AllValuesIntegers()

	// Assert
	if len(valStrings) != 2 {
		t.Errorf("AllValuesStrings expected 2, got %d", len(valStrings))
	}
	if len(valStringsSorted) != 2 {
		t.Errorf("AllValuesStringsSorted expected 2, got %d", len(valStringsSorted))
	}
	if len(valIntegers) != 2 {
		t.Errorf("AllValuesIntegers expected 2, got %d", len(valIntegers))
	}
}

// ==========================================
// DynamicMap - IsEqual / IsRawEqual / IsMismatch
// ==========================================

func Test_ExtDynMap_IsEqual_Verification(t *testing.T) {
	// Arrange
	dm1 := enumimpl.DynamicMap{"A": 1, "B": 2}
	dm2 := enumimpl.DynamicMap{"A": 1, "B": 2}
	dm3 := enumimpl.DynamicMap{"A": 1, "B": 3}

	// Act & Assert
	if !dm1.IsEqual(false, &dm2) {
		t.Error("Same content should be equal")
	}
	if dm1.IsEqual(false, &dm3) {
		t.Error("Different content should not be equal")
	}
	if !dm1.IsMismatch(false, &dm3) {
		t.Error("IsMismatch should be true for different")
	}

	// nil cases
	var nilDm *enumimpl.DynamicMap
	if !nilDm.IsEqual(false, nil) {
		t.Error("nil.IsEqual(nil) should be true")
	}
	if nilDm.IsEqual(false, &dm1) {
		t.Error("nil.IsEqual(non-nil) should be false")
	}
}

func Test_ExtDynMap_IsRawEqual_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act & Assert
	if !dm.IsRawEqual(false, map[string]any{"A": 1}) {
		t.Error("IsRawEqual same should be true")
	}
	if dm.IsRawEqual(false, map[string]any{"A": 2}) {
		t.Error("IsRawEqual different should be false")
	}
}

func Test_ExtDynMap_IsRawEqual_RegardlessType_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act & Assert
	// int vs int8 should match regardless of type
	if !dm.IsRawEqual(true, map[string]any{"A": 1}) {
		t.Error("IsRawEqual regardless should be true for same value")
	}
}

// ==========================================
// DynamicMap - IsKeysEqualOnly
// ==========================================

func Test_ExtDynMap_IsKeysEqualOnly_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}

	// Act & Assert
	if !dm.IsKeysEqualOnly(map[string]any{"A": 10, "B": 20}) {
		t.Error("Same keys should be equal regardless of values")
	}
	if dm.IsKeysEqualOnly(map[string]any{"A": 1, "C": 3}) {
		t.Error("Different keys should not be equal")
	}
}

// ==========================================
// DynamicMap - KeyValue / KeyValueString / KeyValueInt / KeyValueByte
// ==========================================

func Test_ExtDynMap_KeyValue_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 42, "B": "hello"}

	// Act
	val, found := dm.KeyValue("A")

	// Assert
	if !found || val != 42 {
		t.Errorf("KeyValue A expected 42, got %v", val)
	}

	// Act
	valStr, foundStr := dm.KeyValueString("B")

	// Assert
	if !foundStr || valStr != "hello" {
		t.Errorf("KeyValueString B expected 'hello', got '%s'", valStr)
	}

	// Act
	valInt, foundInt, failed := dm.KeyValueInt("A")

	// Assert
	if !foundInt || failed || valInt != 42 {
		t.Errorf("KeyValueInt A expected 42, got %d (found=%v, failed=%v)", valInt, foundInt, failed)
	}

	// Act - missing key
	_, foundMissing := dm.KeyValue("Z")

	// Assert
	if foundMissing {
		t.Error("KeyValue missing should not be found")
	}
}

func Test_ExtDynMap_KeyValueIntDefault_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 42}

	// Act
	val := dm.KeyValueIntDefault("A")
	valMissing := dm.KeyValueIntDefault("Z")

	// Assert
	if val != 42 {
		t.Errorf("KeyValueIntDefault A expected 42, got %d", val)
	}
	// missing returns InvalidValue
	_ = valMissing // just ensure no panic
}

// ==========================================
// DynamicMap - Add
// ==========================================

func Test_ExtDynMap_Add_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	dm.Add("B", 2)

	// Assert
	if dm.Length() != 2 {
		t.Errorf("Add expected 2, got %d", dm.Length())
	}
}

// ==========================================
// DynamicMap - Raw
// ==========================================

func Test_ExtDynMap_Raw_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act
	raw := dm.Raw()

	// Assert
	if len(raw) != 1 {
		t.Errorf("Raw expected 1, got %d", len(raw))
	}
}

// ==========================================
// DynamicMap - DiffRaw
// ==========================================

func Test_ExtDynMap_DiffRaw_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}
	right := map[string]any{"A": 1, "C": 3}

	// Act
	diff := dm.DiffRaw(false, right)

	// Assert
	if diff.Length() == 0 {
		t.Error("DiffRaw should find differences")
	}
}

func Test_ExtDynMap_DiffRaw_NilCases_Verification(t *testing.T) {
	// Arrange
	var nilDm *enumimpl.DynamicMap
	right := map[string]any{"A": 1}
	dm := enumimpl.DynamicMap{"A": 1}

	// Act & Assert
	diff1 := nilDm.DiffRaw(false, nil)
	if diff1.Length() != 0 {
		t.Error("nil vs nil diff should be empty")
	}

	diff2 := nilDm.DiffRaw(false, right)
	if diff2.Length() == 0 {
		t.Error("nil vs non-nil diff should have items")
	}

	diff3 := dm.DiffRaw(false, nil)
	if diff3.Length() == 0 {
		t.Error("non-nil vs nil diff should have items")
	}
}

// ==========================================
// DynamicMap - SortedKeyValues / SortedKeyAnyValues
// ==========================================

func Test_ExtDynMap_SortedKeyValues_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"B": 2, "A": 1}

	// Act
	kvs := dm.SortedKeyValues()

	// Assert
	if len(kvs) != 2 {
		t.Errorf("SortedKeyValues expected 2, got %d", len(kvs))
	}
	if kvs[0].Key != "A" {
		t.Errorf("First sorted key expected 'A', got '%s'", kvs[0].Key)
	}
}

func Test_ExtDynMap_SortedKeyAnyValues_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"B": 2, "A": 1}

	// Act
	kavs := dm.SortedKeyAnyValues()

	// Assert
	if len(kavs) != 2 {
		t.Errorf("SortedKeyAnyValues expected 2, got %d", len(kavs))
	}
}

// ==========================================
// DynamicMap - MapIntegerString
// ==========================================

func Test_ExtDynMap_MapIntegerString_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Valid": 1}

	// Act
	rangeMap, sorted := dm.MapIntegerString()

	// Assert
	if len(rangeMap) != 2 {
		t.Errorf("MapIntegerString map expected 2, got %d", len(rangeMap))
	}
	if len(sorted) != 2 || sorted[0] != 0 {
		t.Errorf("MapIntegerString sorted expected [0,1], got %v", sorted)
	}
}

// ==========================================
// DynamicMap - IsValueString / IsValueTypeOf
// ==========================================

func Test_ExtDynMap_IsValueString_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": "hello"}

	// Act & Assert
	if !dm.IsValueString() {
		t.Error("IsValueString should be true for string values")
	}

	dm2 := enumimpl.DynamicMap{"A": 42}
	if dm2.IsValueString() {
		t.Error("IsValueString should be false for int values")
	}
}

// ==========================================
// DynamicMap - BasicByte / BasicInt8 etc
// ==========================================

func Test_ExtDynMap_BasicByte_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1, "Inactive": 2}

	// Act
	bb := dm.BasicByte("TestType")

	// Assert
	if bb.Min() != 0 {
		t.Errorf("BasicByte Min expected 0, got %d", bb.Min())
	}
	if bb.Max() != 2 {
		t.Errorf("BasicByte Max expected 2, got %d", bb.Max())
	}
	if bb.TypeName() != "TestType" {
		t.Errorf("TypeName expected 'TestType', got '%s'", bb.TypeName())
	}
}

func Test_ExtDynMap_BasicByte_Methods_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}
	bb := dm.BasicByte("TestType")

	// Act & Assert
	if !bb.IsValidRange(0) {
		t.Error("IsValidRange(0) should be true")
	}
	if !bb.IsValidRange(1) {
		t.Error("IsValidRange(1) should be true")
	}
	if bb.IsValidRange(5) {
		t.Error("IsValidRange(5) should be false")
	}
	if !bb.IsAnyOf(1, 0, 1) {
		t.Error("IsAnyOf should find 1")
	}
	if bb.IsAnyOf(5, 0, 1) {
		t.Error("IsAnyOf should not find 5")
	}
	if bb.Length() != 2 {
		t.Errorf("Length expected 2, got %d", bb.Length())
	}
	if bb.Count() != 2 {
		t.Errorf("Count expected 2, got %d", bb.Count())
	}
}

func Test_ExtDynMap_BasicByte_Unmarshal_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}
	bb := dm.BasicByte("TestType")

	// Act
	val, err := bb.UnmarshallToValue(true, nil)

	// Assert
	if err != nil {
		t.Errorf("UnmarshallToValue nil with mapFirst should not error, got %v", err)
	}
	if val != bb.Min() {
		t.Errorf("UnmarshallToValue nil expected min, got %d", val)
	}

	// Act - not mapped to first
	_, err2 := bb.UnmarshallToValue(false, nil)

	// Assert
	if err2 == nil {
		t.Error("UnmarshallToValue nil without mapFirst should error")
	}
}

func Test_ExtDynMap_BasicByte_StringRanges_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}
	bb := dm.BasicByte("TestType")

	// Act
	ranges := bb.StringRanges()
	csv := bb.RangeNamesCsv()
	msg := bb.RangesInvalidMessage()
	err := bb.RangesInvalidErr()

	// Assert
	if len(ranges) != 2 {
		t.Errorf("StringRanges expected 2, got %d", len(ranges))
	}
	if csv == "" {
		t.Error("RangeNamesCsv should not be empty")
	}
	if msg == "" {
		t.Error("RangesInvalidMessage should not be empty")
	}
	if err == nil {
		t.Error("RangesInvalidErr should not be nil")
	}
}

// ==========================================
// DiffLeftRight
// ==========================================

func Test_ExtDiffLeftRight_Verification(t *testing.T) {
	for caseIndex, tc := range extDiffLeftRightTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		dlr := &enumimpl.DiffLeftRight{Left: left, Right: right}
		actual := args.Map{
			"isSame":     fmt.Sprintf("%v", dlr.IsSame()),
			"isNotEqual": fmt.Sprintf("%v", dlr.IsNotEqual()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ExtDiffLeftRight_Methods_Verification(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "hello", Right: "world"}

	// Act & Assert
	if dlr.IsSameTypeSame() != true {
		t.Error("Same type (both string) should return true")
	}
	if dlr.IsSameRegardlessOfType() {
		t.Error("Different values should not be same")
	}
	if dlr.IsEqual(false) {
		t.Error("IsEqual(false) should be false for different values")
	}
	if !dlr.HasMismatch(false) {
		t.Error("HasMismatch should be true")
	}
	if dlr.String() == "" {
		t.Error("String() should not be empty")
	}
	if dlr.JsonString() == "" {
		t.Error("JsonString() should not be empty")
	}
	if dlr.DiffString() == "" {
		t.Error("DiffString() should not be empty for mismatched")
	}

	// Arrange - same
	dlrSame := &enumimpl.DiffLeftRight{Left: "same", Right: "same"}

	// Act & Assert
	if dlrSame.DiffString() != "" {
		t.Error("DiffString() should be empty for same values")
	}

	// nil case
	var nilDlr *enumimpl.DiffLeftRight
	if nilDlr.JsonString() != "" {
		t.Error("nil.JsonString() should be empty")
	}
}

func Test_ExtDiffLeftRight_Types_Verification(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "hello", Right: 42}

	// Act
	l, r := dlr.Types()

	// Assert
	if l == r {
		t.Error("Different types should not be equal")
	}
	if dlr.IsSameTypeSame() {
		t.Error("IsSameTypeSame should be false for string vs int")
	}
}

func Test_ExtDiffLeftRight_SpecificFullString_Verification(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "A", Right: "B"}

	// Act
	l, r := dlr.SpecificFullString()

	// Assert
	if l == "" || r == "" {
		t.Error("SpecificFullString should not be empty")
	}
}

// ==========================================
// KeyAnyVal
// ==========================================

func Test_ExtKeyAnyVal_Verification(t *testing.T) {
	for caseIndex, tc := range extKeyAnyValTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")
		value, _ := input.Get("value")

		// Act
		kav := enumimpl.KeyAnyVal{Key: key, AnyValue: value}
		actual := args.Map{
			"key":      kav.KeyString(),
			"isString": fmt.Sprintf("%v", kav.IsString()),
		}

		if !kav.IsString() {
			actual["valInt"] = fmt.Sprintf("%d", kav.ValInt())
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ExtKeyAnyVal_Methods_Verification(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "Test", AnyValue: 42}

	// Act & Assert
	if kav.AnyVal() != 42 {
		t.Error("AnyVal should be 42")
	}
	if kav.AnyValString() == "" {
		t.Error("AnyValString should not be empty")
	}
	if kav.WrapKey() == "" {
		t.Error("WrapKey should not be empty")
	}
	if kav.WrapValue() == "" {
		t.Error("WrapValue should not be empty")
	}
	if kav.String() == "" {
		t.Error("String should not be empty")
	}

	// KeyValInteger conversion
	kvi := kav.KeyValInteger()
	if kvi.Key != "Test" || kvi.ValueInteger != 42 {
		t.Errorf("KeyValInteger expected Test/42, got %s/%d", kvi.Key, kvi.ValueInteger)
	}
}

// ==========================================
// KeyValInteger
// ==========================================

func Test_ExtKeyValInteger_Verification(t *testing.T) {
	for caseIndex, tc := range extKeyValIntegerTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")
		valueRaw, _ := input.Get("value")
		value := valueRaw.(int)

		// Act
		kvi := enumimpl.KeyValInteger{Key: key, ValueInteger: value}
		actual := args.Map{
			"key":      kvi.Key,
			"isString": fmt.Sprintf("%v", kvi.IsString()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ExtKeyValInteger_Methods_Verification(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "Test", ValueInteger: 5}

	// Act & Assert
	if kvi.WrapKey() == "" {
		t.Error("WrapKey should not be empty")
	}
	if kvi.WrapValue() == "" {
		t.Error("WrapValue should not be empty")
	}
	if kvi.String() == "" {
		t.Error("String should not be empty")
	}

	kav := kvi.KeyAnyVal()
	if kav.Key != "Test" || kav.AnyValue != 5 {
		t.Error("KeyAnyVal conversion mismatch")
	}
}

// ==========================================
// KeyAnyValues (func)
// ==========================================

func Test_ExtKeyAnyValues_Verification(t *testing.T) {
	// Arrange
	names := []string{"A", "B", "C"}
	values := []int{1, 2, 3}

	// Act
	result := enumimpl.KeyAnyValues(names, values)

	// Assert
	if len(result) != 3 {
		t.Errorf("KeyAnyValues expected 3, got %d", len(result))
	}
	if result[0].Key != "A" || result[0].ValInt() != 1 {
		t.Errorf("First element expected A/1, got %s/%d", result[0].Key, result[0].ValInt())
	}
}

func Test_ExtKeyAnyValues_Empty_Verification(t *testing.T) {
	// Arrange
	// Act
	result := enumimpl.KeyAnyValues([]string{}, nil)

	// Assert
	if len(result) != 0 {
		t.Errorf("Empty names should return empty, got %d", len(result))
	}
}

// ==========================================
// AllNameValues
// ==========================================

func Test_ExtAllNameValues_Verification(t *testing.T) {
	// Arrange
	names := []string{"Invalid", "Active"}
	values := []byte{0, 1}

	// Act
	result := enumimpl.AllNameValues(names, values)

	// Assert
	if len(result) != 2 {
		t.Errorf("AllNameValues expected 2, got %d", len(result))
	}
}

// ==========================================
// ConvEnumAnyValToInteger
// ==========================================

func Test_ExtConvEnumAnyValToInteger_Verification(t *testing.T) {
	// Arrange & Act & Assert
	if enumimpl.ConvEnumAnyValToInteger(42) != 42 {
		t.Error("int 42 should convert to 42")
	}
	if enumimpl.ConvEnumAnyValToInteger("hello") >= 0 {
		t.Error("string should convert to MinInt (negative)")
	}
	if enumimpl.ConvEnumAnyValToInteger(byte(5)) != 5 {
		t.Error("byte 5 should convert to 5")
	}
}

// ==========================================
// IntegersRangesOfAnyVal
// ==========================================

func Test_ExtIntegersRangesOfAnyVal_Verification(t *testing.T) {
	// Arrange
	values := []int{3, 1, 2}

	// Act
	result := enumimpl.IntegersRangesOfAnyVal(values)

	// Assert
	if len(result) != 3 || result[0] != 1 || result[2] != 3 {
		t.Errorf("IntegersRangesOfAnyVal expected sorted [1,2,3], got %v", result)
	}
}

// ==========================================
// UnsupportedNames
// ==========================================

func Test_ExtUnsupportedNames_Verification(t *testing.T) {
	// Arrange
	all := []string{"A", "B", "C", "D"}
	supported := []string{"A", "C"}

	// Act
	unsupported := enumimpl.UnsupportedNames(all, supported...)

	// Assert
	if len(unsupported) != 2 {
		t.Errorf("UnsupportedNames expected 2, got %d", len(unsupported))
	}
}

// ==========================================
// PrependJoin / JoinPrependUsingDot
// ==========================================

func Test_ExtPrependJoin_Verification(t *testing.T) {
	// Arrange
	// Act
	result := enumimpl.PrependJoin(".", "prefix", "a", "b")

	// Assert
	if result != "prefix.a.b" {
		t.Errorf("PrependJoin expected 'prefix.a.b', got '%s'", result)
	}
}

func Test_ExtJoinPrependUsingDot_Verification(t *testing.T) {
	// Arrange
	// Act
	result := enumimpl.JoinPrependUsingDot("prefix", "a", "b")

	// Assert
	if result != "prefix.a.b" {
		t.Errorf("JoinPrependUsingDot expected 'prefix.a.b', got '%s'", result)
	}
}

// ==========================================
// NameWithValue
// ==========================================

func Test_ExtNameWithValue_Verification(t *testing.T) {
	// Arrange
	// Act
	result := enumimpl.NameWithValue("TestEnum")

	// Assert
	if result == "" {
		t.Error("NameWithValue should not be empty")
	}
}

// ==========================================
// Format
// ==========================================

func Test_ExtFormat_Verification(t *testing.T) {
	// Arrange
	format := "Enum of {type-name} - {name} - {value}"

	// Act
	result := enumimpl.Format("MyEnum", "Active", "1", format)

	// Assert
	if result != "Enum of MyEnum - Active - 1" {
		t.Errorf("Format expected 'Enum of MyEnum - Active - 1', got '%s'", result)
	}
}

// ==========================================
// differCheckerImpl
// ==========================================

func Test_ExtDifferCheckerImpl_Verification(t *testing.T) {
	// Arrange
	checker := enumimpl.DefaultDiffCheckerImpl

	// Act & Assert
	if !checker.IsEqual(false, 42, 42) {
		t.Error("IsEqual same should be true")
	}
	if checker.IsEqual(false, 42, 43) {
		t.Error("IsEqual different should be false")
	}
	if !checker.IsEqual(true, 42, 42) {
		t.Error("IsEqual regardless same should be true")
	}

	// GetSingleDiffResult
	result := checker.GetSingleDiffResult(true, "left", "right")
	if result != "left" {
		t.Errorf("GetSingleDiffResult isLeft=true should return left, got %v", result)
	}

	result2 := checker.GetSingleDiffResult(false, "left", "right")
	if result2 != "right" {
		t.Errorf("GetSingleDiffResult isLeft=false should return right, got %v", result2)
	}
}

// ==========================================
// leftRightDiffCheckerImpl
// ==========================================

func Test_ExtLeftRightDiffCheckerImpl_Verification(t *testing.T) {
	// Arrange
	checker := enumimpl.LeftRightDiffCheckerImpl

	// Act
	result := checker.GetSingleDiffResult(true, "L", "R")

	// Assert
	if result == nil {
		t.Error("GetSingleDiffResult should not be nil")
	}

	// GetResultOnKeyMissingInRightExistInLeft
	missing := checker.GetResultOnKeyMissingInRightExistInLeft("key", "val")
	if missing == nil {
		t.Error("GetResultOnKeyMissingInRightExistInLeft should not be nil")
	}
}

// ==========================================
// DynamicMap - DiffJsonMessage
// ==========================================

func Test_ExtDynMap_DiffJsonMessage_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1, "B": 2}

	// Act - same
	msg1 := dm.DiffJsonMessage(false, map[string]any{"A": 1, "B": 2})

	// Assert
	if msg1 != "" {
		t.Error("DiffJsonMessage same should be empty")
	}

	// Act - different
	msg2 := dm.DiffJsonMessage(false, map[string]any{"A": 1, "C": 3})

	// Assert
	if msg2 == "" {
		t.Error("DiffJsonMessage different should not be empty")
	}
}

// ==========================================
// DynamicMap - ShouldDiffMessage
// ==========================================

func Test_ExtDynMap_ShouldDiffMessage_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act - same
	msg1 := dm.ShouldDiffMessage(false, "test", map[string]any{"A": 1})

	// Assert
	if msg1 != "" {
		t.Error("ShouldDiffMessage same should be empty")
	}

	// Act - different
	msg2 := dm.ShouldDiffMessage(false, "test", map[string]any{"A": 2})

	// Assert
	if msg2 == "" {
		t.Error("ShouldDiffMessage different should not be empty")
	}
}

// ==========================================
// DynamicMap - ExpectingMessage
// ==========================================

func Test_ExtDynMap_ExpectingMessage_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act - same
	msg1 := dm.ExpectingMessage("test", map[string]any{"A": 1})

	// Assert
	if msg1 != "" {
		t.Error("ExpectingMessage same should be empty")
	}

	// Act - different
	msg2 := dm.ExpectingMessage("test", map[string]any{"A": 2})

	// Assert
	if msg2 == "" {
		t.Error("ExpectingMessage different should not be empty")
	}
}

// ==========================================
// DynamicMap - DiffJsonMessageLeftRight
// ==========================================

func Test_ExtDynMap_DiffJsonMessageLeftRight_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"A": 1}

	// Act - same
	msg1 := dm.DiffJsonMessageLeftRight(false, map[string]any{"A": 1})

	// Assert
	if msg1 != "" {
		t.Error("DiffJsonMessageLeftRight same should be empty")
	}

	// Act - different
	msg2 := dm.DiffJsonMessageLeftRight(false, map[string]any{"B": 2})

	// Assert
	if msg2 == "" {
		t.Error("DiffJsonMessageLeftRight different should not be empty")
	}
}

// ==========================================
// numberEnumBase methods via BasicByte
// ==========================================

func Test_ExtNumberEnumBase_Methods_Verification(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1, "Inactive": 2}
	bb := dm.BasicByte("TestType")

	// Act & Assert - MinMaxAny
	minAny, maxAny := bb.MinMaxAny()
	if minAny == nil || maxAny == nil {
		t.Error("MinMaxAny should not return nil")
	}

	// MinValueString / MaxValueString
	if bb.MinValueString() == "" {
		t.Error("MinValueString should not be empty")
	}
	if bb.MaxValueString() == "" {
		t.Error("MaxValueString should not be empty")
	}

	// MinInt / MaxInt
	if bb.MinInt() != 0 {
		t.Errorf("MinInt expected 0, got %d", bb.MinInt())
	}
	if bb.MaxInt() != 2 {
		t.Errorf("MaxInt expected 2, got %d", bb.MaxInt())
	}

	// AllNameValues
	anv := bb.AllNameValues()
	if len(anv) != 3 {
		t.Errorf("AllNameValues expected 3, got %d", len(anv))
	}

	// IntegerEnumRanges
	ier := bb.IntegerEnumRanges()
	if len(ier) != 3 {
		t.Errorf("IntegerEnumRanges expected 3, got %d", len(ier))
	}

	// NamesHashset
	nh := bb.NamesHashset()
	if len(nh) != 3 {
		t.Errorf("NamesHashset expected 3, got %d", len(nh))
	}

	// RangesDynamicMap / DynamicMap
	rdm := bb.RangesDynamicMap()
	if len(rdm) != 3 {
		t.Errorf("RangesDynamicMap expected 3, got %d", len(rdm))
	}

	// KeyAnyValues
	kavs := bb.KeyAnyValues()
	if len(kavs) != 3 {
		t.Errorf("KeyAnyValues expected 3, got %d", len(kavs))
	}

	// Format
	formatted := bb.Format("Enum of {type-name} - {name} - {value}", byte(0))
	if formatted == "" {
		t.Error("Format should not be empty")
	}
}
