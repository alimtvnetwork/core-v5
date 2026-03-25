package coretaskinfo

import (
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// Info holds metadata about a task including name, description, URLs, examples,
// and exclude/include display options.
//
// Getters and field checks are in InfoGetters.go.
// JSON serialization and deserialization are in InfoJson.go.
// Map builders are in InfoMap.go.
type Info struct {
	RootName       string            `json:"RootName,omitempty"`
	Description    string            `json:"Description,omitempty"`
	Url            string            `json:"Url,omitempty"`
	HintUrl        string            `json:"HintUrl,omitempty"`
	ErrorUrl       string            `json:"ErrorUrl,omitempty"`
	ExampleUrl     string            `json:"ExampleUrl,omitempty"`
	SingleExample  string            `json:"SingleExample,omitempty"`
	Examples       []string          `json:"Examples,omitempty"` // proves sample examples to call things correctly
	ExcludeOptions *ExcludingOptions `json:"ExcludeOptions,omitempty"`
	lazyMap        map[string]string
}

// =============================================================================
// Constructors / Mutators
// =============================================================================

// SetSecure
//
//	on nil creates and returns new info with secure flag
func (it *Info) SetSecure() *Info {
	if it == nil {
		return &Info{
			ExcludeOptions: &ExcludingOptions{
				IsSecureText: true,
			},
		}
	}

	it.ExcludeOptions = it.
		ExcludeOptions.
		SetSecure()

	return it
}

// SetPlain
//
//	on nil creates and returns
//	new info which is plain not secure
func (it *Info) SetPlain() *Info {
	if it == nil {
		return &Info{} // plain text
	}

	it.ExcludeOptions = it.
		ExcludeOptions.
		SetPlainText()

	return it
}

// =============================================================================
// Core Identity
// =============================================================================

func (it *Info) Name() string {
	if it.IsNull() {
		return ""
	}

	return it.RootName
}

func (it *Info) IsNull() bool {
	return it == nil
}

func (it *Info) IsDefined() bool {
	return it != nil && it.RootName != ""
}

func (it *Info) HasAnyName() bool {
	return it != nil && it.RootName != ""
}

func (it *Info) IsName(name string) bool {
	return it != nil && it.RootName == name
}

func (it *Info) IsEmpty() bool {
	return it == nil
}

func (it *Info) HasAnyItem() bool {
	return it != nil
}

func (it *Info) Options() *ExcludingOptions {
	if it == nil {
		return &ExcludingOptions{}
	}

	return it.ExcludeOptions
}

func (it *Info) ExamplesAsSlice() *corestr.SimpleSlice {
	if it.IsNull() {
		return corestr.Empty.SimpleSlice()
	}

	return corestr.New.SimpleSlice.Strings(it.Examples)
}

// =============================================================================
// Clone
// =============================================================================

func (it Info) ToPtr() *Info {
	return &it
}

func (it Info) ToNonPtr() Info {
	return it
}

func (it Info) Clone() Info {
	return Info{
		RootName:       it.RootName,
		Description:    it.Description,
		Url:            it.Url,
		HintUrl:        it.HintUrl,
		ErrorUrl:       it.ErrorUrl,
		ExampleUrl:     it.ExampleUrl,
		SingleExample:  it.SingleExample,
		Examples:       it.Examples,
		ExcludeOptions: it.ExcludeOptions.ClonePtr(),
	}
}

func (it *Info) ClonePtr() *Info {
	if it == nil {
		return nil
	}

	return &Info{
		RootName:       it.RootName,
		Description:    it.Description,
		Url:            it.Url,
		HintUrl:        it.HintUrl,
		ErrorUrl:       it.ErrorUrl,
		ExampleUrl:     it.ExampleUrl,
		SingleExample:  it.SingleExample,
		Examples:       it.Examples,
		ExcludeOptions: it.ExcludeOptions.ClonePtr(),
	}
}
