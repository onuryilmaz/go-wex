package btce

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTicker(t *testing.T) {

	Convey("Ticker data for BTC-USD", t, func() {
		tickers, err := GetTicker([]string{"btc_usd"})

		Convey("No error should occur", func() {
			So(err, ShouldBeNil)
		})

		Convey("One ticker information should be returned", func() {
			So(tickers, ShouldHaveSameTypeAs, Ticker{})
			So(tickers, ShouldContainKey, "btc_usd")
			So(tickers["btc_usd"], ShouldHaveSameTypeAs, TickerPair{})
		})
	})

}

func TestInfo(t *testing.T) {

	Convey("Information data", t, func() {
		information, err := GetInfo()

		Convey("No error should occur", func() {
			So(err, ShouldBeNil)
		})

		Convey("Information data for 'btc_usd' should be returned", func() {
			So(information, ShouldHaveSameTypeAs, Info{})
			So(information.Pairs["btc_usd"], ShouldHaveSameTypeAs, PairInfo{})
		})
	})

}

func TestDepth(t *testing.T) {

	Convey("Depth data", t, func() {
		depth, err := GetDepth([]string{"btc_usd"})

		Convey("No error should occur", func() {
			So(err, ShouldBeNil)
		})

		Convey("Depth data for 'btc_usd' should be returned", func() {
			So(depth, ShouldHaveSameTypeAs, Depth{})
			So(depth["btc_usd"], ShouldHaveSameTypeAs, DepthPair{})
		})
	})

}

func TestTrade(t *testing.T) {

	Convey("Trade data", t, func() {
		trade, err := GetTrade([]string{"btc_usd"})

		Convey("No error should occur", func() {
			So(err, ShouldBeNil)
		})

		Convey("Trade data for 'btc_usd' should be returned", func() {
			So(trade, ShouldHaveSameTypeAs, Trade{})
			So(len(trade["btc_usd"]), ShouldBeGreaterThanOrEqualTo, 0)
			//So(trade["btc_usd"][0].TID, ShouldBeGreaterThanOrEqualTo, 0)
		})
	})

}
