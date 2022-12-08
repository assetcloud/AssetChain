// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testnode

import (
	"io/ioutil"
	"testing"

	ty "github.com/assetcloud/assetchain/plugin/dapp/pos33/types"
	ticketwallet "github.com/assetcloud/assetchain/plugin/dapp/pos33/wallet"
	"github.com/assetcloud/chain/util/testnode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/assetcloud/chain/system"
	"github.com/assetcloud/chain/types"
	_ "github.com/assetcloud/plugin/plugin"
)

func TestWalletPos33Ticket(t *testing.T) {
	minerAddr := "12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"
	t.Log("Begin wallet ticket test")

	strCfg, err := ioutil.ReadFile("./chain.toml")
	if err != nil {
		t.Error(err)
		return
	}
	cfg := types.NewChainConfig(string(strCfg))
	cfg.GetModuleConfig().Consensus.Name = "pos33"
	mock33 := testnode.NewWithConfig(cfg, nil)
	defer mock33.Close()
	err = mock33.WaitHeight(0)
	assert.Nil(t, err)
	msg, err := mock33.GetAPI().Query(ty.Pos33TicketX, "Pos33TicketList", &ty.Pos33TicketList{Addr: minerAddr, Status: 1})
	assert.Nil(t, err)
	// ticketList := msg.(*ty.ReplyPos33TicketList)
	// assert.NotNil(t, ticketList)
	//return
	ticketwallet.FlushPos33Ticket(mock33.GetAPI())
	err = mock33.WaitHeight(2)
	assert.Nil(t, err)
	header, err := mock33.GetAPI().GetLastHeader()
	require.Equal(t, err, nil)
	require.Equal(t, header.Height >= 2, true)

	in := &ty.Pos33TicketClose{MinerAddress: minerAddr}
	msg, err = mock33.GetAPI().ExecWalletFunc(ty.Pos33TicketX, "ClosePos33Tickets", in)
	assert.Nil(t, err)
	hashes := msg.(*types.ReplyHashes)
	assert.NotNil(t, hashes)

	in = &ty.Pos33TicketClose{}
	msg, err = mock33.GetAPI().ExecWalletFunc(ty.Pos33TicketX, "ClosePos33Tickets", in)
	assert.Nil(t, err)
	hashes = msg.(*types.ReplyHashes)
	assert.NotNil(t, hashes)
	t.Log("End wallet ticket test")
}
