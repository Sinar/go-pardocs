run: build
	@./go-pardocs
    
build: 
	@go build ./cmd/go-pardocs

test:
	@go test ./...

# https://vic.demuzere.be/articles/golang-makefile-crosscompile/
PLATFORMS := linux/amd64 windows/amd64 darwin/amd64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

release: $(PLATFORMS)

$(PLATFORMS):
	@mkdir -p dist/$(os)-$(arch)
	@GOOS=$(os) GOARCH=$(arch) go build -o './dist/$(os)-$(arch)/go-pardocs_$(os)-$(arch)' ./cmd/go-pardocs

.PHONY  release: $(PLATFORMS)
