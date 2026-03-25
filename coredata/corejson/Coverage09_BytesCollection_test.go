package corejson

import (
	"testing"
	"time"
)

func TestBytesCollection_Basic(t *testing.T) {
	var nilC *BytesCollection
	if nilC.Length() != 0 { t.Fatal("expected 0") }
	if nilC.LastIndex() != -1 { t.Fatal("expected -1") }
	if !nilC.IsEmpty() { t.Fatal("expected empty") }
	if nilC.HasAnyItem() { t.Fatal("expected false") }
	if nilC.FirstOrDefault() != nil { t.Fatal("expected nil") }
	if nilC.LastOrDefault() != nil { t.Fatal("expected nil") }
}

func TestBytesCollection_AddAndAccess(t *testing.T) {
	c := NewBytesCollection.UsingCap(5)
	c.Add([]byte(`"a"`))
	c.AddSkipOnNil(nil)
	c.AddSkipOnNil([]byte(`"b"`))
	c.AddNonEmpty(nil)
	c.AddNonEmpty([]byte(`"c"`))
	c.AddPtr(nil)
	c.AddPtr([]byte(`"d"`))
	if c.Length() != 4 { t.Fatal("expected 4, got", c.Length()) }

	if c.GetAt(0) == nil { t.Fatal("expected non-nil") }
	jr := c.JsonResultAt(0)
	if jr == nil { t.Fatal("expected non-nil") }
}

func TestBytesCollection_AddResult(t *testing.T) {
	c := NewBytesCollection.UsingCap(2)
	r := NewResult.Any("hello")
	c.AddResult(r)
	c.AddResultPtr(nil)
	c.AddResultPtr(&r)
	if c.Length() != 2 { t.Fatal("expected 2") }
}

func TestBytesCollection_Adds(t *testing.T) {
	c := NewBytesCollection.UsingCap(5)
	c.Adds([]byte(`"a"`), nil, []byte(`"b"`))
	if c.Length() != 2 { t.Fatal("expected 2") }
	c.Adds()
}

func TestBytesCollection_AddAny(t *testing.T) {
	c := NewBytesCollection.UsingCap(2)
	err := c.AddAny("hello")
	if err != nil { t.Fatal("unexpected error") }
	if c.Length() != 1 { t.Fatal("expected 1") }
	err2 := c.AddAnyItems("a", "b")
	if err2 != nil { t.Fatal("unexpected error") }
	if c.Length() != 3 { t.Fatal("expected 3") }
	err3 := c.AddAnyItems()
	if err3 != nil { t.Fatal("unexpected error") }
}

func TestBytesCollection_UnmarshalAt(t *testing.T) {
	c := NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"hello"`))
	var s string
	err := c.UnmarshalAt(0, &s)
	if err != nil || s != "hello" { t.Fatal("unexpected") }
}

func TestBytesCollection_UnmarshalIntoSameIndex(t *testing.T) {
	c := NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"hello"`))
	c.Add([]byte(`42`))
	var s string
	var i int
	errs, _ := c.UnmarshalIntoSameIndex(&s, &i)
	if len(errs) != 2 { t.Fatal("expected 2") }
	c2 := NewBytesCollection.Empty()
	c2.UnmarshalIntoSameIndex(nil)
}

