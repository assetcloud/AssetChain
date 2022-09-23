#!/usr/bin/env bash

user_account=$1
miner_account=$2
amount=$3
user_sk=$4

function init() {
    main_jrpc="http://localhost:9901"

    sleep 1
    echo "=========== # user_account deposit to pos33  ============="
    deposit_hash=$( ./assetchain-cli --rpc_laddr=${main_jrpc} send coins transfer -a $amount -t 0x79a69527dd0d62dcfd3eae90eb6c3134c57489eb -k $user_sk --signAddrType 2)
    echo "${deposit_hash}"
    if [ -z "${deposit_hash}" ]; then
        exit 1
    fi

    sleep 1
    echo "=========== # transf_account entrust mining  ============="
    entrust_hash=$(./assetchain-cli --rpc_laddr=${main_jrpc} send pos33 entrust -a $amount -e $miner_account -r $user_account -k $user_sk --signAddrType 2)
    echo "${entrust_hash}"
    if [ -z "${entrust_hash}" ]; then
        exit 1
    fi
}

init
