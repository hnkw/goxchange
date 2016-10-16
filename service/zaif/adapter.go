package zaif

import "github.com/hnkw/goxchange/goxchange"

// NewAdapter returns a instance of zaif.Adapter
func NewAdapter(exc *Exchange, client *Client) *Adapter {
	return &Adapter{
		exc:    exc,
		client: client,
	}
}

// Adapter translate methods from zaif to Exchange interface
type Adapter struct {
	exc    *Exchange
	client *Client
}

// Deposit is return deposit
func (a *Adapter) Deposit() (*goxchange.Deposit, error) {
	resp, err := a.client.GetInfo()
	if err != nil {
		return nil, err
	}

	return &goxchange.Deposit{
		JPY: resp.Return.Deposit.JPY,
		BTC: resp.Return.Deposit.BTC,
	}, nil
}

// Connect send goxchange.Books continuously
func (a *Adapter) Connect() (<-chan *goxchange.Books, <-chan error, error) {
	streamCh, errorCh, err := a.exc.Connect()

	oc := make(chan *goxchange.Books)
	ec := make(chan error)

	go func() {
		for {
			select {
			case stream := <-streamCh:
				b := goxchange.Books{
					Asks:      a.convert(stream.Asks),
					Bids:      a.convert(stream.Bids),
					TimeStamp: stream.TimeStamp,
				}
				oc <- &b
			case err := <-errorCh:
				ec <- err
				return
			}
		}
	}()

	return oc, ec, err
}

func (a *Adapter) convert(orders [][]float64) goxchange.Orders {
	ret := make(goxchange.Orders, 0, len(orders))

	for _, order := range orders {
		ret = append(ret, goxchange.Order(order))
	}

	return ret
}

// Ask exchange BTC to JPY
func (a *Adapter) Ask(o goxchange.Order) error {
	req := TradeReq{
		CurrencyPair: "btc_jpy",
		Action:       "ask",
		Price:        o.Price(),
		Amount:       o.Amount(),
	}
	_, err := a.client.Trade(&req)
	return err

}

// Bit exchange JPY to BTC
func (a *Adapter) Bit(o goxchange.Order) error {
	req := TradeReq{
		CurrencyPair: "btc_jpy",
		Action:       "bit",
		Price:        o.Price(),
		Amount:       o.Amount(),
	}
	_, err := a.client.Trade(&req)
	return err
}
