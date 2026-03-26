package corejsontests

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"errors"
	"testing"
)

func TestEmptyCreator(t *testing.T) {
	_ = corejson.Empty.Result()
	_ = corejson.Empty.ResultPtr()
	_ = corejson.Empty.ResultWithErr("t", errors.New("e"))
	_ = corejson.Empty.ResultPtrWithErr("t", errors.New("e"))
	_ = corejson.Empty.BytesCollection()
	_ = corejson.Empty.BytesCollectionPtr()
	_ = corejson.Empty.ResultsCollection()
	_ = corejson.Empty.ResultsPtrCollection()
	_ = corejson.Empty.MapResults()
}

func TestNewResultCreator(t *testing.T) {
	_ = corejson.NewResult.UsingBytes([]byte(`"x"`))
	_ = corejson.NewResult.UsingBytesType([]byte(`"x"`), "string")
	_ = corejson.NewResult.UsingBytesTypePtr([]byte(`"x"`), "string")
	_ = corejson.NewResult.UsingTypeBytesPtr("string", []byte(`"x"`))
	_ = corejson.NewResult.UsingBytesPtr([]byte(`"x"`))
	_ = corejson.NewResult.UsingBytesPtr(nil)
	_ = corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "t")
	_ = corejson.NewResult.UsingBytesPtrErrPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingBytesErrPtr([]byte{}, nil, "t")
	_ = corejson.NewResult.UsingBytesErrPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.Ptr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingJsonBytesTypeError([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingJsonBytesError([]byte(`"x"`), nil)
	_ = corejson.NewResult.UsingTypePlusString("t", `"x"`)
	_ = corejson.NewResult.UsingStringWithType(`"x"`, "t")
	_ = corejson.NewResult.UsingString(`"x"`)
	_ = corejson.NewResult.CreatePtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.NonPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.Create([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "t")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, nil, "t")
	_ = corejson.NewResult.PtrUsingBytesPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.Error(errors.New("e"))
	_ = corejson.NewResult.ErrorPtr(errors.New("e"))
	_ = corejson.NewResult.Empty()
	_ = corejson.NewResult.EmptyPtr()
	_ = corejson.NewResult.TypeName("t")
	_ = corejson.NewResult.TypeNameBytes("t")
	_ = corejson.NewResult.Many("a", "b")
	_ = corejson.NewResult.Any("x")
	_ = corejson.NewResult.AnyPtr("x")
	_ = corejson.NewResult.Serialize("x")
	_ = corejson.NewResult.Marshal("x")
	_ = corejson.NewResult.CastingAny("x")
	_ = corejson.NewResult.AnyToCastingResult("x")
}

func TestNewResultCreator_UsingStringPtr(t *testing.T) {
	s := `"hello"`
	_ = corejson.NewResult.UsingStringPtr(&s)
	_ = corejson.NewResult.UsingStringPtr(nil)

	empty := ""
	_ = corejson.NewResult.UsingStringPtr(&empty)
}

func TestNewResultCreator_UsingTypePlusStringPtr(t *testing.T) {
	s := `"x"`
	_ = corejson.NewResult.UsingTypePlusStringPtr("t", &s)
	_ = corejson.NewResult.UsingTypePlusStringPtr("t", nil)
}

func TestNewResultCreator_PtrUsingStringPtr(t *testing.T) {
	s := `"x"`
	_ = corejson.NewResult.PtrUsingStringPtr(&s, "t")
	_ = corejson.NewResult.PtrUsingStringPtr(nil, "t")
}

func TestNewResultCreator_UsingErrorStringPtr(t *testing.T) {
	s := `"x"`
	_ = corejson.NewResult.UsingErrorStringPtr(nil, &s, "t")
	_ = corejson.NewResult.UsingErrorStringPtr(errors.New("e"), &s, "t")
	_ = corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "t")
}

