package engine

import (
	"errors"
	"testing"
	"time"

	"github.com/nbltrust/gocryptotrader/exchanges/order"
	"github.com/nbltrust/gocryptotrader/exchanges/sharedtestvalues"
	"github.com/nbltrust/gocryptotrader/exchanges/ticker"
	"github.com/nbltrust/gocryptotrader/exchanges/websocket/wshandler"
)

func TestWebsocketDataHandlerProcess(t *testing.T) {
	ws := wshandler.New()
	err := ws.Setup(&wshandler.WebsocketSetup{Enabled: true})
	if err != nil {
		t.Error(err)
	}
	ws.DataHandler = sharedtestvalues.GetWebsocketInterfaceChannelOverride()
	go WebsocketDataReceiver(ws)
	ws.DataHandler <- "string"
	time.Sleep(time.Second)
	close(shutdowner)
}

func TestHandleData(t *testing.T) {
	OrdersSetup(t)
	var exchName = "exch"
	var orderID = "testOrder.Detail"
	err := WebsocketDataHandler(exchName, errors.New("error"))
	if err == nil {
		t.Error("Error not handled correctly")
	}
	err = WebsocketDataHandler(exchName, nil)
	if err == nil {
		t.Error("Expected nil data error")
	}
	err = WebsocketDataHandler(exchName, wshandler.TradeData{})
	if err != nil {
		t.Error(err)
	}
	err = WebsocketDataHandler(exchName, wshandler.FundingData{})
	if err != nil {
		t.Error(err)
	}
	err = WebsocketDataHandler(exchName, &ticker.Price{})
	if err != nil {
		t.Error(err)
	}
	err = WebsocketDataHandler(exchName, wshandler.KlineData{})
	if err != nil {
		t.Error(err)
	}
	err = WebsocketDataHandler(exchName, wshandler.WebsocketOrderbookUpdate{})
	if err != nil {
		t.Error(err)
	}
	origOrder := &order.Detail{
		Exchange: fakePassExchange,
		ID:       orderID,
		Amount:   1337,
		Price:    1337,
	}
	err = WebsocketDataHandler(exchName, origOrder)
	if err != nil {
		t.Error(err)
	}
	// Send it again since it exists now
	err = WebsocketDataHandler(exchName, &order.Detail{
		Exchange: fakePassExchange,
		ID:       orderID,
		Amount:   1338,
	})
	if err != nil {
		t.Error(err)
	}
	if origOrder.Amount != 1338 {
		t.Error("Bad pipeline")
	}

	err = WebsocketDataHandler(exchName, &order.Modify{
		Exchange: fakePassExchange,
		ID:       orderID,
		Status:   order.Active,
	})
	if err != nil {
		t.Error(err)
	}
	if origOrder.Status != order.Active {
		t.Error("Expected order to be modified to Active")
	}

	err = WebsocketDataHandler(exchName, &order.Cancel{
		Exchange: fakePassExchange,
		ID:       orderID,
	})
	if err != nil {
		t.Error(err)
	}
	if origOrder.Status != order.Cancelled {
		t.Error("Expected order status to be cancelled")
	}
	// Send some gibberish
	err = WebsocketDataHandler(exchName, order.Stop)
	if err != nil {
		t.Error(err)
	}

	err = WebsocketDataHandler(exchName, wshandler.UnhandledMessageWarning{Message: "there's an issue here's a tissue"})
	if err != nil {
		t.Error(err)
	}

	classificationError := order.ClassificationError{
		Exchange: "test",
		OrderID:  "one",
		Err:      errors.New("lol"),
	}
	err = WebsocketDataHandler(exchName, classificationError)
	if err == nil {
		t.Error("Expected error")
	}
	if err.Error() != classificationError.Error() {
		t.Errorf("Problem formatting error. Expected %v Received %v", classificationError.Error(), err.Error())
	}
}
