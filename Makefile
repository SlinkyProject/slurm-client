##@ General

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build

.PHONY: all
all: build

.PHONY: build
build: fmt tidy vet ## Build project
	go build ./...

.PHONY: clean
clean: ## Clean artifacts
	rm -f cover.out cover.html
	rm -rf bin/

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary (ideally with version)
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f $(1) ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv "$$(echo "$(1)" | sed "s/-$(3)$$//")" $(1) ;\
}
endef

## Tool Binaries
GOLANGCI_LINT ?= $(LOCALBIN)/golangci-lint-$(GOLANGCI_LINT_VERSION)
GOVULNCHECK ?= $(LOCALBIN)/govulncheck-$(GOVULNCHECK_VERSION)

## Tool Versions
GOLANGCI_LINT_VERSION ?= v2.1.6
GOVULNCHECK_VERSION ?= latest

.PHONY: govulncheck-bin
govulncheck-bin: $(GOVULNCHECK) ## Download govulncheck locally if necessary.
$(GOVULNCHECK): $(LOCALBIN)
	$(call go-install-tool,$(GOVULNCHECK),golang.org/x/vuln/cmd/govulncheck,$(GOVULNCHECK_VERSION))


.PHONY: golangci-lint-bin
golangci-lint-bin: $(GOLANGCI_LINT) ## Download golangci-lint locally if necessary.
$(GOLANGCI_LINT): $(LOCALBIN)
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(LOCALBIN) $(GOLANGCI_LINT_VERSION)
	mv $(LOCALBIN)/golangci-lint $(GOLANGCI_LINT)

##@ Development

.PHONY: generate
generate: ## Run all generate targets.
	$(MAKE) generate-api-matrix
	go generate ./...

.PHONY: generate-api-matrix
generate-api-matrix: ## Generate Slurm OpenAPI spec files by matrix.
	declare -A VERSION_MATRIX=( \
		["ghcr.io/slinkyproject/slurmrestd:25.05.0-ubuntu24.04"]="" \
		["ghcr.io/slinkyproject/slurmrestd:24.11.5-ubuntu24.04"]="+inline_enums" \
		["ghcr.io/slinkyproject/slurmrestd:24.05.8-ubuntu24.04"]="+prefer_refs" \
	); \
	for key in $${!VERSION_MATRIX[@]}; do \
		$(MAKE) generate-api SLURM_IMAGE=$${key} SLURM_DATA_PARSER_OPTS=$${VERSION_MATRIX[$${key}]} ; \
	done

CONTAINER_TOOL ?= docker
SLURM_IMAGE ?= ghcr.io/slinkyproject/slurmrestd:25.05.0-ubuntu24.04
SLURM_DATA_PARSER_OPTS ?= +inline_enums

TEMPLATES_DIR = api/.template
OAPI_CODEGEN_VERSION ?= v2.4.1

.PHONY: generate-api
generate-api: ## Generate Slurm OpenAPI spec file.
ifeq ($(SLURM_DATA_PARSER), )
	@$(eval SLURM_DATA_PARSER = $(shell \
		$(CONTAINER_TOOL) run --rm \
			--volume ./hack/etc/slurm:/etc/slurm --volume ./:/workspace --workdir /workspace \
			--env SLURMRESTD_SECURITY=disable_unshare_files,disable_unshare_sysv \
			--user 65534:65534 \
			${SLURM_IMAGE} \
			-d list 2>&1 | grep -Eo 'data_parser/.+' | sort -u | tail -n 1 | sed -e 's/data_parser\///g'))
endif
	@$(eval SLURM_GO_MODULE := $(shell echo ${SLURM_DATA_PARSER} | sed -e 's/\.//g'))
	mkdir -p api/${SLURM_GO_MODULE}
	$(CONTAINER_TOOL) run --rm \
		--volume ./:/workspace --workdir /workspace \
		--env SLURMRESTD_SECURITY=disable_unshare_files,disable_unshare_sysv \
		--user 65534:65534 \
		${SLURM_IMAGE} \
		-f /dev/null -s openapi/slurmdbd,openapi/slurmctld -d data_parser/${SLURM_DATA_PARSER}${SLURM_DATA_PARSER_OPTS} --generate-openapi-spec 2>/dev/null > api/${SLURM_GO_MODULE}/slurm-openapi.gen.json
	$(foreach file, $(wildcard $(TEMPLATES_DIR)/*), sed -e "s/{{SLURM_GO_MODULE}}/${SLURM_GO_MODULE}/g" -e "s/{{OAPI_CODEGEN_VERSION}}/${OAPI_CODEGEN_VERSION}/g" ${file} > api/${SLURM_GO_MODULE}/$(shell basename ${file} .tpl);)

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: tidy
tidy: ## Run go mod tidy against code
	go mod tidy

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: get-u
get-u: ## Run `go get -u`
	go get -u ./...
	$(MAKE) tidy

.PHONY: govulncheck
govulncheck: govulncheck-bin ## Run govulncheck
	$(GOVULNCHECK) ./...

# https://github.com/golangci/golangci-lint/blob/main/.pre-commit-hooks.yaml
.PHONY: golangci-lint
golangci-lint: golangci-lint-bin ## Run golangci-lint.
	$(GOLANGCI_LINT) run --fix

# https://github.com/golangci/golangci-lint/blob/main/.pre-commit-hooks.yaml
.PHONY: golangci-lint-fmt
golangci-lint-fmt: golangci-lint-bin ## Run golangci-lint fmt.
	$(GOLANGCI_LINT) fmt

CODECOV_PERCENT ?= 80.0

.PHONY: test
test: ## Run tests.
	rm -f cover.out cover.html
	go test `go list ./... | grep -v "slurm-client/api"` -v -coverprofile cover.out
	go tool cover -func cover.out
	go tool cover -html cover.out -o cover.html
	@percentage=$$(go tool cover -func=cover.out | grep ^total | awk '{print $$3}' | tr -d '%'); \
		if (( $$(echo "$$percentage < $(CODECOV_PERCENT)" | bc -l) )); then \
			echo "----------"; \
			echo "Total test coverage ($${percentage}%) is less than the coverage threshold ($(CODECOV_PERCENT)%)."; \
			exit 1; \
		fi
