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
	wg.Add(len(exchanges))
	for x := range exchanges {
		go func(x int, wg *sync.WaitGroup) {
			defer wg.Done()
			if !exchanges[x].SupportsREST() {
				return
			}
			assets := exchanges[x].GetAssetTypes()
			for y := range assets {
				p := exchanges[x].GetAvailablePairs(assets[y]).GetRandomPair()
				var ob *orderbook.Base
				ob, err = exchanges[x].UpdateOrderbook(p, assets[y])
				if ob == nil {
					continue
				}
				if err != nil && err.Error() == orderbook.ErrNoOrderbook {
					t.Errorf("%s %s %s orderbook error: empty orderbook\n", exchanges[x].GetName(), assets[y], p)
					continue
				}
				if len(ob.Bids) == 0 {
					t.Errorf("%s %s %s orderbook error: empty bids\n", exchanges[x].GetName(), assets[y], p)
				}
				if len(ob.Asks) == 0 {
					t.Errorf("%s %s %s orderbook error: empty asks\n", exchanges[x].GetName(), assets[y], p)
				}
			}
		}(x, &wg)
	}
	wg.Wait()
}
