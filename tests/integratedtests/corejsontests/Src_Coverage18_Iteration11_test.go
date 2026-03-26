package corejsontests

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"errors"
	"testing"
)

type i11Serializer struct {
	bytes []byte
	err   error
}

func (it i11Serializer) corejson.Serialize() ([]byte, error) {
	return it.bytes, it.err
}

type i11Deserializer struct {
	err error
}

func (it i11Deserializer) corejson.Deserialize(toPtr any) error {
	if it.err != nil {
		return it.err
	}

	s, ok := toPtr.(*string)
	if ok {
		*s = "ok"
	}

	return nil
}

type i11Jsoner struct {
	value any
}

func (it i11Jsoner) Json() corejson.Result {
	return corejson.New(it.value)
}

func (it i11Jsoner) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it.value)
}

func i11MustPanic(t *testing.T, fn func()) {
	t.Helper()

	didPanic := false
	func() {
		defer func() {
			if recover() != nil {
				didPanic = true
			}
		}()

		fn()
	}()

	if !didPanic {
		t.Fatal("expected panic")
	}
}

func Test_I11_CJ_JsonString_Functions(t *testing.T) {
	jsonString, err := corejson.JsonString(map[string]int{"a": 1})
	if err != nil || jsonString == "" {
		t.Fatal("expected valid json string")
	}

	_, invalidErr := corejson.JsonString(make(chan int))
	if invalidErr == nil {
		t.Fatal("expected marshal error")
	}

	ok := corejson.JsonStringOrErrMsg([]int{1, 2, 3})
	if ok == "" {
		t.Fatal("expected json string")
	}

	msg := corejson.JsonStringOrErrMsg(make(chan int))
	if msg == "" {
		t.Fatal("expected error string")
	}
}

func Test_I11_CJ_BytesAndEmptyCreators(t *testing.T) {
	if len(corejson.BytesCloneIf(false, []byte("abc"))) != 0 {
		t.Fatal("expected empty on non-deep-clone")
	}

	cloned := corejson.BytesCloneIf(true, []byte("abc"))
	if string(cloned) != "abc" {
		t.Fatal("expected cloned bytes")
	}

	if corejson.BytesToString(nil) != "" {
		t.Fatal("expected empty string")
	}

	if corejson.BytesToPrettyString(nil) != "" {
		t.Fatal("expected empty pretty string")
	}

	if corejson.Empty.Result().HasAnyItem() {
		t.Fatal("empty result should not have item")
	}

	errResult := corejson.Empty.ResultWithErr("T", errors.New("x"))
	if errResult.Error == nil || errResult.TypeName != "T" {
		t.Fatal("expected populated empty result with error")
	}

	if corejson.Empty.ResultPtrWithErr("TP", errors.New("err")) == nil {
		t.Fatal("expected non-nil ptr")
	}

	if corejson.Empty.ResultPtr() == nil || corejson.Empty.BytesCollectionPtr() == nil || corejson.Empty.MapResults() == nil {
		t.Fatal("expected all empty creator ptr outputs")
	}

	if corejson.Empty.ResultsCollection().Length() != 0 || corejson.Empty.ResultsPtrCollection().Length() != 0 {
		t.Fatal("expected empty collections")
	}
}

func Test_I11_CJ_NewCreators_Branches(t *testing.T) {
	if corejson.NewResult.UsingBytesPtr(nil).Length() != 0 {
		t.Fatal("expected empty bytes ptr result")
	}

	if corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "T").Error == nil {
		t.Fatal("expected error from UsingBytesPtrErrPtr")
	}

	if len(corejson.NewResult.UsingBytesErrPtr(nil, errors.New("e"), "T").Bytes) != 0 {
		t.Fatal("expected empty bytes for len==0 branch")
	}

	if corejson.NewResult.PtrUsingStringPtr(nil, "S").Error == nil {
		t.Fatal("expected nil string ptr error")
	}

	if corejson.NewResult.UsingErrorStringPtr(errors.New("seed"), nil, "S").Error == nil {
		t.Fatal("expected nil string ptr error with seed")
	}

	if corejson.NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "T").Error == nil {
		t.Fatal("expected error branch")
	}

	if len(corejson.NewResult.PtrUsingBytesPtr(nil, nil, "T").Bytes) != 0 {
		t.Fatal("expected nil-bytes branch")
	}

	if len(corejson.NewResult.PtrUsingBytesPtr([]byte("x"), nil, "T").Bytes) == 0 {
		t.Fatal("expected bytes branch")
	}

	if corejson.NewResult.UsingSerializer(nil) != nil {
		t.Fatal("expected nil serializer result")
	}

	serialized := corejson.NewResult.UsingSerializer(i11Serializer{bytes: []byte("\"ok\"")})
	if serialized == nil || len(serialized.Bytes) == 0 {
		t.Fatal("expected serialized result")
	}

	if corejson.NewResult.UsingSerializerFunc(nil) != nil {
		t.Fatal("expected nil serializer func result")
	}

	withFunc := corejson.NewResult.UsingSerializerFunc(func() ([]byte, error) {
		return []byte("\"ok\""), nil
	})
	if withFunc == nil || len(withFunc.Bytes) == 0 {
		t.Fatal("expected serializer func result")
	}

	if corejson.NewResult.UsingJsoner(nil) != nil {
		t.Fatal("expected nil jsoner result")
	}

	if corejson.NewResult.UsingJsoner(i11Jsoner{value: map[string]int{"a": 1}}) == nil {
		t.Fatal("expected jsoner result")
	}

	if corejson.NewResult.AnyToCastingResult("\"ok\"") == nil {
		t.Fatal("expected casting result")
	}
}

