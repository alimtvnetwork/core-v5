package corerangetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── RangeString ──

func Test_RangeString_Valid_Cov2(t *testing.T) {
	rs := corerange.NewRangeString("hello:world", ":")
	actual := args.Map{"isValid": rs.IsValid, "hasStart": rs.HasStart, "hasEnd": rs.HasEnd}
	expected := args.Map{"isValid": true, "hasStart": true, "hasEnd": true}
	expected.ShouldBeEqual(t, 0, "RangeString_Valid", actual)
}

func Test_RangeString_NoSeparator_Cov2(t *testing.T) {
	actual := args.Map{"isValid": corerange.NewRangeString("hello", ":").IsValid}
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "RangeString_NoSeparator", actual)
}

func Test_RangeString_Empty_Cov2(t *testing.T) {
	actual := args.Map{"isValid": corerange.NewRangeString("", ":").IsValid}
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "RangeString_Empty", actual)
}

func Test_RangeString_Methods_Cov2(t *testing.T) {
	rs := corerange.NewRangeString("hello:world", ":")
	actual := args.Map{"stringNotEmpty": rs.String() != "", "start": rs.Start, "end": rs.End}
	expected := args.Map{"stringNotEmpty": true, "start": "hello", "end": "world"}
	expected.ShouldBeEqual(t, 0, "RangeString_Methods", actual)
}

// ── RangeInt8 ──

func Test_RangeInt8_Cov2(t *testing.T) {
	ri8 := corerange.NewRangeIntMinMax("3:7", ":", 0, 10).CreateRangeInt8()
	actual := args.Map{"isValid": ri8.IsValid, "start": int(ri8.Start), "end": int(ri8.End)}
	expected := args.Map{"isValid": true, "start": 3, "end": 7}
	expected.ShouldBeEqual(t, 0, "RangeInt8", actual)
}

