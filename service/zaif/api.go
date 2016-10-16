package zaif

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/fatih/structs"
	cmn "github.com/hnkw/goxchange/common"
)

// Client is trading client for Zaif
type Client struct {
	sign       *Sign
	httpClient *http.Client
	l          cmn.Logger
}

// NewDefaultClient is default setting and instance for production on Zaif
func NewDefaultClient() *Client {
	return &Client{
		sign:       NewDefaultSign(),
		httpClient: http.DefaultClient,
		l:          cmn.DefaultLogger,
	}
}

// GetInfo is API for get_info
// get_info
// 現在の残高（余力および残高）、APIキーの権限、過去のトレード数、
// アクティブな注文数、サーバーのタイムスタンプを取得します。
func (c *Client) GetInfo() (*GetInfoResp, error) {
	req, err := c.sign.createRequest("get_info", nil)
	if err != nil {
		return nil, err
	}

	var resp GetInfoResp
	if err := c.reqest(req, &resp); err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		return nil, fmt.Errorf("resp not success resp[%v]", resp)
	}

	return &resp, nil
}

// TradeHistory is API for trade_history
// trade_history
// ユーザー自身の取引履歴を取得します。
func (c *Client) TradeHistory(params *TradeHistoryReq) (*TradeHistoryResp, error) {
	data := c.marshalUrlValues(params)

	req, err := c.sign.createRequest("trade_history", data)
	if err != nil {
		return nil, err
	}

	var resp TradeHistoryResp
	if err := c.reqest(req, &resp); err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		return nil, fmt.Errorf("resp not success resp[%v]", resp)
	}

	return &resp, nil
}

// Trade is API for trade
// trade
// 注文を行います。
func (c *Client) Trade(params *TradeReq) (*TradeResp, error) {
	data := c.marshalUrlValues(params)

	req, err := c.sign.createRequest("trade", data)
	if err != nil {
		return nil, err
	}

	var resp TradeResp
	if err := c.reqest(req, &resp); err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		return nil, fmt.Errorf("resp not success resp[%v]", resp)
	}

	return &resp, nil
}

// CancelOrder is API for cancel_order
// cancel_order
// 注文の取消
func (c *Client) CancelOrder(params *CancelOrderReq) (*CancelOrderResp, error) {
	data := c.marshalUrlValues(params)

	req, err := c.sign.createRequest("cancel_order", data)
	if err != nil {
		return nil, err
	}

	var resp CancelOrderResp
	if err := c.reqest(req, &resp); err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		return nil, fmt.Errorf("resp not success resp[%v]", resp)
	}

	return &resp, nil
}

func (c *Client) reqest(req *http.Request, apiResp interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.l.Debugf("req[%+v] resp[%+v] respBody[%s]\n", req, resp, string(body))

	return json.Unmarshal(body, apiResp)
}

func (c *Client) marshalUrlValues(params interface{}) url.Values {
	var ret url.Values
	ret = url.Values{}

	s := structs.New(params)
	for _, name := range s.Names() {
		f := s.Field(name)
		v := f.Value()
		key := f.Tag("mapkey")

		var str string
		switch t := v.(type) {
		case string:
			str = t
		case float64, float32:
			str = fmt.Sprintf("%v", t)
		case int64, uint64, int32, uint32, int, uint:
			str = fmt.Sprintf("%d", t)
		case fmt.Stringer:
			str = t.String()
		default:
			c.l.Warnf("name [%s] key [%s] v[%v]\n",
				name, key, v)
			continue
		}

		ret[key] = []string{str}
	}

	return ret
}
