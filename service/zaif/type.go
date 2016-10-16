package zaif

import (
	"encoding/json"
	"time"
)

// Prices is amount of jpy, btc and mona in zaif
type Prices struct {
	JPY  float64 `json:"jpy"`
	BTC  float64 `json:"btc"`
	MONA float64 `json:"mona"`
}

// GetInfoResp is response for get_info API
type GetInfoResp struct {
	Success int `json:"success"`
	Return  struct {
		Funds   Prices `json:"funds"`
		Deposit Prices `json:"deposit"`
		Rights  struct {
			Info     int `json:"info"`
			Trade    int `json:"trade"`
			Withdraw int `json:"withdraw"`
		}
		TradeCount int   `json:"trade_count"`
		OpenOrders int   `json:"open_orders"`
		ServerTime int64 `json:"server_time"`
	} `json:"return"`
}

// TradeHistoryReq is request for trade_history API
// parameter     required description                              data type                       default
// from          No       この順番のレコードから取得               numerical                       0
// count         No       取得するレコード数                       numerical                       1000
// from_id       No       このトランザクションIDのレコードから取得 numerical                       0
// end_id        No       このトランザクションIDのレコードまで取得 numerical                       infinity
// order         No       ソート順                                 ASC (昇順)もしくは DESC (降順)  DESC
// since         No       開始タイムスタンプ                       UNIX time                       0
// end           No       終了タイムスタンプ                       UNIX time                       infinity
// currency_pair No       通貨ペア。指定なしで全ての通貨ペア       (例) btc_jpy                    全ペア
type TradeHistoryReq struct {
	From         uint64 `mapkey:"from"`
	Count        uint64 `mapkey:"count"`
	FromID       uint64 `mapkey:"from_id"`
	EndID        uint64 `mapkey:"end_id"`
	Order        string `mapkey:"order"`
	Since        uint64 `mapkey:"since"`
	End          uint64 `mapkey:"end"`
	CurrencyPair string `mapkey:"currency_pair"`
}

// Trade is a trade information
type Trade struct {
	CurrencyPair string
	Action       string
	Amount       float64
	Price        float64
	Fee          float64
	YourAction   string
	Bonus        float64
	TimeStamp    time.Time
}

// UnmarshalJSON implements unmarshaler for Trade
func (t *Trade) UnmarshalJSON(body []byte) error {
	var raw struct {
		CurrencyPair string  `json:"currency_pair"`
		Action       string  `json:"action"`
		Amount       float64 `json:"amount"`
		Price        float64 `json:"price"`
		Fee          float64 `json:"fee"`
		YourAction   string  `json:"your_action"`
		Bonus        float64 `json:"bonus"`
		TimeStamp    int64   `json:"timestamp"`
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return err
	}

	t.CurrencyPair = raw.CurrencyPair
	t.Action = raw.Action
	t.Amount = raw.Amount
	t.Price = raw.Price
	t.Fee = raw.Fee
	t.YourAction = raw.YourAction
	t.Bonus = raw.Bonus
	t.TimeStamp = time.Unix(raw.TimeStamp, 0)
	return nil
}

// TradeHistoryResp is response for trade_history API
type TradeHistoryResp struct {
	Success int `json:"success"`
	Return  struct {
		Trades map[string]Trade
	} `json:"return"`
}

// ActiveOrdersReq is request for active_orders API
// parameter     required description                                data type    default
// currency_pair No       取得する通貨ペア。指定なしで全ての通貨ペア (例) btc_jpy 全てのペア
type ActiveOrdersReq struct {
	CurrencyPair string `mapkey:"currency_pair"`
}

// ActiveOrdersResp is response for active_orders API
type ActiveOrdersResp struct {
	Success int `json:"success"`
	Return  struct {
		Trades map[string]Trade
	} `json:"return"`
}

// TradeReq is request for trade API
// parameter     required description      data type           default
// currency_pair Yes      発注する通貨ペア (例) btc_jpy        -
// action        Yes      注文の種類       bid もしくは ask    -
// price         Yes      価格             numerical           -
// amount        Yes      数量             (例: 0.3) numerical -
type TradeReq struct {
	CurrencyPair string  `mapkey:"currency_pair "`
	Action       string  `mapkey:"action"`
	Price        float64 `mapkey:"price"`
	Amount       float64 `mapkey:"amount"`
}

// TradeResp is response for trade API
type TradeResp struct {
	Success int `json:"success"`
	Return  struct {
		Received float64 `json:"received"`
		Remains  float64 `json:"remains"`
		OrderID  uint64  `json:"order_id"`
		Funds    Prices  `json:"funds"`
	} `json:"return"`
}

// CancelOrderReq is request for trade_cancel API
// parameter required description                                      data type default
// order_id  Yes      注文ID（tradeまたはactive_ordersで取得できます） numerical -
type CancelOrderReq struct {
	OrderID uint64 `mapkey:"order_id"`
}

// CancelOrderResp is response for trade_cancel API
type CancelOrderResp struct {
	Success int `json:"success"`
	Return  struct {
		OrderID uint64 `json:"order_id"`
		Funds   Prices `json:"funds"`
	} `json:"return"`
}

// WithdrawReq is request for withdraw API
// parameter required description                 data type         default
// currency  Yes      引き出す通貨                btc もしくは mona -
// address   Yes      送信先のアドレス            address string    -
// amount    Yes      引き出す金額(例: 0.3)       numerical         -
// opt_fee   No       採掘者への手数料(例: 0.003) numerical         -
type WithdrawReq struct {
	currency string  `mapkey:"currency"`
	address  string  `mapkey:"address"`
	amount   float64 `mapkey:"amount"`
	optFee   float64 `mapkey:"opt_fee"`
}

// WithdrawResp is response for withdraw API
type WithdrawResp struct {
	Success int `json:"success"`
	Return  struct {
		TxID  string `json:"txid"`
		Funds Prices `json:"funds"`
	} `json:"return"`
}
