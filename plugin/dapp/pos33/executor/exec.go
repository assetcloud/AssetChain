// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"errors"

	ty "github.com/assetcloud/assetchain/plugin/dapp/pos33/types"
	"github.com/assetcloud/chain/types"
)


//Exec_Miner exec miner
func (t *Pos33Ticket) Exec_Miner(payload *ty.Pos33MinerMsg, tx *types.Transaction, index int) (*types.Receipt, error) {
	actiondb := NewAction(t, tx)
	r, err := actiondb.Pos33MinerNew(payload, index)
	if err != nil {
		panic(err)
	}
	return r, nil
}

// // Exec_Tbind exec bind
// func (t *Pos33Ticket) Exec_Tbind(payload *ty.Pos33TicketBind, tx *types.Transaction, index int) (*types.Receipt, error) {
// 	actiondFailed to connect to github.com port 443 after 4077 ms: Network is unreachableb := NewAction(t, tx)
// 	return actiondb.Pos33TicketBind(payload)
// }

// Exec_Entrust exec entrust
func (t *Pos33Ticket) Exec_Entrust(payload *ty.Pos33Entrust, tx *types.Transaction, index int) (*types.Receipt, error) {
	action := NewAction(t, tx)
	chainCfg := action.api.GetConfig()
	if !chainCfg.IsDappFork(action.height, ty.Pos33TicketX, "UseEntrust") {
		return nil, errors.New("config exec.pos33.UseEntrust error")
	}
	return action.Pos33Entrust(payload)
}

// // Exec_Migrate exec migrate
// func (t *Pos33Ticket) Exec_Migrate(payload *ty.Pos33Migrate, tx *types.Transaction, index int) (*types.Receipt, error) {
// 	action := NewAction(t, tx)
// 	return action.Pos33Migrate(payload)
// }

// Exec_BlsBind exec bls bind
func (t *Pos33Ticket) Exec_BlsBind(payload *ty.Pos33BlsBind, tx *types.Transaction, index int) (*types.Receipt, error) {
	action := NewAction(t, tx)
	return action.Pos33BlsBind(payload)
}

// // Exec_WithdrawReward exec withdraw reward
// func (t *Pos33Ticket) Exec_Withdraw(payload *ty.Pos33WithdrawReward, tx *types.Transaction, index int) (*types.Receipt, error) {
// 	action := NewAction(t, tx)
// 	return action.Pos33WithdrawReward(payload)
// }

// // Exec_SetMinerFeeRate exec set miner fee rate
// func (t *Pos33Ticket) Exec_FeeRate(payload *ty.Pos33MinerFeeRate, tx *types.Transaction, index int) (*types.Receipt, error) {
// 	action := NewAction(t, tx)
// 	return action.Pos33SetMinerFeeRate(payload)
// }
