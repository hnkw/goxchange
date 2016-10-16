package coincheck

import (
	"encoding/json"
	"testing"
)

var tickerData = `
{"last":35585,"bid":35578,"ask":35787,"high":35837,"low":34359,"volume":"443.66911669","timestamp":1437797863}
`

var booksData = `
{"asks":[["35787.0","0.016"],["35911.0","10.0"],["35912.0","2.526"],["35940.0","1.57"],["35954.0","4.117"],["35964.0","3.315"],["35975.0","1.568"],["35978.0","1.987"],["35980.0","0.54428413"],["35987.0","5.49"],["35992.0","4.117"],["36000.0","0.201"],["36005.0","50.48752702"],["36012.0","0.2"],["36014.0","2.684"],["36029.0","5.00665554"],["36030.0","0.361"],["36043.0","4.705"],["36054.0","1.764"],["36082.0","3.529"]],"bids":[["35578.0","27.0204"],["35478.0","5.0"],["35473.0","2.526"],["35441.0","4.285"],["35438.0","5.185"],["35430.0","1.481"],["35428.0","0.857"],["35419.0","6.285"],["35408.0","3.333"],["35341.0","48.742992"],["35334.0","0.315"],["35319.0","4.0"],["35318.0","56.18450842"],["35314.0","3.214"],["35301.0","5.892"],["35291.0","1.296"],["35272.0","0.857"],["35269.0","6.0"],["35258.0","54.25395436"],["35247.0","1.578"]]}
`

func Test_Unmarshal(t *testing.T) {
	var ticker Ticker
	if err := json.Unmarshal([]byte(tickerData), &ticker); err != nil {
		t.Errorf("tickerData should be Unmarshaled with test data %+v\n", err)
	}

	var books Books
	if err := json.Unmarshal([]byte(booksData), &books); err != nil {
		t.Errorf("booksData should be Unmarshaled with test data %+v\n", err)
	}
}

func Test_Ticker(t *testing.T) {
	e := NewExchange()
	_, err := e.Ticker()
	if err != nil {
		t.Errorf("Timer should return timer information %+v\n", err)
	}
}

func Test_Books(t *testing.T) {
	e := NewExchange()
	_, err := e.Books()
	if err != nil {
		t.Errorf("Books should return timer information %+v\n", err)
	}
}
