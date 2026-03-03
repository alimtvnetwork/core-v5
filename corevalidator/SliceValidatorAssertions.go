package corevalidator

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

func (it *SliceValidator) AssertAllQuick(
	t *testing.T,
	caseIndex int,
	header string,
	actualElements ...string,
) {
	if it == nil {
		return
	}

	toErr := it.AllVerifyErrorQuick(
		caseIndex,
		header,
		actualElements...,
	)

	convey.Convey(
		header, t, func() {
			convey.So(
				toErr,
				should.BeNil,
			)
		},
	)
}
