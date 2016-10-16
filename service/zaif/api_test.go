package zaif

import (
	"testing"
)

func Test_local_marshalUrlValues(t *testing.T) {
	c := &Client{}

	type testStruct struct {
		A string  `mapkey:"a"`
		B int     `mapkey:"b"`
		C uint64  `mapkey:"c"`
		D float64 `mapkey:"d"`
	}

	test := testStruct{
		A: "AA",
		B: 1,
		C: 2,
		D: 3.3,
	}

	v := c.marshalUrlValues(&test)

	valA, ok := v["a"]
	if !ok {
		t.Errorf("mapkey a not found %+v\n", v)
		return
	}
	if valA[0] != "AA" {
		t.Errorf("valA should be AA %+v\n", valA[0])
	}

	valB, ok := v["b"]
	if !ok {
		t.Errorf("mapkey b not found %+v\n", v)
	}
	if valB[0] != "1" {
		t.Errorf("valB should be 1 %+v\n", valB[0])
	}

	valC, ok := v["c"]
	if !ok {
		t.Errorf("mapkey c not found %+v\n", v)
	}
	if valC[0] != "2" {
		t.Errorf("valC should be 2 %+v\n", valC[0])
	}

	valD, ok := v["d"]
	if !ok {
		t.Errorf("mapkey d not found %+v\n", v)
	}
	if valD[0] != "3.3" {
		t.Errorf("valD should be 3.3 %+v\n", valD[0])
	}
}

func checkEnv(t *testing.T, sign *Sign) {
	if sign.c.ZaifAPIKey == "" {
		t.Error("API key not set")
	}

	if sign.c.ZaifAPISecret == "" {
		t.Error("API Secret not set")
	}
}
