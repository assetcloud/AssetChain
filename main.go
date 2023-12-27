package main

import (
	"flag"
	"fmt"
	"runtime/debug"

	_ "github.com/assetcloud/assetchain/plugin"
	"github.com/assetcloud/assetchain/version"
	_ "github.com/assetcloud/chain/system"

	frameVersion "github.com/assetcloud/chain/common/version"
	"github.com/assetcloud/chain/util/cli"
	pluginVersion "github.com/assetcloud/plugin/version"
)

var (
	percent    = flag.Int("p", 0, "SetGCPercent")
	versionCmd = flag.Bool("version", false, "assetchain detail version")
)

func main() {
	flag.Parse()
	if *versionCmd {
		fmt.Printf("Build time: %s", version.BuildTime)
		fmt.Printf("System version: %s", version.Platform)
		fmt.Printf("Golang version: %s", version.GoVersion)
		fmt.Printf("assetchain version: %s", version.GetVersion())
		fmt.Printf("Chain frame version: %s", frameVersion.GetVersion())
		fmt.Printf("Chain plugin version: %s", pluginVersion.GetVersion())
		return
	}
	if *percent < 0 || *percent > 100 {
		*percent = 0
	}
	if *percent > 0 {
		debug.SetGCPercent(*percent)
	}
	cli.RunChain("assetchain", assetchain)
}
