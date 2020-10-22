.PHONY: test
test:
	go test -v ./... -cover

.PHONY: vendor
vendor:
	go mod vendor