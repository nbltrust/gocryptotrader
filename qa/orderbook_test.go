package qa

import (
	"sync"
	"testing"

	"github.com/thrasher-corp/gocryptotrader/engine"
	exchange "github.com/thrasher-corp/gocryptotrader/exchanges"
	"github.com/thrasher-corp/gocryptotrader/exchanges/orderbook"
)

func TestRESTEmptyOrderbook(t *testing.T) {
	var err error
	engine.Bot, err = engine.New()
	if err != nil {
		t.Fatal(err)
	}

	engine.Bot.Settings = engine.Settings{
		EnableExchangeHTTPRateLimiter: true,
	}

	var wg sync.WaitGroup
	for x := range exchange.Exchanges {
		name := exchange.Exchanges[x]
		err = engine.LoadExchange(exchange.Exchanges[x], true, &wg)
		if err != nil {
			t.Errorf("Failed to load exchange %s. Error: %s", name, err)
		}
	}
	wg.Wait()

	exchanges := engine.GetExchanges()
	for x := range exchanges {
		if !exchanges[x].SupportsREST() {
			continue
		}
		wg.Add(1)
		assets := exchanges[x].GetAssetTypes()
		for y := range assets {
			p := exchanges[x].GetAvailablePairs(assets[y]).GetRandomPair()
			_, err = exchanges[x].UpdateOrderbook(p, assets[y])
			if err != nil && err.Error() == orderbook.ErrNoOrderbook {
				t.Errorf("%s %s %s orderbook error: empty orderbook\n", exchanges[x].GetName(), assets[y], p)
			}
		}
		wg.Done()
	}
	wg.Wait()
}
