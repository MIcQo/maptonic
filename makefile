
build:
	go build -ldflags "-X github.com/MIcQo/maptonic/config.Version=v1.5.0" -o bin/maptonic main.go

hook:
	goimports -w .
	gofmt -w .
	golangci-lint run -v