// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

//only load all plugin and system
import (
	"testing"

	ty "github.com/assetcloud/assetchain/plugin/dapp/pos33/types"
	rpctypes "github.com/assetcloud/chain/rpc/types"
	_ "github.com/assetcloud/chain/system"
	"github.com/assetcloud/chain/types"
	"github.com/assetcloud/chain/util/testnode"
	_ "github.com/assetcloud/plugin/plugin"
	"github.com/stretchr/testify/assert"
)

func TestNewPos33Ticket(t *testing.T) {
	//选票(可以用hotwallet 关闭选票)
	cfg := types.NewChainConfig(cfgstring)
	cfg.GetModuleConfig().Consensus.Name = "pos33"
	mocker := testnode.NewWithConfig(cfg, nil)
	mocker.Listen()
	defer mocker.Close()

	in := &ty.Pos33TicketClose{MinerAddress: mocker.GetHotAddress()}
	var res rpctypes.ReplyHashes
	err := mocker.GetJSONC().Call("pos33.ClosePos33Tickets", in, &res)
	assert.Nil(t, err)
}
