package coincheck

import (
	"os"
	"reflect"
	"strconv"
)

// Config is conincheck configuration
type Config struct {
	CoincheckServiceURL string
	CoincheckAPIKey     string
	CoincheckAPISecret  string
}

// DefaultConf is production configuration
var DefaultConf = Config{
	CoincheckServiceURL: "https://coincheck.jp",
	CoincheckAPIKey:     "",
	CoincheckAPISecret:  "",
}

// OverrideEnvVal accrue value for config from environment variable
func (c *Config) OverrideEnvVal() *Config {
	var ret Config
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
