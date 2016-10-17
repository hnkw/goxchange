package goxchange

// Exchange is the interface for abstraction methods in each exchanges
type Exchange interface {
	Deposit() (*Deposit, error)
	Connect() (<-chan *Books, <-chan error, error)
	Ask(o Order) error
	Bit(o Order) error
}
