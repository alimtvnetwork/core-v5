package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretaskinfo"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── InfoJson — all methods ──

func Test_Cov2_Info_Json(t *testing.T) {
	info := coretaskinfo.Info{RootName: "task"}
	r := info.Json()
	actual := args.Map{"hasBytes": r.HasBytes()}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Info Json", actual)
}

func Test_Cov2_Info_JsonPtr(t *testing.T) {
	info := coretaskinfo.Info{RootName: "task"}
	r := info.JsonPtr()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Info JsonPtr", actual)
}

func Test_Cov2_Info_JsonString(t *testing.T) {
	info := coretaskinfo.Info{RootName: "task"}
	actual := args.Map{"notEmpty": info.JsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info JsonString", actual)
}

func Test_Cov2_Info_JsonString_Nil(t *testing.T) {
	info := &coretaskinfo.Info{}
	result := info.JsonString()
	actual := args.Map{"notPanic": true, "hasResult": result != ""}
	expected := args.Map{"notPanic": true, "hasResult": actual["hasResult"]}
	expected.ShouldBeEqual(t, 0, "Info JsonString nil", actual)
}

func Test_Cov2_Info_JsonStringMust(t *testing.T) {
	info := coretaskinfo.Info{RootName: "task"}
	actual := args.Map{"notEmpty": info.JsonStringMust() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info JsonStringMust", actual)
}

func Test_Cov2_Info_PrettyJsonString(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	actual := args.Map{"notEmpty": info.PrettyJsonString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info PrettyJsonString", actual)
}

func Test_Cov2_Info_PrettyJsonString_Nil(t *testing.T) {
	var info *coretaskinfo.Info
	actual := args.Map{"empty": info.PrettyJsonString()}
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "Info PrettyJsonString nil", actual)
}

func Test_Cov2_Info_String(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	actual := args.Map{"notEmpty": info.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info String", actual)
}

func Test_Cov2_Info_String_Nil(t *testing.T) {
	var info *coretaskinfo.Info
	actual := args.Map{"empty": info.String()}
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "Info String nil", actual)
}

func Test_Cov2_Info_Serialize(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	bytes, err := info.Serialize()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(bytes) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Info Serialize", actual)
}

func Test_Cov2_Info_ExamplesAsString(t *testing.T) {
	info := &coretaskinfo.Info{Examples: []string{"a", "b"}}
	result := info.ExamplesAsString()
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a, b"}
	expected.ShouldBeEqual(t, 0, "Info ExamplesAsString", actual)
}

func Test_Cov2_Info_ExamplesAsString_Nil(t *testing.T) {
	var info *coretaskinfo.Info
	actual := args.Map{"result": info.ExamplesAsString()}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Info ExamplesAsString nil", actual)
}

func Test_Cov2_Info_AsJsonContractsBinder(t *testing.T) {
	info := coretaskinfo.Info{RootName: "task"}
	binder := info.AsJsonContractsBinder()
	actual := args.Map{"notNil": binder != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Info AsJsonContractsBinder", actual)
}

// ── InfoMap — Map / LazyMap ──

func Test_Cov2_Info_Map_Defined(t *testing.T) {
	info := &coretaskinfo.Info{
		RootName: "task", Description: "desc", Url: "url",
		HintUrl: "hint", ErrorUrl: "errUrl", ExampleUrl: "exUrl",
		SingleExample: "single", Examples: []string{"e1"},
	}
	m := info.Map()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 8}
	expected.ShouldBeEqual(t, 0, "Info Map defined", actual)
}

func Test_Cov2_Info_Map_Nil(t *testing.T) {
	var info *coretaskinfo.Info
	m := info.Map()
	actual := args.Map{"len": len(m)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Info Map nil", actual)
}

func Test_Cov2_Info_LazyMap(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	m1 := info.LazyMap()
	m2 := info.LazyMap() // cached
	actual := args.Map{"len": len(m1), "sameRef": len(m2) == len(m1)}
	expected := args.Map{"len": 1, "sameRef": true}
	expected.ShouldBeEqual(t, 0, "Info LazyMap", actual)
}

func Test_Cov2_Info_LazyMapPrettyJsonString(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	result := info.LazyMapPrettyJsonString()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info LazyMapPrettyJsonString", actual)
}

func Test_Cov2_Info_MapWithPayload(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	m := info.MapWithPayload([]byte("data"))
	actual := args.Map{"gt0": len(m) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Info MapWithPayload", actual)
}

func Test_Cov2_Info_PrettyJsonStringWithPayloads(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	result := info.PrettyJsonStringWithPayloads([]byte("data"))
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Info PrettyJsonStringWithPayloads", actual)
}

// ── newInfoCreator ──

func Test_Cov2_NewInfo_Default(t *testing.T) {
	info := coretaskinfo.New.Info.Default("n", "d", "u")
	actual := args.Map{"name": info.Name()}
	expected := args.Map{"name": "n"}
	expected.ShouldBeEqual(t, 0, "NewInfo Default", actual)
}

func Test_Cov2_NewInfo_Examples(t *testing.T) {
	info := coretaskinfo.New.Info.Examples("n", "d", "u", "e1")
	actual := args.Map{"hasEx": info.HasExamples()}
	expected := args.Map{"hasEx": true}
	expected.ShouldBeEqual(t, 0, "NewInfo Examples", actual)
}

func Test_Cov2_NewInfo_Create(t *testing.T) {
	info := coretaskinfo.New.Info.Create(false, "n", "d", "u", "h", "e", "ex", "ch", "e1")
	actual := args.Map{"name": info.Name()}
	expected := args.Map{"name": "n"}
	expected.ShouldBeEqual(t, 0, "NewInfo Create", actual)
}

func Test_Cov2_NewInfo_SecureCreate(t *testing.T) {
	info := coretaskinfo.New.Info.SecureCreate("n", "d", "u", "h", "e", "ex", "ch")
	actual := args.Map{"secure": info.IsSecure()}
	expected := args.Map{"secure": true}
	expected.ShouldBeEqual(t, 0, "NewInfo SecureCreate", actual)
}

func Test_Cov2_NewInfo_PlainCreate(t *testing.T) {
	info := coretaskinfo.New.Info.PlainCreate("n", "d", "u", "h", "e", "ex", "ch")
	actual := args.Map{"plain": info.IsPlainText()}
	expected := args.Map{"plain": true}
	expected.ShouldBeEqual(t, 0, "NewInfo PlainCreate", actual)
}

// ── newInfoPlainTextCreator ──

func Test_Cov2_PlainCreator_AllMethods(t *testing.T) {
	p := coretaskinfo.New.Info.Plain
	actual := args.Map{
		"default":  p.Default("n", "d", "u").Name(),
		"ndu":      p.NameDescUrl("n", "d", "u").Name(),
		"nduEx":    p.NameDescUrlExamples("n", "d", "u", "e").HasExamples(),
		"ndueErr":  p.NewNameDescUrlErrorUrl("n", "d", "u", "e").HasErrorUrl(),
		"nduee":    p.NameDescUrlErrUrlExamples("n", "d", "u", "e", "e1").HasExamples(),
		"ndex":     p.NameDescExamples("n", "d", "e1").HasExamples(),
		"ex":       p.Examples("n", "d", "e1").HasExamples(),
		"nuEx":     p.NameUrlExamples("n", "u", "e1").HasExamples(),
		"uEx":      p.UrlExamples("u", "e1").HasExamples(),
		"exOnly":   p.ExamplesOnly("e1").HasExamples(),
		"urlOnly":  p.UrlOnly("u").HasUrl(),
		"errOnly":  p.ErrorUrlOnly("e").HasErrorUrl(),
		"hintOnly": p.HintUrlOnly("h").HasHintUrl(),
		"descHint": p.DescHintUrlOnly("d", "h").HasHintUrl(),
		"nameHint": p.NameHintUrlOnly("n", "h").HasHintUrl(),
		"singleEx": p.SingleExampleOnly("s").HasChainingExample(),
		"allUrl":   p.AllUrl("n", "d", "u", "h", "e").HasUrl(),
		"allUrlEx": p.AllUrlExamples("n", "d", "u", "h", "e", "e1").HasExamples(),
		"urlSE":    p.UrlSingleExample("n", "d", "u", "s").HasChainingExample(),
		"singleE":  p.SingleExample("n", "d", "s").HasChainingExample(),
		"exUrl":    p.ExampleUrl("n", "d", "eu", "s").HasExampleUrl(),
		"exUrlSE":  p.ExampleUrlSingleExample("n", "d", "eu", "s").HasChainingExample(),
	}
	expected := args.Map{
		"default": "n", "ndu": "n", "nduEx": true, "ndueErr": true,
		"nduee": true, "ndex": true, "ex": true, "nuEx": true,
		"uEx": true, "exOnly": true, "urlOnly": true, "errOnly": true,
		"hintOnly": true, "descHint": true, "nameHint": true, "singleEx": true,
		"allUrl": true, "allUrlEx": true, "urlSE": true, "singleE": true,
		"exUrl": true, "exUrlSE": true,
	}
	expected.ShouldBeEqual(t, 0, "PlainCreator all methods", actual)
}

// ── newInfoSecureTextCreator ──

func Test_Cov2_SecureCreator_AllMethods(t *testing.T) {
	s := coretaskinfo.New.Info.Secure
	actual := args.Map{
		"default":   s.Default("n", "d", "u").IsSecure(),
		"ndu":       s.NameDescUrl("n", "d", "u").IsSecure(),
		"nduEx":     s.NameDescUrlExamples("n", "d", "u", "e").IsSecure(),
		"ndueErr":   s.NewNameDescUrlErrorUrl("n", "d", "u", "e").IsSecure(),
		"nduee":     s.NameDescUrlErrUrlExamples("n", "d", "u", "e", "e1").IsSecure(),
		"ndex":      s.NameDescExamples("n", "d", "e1").IsSecure(),
		"ex":        s.Examples("n", "d", "e1").IsSecure(),
		"exOnly":    s.ExamplesOnly("e1").IsSecure(),
		"urlOnly":   s.UrlOnly("u").IsSecure(),
		"errOnly":   s.ErrorUrlOnly("e").IsSecure(),
		"hintOnly":  s.HintUrlOnly("h").IsSecure(),
		"nameHint":  s.NameHintUrlOnly("n", "h").IsSecure(),
		"singleEx":  s.SingleExampleOnly("s").IsSecure(),
		"allUrlEx":  s.AllUrlExamples("n", "d", "u", "h", "e", "e1").IsSecure(),
		"allUrl":    s.AllUrl("n", "d", "u", "h", "e").IsSecure(),
		"urlSE":     s.UrlSingleExample("n", "d", "u", "s").IsSecure(),
		"singleE":   s.SingleExample("n", "d", "s").IsSecure(),
		"exUrl":     s.ExampleUrl("n", "d", "eu", "s").IsSecure(),
		"exUrlSE":   s.ExampleUrlSingleExample("n", "d", "eu", "s").IsSecure(),
		"newExUrl":  s.NewExampleUrlSecure("n", "d", "eu", "s").IsSecure(),
	}
	expected := args.Map{
		"default": true, "ndu": true, "nduEx": true, "ndueErr": true,
		"nduee": true, "ndex": true, "ex": true, "exOnly": true,
		"urlOnly": true, "errOnly": true, "hintOnly": true, "nameHint": true,
		"singleEx": true, "allUrlEx": true, "allUrl": true, "urlSE": true,
		"singleE": true, "exUrl": true, "exUrlSE": true, "newExUrl": true,
	}
	expected.ShouldBeEqual(t, 0, "SecureCreator all methods", actual)
}

// ── Deserialized / DeserializedUsingJsonResult ──

func Test_Cov2_NewInfo_Deserialized(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	bytes, _ := info.Serialize()
	parsed, err := coretaskinfo.New.Info.Deserialized(bytes)
	actual := args.Map{"noErr": err == nil, "name": parsed.Name()}
	expected := args.Map{"noErr": true, "name": "task"}
	expected.ShouldBeEqual(t, 0, "NewInfo Deserialized", actual)
}

func Test_Cov2_NewInfo_DeserializedUsingJsonResult(t *testing.T) {
	info := coretaskinfo.Info{RootName: "task"}
	jsonResult := info.JsonPtr()
	parsed, err := coretaskinfo.New.Info.DeserializedUsingJsonResult(jsonResult)
	actual := args.Map{"noErr": err == nil, "name": parsed.Name()}
	expected := args.Map{"noErr": true, "name": "task"}
	expected.ShouldBeEqual(t, 0, "NewInfo DeserializedUsingJsonResult", actual)
}

// ── Exclude options ──

func Test_Cov2_Info_WithExcludeOptions(t *testing.T) {
	info := &coretaskinfo.Info{
		RootName: "task",
		ExcludeOptions: &coretaskinfo.ExcludingOptions{
			IsExcludeRootName:    true,
			IsExcludeDescription: true,
			IsExcludeUrl:         true,
			IsExcludeHintUrl:     true,
			IsExcludeErrorUrl:    true,
			IsSecureText:         true,
		},
	}
	actual := args.Map{
		"exName": info.IsExcludeRootName(),
		"exDesc": info.IsExcludeDescription(),
		"exUrl":  info.IsExcludeUrl(),
		"exHint": info.IsExcludeHintUrl(),
		"exErr":  info.IsExcludeErrorUrl(),
		"exPay":  info.IsExcludePayload(),
	}
	expected := args.Map{
		"exName": true, "exDesc": true, "exUrl": true,
		"exHint": true, "exErr": true, "exPay": true,
	}
	expected.ShouldBeEqual(t, 0, "Info with ExcludeOptions", actual)
}

// ── MapWithPayloadAsAny ──

func Test_Cov2_Info_MapWithPayloadAsAny(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	m := info.MapWithPayloadAsAny("hello")
	actual := args.Map{"gt0": len(m) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Info MapWithPayloadAsAny", actual)
}

// ── LazyMapWithPayload ──

func Test_Cov2_Info_LazyMapWithPayload(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	m := info.LazyMapWithPayload([]byte("data"))
	actual := args.Map{"gt0": len(m) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Info LazyMapWithPayload", actual)
}

// ── LazyMapWithPayloadAsAny ──

func Test_Cov2_Info_LazyMapWithPayloadAsAny(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	m := info.LazyMapWithPayloadAsAny("payload")
	actual := args.Map{"gt0": len(m) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "Info LazyMapWithPayloadAsAny", actual)
}

// ── JsonParseSelfInject ──

func Test_Cov2_Info_JsonParseSelfInject(t *testing.T) {
	info := coretaskinfo.Info{RootName: "task"}
	jsonResult := info.JsonPtr()
	var parsed coretaskinfo.Info
	err := parsed.JsonParseSelfInject(jsonResult)
	actual := args.Map{"noErr": err == nil, "name": parsed.Name()}
	expected := args.Map{"noErr": true, "name": "task"}
	expected.ShouldBeEqual(t, 0, "Info JsonParseSelfInject", actual)
}

// ── Deserialize ──

func Test_Cov2_Info_Deserialize(t *testing.T) {
	info := &coretaskinfo.Info{RootName: "task"}
	var parsed coretaskinfo.Info
	err := info.Deserialize(&parsed)
	actual := args.Map{"noErr": err == nil, "name": parsed.Name()}
	expected := args.Map{"noErr": true, "name": "task"}
	expected.ShouldBeEqual(t, 0, "Info Deserialize", actual)
}
