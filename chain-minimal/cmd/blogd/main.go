package main

import (
	"fmt"
	"os"

	clienthelpers "cosmossdk.io/client/v2/helpers"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"blockchainvalley.io/bgchain/app"
	"blockchainvalley.io/bgchain/app/params"
	"blockchainvalley.io/bgchain/cmd/blogd/cmd"
)

func main() {
	params.SetAddressPrefixes()

	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, clienthelpers.EnvPrefix, app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
