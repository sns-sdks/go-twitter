.PHONY: all
all: fmt

.PHONY: fmt
fmt:
	@test -z $$(go fmt ./...)
