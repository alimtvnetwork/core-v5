package corejson

import (
	"errors"
	"testing"
)

func TestEmptyCreator(t *testing.T) {
	_ = Empty.Result()
	_ = Empty.ResultPtr()
	_ = Empty.ResultWithErr("t", errors.New("e"))
	_ = Empty.ResultPtrWithErr("t", errors.New("e"))
	_ = Empty.BytesCollection()
	_ = Empty.BytesCollectionPtr()
	_ = Empty.ResultsCollection()
	_ = Empty.ResultsPtrCollection()
	_ = Empty.MapResults()
}

func TestNewResultCreator(t *testing.T) {
	_ = NewResult.UsingBytes([]byte(`"x"`))
	_ = NewResult.UsingBytesType([]byte(`"x"`), "string")
	_ = NewResult.UsingBytesTypePtr([]byte(`"x"`), "string")
	_ = NewResult.UsingTypeBytesPtr("string", []byte(`"x"`))
	_ = NewResult.UsingBytesPtr([]byte(`"x"`))
	_ = NewResult.UsingBytesPtr(nil)
	_ = NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "t")
	_ = NewResult.UsingBytesPtrErrPtr([]byte(`"x"`), nil, "t")
	_ = NewResult.UsingBytesErrPtr([]byte{}, nil, "t")
	_ = NewResult.UsingBytesErrPtr([]byte(`"x"`), nil, "t")
	_ = NewResult.Ptr([]byte(`"x"`), nil, "t")
	_ = NewResult.UsingJsonBytesTypeError([]byte(`"x"`), nil, "t")
	_ = NewResult.UsingJsonBytesError([]byte(`"x"`), nil)
	_ = NewResult.UsingTypePlusString("t", `"x"`)
	_ = NewResult.UsingStringWithType(`"x"`, "t")
	_ = NewResult.UsingString(`"x"`)
	_ = NewResult.CreatePtr([]byte(`"x"`), nil, "t")
	_ = NewResult.NonPtr([]byte(`"x"`), nil, "t")
	_ = NewResult.Create([]byte(`"x"`), nil, "t")
	_ = NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "t")
	_ = NewResult.PtrUsingBytesPtr(nil, nil, "t")
	_ = NewResult.PtrUsingBytesPtr([]byte(`"x"`), nil, "t")
	_ = NewResult.Error(errors.New("e"))
	_ = NewResult.ErrorPtr(errors.New("e"))
	_ = NewResult.Empty()
	_ = NewResult.EmptyPtr()
	_ = NewResult.TypeName("t")
	_ = NewResult.TypeNameBytes("t")
	_ = NewResult.Many("a", "b")
	_ = NewResult.Any("x")
	_ = NewResult.AnyPtr("x")
	_ = NewResult.Serialize("x")
	_ = NewResult.Marshal("x")
	_ = NewResult.CastingAny("x")
	_ = NewResult.AnyToCastingResult("x")
}

func TestNewResultCreator_UsingStringPtr(t *testing.T) {
	s := `"hello"`
	_ = NewResult.UsingStringPtr(&s)
	_ = NewResult.UsingStringPtr(nil)

	empty := ""
	_ = NewResult.UsingStringPtr(&empty)
}

func TestNewResultCreator_UsingTypePlusStringPtr(t *testing.T) {
	s := `"x"`
	_ = NewResult.UsingTypePlusStringPtr("t", &s)
	_ = NewResult.UsingTypePlusStringPtr("t", nil)
}

func TestNewResultCreator_PtrUsingStringPtr(t *testing.T) {
	s := `"x"`
	_ = NewResult.PtrUsingStringPtr(&s, "t")
	_ = NewResult.PtrUsingStringPtr(nil, "t")
}

func TestNewResultCreator_UsingErrorStringPtr(t *testing.T) {
	s := `"x"`
	_ = NewResult.UsingErrorStringPtr(nil, &s, "t")
	_ = NewResult.UsingErrorStringPtr(errors.New("e"), &s, "t")
	_ = NewResult.UsingErrorStringPtr(errors.New("e"), nil, "t")
}

