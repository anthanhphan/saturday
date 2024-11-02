run_test:
	go test ./... -short -cover

tidy:
	go mod tidy

lint:	
	golangci-lint run
	
vuln_check:
	govulncheck ./...