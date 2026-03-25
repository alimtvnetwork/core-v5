package reflectmodeltests

import (
	"reflect"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// sampleStruct is a test helper struct used to create
// real reflect.Method and reflect.StructField instances.
type sampleStruct struct {
	Name    string
	Age     int
	Active  bool
	private string
}

func (s sampleStruct) PublicMethod(a string, b int) (string, error) {
	return a, nil
}

func (s sampleStruct) NoArgsMethod() string {
	return "hello"
}

func (s sampleStruct) MultiReturn() (int, string, error) {
	return 0, "", nil
}

// newFieldProcessor creates a FieldProcessor from a sampleStruct field by name.
func newFieldProcessor(fieldName string, index int) *reflectmodel.FieldProcessor {
	t := reflect.TypeOf(sampleStruct{})
	field, ok := t.FieldByName(fieldName)
	if !ok {
		return nil
	}

	return &reflectmodel.FieldProcessor{
		Name:      field.Name,
		Index:     index,
		Field:     field,
		FieldType: field.Type,
	}
}

// newMethodProcessor creates a MethodProcessor from sampleStruct by method name.
func newMethodProcessor(methodName string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(sampleStruct{})

	method, ok := t.MethodByName(methodName)
	if !ok {
		return nil
	}

	return &reflectmodel.MethodProcessor{
		Name:          method.Name,
		Index:         method.Index,
		ReflectMethod: method,
	}
}
