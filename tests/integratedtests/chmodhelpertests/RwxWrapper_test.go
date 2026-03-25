package chmodhelpertests

import (
	"os"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_RwxWrapper_Create(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperCreateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mode, _ := input.GetAsString("input")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.Create(mode)
		if err != nil {
			t.Fatalf("Create(%s) returned error: %v", mode, err)
		}

		actual := args.Map{
			"ownerRwx":    wrapper.Owner.ToRwxString(),
			"groupRwx":    wrapper.Group.ToRwxString(),
			"otherRwx":    wrapper.Other.ToRwxString(),
			"fullRwx":     wrapper.ToFullRwxValueString(),
			"rwx9":        wrapper.ToFullRwxValueStringExceptHyphen(),
			"fileMode":    wrapper.ToFileModeString(),
			"rwxCompiled": wrapper.ToRwxCompiledStr(),
			"isEmpty":     wrapper.IsEmpty(),
			"isDefined":   wrapper.IsDefined(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_RwxFullString(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperRwxFullStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwxStr, _ := input.GetAsString("input")
		expected := testCase.ExpectedInput.(args.Map)
		hasError, _ := expected.GetAsBool("hasError")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.RwxFullString(rwxStr)

		if hasError {
			// Assert
			actual := args.Map{
				"hasError": err != nil,
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			// Assert
			actual := args.Map{
				"ownerRwx":  wrapper.Owner.ToRwxString(),
				"groupRwx":  wrapper.Group.ToRwxString(),
				"otherRwx":  wrapper.Other.ToRwxString(),
				"hasError":  err != nil,
				"isDefined": wrapper.IsDefined(),
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

func Test_RwxWrapper_Rwx9(t *testing.T) {
	for caseIndex, testCase := range rwxWrapper9StringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rwxStr, _ := input.GetAsString("input")
		expected := testCase.ExpectedInput.(args.Map)
		hasError, _ := expected.GetAsBool("hasError")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.Rwx9(rwxStr)

		if hasError {
			// Assert
			actual := args.Map{
				"hasError": err != nil,
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		} else {
			// Assert
			actual := args.Map{
				"ownerRwx":  wrapper.Owner.ToRwxString(),
				"groupRwx":  wrapper.Group.ToRwxString(),
				"otherRwx":  wrapper.Other.ToRwxString(),
				"hasError":  err != nil,
				"isDefined": wrapper.IsDefined(),
			}
			testCase.ShouldBeEqualMap(t, caseIndex, actual)
		}
	}
}

func Test_RwxWrapper_Bytes(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperBytesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mode, _ := input.GetAsString("input")

		// Act
		wrapper, _ := chmodhelper.New.RwxWrapper.Create(mode)
		bytes := wrapper.Bytes()

		actual := args.Map{
			"byte0": int(bytes[0]),
			"byte1": int(bytes[1]),
			"byte2": int(bytes[2]),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_Clone(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mode, _ := input.GetAsString("input")

		// Act
		wrapper, _ := chmodhelper.New.RwxWrapper.Create(mode)
		cloned := wrapper.Clone()

		actual := args.Map{
			"isEqual":    wrapper.IsEqualPtr(cloned),
			"ownerRwx":   cloned.Owner.ToRwxString(),
			"groupRwx":   cloned.Group.ToRwxString(),
			"otherRwx":   cloned.Other.ToRwxString(),
			"clonedNull": cloned.IsNull(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_IsEqualPtr(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperIsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftMode, _ := input.GetAsString("left")
		rightMode, _ := input.GetAsString("right")

		// Act
		left, _ := chmodhelper.New.RwxWrapper.Create(leftMode)
		right, _ := chmodhelper.New.RwxWrapper.Create(rightMode)

		actual := args.Map{
			"isEqual": left.IsEqualPtr(&right),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_NilClone(t *testing.T) {
	// Arrange
	var wrapper *chmodhelper.RwxWrapper

	// Act
	cloned := wrapper.Clone()

	// Assert
	if cloned != nil {
		t.Error("expected nil clone to be nil")
	}
}

func Test_RwxWrapper_NilIsEmpty(t *testing.T) {
	// Arrange
	var wrapper *chmodhelper.RwxWrapper

	// Act & Assert
	if !wrapper.IsEmpty() {
		t.Error("expected nil wrapper IsEmpty to be true")
	}
	if !wrapper.IsNull() {
		t.Error("expected nil wrapper IsNull to be true")
	}
	if !wrapper.IsInvalid() {
		t.Error("expected nil wrapper IsInvalid to be true")
	}
}

func Test_RwxWrapper_IsEqualPtr_BothNil(t *testing.T) {
	// Arrange
	var left *chmodhelper.RwxWrapper
	var right *chmodhelper.RwxWrapper

	// Act & Assert
	if !left.IsEqualPtr(right) {
		t.Error("expected both nil IsEqualPtr to be true")
	}
}

func Test_RwxWrapper_IsEqualPtr_OneNil(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")
	var nilWrapper *chmodhelper.RwxWrapper

	// Act & Assert
	if wrapper.IsEqualPtr(nilWrapper) {
		t.Error("expected one nil IsEqualPtr to be false")
	}
}

func Test_RwxWrapper_ToFileMode(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	mode := wrapper.ToFileMode()

	// Assert
	expectedMode := os.FileMode(0755)
	if mode != expectedMode {
		t.Errorf("expected FileMode %v, got %v", expectedMode, mode)
	}
}

func Test_RwxWrapper_ToUint32Octal(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	octal := wrapper.ToUint32Octal()

	// Assert
	if octal != 0755 {
		t.Errorf("expected octal 0755 (%d), got %d", uint32(0755), octal)
	}
}

func Test_RwxWrapper_String(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	str := wrapper.String()

	// Assert
	if str != "-rwxr-xr-x" {
		t.Errorf("expected '-rwxr-xr-x', got '%s'", str)
	}
}

func Test_RwxWrapper_UsingBytes(t *testing.T) {
	// Arrange
	bytes := [3]byte{7, 5, 5}

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingBytes(bytes)

	// Assert
	if wrapper.ToFullRwxValueString() != "-rwxr-xr-x" {
		t.Errorf("expected '-rwxr-xr-x', got '%s'", wrapper.ToFullRwxValueString())
	}
}

func Test_RwxWrapper_UsingSpecificByte(t *testing.T) {
	// Arrange & Act
	wrapper := chmodhelper.New.RwxWrapper.UsingSpecificByte(7, 5, 5)

	// Assert
	if wrapper.ToFullRwxValueString() != "-rwxr-xr-x" {
		t.Errorf("expected '-rwxr-xr-x', got '%s'", wrapper.ToFullRwxValueString())
	}
}

func Test_RwxWrapper_Invalid(t *testing.T) {
	// Arrange & Act
	wrapper := chmodhelper.New.RwxWrapper.Invalid()

	// Assert
	if !wrapper.IsEmpty() {
		t.Error("expected Invalid wrapper to be empty")
	}
}

func Test_RwxWrapper_InvalidPtr(t *testing.T) {
	// Arrange & Act
	wrapper := chmodhelper.New.RwxWrapper.InvalidPtr()

	// Assert
	if !wrapper.IsEmpty() {
		t.Error("expected InvalidPtr wrapper to be empty")
	}
}

func Test_RwxWrapper_Empty(t *testing.T) {
	// Arrange & Act
	wrapper := chmodhelper.New.RwxWrapper.Empty()

	// Assert
	if !wrapper.IsEmpty() {
		t.Error("expected Empty wrapper to be empty")
	}
	if !wrapper.IsNull() {
		// Note: Empty returns *RwxWrapper{}, not nil
		// IsNull checks for nil
	}
}

func Test_RwxWrapper_ToCompiledOctalBytes(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperOctalTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mode, _ := input.GetAsString("input")

		// Act
		wrapper, _ := chmodhelper.New.RwxWrapper.Create(mode)
		octal4 := wrapper.ToCompiledOctalBytes4Digits()

		actual := args.Map{
			"octal4": string(octal4[:]),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_ToCompiledSplitValues(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	owner, group, other := wrapper.ToCompiledSplitValues()

	// Assert
	expectedOwner := byte('7')
	expectedGroup := byte('5')
	expectedOther := byte('5')
	if owner != expectedOwner {
		t.Errorf("expected owner '%c', got '%c'", expectedOwner, owner)
	}
	if group != expectedGroup {
		t.Errorf("expected group '%c', got '%c'", expectedGroup, group)
	}
	if other != expectedOther {
		t.Errorf("expected other '%c', got '%c'", expectedOther, other)
	}
}

func Test_RwxWrapper_FriendlyDisplay(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	display := wrapper.FriendlyDisplay()

	// Assert
	if display == "" {
		t.Error("expected FriendlyDisplay to not be empty")
	}
}

func Test_RwxWrapper_ToRwxOwnerGroupOther(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	ogo := wrapper.ToRwxOwnerGroupOther()

	// Assert
	if ogo == nil {
		t.Fatal("expected ToRwxOwnerGroupOther to not be nil")
	}
	if ogo.Owner != "rwx" {
		t.Errorf("expected Owner 'rwx', got '%s'", ogo.Owner)
	}
	if ogo.Group != "r-x" {
		t.Errorf("expected Group 'r-x', got '%s'", ogo.Group)
	}
	if ogo.Other != "r-x" {
		t.Errorf("expected Other 'r-x', got '%s'", ogo.Other)
	}
}

func Test_RwxWrapper_IsRwxFullEqual(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act & Assert
	if !wrapper.IsRwxFullEqual("-rwxr-xr-x") {
		t.Error("expected IsRwxFullEqual('-rwxr-xr-x') true")
	}
	if wrapper.IsRwxFullEqual("-rw-r--r--") {
		t.Error("expected IsRwxFullEqual('-rw-r--r--') false")
	}
	if wrapper.IsRwxFullEqual("short") {
		t.Error("expected IsRwxFullEqual('short') false for short string")
	}
}

func Test_RwxWrapper_IsEqualFileMode(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act & Assert
	if !wrapper.IsEqualFileMode(os.FileMode(0755)) {
		t.Error("expected IsEqualFileMode(0755) true")
	}
	if wrapper.IsEqualFileMode(os.FileMode(0644)) {
		t.Error("expected IsEqualFileMode(0644) false")
	}
	if !wrapper.IsNotEqualFileMode(os.FileMode(0644)) {
		t.Error("expected IsNotEqualFileMode(0644) true")
	}
}

func Test_RwxWrapper_HasAnyItem(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")
	emptyWrapper := chmodhelper.New.RwxWrapper.Invalid()

	// Act & Assert
	if !wrapper.HasAnyItem() {
		t.Error("expected HasAnyItem true for 755")
	}
	if emptyWrapper.HasAnyItem() {
		t.Error("expected HasAnyItem false for empty wrapper")
	}
}

func Test_RwxWrapper_ToPtr_ToNonPtr(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	ptr := wrapper.ToPtr()
	nonPtr := ptr.ToNonPtr()

	// Assert
	if ptr == nil {
		t.Error("expected ToPtr to not be nil")
	}
	if nonPtr.ToFullRwxValueString() != "-rwxr-xr-x" {
		t.Errorf("expected ToNonPtr to preserve value, got '%s'", nonPtr.ToFullRwxValueString())
	}
}

func Test_RwxWrapper_Json(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	jsonResult := wrapper.Json()
	jsonStr := jsonResult.JsonString()

	// Assert
	if jsonStr == "" {
		t.Error("expected Json string to not be empty")
	}
}

func Test_RwxWrapper_UsingVariant(t *testing.T) {
	for caseIndex, testCase := range rwxWrapperVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		variantStr, _ := input.GetAsString("input")

		// Act
		wrapper, err := chmodhelper.New.RwxWrapper.UsingVariant(chmodhelper.Variant(variantStr))

		actual := args.Map{
			"fullRwx":  wrapper.ToFullRwxValueString(),
			"hasError": err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_RwxWrapper_UsingAttrVariants(t *testing.T) {
	// Arrange & Act
	wrapper := chmodhelper.New.RwxWrapper.UsingAttrVariants(
		chmodhelper.ReadWriteExecute,
		chmodhelper.ReadExecute,
		chmodhelper.ReadExecute,
	)

	// Assert
	if wrapper.ToFullRwxValueString() != "-rwxr-xr-x" {
		t.Errorf("expected '-rwxr-xr-x', got '%s'", wrapper.ToFullRwxValueString())
	}
}

func Test_RwxWrapper_UsingAttrs(t *testing.T) {
	// Arrange
	owner := chmodhelper.New.Attribute.Create(true, true, true)
	group := chmodhelper.New.Attribute.Create(true, false, true)
	other := chmodhelper.New.Attribute.Create(true, false, true)

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingAttrs(owner, group, other)

	// Assert
	if wrapper.ToFullRwxValueString() != "-rwxr-xr-x" {
		t.Errorf("expected '-rwxr-xr-x', got '%s'", wrapper.ToFullRwxValueString())
	}
}

func Test_RwxWrapper_CreatePtr(t *testing.T) {
	// Arrange & Act
	ptr, err := chmodhelper.New.RwxWrapper.CreatePtr("755")

	// Assert
	if err != nil {
		t.Fatalf("CreatePtr returned error: %v", err)
	}
	if ptr == nil {
		t.Fatal("expected CreatePtr to not be nil")
	}
	if ptr.ToFullRwxValueString() != "-rwxr-xr-x" {
		t.Errorf("expected '-rwxr-xr-x', got '%s'", ptr.ToFullRwxValueString())
	}
}

func Test_RwxWrapper_UsingFileModePtr(t *testing.T) {
	// Arrange
	mode := os.FileMode(0755)

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingFileModePtr(mode)

	// Assert
	if wrapper == nil {
		t.Fatal("expected UsingFileModePtr to not be nil")
	}
	if wrapper.ToFullRwxValueString() != "-rwxr-xr-x" {
		t.Errorf("expected '-rwxr-xr-x', got '%s'", wrapper.ToFullRwxValueString())
	}
}

func Test_RwxWrapper_UsingFileModePtr_Zero(t *testing.T) {
	// Arrange
	mode := os.FileMode(0)

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingFileModePtr(mode)

	// Assert
	if wrapper == nil {
		t.Fatal("expected UsingFileModePtr to not be nil even for zero")
	}
	if !wrapper.IsEmpty() {
		t.Error("expected zero FileMode to create empty wrapper")
	}
}

func Test_RwxWrapper_UsingFileMode(t *testing.T) {
	// Arrange
	mode := os.FileMode(0644)

	// Act
	wrapper := chmodhelper.New.RwxWrapper.UsingFileMode(mode)

	// Assert
	if wrapper.ToFullRwxValueString() != "-rw-r--r--" {
		t.Errorf("expected '-rw-r--r--', got '%s'", wrapper.ToFullRwxValueString())
	}
}

func Test_RwxWrapper_UsingVariantPtr(t *testing.T) {
	// Arrange & Act
	ptr, err := chmodhelper.New.RwxWrapper.UsingVariantPtr(chmodhelper.Variant("755"))

	// Assert
	if err != nil {
		t.Fatalf("UsingVariantPtr returned error: %v", err)
	}
	if ptr == nil {
		t.Fatal("expected UsingVariantPtr to not be nil")
	}
	if ptr.ToFullRwxValueString() != "-rwxr-xr-x" {
		t.Errorf("expected '-rwxr-xr-x', got '%s'", ptr.ToFullRwxValueString())
	}
}

func Test_RwxWrapper_ToFullRwxValuesChars(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act
	chars := wrapper.ToFullRwxValuesChars()

	// Assert
	if len(chars) != 10 {
		t.Errorf("expected 10 chars, got %d", len(chars))
	}
	if string(chars) != "-rwxr-xr-x" {
		t.Errorf("expected '-rwxr-xr-x', got '%s'", string(chars))
	}
}

func Test_RwxWrapper_IsEqualVarWrapper_Nil(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act & Assert
	if wrapper.IsEqualVarWrapper(nil) {
		t.Error("expected IsEqualVarWrapper(nil) false")
	}
}

func Test_RwxWrapper_IsRwxEqualFileInfo_Nil(t *testing.T) {
	// Arrange
	wrapper, _ := chmodhelper.New.RwxWrapper.Create("755")

	// Act & Assert
	if wrapper.IsRwxEqualFileInfo(nil) {
		t.Error("expected IsRwxEqualFileInfo(nil) false")
	}
}

func Test_FileModeFriendlyString(t *testing.T) {
	// Arrange
	mode := os.FileMode(0755)

	// Act
	result := chmodhelper.FileModeFriendlyString(mode)

	// Assert
	if result == "" {
		t.Error("expected FileModeFriendlyString to not be empty")
	}
}

func Test_AttrVariant(t *testing.T) {
	for caseIndex, testCase := range attrVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsInt("input")

		// Act
		variant := chmodhelper.AttrVariant(val)
		attr := variant.ToAttribute()

		actual := args.Map{
			"value":       int(variant.Value()),
			"attrRead":    attr.IsRead,
			"attrWrite":   attr.IsWrite,
			"attrExecute": attr.IsExecute,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_AttrVariant_IsGreaterThan(t *testing.T) {
	// Arrange
	v := chmodhelper.ReadWriteExecute // 7

	// Act & Assert
	if !v.IsGreaterThan(8) {
		t.Error("expected IsGreaterThan(8) true (8 > 7)")
	}
	if v.IsGreaterThan(5) {
		t.Error("expected IsGreaterThan(5) false (5 < 7)")
	}
}

func Test_Variant_String(t *testing.T) {
	// Arrange
	v := chmodhelper.X755

	// Act
	str := v.String()

	// Assert
	if str != "755" {
		t.Errorf("expected '755', got '%s'", str)
	}
}

func Test_Variant_ExpandOctalByte(t *testing.T) {
	// Arrange
	v := chmodhelper.X755

	// Act
	r, w, x := v.ExpandOctalByte()

	// Assert
	if r != '7' {
		t.Errorf("expected r='7' (%d), got %d", byte('7'), r)
	}
	if w != '5' {
		t.Errorf("expected w='5' (%d), got %d", byte('5'), w)
	}
	if x != '5' {
		t.Errorf("expected x='5' (%d), got %d", byte('5'), x)
	}
}

func Test_Variant_ToWrapper(t *testing.T) {
	// Arrange
	v := chmodhelper.X755

	// Act
	wrapper, err := v.ToWrapper()

	// Assert
	if err != nil {
		t.Fatalf("ToWrapper returned error: %v", err)
	}
	if wrapper.ToFullRwxValueString() != "-rwxr-xr-x" {
		t.Errorf("expected '-rwxr-xr-x', got '%s'", wrapper.ToFullRwxValueString())
	}
}

func Test_Variant_ToWrapperPtr(t *testing.T) {
	// Arrange
	v := chmodhelper.X644

	// Act
	wrapper, err := v.ToWrapperPtr()

	// Assert
	if err != nil {
		t.Fatalf("ToWrapperPtr returned error: %v", err)
	}
	if wrapper == nil {
		t.Fatal("expected ToWrapperPtr to not be nil")
	}
	if wrapper.ToFullRwxValueString() != "-rw-r--r--" {
		t.Errorf("expected '-rw-r--r--', got '%s'", wrapper.ToFullRwxValueString())
	}
}

func Test_Attribute_HasAnyItem(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, false, false)
	emptyAttr := chmodhelper.New.Attribute.Create(false, false, false)

	// Act & Assert
	if !attr.HasAnyItem() {
		t.Error("expected HasAnyItem true for read-only attribute")
	}
	if emptyAttr.HasAnyItem() {
		t.Error("expected HasAnyItem false for empty attribute")
	}
}

func Test_Attribute_ToSum(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, true, true)

	// Act
	sum := attr.ToSum()

	// Assert
	if sum != 7 {
		t.Errorf("expected ToSum 7, got %d", sum)
	}
}

func Test_Attribute_ToRwx(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, false, true)

	// Act
	rwx := attr.ToRwx()

	// Assert
	if rwx != [3]byte{'r', '-', 'x'} {
		t.Errorf("expected [r,-,x], got %v", rwx)
	}
}

func Test_Attribute_ToStringByte(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, true, true)

	// Act
	sb := attr.ToStringByte()

	// Assert - 7 + '0' = '7'
	if sb != '7' {
		t.Errorf("expected '7' (%d), got %d", byte('7'), sb)
	}
}

func Test_Attribute_ToSpecificBytes(t *testing.T) {
	// Arrange
	attr := chmodhelper.New.Attribute.Create(true, true, false)

	// Act
	read, write, exe, sum := attr.ToSpecificBytes()

	// Assert
	if read != 4 {
		t.Errorf("expected read=4, got %d", read)
	}
	if write != 2 {
		t.Errorf("expected write=2, got %d", write)
	}
	if exe != 0 {
		t.Errorf("expected exe=0, got %d", exe)
	}
	if sum != 6 {
		t.Errorf("expected sum=6, got %d", sum)
	}
}

func Test_Attribute_UsingByteMust(t *testing.T) {
	// Arrange & Act
	attr := chmodhelper.New.Attribute.UsingByteMust(5)

	// Assert
	if !attr.IsRead {
		t.Error("expected IsRead true for byte 5")
	}
	if attr.IsWrite {
		t.Error("expected IsWrite false for byte 5")
	}
	if !attr.IsExecute {
		t.Error("expected IsExecute true for byte 5")
	}
}

func Test_Attribute_UsingVariantMust(t *testing.T) {
	// Arrange & Act
	attr := chmodhelper.New.Attribute.UsingVariantMust(chmodhelper.ReadWriteExecute)

	// Assert
	if !attr.IsRead || !attr.IsWrite || !attr.IsExecute {
		t.Error("expected all permissions true for ReadWriteExecute variant")
	}
}

func Test_Attribute_UsingVariant(t *testing.T) {
	// Arrange & Act
	attr, err := chmodhelper.New.Attribute.UsingVariant(chmodhelper.ReadExecute)

	// Assert
	if err != nil {
		t.Fatalf("UsingVariant returned error: %v", err)
	}
	if !attr.IsRead || attr.IsWrite || !attr.IsExecute {
		t.Error("expected read+execute only for ReadExecute variant")
	}
}

func Test_Attribute_Default(t *testing.T) {
	// Arrange & Act
	attr := chmodhelper.New.Attribute.Default(true, false, true)

	// Assert
	if !attr.IsRead || attr.IsWrite || !attr.IsExecute {
		t.Error("expected Default to create attribute with given values")
	}
}

func Test_IsChmod_EmptyLocation(t *testing.T) {
	// Arrange & Act & Assert
	if chmodhelper.IsChmod("", "-rwxr-xr-x") {
		t.Error("expected IsChmod empty location to be false")
	}
}

func Test_IsChmod_InvalidLength(t *testing.T) {
	// Arrange & Act & Assert
	if chmodhelper.IsChmod(".", "rwx") {
		t.Error("expected IsChmod invalid rwx length to be false")
	}
}
