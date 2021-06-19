tests:
	go test -v ./internal/... -coverprofile coverage.out

coverage: tests
	go tool cover -html=coverage.out