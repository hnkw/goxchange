package coincheck

import (
	"time"
)

// CreateOrderReq is request body for POST /api/exchange/orders
// parameter     required description        data type                          default
// rate          true     注文のレート       （例）28000                        -
// amount        true     注文での量         （例）0.1                          -
// order_type    true     売りか買いかを指定 売りならば "sell" 買いならば "buy" -
// pair          true     取引ペアを指定     現在は "btc_jpy" のみです。        -
type CreateOrderReq struct {
	Rate      float64 `json:"rate"`
	Amount    float64 `json:"amount"`
	OrderType string  `json:"order_type"`
	Pair      string  `json:"pair"`
}

// OrderResp is a order response in coincheck API
type OrderResp struct {
	ID            uint64    `json:"id"`
	Rate          float64   `json:"rate,string"`
	Amount        float64   `json:"amount,string"`
	PendingAmount float64   `json:"pending_amount,string"`
	OrderType     string    `json:"order_type"`
	Pair          string    `json:"pair"`
	CreatedAt     time.Time `json:"created_at"`
}

// CreateOrderResp is response body for POST /api/exchange/orders
// id          新規注文のID
// rate        注文のレート
// amount      注文の量
// order_type  注文のタイプ（"sell" or "buy"）
// pair        取引ぺア
// created_at 注文の作成日時
type CreateOrderResp struct {
	Success bool `json:"success"`
	OrderResp
}

// OpenOrdersResp is response body for GET /api/exchange/orders/opens
// id             注文のID（新規注文でのIDと同一です）
// rate           注文のレート
// pending_amount 注文の未決済の量
// order_type     注文のタイプ（"sell" or "buy"）
// pair           取引ペア
// created_at     注文の作成日時
type OpenOrdersResp struct {
	Success bool        `json:"success"`
	Orders  []OrderResp `json:"orders"`
}

// DeleteOrderResp is response body for DELETE /api/exchange/orders/{id}
type DeleteOrderResp struct {
	Success bool   `json:"success"`
	ID      uint64 `json:"id"`
}

// Transaction is a trade history returned GET /api/exchange/orders/transactions
// id ID
// order_id 注文のID
// created_at 取引が行われた時間
// funds 各残高の増減分
type Transaction struct {
	ID        uint64 `json:"id"`
	OrderID   uint64 `json:"order_id"`
	CreatedAt time.Time
	Funds     struct {
		BTC float64 `json:"btc,string"`
		JPY float64 `json:"jpy,string"`
	} `json:"funds"`
}

// TransactionsResp is response body for GET /api/exchange/orders/transactions
type TransactionsResp struct {
	Success      bool          `json:"success"`
	Transactions []Transaction `json:"transactions"`
}

// BalanceResp is response body for GET /api/accounts/balance
// jpy             日本円の残高
// btc             ビットコインの残高
// jpy_reserved    未決済の買い注文に利用している日本円の合計
// btc_reserved    未決済の売り注文に利用しているビットコインの合計
// jpy_lend_in_use 貸出申請をしている日本円の合計（現在は日本円貸出の機能を提供していません）
// btc_lend_in_use 貸出申請をしているビットコインの合計
// jpy_lent        貸出をしている日本円の合計（現在は日本円貸出の機能を提供していません）
// btc_lent        貸出をしているビットコインの合計
// jpy_debt        借りている日本円の合計（現在は日本円を借りる機能を提供していません）
// btc_debt        借りているビットコインの合計
type BalanceResp struct {
	Success      bool    `json:"success"`
	JPY          float64 `json:"jpy,string"`
	BTC          float64 `json:"btc,string"`
	JPYReserved  float64 `json:"jpy_reserved,string"`
	BTCReserved  float64 `json:"btc_reserved,string"`
	JPYLendInUse float64 `json:"jpy_lend_in_use,string"`
	BTCLendInUse float64 `json:"btc_lend_in_use,string"`
	JPYLent      float64 `json:"jpy_lent,string"`
	BTCLent      float64 `json:"btc_lent,string"`
	JPYDebt      float64 `json:"jpy_debt,string"`
	BTCDebt      float64 `json:"btc_debt,string"`
}

// SendMoneyReq is request body for POST /api/send_money
// address 送り先のビットコインアドレス
// amount  送りたいビットコインの量
type SendMoneyReq struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount,string"`
}

// SendMoneyResp is response body for POST /api/send_money
// id      送金のIDです
// address 送った先のbitcoinアドレス
// amount  送ったbitcoinの量
// fee     手数料
type SendMoneyResp struct {
	Success bool    `json:"success"`
	ID      uint64  `json:"id,string"`
	Address string  `json:"address"`
	Amount  float64 `json:"amount,string"`
	Fee     float64 `json:"fee,string"`
}

// AccountsResp is response body for GET /api/accounts
// id              アカウントのID。日本円入金の際に指定するIDと一致します
// email           登録されたメールアドレス
// bitcoin_address あなたのデポジット用ビットコインのアドレス
type AccountsResp struct {
	Success        bool   `json:"success"`
	ID             uint64 `json:"id"`
	Email          string `json:"email"`
	BitcoinAddress string `json:"bitcoin_address"`
}
