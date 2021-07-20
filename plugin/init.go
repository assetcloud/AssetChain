package plugin

import (
	_ "github.com/assetcloud/OCIAChain/plugin/consensus/init" //register consensus init package
	_ "github.com/assetcloud/OCIAChain/plugin/crypto/init"
	_ "github.com/assetcloud/OCIAChain/plugin/dapp/init"
	_ "github.com/assetcloud/OCIAChain/plugin/mempool/init"
	_ "github.com/assetcloud/OCIAChain/plugin/store/init"
)
