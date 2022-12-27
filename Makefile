.PHONY: compile
compile:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(PWD)/build/desafio .

.PHONY: compress
compress:
	@upx $(PWD)/build/desafio

.PHONY: build
build: compile compress

.PHONY: test
test:
	@echo "running tests"
	go test -v ./...

.PHONY: clean
clean:
	@echo "cleaning releases"
	@GOOS=linux go clean -i -x ./...
	@rm -rf build/

.PHONY: generate
generate:
	go install github.com/swaggo/swag/cmd/swag@latest
	go generate ./...
	go mod tidy

.PHONY: swagger
swagger:
	swag init -g ./cmd/server/main.go -o /docs/swagger --parseDependency
