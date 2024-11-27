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
	rm -f govulnreport.txt
	rm -rf bin/

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

##@ Development

.PHONY: generate
generate: ## Run all generate targets.
	$(MAKE) generate-api
	go generate ./...

CONTAINER_TOOL ?= docker
SLURM_IMAGE ?= ghcr.io/slinkyproject/slurmrestd:24.05-ubuntu-24.04

TEMPLATES_DIR = api/.template
OAPI_CODEGEN_VERSION ?= v2.3.0

.PHONY: generate-api
generate-api: ## Generate Slurm OpenAPI spec file.
ifeq ($(SLURM_DATA_PARSER), )
	@$(eval SLURM_DATA_PARSER = $(shell \
		$(CONTAINER_TOOL) run --rm -t \
		--volume ./hack/etc/slurm:/etc/slurm --volume ./:/workspace --workdir /workspace \
		--env SLURMRESTD_SECURITY=disable_unshare_files,disable_unshare_sysv,disable_user_check \
		${SLURM_IMAGE} \
		-d list 2>&1 | grep -Eo 'data_parser/.+' | sort -u | tail -n 1 | sed -e 's/data_parser\///g'))
endif
	@$(eval SLURM_GO_MODULE := $(shell echo ${SLURM_DATA_PARSER} | sed -e 's/\.//g'))
	mkdir -p api/${SLURM_GO_MODULE}
	$(CONTAINER_TOOL) run --rm -t \
		--volume ./:/workspace --workdir /workspace \
		--env SLURMRESTD_SECURITY=disable_unshare_files,disable_unshare_sysv,disable_user_check \
		${SLURM_IMAGE} \
		-f /dev/null -s openapi/slurmdbd,openapi/slurmctld -d data_parser/${SLURM_DATA_PARSER}+prefer_refs --generate-openapi-spec 2>/dev/null > api/${SLURM_GO_MODULE}/slurm-openapi.gen.json
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

.PHONY: test
test: ## Run tests.
	rm -f cover.out cover.html
	go test `go list ./... | grep -v "slurm-client/api"` -v -coverprofile cover.out
	go tool cover -html cover.out -o cover.html

CODECOV_PERCENT ?= 56.0

.PHONY: codecov
codecov: test ## Run codecov checking
	@percentage=$$(go tool cover -func=cover.out | grep ^total | awk '{print $$3}' | tr -d '%'); \
	if [ $$(echo "$$percentage < $(CODECOV_PERCENT)" | bc -l) -eq 1 ]; then \
		echo "The total percentage ($$percentage%) is less than $(CODECOV_PERCENT)%.";  \
	exit 1; \
	else \
		echo "The total percentage ($$percentage%) is greater than or equal to $(CODECOV_PERCENT)%."; \
	fi

.PHONY: vuln-scan
vuln-scan: ## Run vulnerability scanning tool
	GOBIN=$(LOCALBIN) go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./... > govulnreport.txt 2>&1 || echo "Found vulnerabilities.  Details in govulnreport.txt"

.PHONY: audit
audit: fmt tidy vet vuln-scan ## Run testing tools
