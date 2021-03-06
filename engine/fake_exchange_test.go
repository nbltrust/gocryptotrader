package engine

import (
	"sync"
	"time"

	"github.com/nbltrust/gocryptotrader/config"
	"github.com/nbltrust/gocryptotrader/currency"
	exchange "github.com/nbltrust/gocryptotrader/exchanges"
	"github.com/nbltrust/gocryptotrader/exchanges/account"
	"github.com/nbltrust/gocryptotrader/exchanges/asset"
	"github.com/nbltrust/gocryptotrader/exchanges/kline"
	"github.com/nbltrust/gocryptotrader/exchanges/order"
	"github.com/nbltrust/gocryptotrader/exchanges/orderbook"
	"github.com/nbltrust/gocryptotrader/exchanges/ticker"
	"github.com/nbltrust/gocryptotrader/exchanges/websocket/wshandler"
	"github.com/nbltrust/gocryptotrader/portfolio/withdraw"
)

const (
	fakePassExchange = "FakePassExchange"
)

// FakePassingExchange is used to override IBotExchange responses in tests
// In this context, we don't care what FakePassingExchange does as we're testing
// the engine package
type FakePassingExchange struct {
	exchange.Base
}

// addPassingFakeExchange adds an exchange to engine tests where all funcs return a positive result
func addPassingFakeExchange(baseExchangeName string) error {
	testExch := GetExchangeByName(baseExchangeName)
	if testExch == nil {
		return ErrExchangeNotFound
	}
	base := testExch.GetBase()
	Bot.Config.Exchanges = append(Bot.Config.Exchanges, config.ExchangeConfig{
		Name:    fakePassExchange,
		Enabled: true,
		Verbose: false,
	})

	Bot.exchangeManager.add(&FakePassingExchange{
		Base: exchange.Base{
			Name:                          fakePassExchange,
			Enabled:                       true,
			LoadedByConfig:                true,
			SkipAuthCheck:                 true,
			API:                           base.API,
			Features:                      base.Features,
			HTTPTimeout:                   base.HTTPTimeout,
			HTTPUserAgent:                 base.HTTPUserAgent,
			HTTPRecording:                 base.HTTPRecording,
			HTTPDebugging:                 base.HTTPDebugging,
			WebsocketResponseCheckTimeout: base.WebsocketResponseCheckTimeout,
			WebsocketResponseMaxLimit:     base.WebsocketResponseMaxLimit,
			WebsocketOrderbookBufferLimit: base.WebsocketOrderbookBufferLimit,
			Websocket:                     base.Websocket,
			Requester:                     base.Requester,
			Config:                        base.Config,
		},
	})
	return nil
}

func (h *FakePassingExchange) Setup(_ *config.ExchangeConfig) error { return nil }
func (h *FakePassingExchange) Start(_ *sync.WaitGroup)              {}
func (h *FakePassingExchange) SetDefaults()                         {}
func (h *FakePassingExchange) GetName() string                      { return fakePassExchange }
func (h *FakePassingExchange) IsEnabled() bool                      { return true }
func (h *FakePassingExchange) SetEnabled(bool)                      {}
func (h *FakePassingExchange) ValidateCredentials() error           { return nil }

func (h *FakePassingExchange) FetchTicker(_ currency.Pair, _ asset.Item) (*ticker.Price, error) {
	return nil, nil
}
func (h *FakePassingExchange) UpdateTicker(_ currency.Pair, _ asset.Item) (*ticker.Price, error) {
	return nil, nil
}
func (h *FakePassingExchange) FetchOrderbook(_ currency.Pair, _ asset.Item) (*orderbook.Base, error) {
	return nil, nil
}
func (h *FakePassingExchange) UpdateOrderbook(_ currency.Pair, _ asset.Item) (*orderbook.Base, error) {
	return nil, nil
}
func (h *FakePassingExchange) FetchTradablePairs(_ asset.Item) ([]string, error) {
	return nil, nil
}
func (h *FakePassingExchange) UpdateTradablePairs(_ bool) error { return nil }

func (h *FakePassingExchange) GetEnabledPairs(_ asset.Item) currency.Pairs {
	return currency.Pairs{}
}
func (h *FakePassingExchange) GetAvailablePairs(_ asset.Item) currency.Pairs {
	return currency.Pairs{}
}
func (h *FakePassingExchange) FetchAccountInfo() (account.Holdings, error) {
	return account.Holdings{}, nil
}

