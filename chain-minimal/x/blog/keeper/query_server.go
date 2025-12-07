package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"blockchainvalley.io/bvchain/x/blog"
	"cosmossdk.io/collections"
)

type queryServer struct {
	k Keeper
}

var _ blog.QueryServer = queryServer{}

func NewQueryServerImpl(k Keeper) blog.QueryServer {
	return queryServer{k: k}
}

func (qs queryServer) GetPost(
	ctx context.Context,
	req *blog.QueryGetPostRequest,
) (*blog.QueryGetPostResponse, error) {
	post, err := qs.k.GetPost(ctx, req.Id)
	if err == collections.ErrNotFound {
		return nil, status.Error(codes.NotFound, "post not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &blog.QueryGetPostResponse{Post: &post}, nil
}

