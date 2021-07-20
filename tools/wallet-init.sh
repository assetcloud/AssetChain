#!/usr/bin/env bash

function init() {
    echo "=========== # start set wallet 1 ============="
    echo "=========== # save seed to wallet ============="
    result=$(./ocia-cli seed generate -l 0)
    result=$(./ocia-cli seed save -p 1314fuzamei -s "${result}" | jq ".isok")
    if [ "${result}" = "false" ]; then
        echo "save seed to wallet error seed, result: ${result}"
        exit 1
    fi

    sleep 1

    echo "=========== # unlock wallet ============="
    result=$(./ocia-cli wallet unlock -p 1314fuzamei -t 0 | jq ".isok")
    if [ "${result}" = "false" ]; then
        exit 1
    fi

    sleep 1

    echo "=========== # create new key for mining ============="
    result=$(./ocia-cli account create -l mining | jq ".acc")
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi

    sleep 1
    echo "=========== # set auto mining ============="
    result=$(./ocia-cli pos33 auto_mine -f 1 | jq ".isok")
    if [ "${result}" = "false" ]; then
        exit 1
    fi

    echo "=========== # end set wallet 1 ============="

}

init
