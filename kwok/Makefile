BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD)
REPO_ROOT	    := $(shell dirname ${PWD})
REPO_ROOT_BIN	:= $(REPO_ROOT)/bin

# Image URL to use all building/pushing image targets
IMG_REPO ?= quay.io/kuberay/kwok
IMG_TAG ?= latest
IMG ?= $(IMG_REPO):$(IMG_TAG)

# Allow for additional test flags (-v, etc)
GO_TEST_FLAGS ?=

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (, $(shell go env GOBIN))
	GOBIN = $(shell go env GOPATH)/bin
else
	GOBIN = $(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

# Container Engine to be used for building images
ENGINE ?= docker

all: build

##@ General

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

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: fumpt
fumpt: gofumpt ## Run gofmtumpt against code.
	$(GOFUMPT) -l -w .

.PHONY: imports
imports: goimports ## Run goimports against code.
	$(GOIMPORTS) -l -w .

.PHONY: lint
lint: golangci-lint fmt vet fumpt imports ## Run the linter.
	# exclude the SA1019 check which checks the usage of deprecated fields.
	test -s $(GOLANGCI_LINT) || ($(GOLANGCI_LINT) run --timeout=3m --exclude='SA1019' --no-config --allow-parallel-runners)

build: fmt vet fumpt imports ## Build kwok server binary.
	go build -o ${REPO_ROOT_BIN}/kuberay-kwok cmd/main.go

##@ Docker Build
docker-image: ## Build image for the kwok server.
	$(ENGINE) build -t ${IMG} -f Dockerfile ..

docker-push: ## Push image for the kwok server.
	$(ENGINE) push ${IMG}

##@ Development Tools Setup

## Location to install dependencies to
$(REPO_ROOT_BIN):
	mkdir -p $(REPO_ROOT_BIN)

## Tool Binaries
GOIMPORTS ?= $(REPO_ROOT_BIN)/goimports
GOFUMPT ?= $(REPO_ROOT_BIN)/gofumpt
GOLANGCI_LINT ?= $(REPO_ROOT_BIN)/golangci-lint

## Tool Versions
GOFUMPT_VERSION ?= v0.3.1
GOIMPORTS_VERSION ?= v0.14.0
GOLANGCI_LINT_VERSION ?= v1.64.8

.PHONY: goimports
goimports: $(GOIMPORTS) ## Download goimports locally if necessary
$(GOIMPORTS): $(REPO_ROOT_BIN)
	test -s $(GOIMPORTS) || GOBIN=$(REPO_ROOT_BIN) go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)

.PHONY: gofumpt
gofumpt: $(GOFUMPT) ## Download gofumpt locally if necessary.
$(GOFUMPT): $(REPO_ROOT_BIN)
	test -s $(GOFUMPT) || GOBIN=$(REPO_ROOT_BIN) go install mvdan.cc/gofumpt@$(GOFUMPT_VERSION)

.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT) ## Download golangci_lint locally if necessary.
$(GOLANGCI_LINT): $(REPO_ROOT_BIN)
	test -s $(GOLANGCI_LINT) || (curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- -b $(REPO_ROOT_BIN)/ $(GOLANGCI_LINT_VERSION))

