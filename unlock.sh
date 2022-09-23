#!/usr/bin/env bash

function init() {
    main_jrpc="http://localhost:9901"
    echo "=========== # unlock wallet ============="
    result=$(./assetchain-cli --rpc_laddr=${main_jrpc} wallet unlock -p 1314fuzamei -t 0 | jq ".isok")
    if [ "${result}" = "false" ]; then
        exit 1
    fi

}

init
