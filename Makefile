.PHONY: all
all: fmt tests tests-cov

fmt:
	@test -z $$(go fmt ./...)

tests:
	go test -v ./twitter

tests-cov:
	go test -v -race -coverprofile coverage.out -covermode atomic ./...
	go tool cover -html=coverage.out

