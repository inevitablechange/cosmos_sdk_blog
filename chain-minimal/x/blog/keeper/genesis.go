package keeper

import (
	"context"

	"blockchainvalley.io/bvchain/x/blog"
	"cosmossdk.io/collections"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *blog.GenesisState) error {
	// GenesisState에 들어 있는 모든 글을 state(BlogPosts map)에 세팅
	for _, indexed := range data.IndexedBlogPostList {
		if err := k.BlogPosts.Set(ctx, indexed.Index, indexed.BlogPost); err != nil {
			return err
		}
	}
	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*blog.GenesisState, error) {
	var indexedPosts []blog.IndexedBlogPost

	// state에 있는 모든 글을 돌면서 GenesisState 형식으로 변환
	if err := k.BlogPosts.Walk(ctx, nil, func(id string, post blog.BlogPost) (bool, error) {
		indexedPosts = append(indexedPosts, blog.IndexedBlogPost{
			Index:    id,
			BlogPost: post,
		})
		return false, nil // false = 계속 순회
	}); err != nil {
		if err != collections.ErrNotFound {
			return nil, err
		}
	}

	return &blog.GenesisState{
		IndexedBlogPostList: indexedPosts,
	}, nil
}
