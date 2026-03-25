package coretests

import "github.com/alimtvnetwork/core/internal/msgcreator"

// IsEqualMessage
//
// Gives generic and consistent
// test message using msgformats.IsEqualMessageFormat
func (it getAssert) IsEqualMessage(
	when,
	actual,
	expected any,
) string {
	return msgcreator.Assert.IsEqualMessage(
		when,
		actual,
		expected,
	)
}

// IsNotEqualMessage
//
// Gives generic and consistent
// test message using msgformats.IsNotEqualMessageFormat
func (it getAssert) IsNotEqualMessage(
	when,
	actual,
	expected any,
) string {
	return msgcreator.Assert.IsNotEqualMessage(
		when,
		actual,
		expected,
	)
}

// IsTrueMessage
//
// Gives generic and consistent
// test message using msgformats.IsTrueMessageFormat
func (it getAssert) IsTrueMessage(
	when,
	actual any,
) string {
	return msgcreator.Assert.IsTrueMessage(
		when,
		actual,
	)
}

// IsFalseMessage
//
// Gives generic and consistent
// test message using msgformats.IsFalseMessageFormat
func (it getAssert) IsFalseMessage(
	when,
	actual any,
) string {
	return msgcreator.Assert.IsFalseMessage(
		when,
		actual,
	)
}

// IsNilMessage
//
// Gives generic and consistent
// test message using msgformats.IsNilMessageFormat
func (it getAssert) IsNilMessage(
	when,
	actual any,
) string {
	return msgcreator.Assert.IsNilMessage(
		when,
		actual,
	)
}

// IsNotNilMessage
//
// Gives generic and consistent
// test message using msgformats.IsNotNilMessageFormat
func (it getAssert) IsNotNilMessage(
	when,
	actual any,
) string {
	return msgcreator.Assert.IsNotNilMessage(
		when,
		actual,
	)
}

// ShouldBeMessage
//
// Gives generic and consistent
// test message using msgformats.ShouldBeMessageFormat
func (it getAssert) ShouldBeMessage(
	title string,
	actual,
	expected any,
) string {
	return msgcreator.Assert.ShouldBeMessage(
		title,
		actual,
		expected,
	)
}

// ShouldNotBeMessage
//
// Gives generic and consistent
// test message using msgformats.ShouldNotBeMessageFormat
func (it getAssert) ShouldNotBeMessage(
	title string,
	actual,
	expected any,
) string {
	return msgcreator.Assert.ShouldNotBeMessage(
		title,
		actual,
		expected,
	)
}
