package coincheck

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Sign --
type Sign struct {
	c *Config
}

// NewSign --
func NewSign(c *Config) *Sign {
	return &Sign{
		c: c,
	}
}

// NewDefaultSign --
func NewDefaultSign() *Sign {
	return &Sign{
		c: DefaultConf.OverrideEnvVal(),
	}
}

func (s Sign) calcSign(nonce, url string, body []byte) (string, error) {
	hash := hmac.New(sha256.New, []byte(s.c.CoincheckAPISecret))

	if _, err := hash.Write([]byte(nonce)); err != nil {
		return "", err
	}

	if _, err := hash.Write([]byte(url)); err != nil {
		return "", err
	}

	if body != nil {
		if _, err := hash.Write(body); err != nil {
			return "", err
		}
	}

	signed := hash.Sum(nil)
	return hex.EncodeToString(signed), nil
}

func (s Sign) nonce() string {
	time.Sleep(time.Second)                         // work around issue same nonce wait 1 second
	return strconv.FormatInt(time.Now().Unix(), 10) // Millie sec
}

func (s Sign) createRequest(method, path string, body interface{}) (*http.Request, error) {
	nonce := s.nonce()
	url := s.c.CoincheckServiceURL + path

	var bodyReader io.Reader
	if body != nil {
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(bodyJSON)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ACCESS-KEY", s.c.CoincheckAPIKey)
	req.Header.Set("ACCESS-NONCE", nonce)
	signed, err := s.calcSign(nonce, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("ACCESS-SIGNATURE", signed)

	return req, nil
}
