default:
	go mod tidy
	go mod download
	go build -v
	go fmt ./...
	go build -v
	@echo "success..."