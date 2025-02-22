run:
	go run cmd/main.go

deps:
	go mod tidy

test:
	go test ./...  -cover -coverprofile=coverage.out 

gen-mocks:
	mockery --dir=./internal --all 

# formats project with go's style guidelines
fmt:
	gofmt -w -s ./internal ./cmd

watch-coverage:
	go tool cover -html=coverage.out

check-linter:
	golangci-lint run --path-prefix=back/auth -v --config=../../infra/lib/cloud-lint/golangci-lint.yaml --timeout 5m
