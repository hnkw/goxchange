package goxchange

import (
	"testing"
)

func Test_Order(t *testing.T) {
	o := NewOrder(100, 10)
	if o.Price() != 100 {
		t.Errorf("price should be 100")
	}
	if o.Amount() != 10 {
		t.Errorf("Amount should be 10")
	}
}

func Test_Shift(t *testing.T) {
	var (
		o  Order
		os Orders
	)
	os = Orders{
		Order{20, 2},
		Order{19, 2},
	}

	o, os = os.Shift()
	if o.Price() != 20 ||
		o.Amount() != 2 {
		t.Errorf("o should return head order")
	}
	if len(os) != 1 ||
		os[0].Price() != 19 ||
		os[0].Amount() != 2 {
		t.Errorf("os should remain tail")
	}
}

func Test_ShiftTwice(t *testing.T) {
	os := Orders{
		Order{20, 2},
		Order{19, 3},
		Order{18, 4},
		Order{17, 5},
	}

	var firstPrice float64
	for o1, remain1 := os.Shift(); o1 != nil; o1, remain1 = remain1.Shift() {
		firstPrice += o1.Price()
	}

	var secondPrice float64
	for o2, remain2 := os.Shift(); o2 != nil; o2, remain2 = remain2.Shift() {
		secondPrice += o2.Price()
	}

	if firstPrice != secondPrice {
		t.Errorf("shift func should not effect firstPrice[%v] secondPrice[%v]", firstPrice, secondPrice)
	}
}
