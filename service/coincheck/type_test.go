package coincheck

import (
	"encoding/json"
	"testing"
)

func Test_local_CreateOrderResp(t *testing.T) {
	var respStr = `
{
  "success": true,
  "id": 12345,
  "rate": "30010.0",
  "amount": "1.3",
  "order_type": "sell",
  "pair": "btc_jpy",
  "created_at": "2015-01-10T05:55:38.000Z"
}
`
	var resp CreateOrderResp
	err := json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		t.Errorf("err should be nil, but [%+v]", err)
	}
}

func Test_local_OpenOrdersResp(t *testing.T) {
	var respStr = `
{
  "success": true,
  "orders": [
    {
      "id": 202835,
      "order_type": "buy",
      "rate": "26890",
      "pair": "btc_jpy",
      "pending_amount": "0.5527",
      "created_at": "2015-01-10T05:55:38.000Z"
    },
    {
      "id": 202836,
      "order_type": "sell",
      "rate": "26990",
      "pair": "btc_jpy",
      "pending_amount": "0.77",
      "created_at": "2015-01-10T05:55:38.000Z"
    }
  ]
}
`
	var resp OpenOrdersResp
	err := json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		t.Errorf("err should be nil, but [%+v]", err)
	}
}

func Test_local_DeleteOrderResp(t *testing.T) {
	var respStr = `
{
  "success": true,
  "id": 12345
}
`
	var resp DeleteOrderResp
	err := json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		t.Errorf("err should be nil, but [%+v]", err)
	}
}

func Test_local_Transactions(t *testing.T) {
	var respStr = `
{
  "success": true,
  "transactions": [
    {
      "id": 5,
      "order_id": 200,
      "created_at": "2015-02-14T05:49:46.000Z",
      "funds": {
        "btc": "-0.4",
        "jpy": "11520.0"
      }
    },
    {
      "id": 7,
      "order_id": 201,
      "created_at": "2015-02-14T05:51:02.000Z",
      "funds": {
        "btc": "-0.1",
        "jpy": "2880.0"
      }
    }
  ]
}
`
	var resp TransactionsResp
	err := json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		t.Errorf("err should be nil, but [%+v]", err)
	}
}

func Test_local_BalanceResp(t *testing.T) {
	var respStr = `
{
  "success": true,
  "jpy": "0.8401",
  "btc": "7.75052654",
  "jpy_reserved": "3000.0",
  "btc_reserved": "3.5002",
  "jpy_lend_in_use": "0",
  "btc_lend_in_use": "0.3",
  "jpy_lent": "0",
  "btc_lent": "1.2",
  "jpy_debt": "0",
  "btc_debt": "0"
}
`
	var resp TransactionsResp
	err := json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		t.Errorf("err should be nil, but [%+v]", err)
	}
}

func Test_local_SendMoneyResp(t *testing.T) {
	var respStr = `
{
  "success": true,
  "id": "276",
  "address": "1v6zFvyNPgdRvhUufkRoTtgyiw1xigncc",
  "amount": "1.5",
  "fee": "0.002"
}
`
	var resp SendMoneyResp
	err := json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		t.Errorf("err should be nil, but [%+v]", err)
	}
}

func Test_local_AccountsResp(t *testing.T) {
	var respStr = `
{
  "success": true,
  "id": 10000,
  "email": "test@gmail.com",
  "bitcoin_address": "1v6zFvyNPgdRvhUufkRoTtgyiw1xigncc"
}
`
	var resp AccountsResp
	err := json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		t.Errorf("err should be nil, but [%+v]", err)
	}
}
