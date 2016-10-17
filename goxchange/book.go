package goxchange

import "time"

// Books is all exchange books
type Books struct {
	Asks Orders
	Bids Orders

	TimeStamp time.Time
}
