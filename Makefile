run: build
	@./go-pardocs
    
build: 
	@go build ./cmd/go-pardocs

test:
	@go test .
