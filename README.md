# goxchange

Multi-service bitcoin exchange client. Currently [goxchange](https://github.com/hnkw/goxchange) provides [zaif](https://zaif.jp) and [coincheck](https://coincheck.com).

## Installation

```go
go get github.com/hnkw/goxchange
```

## Example

### preparation

Set your credencials to enviroment valiables

#### Zaif

```shell
export ZaifAPIKey=<your-zaif-api-key>
export ZaifAPISecret=<your-zaif-api-scecret>
```

#### Coincheck

```shell
export CoincheckAPIKey=<your-coincheck-api-key>
export CoincheckAPISecret=<your-coincheck-api-scecret>
```

### use library

see [detail](https://github.com/hnkw/goxchange/blob/master/example/example.go)

```go
c := zaif.NewDefaultClient()
e := zaif.NewDefaultExchange()
exc := zaif.NewAdapter(e, c)

// use as goxchange.Exchange interface
depo, err := exc.Deposit()
if err != nil {
	// handle error
}
fmt.Printf("depo %+v", depo)
```

## License

[MIT](https://github.com/hnkw/goxchange/blob/master/LICENCE)

## Author

[Kenji Hanakawa](https://github.com/hnkw)
