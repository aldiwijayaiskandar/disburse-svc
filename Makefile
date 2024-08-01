# ==============================================================================
# Main

main:
	go run ./cmd/disburse/main.go
	go run ./cmd/user/main.go
	go run ./cmd/wallet/main.go

build:
	go build ./cmd/api_gateway/main.go
	go build ./cmd/disburse/main.go
	go build ./cmd/user/main.go
	go build ./cmd/wallet/main.go

test:
	go test -cover ./...

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

# ==============================================================================
# Docker compose commands

compose:
	echo "Starting docker environment"
	docker-compose -f docker-compose.yml up --build