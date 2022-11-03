//go:build go1.8
// +build go1.8

package main

import (
	_ "github.com/assetcloud/assetchain/plugin"
	_ "github.com/assetcloud/chain/system"

	"github.com/assetcloud/assetchain/cli/buildflags"
	"github.com/assetcloud/chain/util/cli"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:8801"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName, "xxx")
}
