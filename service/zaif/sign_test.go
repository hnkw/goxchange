package zaif

import (
	"net/url"
	"testing"
)

func Test_local_createRequest(t *testing.T) {
	s := NewDefaultSign()

	value := url.Values{
		"BB": []string{
			"BB Value",
		},
	}

	_, err := s.createRequest("get_info", value)
	if err != nil {
		t.Errorf("s.createRequest should not return err %+v\n", err)
	}
}