func TestNewResultCreator_UsingSerializer(t *testing.T) {
	r := NewResult.UsingSerializer(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func TestNewResultCreator_UsingSerializerFunc(t *testing.T) {
	r := NewResult.UsingSerializerFunc(nil)
	if r != nil {
		t.Fatal("expected nil")
	}

	r2 := NewResult.UsingSerializerFunc(func() ([]byte, error) {
		return []byte(`"x"`), nil
	})
	if r2 == nil || r2.HasError() {
		t.Fatal("expected non-nil")
	}
}

func TestNewResultCreator_UsingJsoner(t *testing.T) {
	r := NewResult.UsingJsoner(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func TestNewResultCreator_Deserialize(t *testing.T) {
	orig := NewResult.Any("hello")
	b, _ := orig.Serialize()
	_ = NewResult.DeserializeUsingBytes(b)
}

func TestNewResultCreator_DeserializeUsingResult(t *testing.T) {
	orig := NewResult.Any("hello")
	b, _ := orig.Serialize()
	jr := NewResult.UsingBytes(b)
	_ = NewResult.DeserializeUsingResult(jr.Ptr())
}

func TestNewResultsCollectionCreator(t *testing.T) {
	_ = NewResultsCollection.Empty()
	_ = NewResultsCollection.Default()
	_ = NewResultsCollection.UsingCap(5)
	_ = NewResultsCollection.AnyItems("a", "b")
	_ = NewResultsCollection.AnyItemsPlusCap(2, "a")
	_ = NewResultsCollection.AnyItemsPlusCap(2)
	_ = NewResultsCollection.UsingResults(NewResult.Any("a"))
	_ = NewResultsCollection.UsingResultsPtr(NewResult.AnyPtr("a"))
	_ = NewResultsCollection.UsingResultsPlusCap(2, NewResult.Any("a"))
	_ = NewResultsCollection.UsingResultsPtrPlusCap(2, NewResult.AnyPtr("a"))
}

func TestNewResultsPtrCollectionCreator(t *testing.T) {
	_ = NewResultsPtrCollection.Empty()
	_ = NewResultsPtrCollection.Default()
	_ = NewResultsPtrCollection.UsingCap(5)
	_ = NewResultsPtrCollection.AnyItems("a", "b")
	_ = NewResultsPtrCollection.AnyItemsPlusCap(2, "a")
	_ = NewResultsPtrCollection.UsingResults(NewResult.AnyPtr("a"))
	_ = NewResultsPtrCollection.UsingResultsPlusCap(2, NewResult.AnyPtr("a"))
}

func TestNewBytesCollectionCreator(t *testing.T) {
	_ = NewBytesCollection.Empty()
	_ = NewBytesCollection.UsingCap(5)
	_, _ = NewBytesCollection.AnyItems("a", "b")
}

func TestNewMapResultsCreator(t *testing.T) {
	_ = NewMapResults.Empty()
	_ = NewMapResults.UsingCap(5)
	_ = NewMapResults.UsingMap(map[string]Result{})
	_ = NewMapResults.UsingMap(map[string]Result{"k": NewResult.Any("v")})
	_ = NewMapResults.UsingMapPlusCap(2, map[string]Result{})
	_ = NewMapResults.UsingMapPlusCap(2, map[string]Result{"k": NewResult.Any("v")})
	_ = NewMapResults.UsingMapPlusCapClone(2, map[string]Result{"k": NewResult.Any("v")})
	_ = NewMapResults.UsingMapPlusCapDeepClone(2, map[string]Result{"k": NewResult.Any("v")})
	_ = NewMapResults.UsingMapAnyItems(map[string]any{"k": "v"})
	_ = NewMapResults.UsingMapAnyItemsPlusCap(2, map[string]any{})
	_ = NewMapResults.UsingKeyWithResults(KeyWithResult{Key: "k", Result: NewResult.Any("v")})
	_ = NewMapResults.UsingKeyWithResultsPlusCap(2, KeyWithResult{Key: "k", Result: NewResult.Any("v")})
	_ = NewMapResults.UsingKeyAnyItems(2, KeyAny{Key: "k", AnyInf: "v"})
}

func TestBytesCloneIf_Func(t *testing.T) {
	b := BytesCloneIf(true, []byte("hello"))
	_ = b
	b2 := BytesCloneIf(false, []byte("hello"))
	_ = b2
	b3 := BytesCloneIf(true, []byte{})
	_ = b3
}

func TestBytesDeepClone_Func(t *testing.T) {
	b := BytesDeepClone([]byte("hello"))
	if len(b) == 0 {
		t.Fatal("expected bytes")
	}
	b2 := BytesDeepClone([]byte{})
	if len(b2) != 0 {
		t.Fatal("expected empty")
	}
}

func TestBytesToString_Func(t *testing.T) {
	s := BytesToString([]byte("hello"))
	if s != "hello" {
		t.Fatal("unexpected")
	}
	s2 := BytesToString([]byte{})
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func TestBytesToPrettyString_Func(t *testing.T) {
	s := BytesToPrettyString([]byte(`{"a":1}`))
	if s == "" {
		t.Fatal("expected non-empty")
	}
	s2 := BytesToPrettyString([]byte{})
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func TestJsonString_Func(t *testing.T) {
	s, err := JsonString("hello")
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func TestJsonStringOrErrMsg_Func(t *testing.T) {
	s := JsonStringOrErrMsg("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}

	ch := make(chan int)
	s2 := JsonStringOrErrMsg(ch)
	if s2 == "" {
		t.Fatal("expected error message")
	}
}
