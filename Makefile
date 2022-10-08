.PHONY: build
build: clean
	CGO_ENABLED=0 go build -o output/bcrypt-tool

.PHONY: clean
clean:
	rm -f output/bcrypt-tool

.PHONY: test
test:
	go test -race ./...

.PHONY: vet
vet:
	go vet ./...

default: build
