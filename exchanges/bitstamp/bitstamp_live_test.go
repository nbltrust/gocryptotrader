//+build mock_test_off

// This will build if build tag mock_test_off is parsed and will do live testing
// using all tests in (exchange)_test.go
package bitstamp

import (
	"log"
	"os"
	"testing"

	"github.com/nbltrust/gocryptotrader/config"
	"github.com/nbltrust/gocryptotrader/exchanges/sharedtestvalues"
)

var mockTests = false

func TestMain(m *testing.M) {
	cfg := config.GetConfig()
	err := cfg.LoadConfig("../../testdata/configtest.json", true)
	if err != nil {
		log.Fatal("Bitstamp load config error", err)
	}
	bitstampConfig, err := cfg.GetExchangeConfig("Bitstamp")
	if err != nil {
		log.Fatal("Bitstamp Setup() init error", err)
	}
	bitstampConfig.API.AuthenticatedSupport = true
	bitstampConfig.API.Credentials.Key = apiKey
	bitstampConfig.API.Credentials.Secret = apiSecret
	bitstampConfig.API.Credentials.ClientID = customerID
	b.SetDefaults()
	err = b.Setup(bitstampConfig)
	if err != nil {
		log.Fatal("Bitstamp setup error", err)
	}
	b.Websocket.DataHandler = sharedtestvalues.GetWebsocketInterfaceChannelOverride()
	b.Websocket.TrafficAlert = sharedtestvalues.GetWebsocketStructChannelOverride()
	log.Printf(sharedtestvalues.LiveTesting, b.Name, b.API.Endpoints.URL)
	os.Exit(m.Run())
}
