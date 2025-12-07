package keeper

import (
	"context"

	"blockchainvalley.io/bvchain/x/blog"
)

type msgServer struct {
	k Keeper
}

var _ blog.MsgServer = msgServer{}

func NewMsgServerImpl(k Keeper) blog.MsgServer {
	return msgServer{k: k}
}

func (ms msgServer) CreatePost(
	ctx context.Context,
	msg *blog.MsgCreatePost,
) (*blog.MsgCreatePostResponse, error) {
	if err := ms.k.CreatePost(ctx, msg); err != nil {
		return nil, err
	}
	return &blog.MsgCreatePostResponse{}, nil
}

