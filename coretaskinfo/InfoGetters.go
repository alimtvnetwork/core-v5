package coretaskinfo

// =============================================================================
// IsInclude checks — returns true when field is present AND not excluded
// =============================================================================

func (it *Info) IsIncludeRootName() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeRootName() &&
		it.RootName != ""
}

func (it *Info) IsIncludeDescription() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeDescription() &&
		it.Description != ""
}

func (it *Info) IsIncludeUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeUrl() &&
		it.Url != ""
}

func (it *Info) IsIncludeHintUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeHintUrl() &&
		it.HintUrl != ""
}

func (it *Info) IsIncludeErrorUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeErrorUrl() &&
		it.ErrorUrl != ""
}

// IsIncludeAdditionalErrorWrap
//
//	returns true on null or it.ExcludeOptions.IsIncludeAdditionalErrorWrap
func (it *Info) IsIncludeAdditionalErrorWrap() bool {
	return it == nil ||
		it.ExcludeOptions.IsIncludeAdditionalErrorWrap()
}

func (it *Info) IsIncludeExampleUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeExampleUrl() &&
		it.ExampleUrl != ""
}

func (it *Info) IsIncludeSingleExample() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeSingleExample() &&
		it.SingleExample != ""
}

func (it *Info) IsIncludeExamples() bool {
	return it != nil &&
		it.ExcludeOptions.IsIncludeExamples() &&
		len(it.Examples) > 0
}

// =============================================================================
// Secure / Plain / Payload visibility
// =============================================================================

func (it *Info) IsSecure() bool {
	return it != nil && it.ExcludeOptions.IsSafeSecureText()
}

func (it *Info) IsPlainText() bool {
	return it == nil || it.ExcludeOptions.IsIncludePayloads()
}

func (it *Info) IsIncludePayloads() bool {
	return it == nil || it.ExcludeOptions.IsIncludePayloads()
}

func (it *Info) IsExcludePayload() bool {
	return it != nil && it.ExcludeOptions.IsSafeSecureText()
}

// =============================================================================
// IsExclude checks — returns true on defined (not null) and excluded
// =============================================================================

// IsExcludeRootName
//
//	returns true on defined (not null) and
//	it.ExcludeOptions.IsExcludeRootName
//
//	return false on null
func (it *Info) IsExcludeRootName() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeRootName()
}

// IsExcludeDescription
//
//	returns true on defined (not null) and
//	it.ExcludeOptions.IsExcludeDescription
//
//	return false on null
func (it *Info) IsExcludeDescription() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeDescription()
}

// IsExcludeUrl
//
//	returns true on defined (not null) and
//	it.ExcludeOptions.IsExcludeUrl
//
//	return false on null
func (it *Info) IsExcludeUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeUrl()
}

// IsExcludeHintUrl
//
//	returns true on defined (not null) and
//	it.ExcludeOptions.IsExcludeHintUrl
//
//	return false on null
func (it *Info) IsExcludeHintUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeHintUrl()
}

// IsExcludeErrorUrl
//
//	returns true on defined (not null) and
//	it.ExcludeOptions.IsExcludeErrorUrl
//
//	return false on null
func (it *Info) IsExcludeErrorUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeErrorUrl()
}

// IsExcludeAdditionalErrorWrap
//
//	returns true on defined (not null) and
//	it.ExcludeOptions.IsExcludeAdditionalErrorWrap
//
//	return false on null
func (it *Info) IsExcludeAdditionalErrorWrap() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeAdditionalErrorWrap()
}

// IsExcludeExampleUrl
//
//	return true on null
func (it *Info) IsExcludeExampleUrl() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeExampleUrl()
}

// IsExcludeSingleExample
//
//	return true on null
func (it *Info) IsExcludeSingleExample() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeSingleExample()
}

// IsExcludeExamples
//
//	return true on null
func (it *Info) IsExcludeExamples() bool {
	return it != nil &&
		it.ExcludeOptions.IsSafeExcludeExamples()
}

// =============================================================================
// Safe* getters — nil-safe field accessors
// =============================================================================

func (it *Info) SafeName() string {
	if it.IsNull() {
		return ""
	}

	return it.RootName
}

func (it *Info) SafeDescription() string {
	if it.IsNull() {
		return ""
	}

	return it.Description
}

func (it *Info) SafeUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.Url
}

func (it *Info) SafeHintUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.HintUrl
}

func (it *Info) SafeErrorUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.ErrorUrl
}

func (it *Info) SafeExampleUrl() string {
	if it.IsNull() {
		return ""
	}

	return it.ExampleUrl
}

func (it *Info) SafeChainingExample() string {
	if it.IsNull() {
		return ""
	}

	return it.ExampleUrl
}

// =============================================================================
// Has* checks — returns true if field is non-empty
// =============================================================================

func (it *Info) HasRootName() bool {
	return it != nil && it.RootName != ""
}

func (it *Info) HasDescription() bool {
	return it != nil && it.Description != ""
}

func (it *Info) HasUrl() bool {
	return it != nil && it.Url != ""
}

func (it *Info) HasHintUrl() bool {
	return it != nil && it.HintUrl != ""
}

func (it *Info) HasErrorUrl() bool {
	return it != nil && it.ErrorUrl != ""
}

func (it *Info) HasExampleUrl() bool {
	return it != nil && it.ExampleUrl != ""
}

func (it *Info) HasChainingExample() bool {
	return it != nil && it.SingleExample != ""
}

func (it *Info) HasExamples() bool {
	return it != nil && len(it.Examples) > 0
}

func (it *Info) HasExcludeOptions() bool {
	return it != nil && !it.ExcludeOptions.IsEmpty()
}

// =============================================================================
// IsEmpty* checks — returns true if field is absent or nil
// =============================================================================

func (it *Info) IsEmptyName() bool {
	return it == nil || it.RootName == ""
}

func (it *Info) IsEmptyDescription() bool {
	return it == nil || it.Description == ""
}

func (it *Info) IsEmptyUrl() bool {
	return it == nil || it.Url == ""
}

func (it *Info) IsEmptyHintUrl() bool {
	return it == nil || it.HintUrl == ""
}

func (it *Info) IsEmptyErrorUrl() bool {
	return it == nil || it.ErrorUrl == ""
}

func (it *Info) IsEmptyExampleUrl() bool {
	return it == nil || it.ExampleUrl == ""
}

func (it *Info) IsEmptySingleExample() bool {
	return it == nil || it.SingleExample == ""
}

func (it *Info) IsEmptyExamples() bool {
	return it == nil || len(it.Examples) == 0
}

func (it *Info) IsEmptyExcludeOptions() bool {
	return it == nil || it.ExcludeOptions.IsZero()
}
