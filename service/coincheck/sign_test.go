package coincheck

import (
	"testing"
)

func Test_local_calcSign(t *testing.T) {
	sign := NewDefaultSign()

	_, err := sign.calcSign("1440079469", "https://coincheck.jp/api/accounts/balance", []byte("hoge=foo"))
	if err != nil {
		t.Errorf("err should be nil, but[%v]", err)
	}
}
