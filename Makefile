
export GO111MODULE=on
export chain_PATH=$(shell go list -f {{.Dir}} github.com/assetcloud/chain)
export PLUGIN_PATH=$(shell go list -f {{.Dir}} github.com/assetcloud/plugin)
PKG_LIST_VET := `go list ./... | grep -v "vendor" | grep -v plugin/dapp/evm/executor/vm/common/crypto/bn256`
PKG_LIST_INEFFASSIGN= `go list -f {{.Dir}} ./... | grep -v "vendor"`
BUILD_FLAGS = -ldflags "-X github.com/assetcloud/chain/common/version.GitCommit=`git rev-parse --short=8 HEAD`"
APP = assetchain
CLI = assetchain-cli
SRC_CLI = github.com/assetcloud/assetchain/cli

.PHONY: default build

default: build

all:  build

build: toolimport
	go build ${BUILD_FLAGS} -v  -o assetchain
	go build ${BUILD_FLAGS} -v  -o assetchain-cli github.com/assetcloud/assetchain/cli

pkg:
	rm assetchain-pkg assetchain-pkg.tgz -rf
	mkdir assetchain-pkg
	cp assetchain assetchain-cli assetchain.toml wallet-init.sh assetchain-pkg
	tar zcfv assetchain-pkg.tgz assetchain-pkg

GOBUILD=go build $(BUILD_FLAGS) $(LDFLAGS)

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(APP)
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(CLI) $(SRC_CLI)
	rm assetchain-pkg assetchain-$@.tgz -rf
	mkdir assetchain-pkg
	cp assetchain assetchain-cli assetchain.toml wallet-init.sh assetchain-pkg
	tar zcvf assetchain-$@.tgz assetchain-pkg

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(APP)
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(CLI) $(SRC_CLI)
	rm assetchain-pkg assetchain-$@.tgz -rf
	mkdir assetchain-pkg
	cp assetchain assetchain-cli assetchain.toml wallet-init.sh assetchain-pkg
	tar zcvf assetchain-$@.tgz assetchain-pkg

linux-arm64:
	GOARCH=arm64 GOOS=linux $(GOBUILD) -o $(APP)
	GOARCH=arm64 GOOS=linux $(GOBUILD) -o $(CLI) $(SRC_CLI)
	rm assetchain-pkg assetchain-$@.tgz -rf
	mkdir assetchain-pkg
	cp assetchain assetchain-cli assetchain.toml wallet-init.sh assetchain-pkg
	tar zcvf assetchain-$@.tgz assetchain-pkg

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(APP).exe
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(CLI).exe $(SRC_CLI)
	rm assetchain-pkg assetchain-$@.tgz -rf
	mkdir assetchain-pkg
	cp assetchain.exe assetchain-cli.exe assetchain.toml wallet-init.sh assetchain-pkg
	tar zcvf assetchain-$@.tgz assetchain-pkg


#make updateplugin version=xxx
#单独更新plugin或chain, version可以是tag或者commit哈希(tag必须是--vMajor.Minor.Patch--规范格式)
updateplugin:
	@if [ -n "$(version)" ]; then   \
    go get github.com/assetcloud/plugin@${version}; \
    else \
    go get github.com/assetcloud/plugin@master;fi
updatechain:
	@if [ -n "$(version)" ]; then   \
	go get github.com/assetcloud/chain@${version}; \
	else \
	go get github.com/assetcloud/chain@master;fi

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


buildtool: ## chain tool
	@go build -o tool `go list -f {{.Dir}} github.com/assetcloud/chain`/cmd/tools

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