func TestNewResultCreator_UsingSerializer(t *testing.T) {
	r := corejson.NewResult.UsingSerializer(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func TestNewResultCreator_UsingSerializerFunc(t *testing.T) {
	r := corejson.NewResult.UsingSerializerFunc(nil)
	if r != nil {
		t.Fatal("expected nil")
	}

	r2 := corejson.NewResult.UsingSerializerFunc(func() ([]byte, error) {
		return []byte(`"x"`), nil
	})
	if r2 == nil || r2.HasError() {
		t.Fatal("expected non-nil")
	}
}

func TestNewResultCreator_UsingJsoner(t *testing.T) {
	r := corejson.NewResult.UsingJsoner(nil)
	if r != nil {
		t.Fatal("expected nil")
	}
}

func TestNewResultCreator_Deserialize(t *testing.T) {
	orig := corejson.NewResult.Any("hello")
	b, _ := orig.Serialize()
	_ = corejson.NewResult.DeserializeUsingBytes(b)
}

func TestNewResultCreator_DeserializeUsingResult(t *testing.T) {
	orig := corejson.NewResult.Any("hello")
	b, _ := orig.Serialize()
	jr := corejson.NewResult.UsingBytes(b)
	_ = corejson.NewResult.DeserializeUsingResult(jr.Ptr())
}

func TestNewResultsCollectionCreator(t *testing.T) {
	_ = corejson.NewResultsCollection.Empty()
	_ = corejson.NewResultsCollection.Default()
	_ = corejson.NewResultsCollection.UsingCap(5)
	_ = corejson.NewResultsCollection.AnyItems("a", "b")
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(2, "a")
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(2)
	_ = corejson.NewResultsCollection.UsingResults(corejson.NewResult.Any("a"))
	_ = corejson.NewResultsCollection.UsingResultsPtr(corejson.NewResult.AnyPtr("a"))
	_ = corejson.NewResultsCollection.UsingResultsPlusCap(2, corejson.NewResult.Any("a"))
	_ = corejson.NewResultsCollection.UsingResultsPtrPlusCap(2, corejson.NewResult.AnyPtr("a"))
}

func TestNewResultsPtrCollectionCreator(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Empty()
	_ = corejson.NewResultsPtrCollection.Default()
	_ = corejson.NewResultsPtrCollection.UsingCap(5)
	_ = corejson.NewResultsPtrCollection.AnyItems("a", "b")
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(2, "a")
	_ = corejson.NewResultsPtrCollection.UsingResults(corejson.NewResult.AnyPtr("a"))
	_ = corejson.NewResultsPtrCollection.UsingResultsPlusCap(2, corejson.NewResult.AnyPtr("a"))
}

func TestNewBytesCollectionCreator(t *testing.T) {
	_ = corejson.NewBytesCollection.Empty()
	_ = corejson.NewBytesCollection.UsingCap(5)
	_, _ = corejson.NewBytesCollection.AnyItems("a", "b")
}

func TestNewMapResultsCreator(t *testing.T) {
	_ = corejson.NewMapResults.Empty()
	_ = corejson.NewMapResults.UsingCap(5)
	_ = corejson.NewMapResults.UsingMap(map[string]corejson.Result{})
	_ = corejson.NewMapResults.UsingMap(map[string]corejson.Result{"k": corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingMapPlusCap(2, map[string]corejson.Result{})
	_ = corejson.NewMapResults.UsingMapPlusCap(2, map[string]corejson.Result{"k": corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingMapPlusCapClone(2, map[string]corejson.Result{"k": corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(2, map[string]corejson.Result{"k": corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingMapAnyItems(map[string]any{"k": "v"})
	_ = corejson.NewMapResults.UsingMapAnyItemsPlusCap(2, map[string]any{})
	_ = corejson.NewMapResults.UsingKeyWithResults(corejson.KeyWithResult{Key: "k", corejson.Result: corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(2, corejson.KeyWithResult{Key: "k", corejson.Result: corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingKeyAnyItems(2, corejson.KeyAny{Key: "k", AnyInf: "v"})
}

func TestBytesCloneIf_Func(t *testing.T) {
	b := corejson.BytesCloneIf(true, []byte("hello"))
	_ = b
	b2 := corejson.BytesCloneIf(false, []byte("hello"))
	_ = b2
	b3 := corejson.BytesCloneIf(true, []byte{})
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
	s := corejson.BytesToString([]byte("hello"))
	if s != "hello" {
		t.Fatal("unexpected")
	}
	s2 := corejson.BytesToString([]byte{})
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func TestBytesToPrettyString_Func(t *testing.T) {
	s := corejson.BytesToPrettyString([]byte(`{"a":1}`))
	if s == "" {
		t.Fatal("expected non-empty")
	}
	s2 := corejson.BytesToPrettyString([]byte{})
	if s2 != "" {
		t.Fatal("expected empty")
	}
}

func TestJsonString_Func(t *testing.T) {
	s, err := corejson.JsonString("hello")
	if err != nil || s == "" {
		t.Fatal("unexpected")
	}
}

func TestJsonStringOrErrMsg_Func(t *testing.T) {
	s := corejson.JsonStringOrErrMsg("hello")
	if s == "" {
		t.Fatal("expected non-empty")
	}

	ch := make(chan int)
	s2 := corejson.JsonStringOrErrMsg(ch)
	if s2 == "" {
		t.Fatal("expected error message")
	}
}
