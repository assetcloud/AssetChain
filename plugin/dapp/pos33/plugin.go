// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pos33

import (
	"github.com/33cn/chain33/pluginmgr"
        "github.com/assetcloud/AssetChain/plugin/dapp/pos33/commands"
        "github.com/assetcloud/AssetChain/plugin/dapp/pos33/executor"
        "github.com/assetcloud/AssetChain/plugin/dapp/pos33/rpc"
        "github.com/assetcloud/AssetChain/plugin/dapp/pos33/types"
 
	// init wallet
	_ "github.com/assetcloud/AssetChain/plugin/dapp/pos33/wallet"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.Pos33TicketX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.Pos33TicketCmd,
		RPC:      rpc.Init,
	})
}
