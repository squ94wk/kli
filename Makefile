.PHONY: build
build:
	go build -o kli cmd/kli/kli.go

.PHONY: install
install:
	go install cmd/kli/kli.go
