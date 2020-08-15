.PHONY: build
build:
	go build -o kli cmd/kli/main.go

.PHONY: install
install:
	go install -o kli cmd/kli/main.go
