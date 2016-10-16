package zaif

import (
	"encoding/json"
	"testing"
)

var unmarshalData = `
{"asks": [[34985.0, 0.0129], [34990.0, 0.0004], [34995.0, 0.0194], [35000.0, 0.1735], [35005.0, 0.1769], [35010.0, 0.0377], [35015.0, 0.1276], [35025.0, 14.528], [35030.0, 0.3511], [35035.0, 0.2418], [35040.0, 9.3991], [35045.0, 0.1943], [35050.0, 0.157], [35055.0, 6.0407], [35060.0, 0.6313], [35065.0, 0.6741], [35070.0, 0.4561], [35075.0, 0.3877], [35080.0, 0.2353], [35085.0, 0.1557]], "last_price": {"action": "bid", "price": 34990.0}, "trades": [{"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34990.0, "tid": 394923, "amount": 0.0002, "date": 1437485523}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34985.0, "tid": 394922, "amount": 0.0002, "date": 1437485523}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34995.0, "tid": 394921, "amount": 0.0029, "date": 1437485329}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34995.0, "tid": 394920, "amount": 0.0029, "date": 1437485329}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34995.0, "tid": 394919, "amount": 0.0029, "date": 1437485329}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34995.0, "tid": 394918, "amount": 0.0029, "date": 1437485328}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34995.0, "tid": 394917, "amount": 0.0029, "date": 1437485328}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34995.0, "tid": 394916, "amount": 0.0029, "date": 1437485328}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34995.0, "tid": 394915, "amount": 0.0208, "date": 1437485328}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34990.0, "tid": 394914, "amount": 0.0002, "date": 1437485328}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34990.0, "tid": 394913, "amount": 0.0002, "date": 1437485328}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34985.0, "tid": 394912, "amount": 0.0002, "date": 1437485328}, {"currenty_pair": "btc_jpy", "trade_type": "bid", "price": 34985.0, "tid": 394911, "amount": 0.0002, "date": 1437485328}, {"currenty_pair": "btc_jpy", "trade_type": "ask", "price": 34960.0, "tid": 394909, "amount": 0.0206, "date": 1437484236}, {"currenty_pair": "btc_jpy", "trade_type": "ask", "price": 34960.0, "tid": 394908, "amount": 0.016, "date": 1437484236}, {"currenty_pair": "btc_jpy", "trade_type": "ask", "price": 34970.0, "tid": 394907, "amount": 0.0004, "date": 1437484235}, {"currenty_pair": "btc_jpy", "trade_type": "ask", "price": 34970.0, "tid": 394906, "amount": 0.0004, "date": 1437484235}, {"currenty_pair": "btc_jpy", "trade_type": "ask", "price": 34970.0, "tid": 394905, "amount": 0.0004, "date": 1437484235}, {"currenty_pair": "btc_jpy", "trade_type": "ask", "price": 34970.0, "tid": 394904, "amount": 0.0004, "date": 1437484235}, {"currenty_pair": "btc_jpy", "trade_type": "ask", "price": 34970.0, "tid": 394903, "amount": 0.0004, "date": 1437484235}, {"currenty_pair": "btc_jpy", "trade_type": "ask", "price": 34970.0, "tid": 394902, "amount": 0.0004, "date": 1437484235}], "bids": [[34940.0, 0.002], [34935.0, 0.0004], [34930.0, 0.0152], [34905.0, 0.016], [34900.0, 0.128], [34895.0, 0.4358], [34890.0, 0.2139], [34885.0, 0.1248], [34880.0, 0.0812], [34875.0, 0.1292], [34870.0, 7.1215], [34865.0, 0.0852], [34855.0, 10.2209], [34850.0, 4.7316], [34845.0, 0.0672], [34840.0, 0.0968], [34835.0, 0.0192], [34830.0, 0.0465], [34825.0, 0.0448], [34820.0, 0.0384]], "currency_pair": "btc_jpy", "timestamp": "2015-07-21 22:33:19.851255"}
`

func Test_StreamUnmarshal(t *testing.T) {
	var e Stream
	if err := json.Unmarshal([]byte(unmarshalData), &e); err != nil {
		t.Errorf("Stream should be Unmarshaled with test data %+v\n", err)
	}
}
