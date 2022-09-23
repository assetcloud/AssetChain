package plugin

import (
	_ "github.com/assetcloud/assetchain/plugin/consensus/init" //register consensus init package
	_ "github.com/assetcloud/assetchain/plugin/crypto/init"
	_ "github.com/assetcloud/assetchain/plugin/dapp/init"
	_ "github.com/assetcloud/assetchain/plugin/mempool/init"
	_ "github.com/assetcloud/assetchain/plugin/p2p/init"
	_ "github.com/assetcloud/assetchain/plugin/store/init"
)
