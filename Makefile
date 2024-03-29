
export GO111MODULE=on
export CHAIN_PATH=$(shell go list -f {{.Dir}} github.com/assetcloud/chain)
export PLUGIN_PATH=$(shell go list -f {{.Dir}} github.com/assetcloud/plugin)
PKG_LIST_VET := `go list ./... | grep -v "vendor" | grep -v plugin/dapp/evm/executor/vm/common/crypto/bn256`
PKG_LIST_INEFFASSIGN= `go list -f {{.Dir}} ./... | grep -v "vendor"`
LDFLAGS := ' -w -s'
BUILDTIME:=$(shell date +"%Y-%m-%d %H:%M:%S %A")
VERSION=$(shell git describe --tags || git rev-parse --short=8 HEAD)
GitCommit=$(shell git rev-parse --short=8 HEAD)
BUILD_FLAGS := -ldflags '-X "github.com/assetcloud/assetchain/version.GitCommit=$(GitCommit)" \
                         -X "github.com/assetcloud/chain/common/version.GitCommit=$(GitCommit)" \
                         -X "github.com/assetcloud/assetchain/version.BuildTime=$(BUILDTIME)"'

.PHONY: default build

default: build

all:  build

build: toolimport
	CGO_ENABLED=1 go build ${BUILD_FLAGS} -v  -o assetchain
	CGO_ENABLED=1 go build ${BUILD_FLAGS} -v  -o assetchain-cli github.com/assetcloud/assetchain/cli


PLATFORM_LIST = \
	darwin-amd64 \
	linux-amd64

WINDOWS_ARCH_LIST = \
	windows-amd64

GOBUILD=CGO_ENABLED=1 go build $(BUILD_FLAGS)' -X "github.com/assetcloud/assetchain/version.Version=$(VERSION)"  -w -s'
SRC_CLI := ./cli
SRC := ./
APP := assetchain
CLI := assetchain-cli

linux-action-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(APP)-linux-amd64 $(SRC)
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(CLI)-linux-amd64 $(SRC_CLI)
	chmod +x $(APP)-linux-amd64 $(CLI)-linux-amd64
	tar -zcvf build/$(APP)-linux-amd64.tar.gz $(APP)-linux-amd64  $(CLI)-linux-amd64 CHANGELOG.md assetchain-fullnode.toml assetchain.toml

_GOBUILD := CGO_ENABLED=1 go build $(BUILD_FLAGS)' -w -s'
linux-amd64:
	GOARCH=amd64 GOOS=linux $(_GOBUILD) -o $(APP)-$@ $(SRC)
	GOARCH=amd64 GOOS=linux $(_GOBUILD) -o $(CLI)-$@ $(SRC_CLI)
	chmod +x $(APP)-$@ $(CLI)-$@
	tar -zcvf build/$(APP)-$@.tar.gz $(APP)-$@  $(CLI)-$@ CHANGELOG.md assetchain-fullnode.toml assetchain.toml

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(_GOBUILD) -o $(APP)-$@ $(SRC)
	GOARCH=amd64 GOOS=darwin $(_GOBUILD) -o $(CLI)-$@ $(SRC_CLI)
	chmod +x $(APP)-$@ $(CLI)-$@
	tar -zcvf build/$(APP)-$@.tar.gz $(APP)-$@  $(CLI)-$@ CHANGELOG.md assetchain-fullnode.toml assetchain.toml

windows-amd64:
	GOARCH=amd64 GOOS=windows $(_GOBUILD) -o $(APP)-$@.exe $(SRC)
	GOARCH=amd64 GOOS=windows $(_GOBUILD) -o $(CLI)-$@.exe $(SRC_CLI)
	#zip -j build/$(APP)-$@.zip $(APP)-$@.exe  $(CLI)-$@.exe CHANGELOG.md assetchain-fullnode.toml assetchain.toml

#make updateplugin version=xxx
#单独更新plugin或chain, version可以是tag或者commit哈希(tag必须是--vMajor.Minor.Patch--规范格式)
updateplugin:
	@if [ -n "$(version)" ]; then   \
    go get -d github.com/assetcloud/plugin@${version}; \
    else \
    go get -d github.com/assetcloud/plugin@master;fi
updatechain:
	@if [ -n "$(version)" ]; then   \
	go get -d github.com/assetcloud/chain@${version}; \
	else \
	go get -d github.com/assetcloud/chain@master;fi

#make update version=xxx, 同时更新chain和plugin, 两个项目必须有相同的tag(tag必须是--vMajor.Minor.Patch--规范格式)
update:updatechain updateplugin

vet:
	@go vet ${PKG_LIST_VET}

ineffassign:
	@golangci-lint  run --no-config --issues-exit-code=1  --deadline=2m --disable-all   --enable=ineffassign -n ${PKG_LIST_INEFFASSIGN}

linter: vet ineffassign ## Use gometalinter check code, ignore some unserious warning
	@./golinter.sh "filter"

.PHONY: checkgofmt
checkgofmt: ## get all go files and run go fmt on them
	@files=$$(find . -name '*.go' -not -path "./vendor/*" | xargs gofmt -l -s); if [ -n "$$files" ]; then \
		  echo "Error: 'make fmt' needs to be run on:"; \
		  echo "${files}"; \
		  exit 1; \
		  fi;
	@files=$$(find . -name '*.go' -not -path "./vendor/*" | xargs goimports -l -w); if [ -n "$$files" ]; then \
		  echo "Error: 'make fmt' needs to be run on:"; \
		  echo "${files}"; \
		  exit 1; \
		  fi;

fmt_shell: ## check shell file
	@find . -name '*.sh' -not -path "./vendor/*" | xargs shfmt -w -s -i 4 -ci -bn

fmt: fmt_shell ## go fmt
	@go fmt ./...
	@find . -name '*.go' -not -path "./vendor/*" | xargs goimports -l -w

largefile-check:
	git gc
	./findlargefile.sh

autotest: ## build autotest binary
	@cd build/autotest && bash ./run.sh build && cd ../../
	@if [ -n "$(dapp)" ]; then \
		cd build/autotest && bash ./run.sh local $(dapp) && cd ../../; fi

buildtool: ## chain tool
	@go build  -o tool `go list -f {{.Dir}} github.com/assetcloud/chain`/cmd/tools

toolimport: buildtool ## update plugin import
	@./tool import --path "plugin" --packname "github.com/assetcloud/assetchain/plugin" --conf "plugin/plugin.toml"

clean:
	@rm -rf datadir
	@rm -rf logs
	@rm -rf wallet
	@rm -rf grpc33.log
	@rm -rf assetchain
	@rm -rf assetchain-cli
	@rm -rf tool
