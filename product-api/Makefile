swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

check_install:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: fmt
fmt:
	go fmt $(shell go list ./... | grep -v /vendor/)
	go run golang.org/x/tools/cmd/goimports -w --local github.com/izumiya/working/product-api .