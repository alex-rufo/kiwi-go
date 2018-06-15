# Go Kiwi

## Summary

Unofficial [Kiwi][kiwi] Go client library.

## Installation

```sh
go get github.com/alex-rufo/kiwi-go
```

## Documentation

### Get flights

```go
k := kiwi.New()
params := kiwi.Parameters{
    FlyFrom: "BCN",
    To: "MAD",
}
resp, err := k.GetFlights(&params)
```