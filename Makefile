all: test bench vet

test:
	go test -race ./...

bench:
	go test -race -bench . -run "Benchmark"

vet:
	go vet ./...