func Test_I11_CJ_CollectionCreators_Branches(t *testing.T) {
	mapFromBytes, err := corejson.NewMapResults.DeserializeUsingBytes([]byte(`{"JsonResultsMap":{"k":{"Bytes":"Im8i","TypeName":"string"}}}`))
	if err != nil || mapFromBytes == nil {
		t.Fatal("expected map deserialization success")
	}

	if _, invalidErr := corejson.NewMapResults.DeserializeUsingBytes([]byte("{")); invalidErr == nil {
		t.Fatal("expected invalid json error")
	}

	if _, issueErr := corejson.NewMapResults.DeserializeUsingResult(&corejson.Result{}); issueErr == nil {
		t.Fatal("expected issue branch on empty result")
	}

	rawMap := map[string]corejson.Result{"a": corejson.NewResult.Any(1)}
	mapped := corejson.NewMapResults.UsingMapOptions(false, false, 0, rawMap)
	mapped.Items["b"] = corejson.NewResult.Any(2)
	if _, has := rawMap["b"]; !has {
		t.Fatal("expected no-clone map reuse")
	}

	if corejson.NewMapResults.UsingKeyAnyItems(0).Length() != 0 {
		t.Fatal("expected empty key-any items map")
	}

	if corejson.NewResultsCollection.UsingJsonersOption(true, 0).Length() != 0 {
		t.Fatal("expected empty jsoners option")
	}

	if _, rcErr := corejson.NewResultsCollection.DeserializeUsingResult(&corejson.Result{}); rcErr == nil {
		t.Fatal("expected collection issue error")
	}

	if corejson.NewResultsCollection.SerializerFunctions().Length() != 0 {
		t.Fatal("expected empty serializer functions collection")
	}

	if corejson.NewResultsPtrCollection.AnyItemsPlusCap(0).Length() != 0 {
		t.Fatal("expected empty any-items ptr collection")
	}

	if _, rpcErr := corejson.NewResultsPtrCollection.DeserializeUsingResult(&corejson.Result{}); rpcErr == nil {
		t.Fatal("expected ptr collection issue error")
	}

	if corejson.NewResultsPtrCollection.Serializers().Length() != 0 {
		t.Fatal("expected empty serializers ptr collection")
	}

	if _, bcErr := corejson.NewBytesCollection.DeserializeUsingBytes([]byte("{")); bcErr == nil {
		t.Fatal("expected bytes collection deserialize error")
	}

	if corejson.NewBytesCollection.JsonersPlusCap(true, 0).Length() != 0 {
		t.Fatal("expected empty jsoners-plus-cap")
	}

	if corejson.NewBytesCollection.Serializers().Length() != 0 {
		t.Fatal("expected empty bytes serializers")
	}
}

