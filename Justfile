set shell := ["bash", "-u", "-c"]

export scripts := ".github/workflows/scripts"
export GOBIN := `echo $PWD/.bin`
export TAG := `git describe --tags $(git rev-list --tags --max-count=1)`

# show available commands
[private]
default:
    @just --list

# tidy up Go modules
[group('build')]
tidy:
    go mod tidy

# run tests across source tree
[group('testing')]
tests:
    go test -v -race -count=1 ./...

# apply go vet command on source tree
[group('lint')]
vet:
    go vet ./...

# apply golangci-lint linters on source tree
[group('lint')]
lint: vet
    $GOBIN/golangci-lint run --config {{scripts}}/golangci.yaml

# show host system information
[group('build')]
@sysinfo:
    echo "{{os()/arch()}} {{num_cpus()}}c"

# locally install build dependencies
[group('build')]
init:
    go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4

# create release binaries
[group('release')]
release:
    export GORELEASER_CURRENT_TAG={{TAG}}
    envy exec gh-release goreleaser release --clean --config {{scripts}}/goreleaser.yaml
    rm -rf dist
