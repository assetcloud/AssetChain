#!/usr/bin/env bash

set -e
set -o pipefail
#set -o verbose
#set -o xtrace

# os: ubuntu18.04 x64

sedfix=""
if [ "$(uname)" == "Darwin" ]; then
    sedfix=".bak"
fi

## get chain path
CHAIN_PATH=$(go list -f "{{.Dir}}" github.com/assetcloud/chain)
PlUGIN_PATH=$(go list -f "{{.Dir}}" github.com/assetcloud/plugin)

function build_auto_test() {

    trap "rm -f ../autotest/main.go" INT TERM EXIT
    cp ../../bityuan ../chain
    cp ../../bityuan-cli ../chain-cli
    cp ../../bityuan.toml ../chain.toml
    local AutoTestMain="${CHAIN_PATH}/cmd/autotest/main.go"
    cp "${AutoTestMain}" ./
    sed -i $sedfix '/^package/a import _ \"github.com\/assetcloud\/plugin\/plugin\"' main.go
    go build -v -i -o autotest
}

function copyAutoTestConfig() {

    declare -a ChainAutoTestDirs=("${CHAIN_PATH}/system" "$PlUGIN_PATH/plugin")
    echo "#copy auto test config to path \"$1\""
    local AutoTestConfigFile="$1/autotest.toml"

    #pre config auto test
    {

        echo 'cliCmd="./chain-cli"'
        echo "checkTimeout=60"
    } >"${AutoTestConfigFile}"

    #copy all the dapp test case config file
    for rootDir in "${ChainAutoTestDirs[@]}"; do

        if [[ ! -d ${rootDir} ]]; then
            continue
        fi

        testDirArr=$(find "${rootDir}" -type d -name autotest)

        for autotest in ${testDirArr}; do

            dapp=$(basename "$(dirname "${autotest}")")
            dappConfig=${autotest}/${dapp}.toml

            #make sure dapp have auto test config
            if [ -e "${dappConfig}" ]; then

                cp "${dappConfig}" "$1"/

                #add dapp test case config
                {
                    echo "[[TestCaseFile]]"
                    echo "dapp=\"$dapp\""
                    echo "filename=\"$dapp.toml\""
                } >>"${AutoTestConfigFile}"

            fi

        done
    done
}

function copyChain() {

    echo "# copy chain bin to path \"$1\", make sure build chain"
    cp ../chain ../chain-cli ../chain.toml "$1"
    cp "${CHAIN_PATH}"/cmd/chain/chain.test.toml "$1"
}

function copyAll() {

    dir="$1"
    #check dir exist
    if [[ ! -d ${dir} ]]; then
        mkdir "${dir}"
    fi
    cp autotest "${dir}"
    copyAutoTestConfig "${dir}"
    copyChain "${dir}"
    echo "# all copy have done!"
}

function main() {

    if [[ $1 == "build" ]]; then #build autotest
        build_auto_test
    else
        dir="$1"
        echo "$dir"
        rm -rf ../autotest/"$dir" && mkdir "$dir"
        cp -r "$CHAIN_PATH"/build/autotest/"$dir"/* ./"$dir"/ && copyAll "$dir"
        chmod -R 755 "$dir" && cd "$dir" && ./autotest.sh "${@:2}" && cd ../
    fi
}

main "$@"
