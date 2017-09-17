package wex

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var wex = API{}

func TestWEX(t *testing.T) {

	Convey("WEX instance created", t, func() {

		Convey("Public API should be available", func() {
			So(wex.Public, ShouldNotBeNil)
		})

		Convey("Trade API should be available", func() {
			So(wex.Trade, ShouldNotBeNil)
		})
	})
}