func Test_RangeInt8_Methods_Cov2(t *testing.T) {
	ri8 := corerange.NewRangeIntMinMax("2:5", ":", 0, 10).CreateRangeInt8()
	actual := args.Map{
		"rangeLength": int(ri8.RangeLength()), "difference": int(ri8.Difference()),
		"isWithin3": ri8.IsWithinRange(3), "isWithin8": ri8.IsWithinRange(8),
		"isInvalid8": ri8.IsInvalidValue(8), "stringNotEmpty": ri8.String() != "",
		"rangesLen": len(ri8.Ranges()), "rangesIntLen": len(ri8.RangesInt8()),
	}
	expected := args.Map{
		"rangeLength": 4, "difference": 3,
		"isWithin3": true, "isWithin8": false,
		"isInvalid8": true, "stringNotEmpty": true,
		"rangesLen": 4, "rangesIntLen": 4,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt8_Methods", actual)
}

// ── RangeInt16 ──

func Test_RangeInt16_Cov2(t *testing.T) {
	ri16 := corerange.NewRangeIntMinMax("3:7", ":", 0, 10).CreateRangeInt16()
	actual := args.Map{"isValid": ri16.IsValid, "start": int(ri16.Start), "end": int(ri16.End)}
	expected := args.Map{"isValid": true, "start": 3, "end": 7}
	expected.ShouldBeEqual(t, 0, "RangeInt16", actual)
}

func Test_RangeInt16_Methods_Cov2(t *testing.T) {
	ri16 := corerange.NewRangeIntMinMax("2:5", ":", 0, 10).CreateRangeInt16()
	actual := args.Map{
		"rangeLength": int(ri16.RangeLength()), "difference": int(ri16.Difference()),
		"isWithin3": ri16.IsWithinRange(3), "stringNotEmpty": ri16.String() != "",
	}
	expected := args.Map{
		"rangeLength": 4, "difference": 3,
		"isWithin3": true, "stringNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "RangeInt16_Methods", actual)
}

// ── RangeByte ──

func Test_RangeByte_Cov2(t *testing.T) {
	rb := corerange.NewRangeIntMinMax("3:7", ":", 0, 10).CreateRangeByte()
	actual := args.Map{"isValid": rb.IsValid, "start": int(rb.Start), "end": int(rb.End)}
	expected := args.Map{"isValid": true, "start": 3, "end": 7}
	expected.ShouldBeEqual(t, 0, "RangeByte", actual)
}

func Test_RangeByte_Methods_Cov2(t *testing.T) {
	rb := corerange.NewRangeIntMinMax("2:5", ":", 0, 10).CreateRangeByte()
	actual := args.Map{
		"rangeLength": int(rb.RangeLength()), "difference": int(rb.Difference()),
		"isWithin3": rb.IsWithinRange(3), "stringNotEmpty": rb.String() != "",
	}
	expected := args.Map{
		"rangeLength": 4, "difference": 3,
		"isWithin3": true, "stringNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "RangeByte_Methods", actual)
}

// ── BaseRange ──

func Test_BaseRange_Clone_Cov2(t *testing.T) {
	br := &corerange.BaseRange{RawInput: "1:5", Separator: ":", IsValid: true, HasStart: true, HasEnd: true}
	cloned := br.BaseRangeClone()
	actual := args.Map{
		"rawInput": cloned.RawInput, "separator": cloned.Separator,
		"isValid": cloned.IsValid, "isInvalid": cloned.IsInvalid(),
	}
	expected := args.Map{
		"rawInput": "1:5", "separator": ":",
		"isValid": true, "isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "BaseRange_Clone", actual)
}

func Test_BaseRange_String_Cov2(t *testing.T) {
	br := &corerange.BaseRange{Separator: ":"}
	actual := args.Map{"notEmpty": br.String(1, 5) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BaseRange_String", actual)
}

func Test_BaseRange_CreateRangeInt_Cov2(t *testing.T) {
	br := &corerange.BaseRange{RawInput: "3:7", Separator: ":", IsValid: true}
	mm := &corerange.MinMaxInt{Min: 0, Max: 10}
	actual := args.Map{"notNil": br.CreateRangeInt(mm) != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BaseRange_CreateRangeInt", actual)
}

// ── MinMaxInt16 ──

func Test_MinMaxInt16_Cov2(t *testing.T) {
	mm := &corerange.MinMaxInt16{Min: 2, Max: 8}
	actual := args.Map{
		"difference": int(mm.Difference()), "rangeLength": int(mm.RangeLength()),
		"isWithin5": mm.IsWithinRange(5), "isWithin10": mm.IsWithinRange(10),
		"isOutOfR10": mm.IsOutOfRange(10), "isInvalid10": mm.IsInvalidValue(10),
	}
	expected := args.Map{
		"difference": 6, "rangeLength": 7,
		"isWithin5": true, "isWithin10": false,
		"isOutOfR10": true, "isInvalid10": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt16", actual)
}

// ── MinMaxInt64 ──

func Test_MinMaxInt64_Cov2(t *testing.T) {
	mm := &corerange.MinMaxInt64{Min: 2, Max: 8}
	actual := args.Map{
		"difference": int(mm.Difference()), "rangeLength": int(mm.RangeLength()),
		"isWithin5": mm.IsWithinRange(5), "isWithin10": mm.IsWithinRange(10),
		"isOutOfR10": mm.IsOutOfRange(10),
	}
	expected := args.Map{
		"difference": 6, "rangeLength": 7,
		"isWithin5": true, "isWithin10": false,
		"isOutOfR10": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt64", actual)
}

// ── MinMaxInt8 ──

func Test_MinMaxInt8_Cov2(t *testing.T) {
	mm := &corerange.MinMaxInt8{Min: 2, Max: 8}
	actual := args.Map{
		"difference": int(mm.Difference()), "rangeLength": int(mm.RangeLength()),
		"isWithin5": mm.IsWithinRange(5), "isWithin10": mm.IsWithinRange(10),
		"isOutOfR10": mm.IsOutOfRange(10),
	}
	expected := args.Map{
		"difference": 6, "rangeLength": 7,
		"isWithin5": true, "isWithin10": false,
		"isOutOfR10": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxInt8", actual)
}

// ── StartEndSimpleString ──

func Test_StartEndSimpleString_Cov2(t *testing.T) {
	se := &corerange.StartEndSimpleString{Start: "hello", End: "world"}
	actual := args.Map{"hasStart": se.HasStart(), "hasEnd": se.HasEnd(), "isStartEmpty": se.Start == "", "isEndEmpty": se.End == ""}
	expected := args.Map{"hasStart": true, "hasEnd": true, "isStartEmpty": false, "isEndEmpty": false}
	expected.ShouldBeEqual(t, 0, "StartEndSimpleString", actual)
}

// ── StartEndString ──

func Test_StartEndString_Cov2(t *testing.T) {
	se := corerange.NewStartEndString("hello:world", ":")
	actual := args.Map{"isValid": se.IsValid, "hasStart": se.HasStart, "hasEnd": se.HasEnd}
	expected := args.Map{"isValid": true, "hasStart": true, "hasEnd": true}
	expected.ShouldBeEqual(t, 0, "StartEndString", actual)
}

func Test_StartEndString_Methods_Cov2(t *testing.T) {
	se := corerange.NewStartEndString("hello:world", ":")
	actual := args.Map{"start": se.Start, "end": se.End, "stringNotEmpty": se.String() != ""}
	expected := args.Map{"start": "hello", "end": "world", "stringNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "StartEndString_Methods", actual)
}

// ── RangeAny ──

func Test_RangeAny_Cov2(t *testing.T) {
	ra := &corerange.RangeAny{BaseRange: &corerange.BaseRange{RawInput: "hello:world", Separator: ":", IsValid: true, HasStart: true, HasEnd: true}}
	actual := args.Map{"isValid": ra.IsValid}
	expected := args.Map{"isValid": true}
	expected.ShouldBeEqual(t, 0, "RangeAny", actual)
}

func Test_RangeAny_NoSeparator_Cov2(t *testing.T) {
	ra := &corerange.RangeAny{BaseRange: &corerange.BaseRange{RawInput: "hello", Separator: ":", IsValid: false}}
	actual := args.Map{"isValid": ra.IsValid}
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "RangeAny_NoSep", actual)
}

// ── MinMaxInt boundary ──

func Test_MinMaxInt_IsWithinRange_Boundary_Cov2(t *testing.T) {
	mm := &corerange.MinMaxInt{Min: 3, Max: 7}
	actual := args.Map{"exactMin": mm.IsWithinRange(3), "exactMax": mm.IsWithinRange(7), "below": mm.IsWithinRange(2), "above": mm.IsWithinRange(8)}
	expected := args.Map{"exactMin": true, "exactMax": true, "below": false, "above": false}
	expected.ShouldBeEqual(t, 0, "MinMaxInt_Within_Boundary", actual)
}