func TestBytesCollection_GetAtSafe(t *testing.T) {
	c := NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"x"`))
	if c.GetAtSafe(0) == nil { t.Fatal("expected non-nil") }
	if c.GetAtSafe(-1) != nil { t.Fatal("expected nil") }
	if c.GetAtSafe(5) != nil { t.Fatal("expected nil") }
	if c.GetAtSafePtr(0) == nil { t.Fatal("expected non-nil") }
	if c.GetResultAtSafe(0) == nil { t.Fatal("expected non-nil") }
	if c.GetResultAtSafe(5) != nil { t.Fatal("expected nil") }
	if c.GetAtSafeUsingLength(0, 1) == nil { t.Fatal("expected non-nil") }
}

func TestBytesCollection_Serializers(t *testing.T) {
	c := NewBytesCollection.UsingCap(2)
	c.AddSerializer(nil)
	c.AddSerializers()
	c.AddSerializerFunc(nil)
	c.AddSerializerFunctions()
}

func TestBytesCollection_MapResults(t *testing.T) {
	c := NewBytesCollection.UsingCap(2)
	mr := NewMapResults.Empty()
	c.AddMapResults(mr)
	c.AddRawMapResults(nil)
	c.AddRawMapResults(map[string]Result{"a": NewResult.Any("x")})
}

func TestBytesCollection_AddsPtr(t *testing.T) {
	c := NewBytesCollection.UsingCap(2)
	r := NewResult.Any("x")
	c.AddsPtr(nil, &r)
	if c.Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_AddBytesCollection(t *testing.T) {
	c := NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	c2 := NewBytesCollection.UsingCap(1)
	c2.Add([]byte(`"b"`))
	c.AddBytesCollection(c2)
	if c.Length() != 2 { t.Fatal("expected 2") }
	c.AddBytesCollection(NewBytesCollection.Empty())
}

func TestBytesCollection_ClearDispose(t *testing.T) {
	c := NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"x"`))
	c.Clear()
	time.Sleep(10 * time.Millisecond)
	if c.Length() != 0 { t.Fatal("expected 0") }
	c.Dispose()
	var nilC *BytesCollection
	nilC.Clear()
	nilC.Dispose()
}

func TestBytesCollection_Strings(t *testing.T) {
	c := NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"a"`))
	strs := c.Strings()
	if len(strs) != 1 { t.Fatal("expected 1") }
	_ = c.StringsPtr()
	empty := NewBytesCollection.Empty()
	if len(empty.Strings()) != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_Paging(t *testing.T) {
	c := NewBytesCollection.UsingCap(10)
	for i := 0; i < 10; i++ {
		c.Add([]byte(`"x"`))
	}
	if c.GetPagesSize(3) != 4 { t.Fatal("expected 4") }
	if c.GetPagesSize(0) != 0 { t.Fatal("expected 0") }

	paged := c.GetPagedCollection(3)
	if len(paged) != 4 { t.Fatal("expected 4") }

	single := c.GetSinglePageCollection(3, 1)
	if single.Length() != 3 { t.Fatal("expected 3") }

	small := NewBytesCollection.UsingCap(1)
	small.Add([]byte(`"x"`))
	if len(small.GetPagedCollection(5)) != 1 { t.Fatal("expected 1") }
	if small.GetSinglePageCollection(5, 1).Length() != 1 { t.Fatal("expected 1") }
}

func TestBytesCollection_JsonMethods(t *testing.T) {
	c := NewBytesCollection.UsingCap(1)
	c.Add([]byte(`"x"`))
	_ = c.JsonModel()
	_ = c.JsonModelAny()
	b, err := c.MarshalJSON()
	if err != nil || len(b) == 0 { t.Fatal("unexpected") }
	_ = c.Json()
	_ = c.JsonPtr()
	_ = c.AsJsonContractsBinder()
	_ = c.AsJsoner()
	_ = c.AsJsonParseSelfInjector()
}

func TestBytesCollection_Clone(t *testing.T) {
	c := NewBytesCollection.UsingCap(2)
	c.Add([]byte(`"x"`))
	sc := c.ShadowClone()
	if sc.Length() != 1 { t.Fatal("expected 1") }
	dc := c.Clone(true)
	if dc.Length() != 1 { t.Fatal("expected 1") }
	cp := c.ClonePtr(true)
	if cp == nil || cp.Length() != 1 { t.Fatal("expected 1") }
	var nilC *BytesCollection
	if nilC.ClonePtr(true) != nil { t.Fatal("expected nil") }

	empty := NewBytesCollection.Empty()
	ec := empty.Clone(true)
	if ec.Length() != 0 { t.Fatal("expected 0") }
}

func TestBytesCollection_Creators(t *testing.T) {
	_, _ = NewBytesCollection.AnyItems("a", "b")
	_ = NewBytesCollection.JsonersPlusCap(true, 0)
	_ = NewBytesCollection.Serializers()
	_, _ = NewBytesCollection.UnmarshalUsingBytes([]byte(`{}`))
}