func (h *FakePassingExchange) UpdateAccountInfo() (account.Holdings, error) {
	return account.Holdings{}, nil
}
func (h *FakePassingExchange) GetAuthenticatedAPISupport(_ uint8) bool { return true }
func (h *FakePassingExchange) SetPairs(_ currency.Pairs, _ asset.Item, _ bool) error {
	return nil
}
func (h *FakePassingExchange) GetAssetTypes() asset.Items { return asset.Items{asset.Spot} }
func (h *FakePassingExchange) GetExchangeHistory(_ currency.Pair, _ asset.Item) ([]exchange.TradeHistory, error) {
	return nil, nil
}
func (h *FakePassingExchange) SupportsAutoPairUpdates() bool        { return true }
func (h *FakePassingExchange) SupportsRESTTickerBatchUpdates() bool { return true }
func (h *FakePassingExchange) GetFeeByType(_ *exchange.FeeBuilder) (float64, error) {
	return 0, nil
}
func (h *FakePassingExchange) GetLastPairsUpdateTime() int64                      { return 0 }
func (h *FakePassingExchange) GetWithdrawPermissions() uint32                     { return 0 }
func (h *FakePassingExchange) FormatWithdrawPermissions() string                  { return "" }
func (h *FakePassingExchange) SupportsWithdrawPermissions(_ uint32) bool          { return true }
func (h *FakePassingExchange) GetFundingHistory() ([]exchange.FundHistory, error) { return nil, nil }
func (h *FakePassingExchange) SubmitOrder(_ *order.Submit) (order.SubmitResponse, error) {
	return order.SubmitResponse{
		IsOrderPlaced: true,
		FullyMatched:  true,
		OrderID:       "FakePassingExchangeOrder",
	}, nil
}
func (h *FakePassingExchange) ModifyOrder(_ *order.Modify) (string, error) { return "", nil }
func (h *FakePassingExchange) CancelOrder(_ *order.Cancel) error           { return nil }
func (h *FakePassingExchange) CancelAllOrders(_ *order.Cancel) (order.CancelAllResponse, error) {
	return order.CancelAllResponse{}, nil
}
func (h *FakePassingExchange) GetOrderInfo(_ string) (order.Detail, error) {
	return order.Detail{}, nil
}
func (h *FakePassingExchange) GetDepositAddress(_ currency.Code, _ string) (string, error) {
	return "", nil
}
func (h *FakePassingExchange) GetOrderHistory(_ *order.GetOrdersRequest) ([]order.Detail, error) {
	return nil, nil
}
func (h *FakePassingExchange) GetActiveOrders(_ *order.GetOrdersRequest) ([]order.Detail, error) {
	return []order.Detail{
		{
			Price:     1337,
			Amount:    1337,
			Exchange:  fakePassExchange,
			ID:        "fakeOrder",
			Type:      order.Market,
			Side:      order.Buy,
			Status:    order.Active,
			AssetType: asset.Spot,
			Date:      time.Now(),
			Pair:      currency.NewPairFromString("BTCUSD"),
		},
	}, nil
}
func (h *FakePassingExchange) SetHTTPClientUserAgent(_ string)             {}
func (h *FakePassingExchange) GetHTTPClientUserAgent() string              { return "" }
func (h *FakePassingExchange) SetClientProxyAddress(_ string) error        { return nil }
func (h *FakePassingExchange) SupportsWebsocket() bool                     { return true }
func (h *FakePassingExchange) SupportsREST() bool                          { return true }
func (h *FakePassingExchange) IsWebsocketEnabled() bool                    { return true }
func (h *FakePassingExchange) GetWebsocket() (*wshandler.Websocket, error) { return nil, nil }
func (h *FakePassingExchange) SubscribeToWebsocketChannels(_ []wshandler.WebsocketChannelSubscription) error {
	return nil
}
func (h *FakePassingExchange) UnsubscribeToWebsocketChannels(_ []wshandler.WebsocketChannelSubscription) error {
	return nil
}
func (h *FakePassingExchange) AuthenticateWebsocket() error { return nil }
func (h *FakePassingExchange) GetSubscriptions() ([]wshandler.WebsocketChannelSubscription, error) {
	return nil, nil
}
func (h *FakePassingExchange) GetDefaultConfig() (*config.ExchangeConfig, error) { return nil, nil }
func (h *FakePassingExchange) GetBase() *exchange.Base                           { return nil }
func (h *FakePassingExchange) SupportsAsset(_ asset.Item) bool                   { return true }
func (h *FakePassingExchange) GetHistoricCandles(_ currency.Pair, _ asset.Item, _, _ time.Time, _ time.Duration) (kline.Item, error) {
	return kline.Item{}, nil
}
func (h *FakePassingExchange) DisableRateLimiter() error { return nil }
func (h *FakePassingExchange) EnableRateLimiter() error  { return nil }
func (h *FakePassingExchange) WithdrawCryptocurrencyFunds(_ *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	return nil, nil
}
func (h *FakePassingExchange) WithdrawFiatFunds(_ *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	return nil, nil
}
func (h *FakePassingExchange) WithdrawFiatFundsToInternationalBank(_ *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	return nil, nil
}
