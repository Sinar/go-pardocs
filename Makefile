run: build
	@./go-pardocs
    
build: test
	@go build ./cmd/go-pardocs

test:
	@go test .
