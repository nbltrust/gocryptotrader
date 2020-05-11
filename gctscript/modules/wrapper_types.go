package modules

import (
	"time"

	"github.com/nbltrust/gocryptotrader/currency"
	"github.com/nbltrust/gocryptotrader/exchanges/account"
	"github.com/nbltrust/gocryptotrader/exchanges/asset"
	"github.com/nbltrust/gocryptotrader/exchanges/kline"
	"github.com/nbltrust/gocryptotrader/exchanges/order"
	"github.com/nbltrust/gocryptotrader/exchanges/orderbook"
	"github.com/nbltrust/gocryptotrader/exchanges/ticker"
	"github.com/nbltrust/gocryptotrader/portfolio/withdraw"
)

const (
	// ErrParameterConvertFailed error to return when type conversion fails
	ErrParameterConvertFailed             = "%v failed conversion"
	ErrParameterWithPositionConvertFailed = "%v at position %v failed conversion"
)

// Wrapper instance of GCT to use for modules
var Wrapper GCT

// GCT interface requirements
type GCT interface {
	Exchange
}

// Exchange interface requirements
type Exchange interface {
	Exchanges(enabledOnly bool) []string
	IsEnabled(exch string) bool
	Orderbook(exch string, pair currency.Pair, item asset.Item) (*orderbook.Base, error)
	Ticker(exch string, pair currency.Pair, item asset.Item) (*ticker.Price, error)
	Pairs(exch string, enabledOnly bool, item asset.Item) (*currency.Pairs, error)
	QueryOrder(exch, orderid string) (*order.Detail, error)
	SubmitOrder(submit *order.Submit) (*order.SubmitResponse, error)
	CancelOrder(exch, orderid string) (bool, error)
	AccountInformation(exch string) (account.Holdings, error)
	DepositAddress(exch string, currencyCode currency.Code) (string, error)
	WithdrawalFiatFunds(exch, bankAccountID string, request *withdraw.Request) (out string, err error)
	WithdrawalCryptoFunds(exch string, request *withdraw.Request) (out string, err error)
	OHLCV(exch string, pair currency.Pair, item asset.Item, start, end time.Time, interval time.Duration) (kline.Item, error)
}

// SetModuleWrapper link the wrapper and interface to use for modules
func SetModuleWrapper(wrapper GCT) {
	Wrapper = wrapper
}
