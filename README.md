# Go Kiwi

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/alex-rufo/kiwi-go)
[![Build Status](https://travis-ci.org/alex-rufo/kiwi-go.svg?branch=master)](https://travis-ci.org/alex-rufo/kiwi-go)

## Summary

Unofficial [Kiwi](https://kiwi.com) Go client library.

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