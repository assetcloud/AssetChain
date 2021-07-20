// +build go1.8

package main

import (
	_ "github.com/33cn/chain33/system"

	_ "github.com/assetcloud/OCIAChain/plugin"

	"github.com/33cn/chain33/util/cli"
	"github.com/assetcloud/OCIAChain/cli/buildflags"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:9901"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName, "ocia")
}
