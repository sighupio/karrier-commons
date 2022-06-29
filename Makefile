_PROJECT_DIRECTORY = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
_GOLANG_IMAGE = quay.io/sighup/golang:1.18.1
_PROJECTNAME = karrier-commons
_GOARCH = "amd64"

ifeq ("$(shell uname -m)", "arm64")
	_GOARCH = "arm64"
endif

#1: docker image
#2: make target
define run-docker
	@docker run --rm \
		-e CGO_ENABLED=0 \
		-e GOARCH=${_GOARCH} \
		-e GOOS=linux \
		-e GOPRIVATE="github.com/sighupio/*" \
		-w /app \
		-v ${NETRC_FILE}:/root/.netrc \
		-v ${_PROJECT_DIRECTORY}:/app \
		$(1) $(2)
endef

.PHONY: env

env:
	@echo 'export CGO_ENABLED=0'
	@echo 'export GOARCH=${_GOARCH}'
	@echo 'export GOPRIVATE=github.com/sighupio/*'

.PHONY: mod-download mod-tidy mod-verify

mod-download:
	@go mod download

mod-tidy:
	@go mod tidy

mod-verify:
	@go mod verify

.PHONY: mod-check-upgrades mod-upgrade

mod-check-upgrades:
	@go list -mod=readonly -u -f "{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}" -m all

mod-upgrade:
	@go get -u ./... && go mod tidy

.PHONY: generate license

generate:
	@go generate ./...

license-add:
	@addlicense -c "SIGHUP s.r.l" -y 2017-present -v -l bsd \
	-ignore "scripts/e2e/libs/**/*" \
	-ignore "vendor/**/*" \
	-ignore "*.gen.go" \
	-ignore ".idea/*" \
	-ignore ".vscode/*" \
	-ignore "*.js" \
	-ignore "kind-config.yaml" \
	-ignore ".husky/**/*" \
	-ignore ".go/**/*" \
	.

license-check:
	@addlicense -c "SIGHUP s.r.l" -y 2017-present -v -l bsd \
	-ignore "scripts/e2e/libs/**/*" \
	-ignore "vendor/**/*" \
	-ignore "*.gen.go" \
	-ignore ".idea/*" \
	-ignore ".vscode/*" \
	-ignore "*.js" \
	-ignore "kind-config.yaml" \
	-ignore ".husky/**/*" \
	-ignore ".go/**/*" \
	--check .

.PHONY: fmt fumpt

fmt:
	@find . -name *.go -type f -not -path '*/vendor/*' \
	| sed 's/^\.\///g' \
	| xargs -I {} sh -c 'echo "formatting {}.." && gofmt -w -s {}'

fumpt:
	@find . -name *.go -type f -not -path '*/vendor/*' \
	| sed 's/^\.\///g' \
	| xargs -I {} sh -c 'echo "formatting {}.." && gofumpt -w -extra {}'

.PHONY: lint lint-go

lint: lint-go

lint-go:
	@golangci-lint -v run --color=always --config=${_PROJECT_DIRECTORY}/.rules/.golangci.yml ./...

.PHONY: test-unit

test-unit:
	@gotestsum --no-color=false -- -tags=unit ./...

# Helpers

%-docker: check-variable-NETRC_FILE
	$(call run-docker,${_GOLANG_IMAGE},make $*)

check-variable-%: # detection of undefined variables.
	@[[ "${${*}}" ]] || (echo '*** Please define variable `${*}` ***' && exit 1)
