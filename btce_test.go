package btce

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var btce = API{}

func TestBTCE(t *testing.T) {

	Convey("BTCE instance created", t, func() {

		Convey("Public API should be available", func() {
			So(btce.Public, ShouldNotBeNil)
		})

		Convey("Trade API should be available", func() {
			So(btce.Trade, ShouldNotBeNil)
		})
	})
}
