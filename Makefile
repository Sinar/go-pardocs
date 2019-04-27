run: build
	@./go-pardocs
    
build: test
	@go build .

test:
	@go test .
