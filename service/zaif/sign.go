package zaif

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Sign --
type Sign struct {
	c *Confing
}

// NewSign --
func NewSign(c *Confing) *Sign {
	return &Sign{
		c: c,
	}
}

// NewDefaultSign --
func NewDefaultSign() *Sign {
	return &Sign{
		c: overrideEnvVal(&DefaultConf),
	}
}

func (s Sign) calcSign(body []byte) (string, error) {
	hash := hmac.New(sha512.New, []byte(s.c.ZaifAPISecret))
	if _, err := hash.Write(body); err != nil {
		return "", err
	}

	signed := hash.Sum(nil)
	return hex.EncodeToString(signed), nil
}

func (s Sign) nonce() string {
	time.Sleep(time.Second)                         // work around issue same nonce wait 1 second
	return strconv.FormatInt(time.Now().Unix(), 10) // Millie sec
}

func (s Sign) appendNonceMethod(nonce, method string, data url.Values) url.Values {
	ret := url.Values{
		"nonce":  []string{nonce},
		"method": []string{method},
	}

	if data != nil {
		for k, v := range data {
			ret[k] = v
		}
	}
	return ret
}

func (s Sign) createRequest(method string, data url.Values) (*http.Request, error) {
	formEncoded := s.appendNonceMethod(s.nonce(), method, data).Encode()
	body := strings.NewReader(formEncoded)

	req, err := http.NewRequest("POST", s.c.ZaifPrivateServiceURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Key", s.c.ZaifAPIKey)

	signed, err := s.calcSign([]byte(formEncoded))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Sign", signed)
	return req, nil
}
