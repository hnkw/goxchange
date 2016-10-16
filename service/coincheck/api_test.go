package coincheck

import (
	"testing"
)

func Test_cloud_api_OpenOrders(t *testing.T) {
	c := NewDefaultClient()
	resp, err := c.OpenOrders()
	if err != nil {
		t.Errorf("err should be nil, but err[%+v]", err)
	}

	if !resp.Success {
		t.Errorf("resp.Success should be true, but resp.Success[%+v]", resp.Success)
	}
}

func Test_cloud_api_Transactions(t *testing.T) {
	c := NewDefaultClient()
	resp, err := c.Transactions()
	if err != nil {
		t.Errorf("err should be nil, but err[%+v]", err)
	}

	if !resp.Success {
		t.Errorf("resp.Success should be true, but resp.Success[%+v]", resp.Success)
	}
}

func Test_cloud_api_Balance(t *testing.T) {
	c := NewDefaultClient()
	resp, err := c.Balance()
	if err != nil {
		t.Errorf("err should be nil, but err[%+v]", err)
	}

	if !resp.Success {
		t.Errorf("resp.Success should be true, but resp.Success[%+v]", resp.Success)
	}
}

func Test_cloud_api_Accounts(t *testing.T) {
	c := NewDefaultClient()
	resp, err := c.Accounts()
	if err != nil {
		t.Errorf("err should be nil, but err[%+v]", err)
	}

	if !resp.Success {
		t.Errorf("resp.Success should be true, but resp.Success[%+v]", resp.Success)
	}
}
