#!/usr/bin/env bash

function init() {
    echo "=========== # start set wallet 1 ============="
    echo "=========== # save seed to wallet ============="
    result=$(./assetchain-cli seed generate -l 0)
    result=$(./assetchain-cli seed save -p asset1314 -s "${result}" | jq ".isok")
    if [ "${result}" = "false" ]; then
        echo "save seed to wallet error seed, result: ${result}"
        exit 1
    fi

    sleep 1

    echo "=========== # unlock wallet ============="
    result=$(./assetchain-cli wallet unlock -p asset1314 -t 0 | jq ".isok")
    if [ "${result}" = "false" ]; then
        exit 1
    fi

    sleep 1

    echo "=========== # create new key for mining ============="
    result=$(./assetchain-cli account create -t 2 -l mining | jq ".acc")
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi

    echo "=========== # end set wallet 1 ============="

}

init
