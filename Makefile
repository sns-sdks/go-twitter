.PHONY: all
all: lint tests tests-cov

lint:
	@test -z $$(go fmt ./...)

tests:
	go test -v ./twitter

tests-cov:
	go test -v -race -coverprofile coverage.out -covermode atomic ./...

tests-html: tests-cov
	go tool cover -html=coverage.out

