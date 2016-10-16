package example

import (
	"fmt"

	"github.com/hnkw/goxchange/service/coincheck"
	"github.com/hnkw/goxchange/service/zaif"
)

func zaifExample() {
	// use Zaif API
	c := zaif.NewDefaultClient()
	info, err := c.GetInfo()
	if err != nil {
		// handle error
	}
	fmt.Printf("zaif GetInfo %+v", info)

	e := zaif.NewDefaultExchange()
	streamCh, _, err := e.Connect()
	if err != nil {
		// handle error
	}

	go func() {
		for {
			select {
			case stream := <-streamCh:
				// receive books information
				fmt.Printf("stream %+v", stream)
			}
		}
	}()
}

func coincheckExample() {
	// use Coincheck API
	cc := coincheck.NewDefaultClient()
	accounts, err := cc.Accounts()
	if err != nil {
		// handle error
	}
	fmt.Printf("coincheck Accounts %+v", accounts)

	e := coincheck.NewExchange()
	booksCh, _, err := e.Connect()
	if err != nil {
		// handle error
	}

	go func() {
		for {
			select {
			case books := <-booksCh:
				// receive books information
				fmt.Printf("books %+v", books)
			}
		}
	}()
}

func useAbstructInterface() {
	c := zaif.NewDefaultClient()
	e := zaif.NewDefaultExchange()

	exc := zaif.NewAdapter(e, c)

	// use as goxchange.Exchange interface
	depo, err := exc.Deposit()
	if err != nil {
		// handle error
	}
	fmt.Printf("depo %+v", depo)
}

func main() {
	zaifExample()
	coincheckExample()
	useAbstructInterface()
}
