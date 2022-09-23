//go:build go1.13
// +build go1.13

package main

import (
	_ "github.com/assetcloud/assetchain/plugin"
	_ "github.com/assetcloud/chain/system"

	"flag"
	"runtime/debug"

	"github.com/assetcloud/chain/util/cli"
)

var percent = flag.Int("p", 0, "SetGCPercent")

func main() {
	flag.Parse()
	if *percent < 0 || *percent > 100 {
		*percent = 0
	}
	if *percent > 0 {
		debug.SetGCPercent(*percent)
	}
	cli.RunChain("assetchain", assetChainConfig)
}
