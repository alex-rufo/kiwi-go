# Go Kiwi
## Summary

Unofficial [Kiwi](https://kiwi.com) Go client library.

## Installation

```sh
go get github.com/greengeko/kiwi-go
```

## Documentation

### Get flights

```go
k := kiwi.New()
params := kiwi.Parameters{
	Partner: "",
	FlyFrom: "BCN",
	To: "MAD",
}
resp, err := k.GetFlights(&params)
```