func Test_I11_CJ_DeserializerLogic_Branches(t *testing.T) {
	var out string

	if err := corejson.Deserialize.UsingStringPtr(nil, &out); err == nil {
		t.Fatal("expected nil-string-ptr error")
	}

	if err := corejson.Deserialize.UsingError(nil, &out); err != nil {
		t.Fatal("expected nil error pass-through")
	}

	if err := corejson.Deserialize.UsingError(errors.New(`"ok"`), &out); err != nil || out != "ok" {
		t.Fatal("expected error-string deserialization")
	}

	if err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &out); err != nil {
		t.Fatal("expected nil error pass-through for json-result error")
	}

	if err := corejson.Deserialize.UsingErrorWhichJsonResult(errors.New(`"ok2"`), &out); err != nil || out != "ok2" {
		t.Fatal("expected json-result error deserialization")
	}

	if err := corejson.Deserialize.UsingStringOption(true, "", &out); err != nil {
		t.Fatal("expected ignored empty string")
	}

	if err := corejson.Deserialize.UsingStringIgnoreEmpty("", &out); err != nil {
		t.Fatal("expected ignored empty string")
	}

	if err := corejson.Deserialize.UsingBytesPointer(nil, &out); err == nil {
		t.Fatal("expected bytes-pointer nil error")
	}

	corejson.Deserialize.UsingBytesPointerMust([]byte(`"x"`), &out)
	if out != "x" {
		t.Fatal("expected using-bytes-pointer-must success")
	}

	if err := corejson.Deserialize.UsingBytesIf(false, []byte(`"y"`), &out); err != nil {
		t.Fatal("expected skipped bytes-if")
	}

	if err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"y"`), &out); err != nil {
		t.Fatal("expected skipped bytes-pointer-if")
	}

	if err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &out); err != nil {
		t.Fatal("expected skip-nil deserializer")
	}

	if err := corejson.Deserialize.UsingDeserializerToOption(false, nil, &out); err == nil {
		t.Fatal("expected explicit nil deserializer error")
	}

	if err := corejson.Deserialize.UsingDeserializerToOption(false, i11Deserializer{}, &out); err != nil || out != "ok" {
		t.Fatal("expected custom deserializer success")
	}

	if err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &out); err == nil {
		t.Fatal("expected nil deserializer func error")
	}

	if err := corejson.Deserialize.UsingJsonerToAny(true, nil, &out); err != nil {
		t.Fatal("expected skipped nil jsoner")
	}

	if err := corejson.Deserialize.UsingJsonerToAny(false, nil, &out); err == nil {
		t.Fatal("expected nil jsoner error")
	}

	if err := corejson.Deserialize.UsingJsonerToAny(false, i11Jsoner{value: "ok3"}, &out); err != nil || out != "ok3" {
		t.Fatal("expected jsoner deserialize success")
	}
}

func Test_I11_CJ_DeserializerAndSerializer_Wrappers(t *testing.T) {
	if val, err := corejson.Deserialize.BytesTo.String([]byte(`"hello"`)); err != nil || val != "hello" {
		t.Fatal("expected bytes-to string success")
	}

	i11MustPanic(t, func() {
		_ = corejson.Deserialize.BytesTo.StringsMust([]byte("{"))
	})

	jr := corejson.NewResult.Any("abc")
	if val, err := corejson.Deserialize.ResultTo.String(&jr); err != nil || val != "abc" {
		t.Fatal("expected result-to string success")
	}

	i11MustPanic(t, func() {
		_ = corejson.Deserialize.ResultTo.StringMust(&corejson.Result{})
	})

	if corejson.Serialize.FromString("v").JsonString() == "" {
		t.Fatal("expected serialized string")
	}

	if corejson.Serialize.UsingAny(map[string]int{"a": 1}).JsonString() == "" {
		t.Fatal("expected using-any json")
	}

	if _, err := corejson.Serialize.Raw(map[string]string{"k": "v"}); err != nil {
		t.Fatal("expected raw serialization")
	}

	if _, err := corejson.Serialize.Marshal([]int{1, 2}); err != nil {
		t.Fatal("expected marshal serialization")
	}

	i11MustPanic(t, func() {
		_ = corejson.Serialize.ApplyMust(make(chan int))
	})

	i11MustPanic(t, func() {
		_ = corejson.Serialize.ToBytesMust(make(chan int))
	})

	i11MustPanic(t, func() {
		_ = corejson.Serialize.ToSafeBytesMust(make(chan int))
	})

	if corejson.Serialize.ToString([]int{1, 2}) == "" {
		t.Fatal("expected string serialization")
	}

	i11MustPanic(t, func() {
		_ = corejson.Serialize.ToStringMust(make(chan int))
	})

	if _, err := corejson.Serialize.ToStringErr("abc"); err != nil {
		t.Fatal("expected to-string-err success")
	}

	if _, err := corejson.Serialize.ToPrettyStringErr(map[string]int{"a": 1}); err != nil {
		t.Fatal("expected pretty-string-err success")
	}

	_ = corejson.Serialize.ToPrettyStringIncludingErr(make(chan int))
	_ = corejson.Serialize.Pretty(map[string]int{"a": 1})
}
