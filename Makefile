OUTPUT_BINARY ?= dist/depot
BUILD_VERSION ?= $(shell git rev-parse HEAD)
BUILD_TIME ?= $(shell date +"%Y-%m-%dT%H:%M:%S")
LDFLAGS := -s -w
LDFLAGS += -X main.buildTime=$(BUILD_TIME)
LDFLAGS += -X main.buildVersion=$(BUILD_VERSION)

DATABASE_URI ?= sqlite://local.db?_fk=1
DEV_DATABASE_URI ?= sqlite://dev?mode=memory&_fk=1
SCHEMA_URI ?= file://schema/tables.sql

.PHONY: build
build: deps generate check
	go build -ldflags '$(LDFLAGS)' -o $(OUTPUT_BINARY) ./main.go

deps:
	mkdir -p $(shell dirname $(OUTPUT_BINARY))
	command -v templ > /dev/null || go install github.com/a-h/templ/cmd/templ@latest
	command -v sqlc > /dev/null || go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	command -v modd > /dev/null || go install github.com/cortesi/modd/cmd/modd@latest
	command -v atlas > /dev/null || curl -sSf https://atlasgo.sh | sh

check:
	go vet ./...

generate: gen-views gen-queries

gen-views:
	templ generate

gen-queries:
	sqlc generate

migrate-diff:
	atlas schema diff --from "$(DATABASE_URI)" --to "$(SCHEMA_URI)" --dev-url "$(DEV_DATABASE_URI)"

migrate:
	atlas schema apply --url $(DATABASE_URI) --to $(SCHEMA_URI) --dev-url "$(DEV_DATABASE_URI)" --auto-approve

clean:
	rm -rf dist
	rm pkg/view/*_templ.go
	rm pkg/db/*.gen.go
