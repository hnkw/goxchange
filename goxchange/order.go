package goxchange

// Order is one order of ask or bid
type Order []float64

// NewOrder construct order struct
func NewOrder(price, amount float64) Order {
	return Order{price, amount}
}

// Price is price of order
func (o Order) Price() float64 {
	return o[0]
}

// Amount is amount of order
func (o Order) Amount() float64 {
	return o[1]
}

// Orders are contains orders as asks and bids
type Orders []Order

// Shift returns head of order and remain orders
func (orders Orders) Shift() (Order, Orders) {
	if len(orders) == 0 {
		return nil, orders
	}

	return orders[0], orders[1:]
}
