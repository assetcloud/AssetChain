package plugin

import (
	_ "github.com/assetcloud/AssetChain/plugin/consensus/init" //register consensus init package
	_ "github.com/assetcloud/AssetChain/plugin/crypto/init"
	_ "github.com/assetcloud/AssetChain/plugin/dapp/init"
	_ "github.com/assetcloud/AssetChain/plugin/mempool/init"
	_ "github.com/assetcloud/AssetChain/plugin/store/init"
)
