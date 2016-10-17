package goxchange

// ExcName is exchange name
type ExcName uint64

const (
	// Zaif is exchange name
	Zaif ExcName = iota + 1
	// CoinCheck is exchange name
	CoinCheck
)

func (n ExcName) String() string {
	switch n {
	case Zaif:
		return "Zaif"
	case CoinCheck:
		return "CoinCheck"
	default:
		return "unknown"
	}
}
