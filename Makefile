run_test:
	go test ./... -short -cover -count=1

tidy:
	go mod tidy

lint:	
	golangci-lint run
	
vuln_check:
	govulncheck ./...