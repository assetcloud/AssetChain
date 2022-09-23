// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pos33

import (
	"github.com/assetcloud/assetchain/plugin/dapp/pos33/commands"
	"github.com/assetcloud/assetchain/plugin/dapp/pos33/executor"
	"github.com/assetcloud/assetchain/plugin/dapp/pos33/rpc"
	"github.com/assetcloud/assetchain/plugin/dapp/pos33/types"
	"github.com/assetcloud/chain/pluginmgr"

	// init wallet
	_ "github.com/assetcloud/assetchain/plugin/dapp/pos33/wallet"
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
