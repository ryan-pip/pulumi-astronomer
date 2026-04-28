ROOT_DIR         := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
SHELL            := /bin/bash
PROJECT          := github.com/ryan-pip/pulumi-astronomer
NODE_MODULE_NAME := @ryan-pip/pulumi_astronomer
TF_NAME          := astronomer
PROVIDER_PATH    := provider
VERSION_PATH     := ${PROVIDER_PATH}/pkg/version.Version

JAVA_GEN         := pulumi-java-gen
JAVA_GEN_VERSION := v0.9.9
TFGEN            := pulumi-tfgen-astronomer
PROVIDER         := pulumi-resource-astronomer
VERSION          := $(shell pulumictl get version)

TESTPARALLELISM  := 4

WORKING_DIR      := $(shell pwd)

.PHONY: development provider build_sdks build_nodejs build_dotnet build_go build_python cleanup

development:: install_plugins provider lint_provider build_sdks install_sdks cleanup # Build the provider & SDKs for a development environment

# Required for the codegen action that runs in pulumi/pulumi and pulumi/pulumi-terraform-bridge
build:: install_plugins provider build_sdks install_sdks
only_build:: build

tfgen:: install_plugins
	(cd provider && go build -o $(WORKING_DIR)/bin/${TFGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${TFGEN})
	$(WORKING_DIR)/bin/${TFGEN} schema --out provider/cmd/${PROVIDER}
	(cd provider && VERSION=$(VERSION) go generate cmd/${PROVIDER}/main.go)

provider:: tfgen install_plugins # build the provider binary
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER})

build_sdks:: install_plugins provider build_nodejs build_python build_go # build_dotnet # build all the sdks

build_nodejs:: VERSION := $(shell pulumictl get version --language javascript)
build_nodejs:: install_plugins tfgen # build the node sdk
	$(WORKING_DIR)/bin/$(TFGEN) nodejs --overlays provider/overlays/nodejs --out sdk/nodejs/
	cd sdk/nodejs/ && \
        yarn install && \
        yarn run tsc && \
        cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json

build_python:: install_plugins tfgen # build the python sdk
	$(WORKING_DIR)/bin/$(TFGEN) python --out sdk/python/
	cd sdk/python/ && \
		cp ../../README.md . && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		python3 -m venv ./bin/venv && \
		./bin/venv/bin/python -m pip install --quiet build==1.2.1 && \
		cd ./bin && ./venv/bin/python -m build .

build_dotnet:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
build_dotnet:: install_plugins tfgen # build the dotnet sdk
	pulumictl get version --language dotnet
	$(WORKING_DIR)/bin/$(TFGEN) dotnet --overlays provider/overlays/dotnet --out sdk/dotnet/
	cd sdk/dotnet/ && \
		echo "${DOTNET_VERSION}" >version.txt && \
        dotnet build /p:Version=${DOTNET_VERSION}

build_go:: install_plugins tfgen # build the go sdk
	$(WORKING_DIR)/bin/$(TFGEN) go --overlays provider/overlays/go --out sdk/go/
	cd sdk/go/ && \
		go mod tidy

build_java:: PACKAGE_VERSION := $(shell pulumictl get version --language generic)
build_java:: $(WORKING_DIR)/bin/$(JAVA_GEN)
	$(WORKING_DIR)/bin/$(JAVA_GEN) generate --schema provider/cmd/$(PROVIDER)/schema.json --out sdk/java  --build gradle-nexus
	cd sdk/java/ && \
		echo "module fake_java_module // Exclude this directory from Go tools\n\ngo 1.17" > go.mod && \
		gradle --console=plain build

$(WORKING_DIR)/bin/$(JAVA_GEN)::
	$(shell pulumictl download-binary -n pulumi-language-java -v $(JAVA_GEN_VERSION) -r pulumi/pulumi-java)

lint_provider:: provider # lint the provider code
	cd provider && golangci-lint run -c ../.golangci.yml

tidy:: # call go mod tidy in relevant directories
	find ./provider -name go.mod -execdir go mod tidy \;

cleanup:: # cleans up the temporary directory
	rm -r $(WORKING_DIR)/bin
	rm -f provider/cmd/${PROVIDER}/schema.go

help::
	@grep '^[^.#]\+:\s\+.*#' Makefile | \
 	sed "s/\(.\+\):\s*\(.*\) #\s*\(.*\)/`printf "\033[93m"`\1`printf "\033[0m"`	\3 [\2]/" | \
 	expand -t20

clean::
	rm -rf sdk/{dotnet,nodejs,go,python} sdk/go.sum

.PHONY: fmt
fmt::
	@echo "Fixing source code with gofmt..."
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

install_plugins::
	[ -x $(shell which pulumi) ] || curl -fsSL https://get.pulumi.com | sh
	pulumi plugin install resource random 4.3.1

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::

install_go_sdk::

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk

test:: test_python # run all hermetic tests (no live API calls)
	cd provider && go test -v -short ./...

test_python_smoke:: build_python # mocked Python SDK smoke tests; never touches system Python
	cd tests/python && \
		uv venv --clear --quiet .venv && \
		. .venv/bin/activate && \
		uv pip install --quiet pytest "pulumi>=3.0.0,<4.0.0" && \
		uv pip install --quiet ../../sdk/python/bin/dist/pulumi_astronomer-*.whl && \
		pytest -v

# Live Python integration tests against the Astronomer API.
# Requires ASTRO_API_TOKEN + ASTRO_WORKSPACE_ID. Tests skip themselves when unset.
# GOWORK=off because examples/ is an independent module not listed in ./go.work.
# PATH prepend makes the locally-built `pulumi-resource-astronomer` binary
# discoverable by the test's pulumi engine — otherwise it tries to download
# the plugin from GitHub Releases at the build version, which doesn't exist.
test_python:: export PATH := $(WORKING_DIR)/bin:$(PATH)
test_python:: provider build_python
	cd examples && GOWORK=off go test -v -tags=python -parallel 1 -timeout 30m -run 'TestAcc.*Py'
