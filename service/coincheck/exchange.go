package coincheck

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	host = "https://coincheck.jp"
)

// Ticker information
type Ticker struct {
	Last      uint64
	Bid       uint64
	Ask       uint64
	High      uint64
	Low       uint64
	Volume    float64
	TimeStamp time.Time
}

// UnmarshalJSON implements unmarshaler for Ticker
func (t *Ticker) UnmarshalJSON(body []byte) error {
	var f struct {
		Last      uint64 `json:"last"`
		Bid       uint64 `json:"bid"`
		Ask       uint64 `json:"ask"`
		High      uint64 `json:"high"`
		Low       uint64 `json:"low"`
		Volume    string `json:"volume"`
		TimeStamp int64  `json:"timestamp"`
	}

	if err := json.Unmarshal(body, &f); err != nil {
		return err
	}

	t.Last = f.Last
	t.Bid = f.Bid
	t.Ask = f.Ask
	t.High = f.High
	t.Low = f.Low

	volume, err := strconv.ParseFloat(f.Volume, 64)
	if err != nil {
		return err
	}
	t.Volume = volume

	t.TimeStamp = time.Unix(f.TimeStamp, 0)

	return nil
}

// Books information
type Books struct {
	Asks [][]float64
	Bids [][]float64
}

func strVectorToFloat(sv [][]string) ([][]float64, error) {
	ret := make([][]float64, 0, len(sv))
	for _, strLine := range sv {
		floatLine := make([]float64, 0, len(strLine))
		for _, strColumn := range strLine {
			floatColum, err := strconv.ParseFloat(strColumn, 64)
			if err != nil {
				return nil, err
			}
			floatLine = append(floatLine, floatColum)
		}
		ret = append(ret, floatLine)
	}
	return ret, nil
}

// UnmarshalJSON implements unmarshaler for Books
func (b *Books) UnmarshalJSON(body []byte) error {
	var f struct {
		Asks [][]string `json:"asks"`
		Bids [][]string `json:"bids"`
	}

	if err := json.Unmarshal(body, &f); err != nil {
		return err
	}

	floatAsks, err := strVectorToFloat(f.Asks)
	if err != nil {
		return err
	}
	b.Asks = floatAsks

	floatBids, err := strVectorToFloat(f.Bids)
	if err != nil {
		return err
	}
	b.Bids = floatBids

	return nil
}

// Exchange is object as coincheck exchange
type Exchange struct {
	client *http.Client
}

// NewExchange returns exchange object
func NewExchange() *Exchange {
	return &Exchange{
		client: http.DefaultClient,
	}
}

func (e *Exchange) get(path string, v interface{}) error {
	url := host + path
	resp, err := e.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}

// Ticker returns books information
func (e *Exchange) Ticker() (*Ticker, error) {
	var ticker Ticker
	if err := e.get("/api/ticker", &ticker); err != nil {
		return nil, err
	}

	return &ticker, nil
}

// Books returns books information
func (e *Exchange) Books() (*Books, error) {
	var books Books
	if err := e.get("/api/order_books", &books); err != nil {
		return nil, err
	}

	return &books, nil
}

// Connect connects coincheck exchange and return channel of exchange information
func (e *Exchange) Connect() (<-chan *Books, <-chan error, error) {

	bc := make(chan *Books)
	ec := make(chan error)
	go func() {
		for {
			b, err := e.Books()
			if err != nil {
				log.Fatal(err)
				ec <- err
				return
			}
			bc <- b

			time.Sleep(time.Second * 5)
		}
	}()

	return bc, ec, nil
}
