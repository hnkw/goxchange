package zaif

import (
	"encoding/json"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

const (
	origin = "http://localhost/"
)

// StreamTrade information
type StreamTrade struct {
	CurrencyPair string  `json:"currenty_pair"`
	TradeType    string  `json:"trade_type"`
	Price        float64 `json:"price"`
	TID          uint64  `json:"tid"`
	Amount       float64 `json:"amount"`
	Date         uint64  `json:"date"`
}

// LastPrice is last price information
type LastPrice struct {
	Action string  `json:"action"`
	Price  float64 `json:"price"`
}

// Stream is current trades information
type Stream struct {
	Asks         [][]float64
	Bids         [][]float64
	Trades       []StreamTrade
	TimeStamp    time.Time
	LastPrice    LastPrice
	CurrencyPair string
}

// UnmarshalJSON implements unmarshaler for Stream
func (e *Stream) UnmarshalJSON(body []byte) error {
	var f struct {
		Asks         [][]float64   `json:"asks"`
		Bids         [][]float64   `json:"bids"`
		Trades       []StreamTrade `json:"trades"`
		TimeStamp    string        `json:"timestamp"`
		LastPrice    LastPrice     `json:"last_price"`
		CurrencyPair string        `json:"currency_pair"`
	}

	if err := json.Unmarshal(body, &f); err != nil {
		return err
	}

	e.Asks = f.Asks
	e.Bids = f.Bids
	e.Trades = f.Trades

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	t, err := time.ParseInLocation("2006-01-02 15:04:05.999999", f.TimeStamp, jst)
	if err != nil {
		return err
	}

	e.TimeStamp = t.UTC()

	e.LastPrice = f.LastPrice
	e.CurrencyPair = f.CurrencyPair

	return nil
}

// Exchange is exchange API for zaif
type Exchange struct {
	c  *Confing
	ws *websocket.Conn
}

// NewExchange construct and initialize Exchange struct
func NewExchange(c *Confing) *Exchange {
	return &Exchange{
		c: c,
	}
}

// NewDefaultExchange create Exchange instance with default settings
func NewDefaultExchange() *Exchange {
	return &Exchange{
		c: overrideEnvVal(&DefaultConf),
	}
}

// Connect connects zaif exchange and return channel of exchange information
func (e *Exchange) Connect() (<-chan *Stream, <-chan error, error) {
	var err error
	e.ws, err = websocket.Dial(e.c.ZaifPublicStreamURL, "", origin)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	log.Printf("Connect! %v\n", e.ws)

	sc := make(chan *Stream)
	ec := make(chan error)

	dec := json.NewDecoder(e.ws)
	go func() {
		for {
			var current Stream
			if err := dec.Decode(&current); err != nil {
				log.Fatal(err)
				ec <- err
				return
			}

			sc <- &current
		}
	}()

	return sc, ec, nil
}
