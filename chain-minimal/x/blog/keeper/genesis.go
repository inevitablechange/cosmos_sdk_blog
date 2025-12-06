package keeper

import (
	"context"

	"blockchainvalley.io/bvchain/x/blog"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *blog.GenesisState) error {
    if err := k.Params.Set(ctx, data.Params); err != nil {
        return err
    }

    return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*blog.GenesisState, error) {
    params, err := k.Params.Get(ctx)
    if err != nil {
        return nil, err
    }

    return &blog.GenesisState{
        Params: params,
    }, nil
}

