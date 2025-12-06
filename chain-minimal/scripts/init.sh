#!/usr/bin/env bash

rm -rf $HOME/.bvchaind
BVCHAIND_BIN=$(which bvchaind)
if [ -z "$BVCHAIND_BIN" ]; then
    GOBIN=$(go env GOPATH)/bin
    BVCHAIND_BIN=$(which $GOBIN/bvchaind)
fi

if [ -z "$BVCHAIND_BIN" ]; then
    echo "please verify bvchaind is installed"
    exit 1
fi

# configure bvchaind
$BVCHAIND_BIN config set client chain-id demo
$BVCHAIND_BIN config set client keyring-backend test
$BVCHAIND_BIN keys add alice
$BVCHAIND_BIN keys add bob
$BVCHAIND_BIN init test --chain-id demo --default-denom bvchain
# update genesis
$BVCHAIND_BIN genesis add-genesis-account alice 10000000bvchain --keyring-backend test
$BVCHAIND_BIN genesis add-genesis-account bob 1000bvchain --keyring-backend test
# create default validator
$BVCHAIND_BIN genesis gentx alice 1000000bvchain --chain-id demo
$BVCHAIND_BIN genesis collect-gentxs
