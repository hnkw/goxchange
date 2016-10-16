package zaif

import (
	"os"
	"reflect"
	"strconv"
)

// Confing --
type Confing struct {
	ZaifPublicServiceURL  string
	ZaifPublicStreamURL   string
	ZaifPrivateServiceURL string
	ZaifAPIKey            string
	ZaifAPISecret         string
}

// DefaultConf --
var DefaultConf = Confing{
	ZaifPublicServiceURL:  "https://api.zaif.jp",
	ZaifPublicStreamURL:   "ws://api.zaif.jp:8888/stream?currency_pair=btc_jpy",
	ZaifPrivateServiceURL: "https://api.zaif.jp/tapi",
	ZaifAPIKey:            "",
	ZaifAPISecret:         "",
}

func overrideEnvVal(c *Confing) *Confing {
	var ret Confing
	ret = *c // copy

	v := reflect.ValueOf(&ret).Elem()
	t := v.Type()

	// Override default values with environment variables
	for i := 0; i < v.NumField(); i++ {
		envVal := os.Getenv(t.Field(i).Name)
		if len(envVal) <= 0 {
			// Environment variable not found
			continue
		}

		field := v.Field(i)
		switch field.Kind() {
		case reflect.Int:
			v, err := strconv.Atoi(envVal)
			if err != nil {
			}
			field.Set(reflect.ValueOf(v))

		case reflect.String:
			field.Set(reflect.ValueOf(envVal))
		}
	}

	return &ret
}
