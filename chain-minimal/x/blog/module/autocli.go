package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	blogv1 "blockchainvalley.io/bvchain/x/blog/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		// Query 쪽 CLI
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: blogv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetPost",
					Use:       "get-post [id]",
					Short:     "Get a blog post by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
			},
		},

		// Tx 쪽 CLI
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: blogv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreatePost",
					Use:       "create-post [id] [title] [contents]",
					Short:     "Create a new blog post with given id, title, and contents",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
						{ProtoField: "title"},
						{ProtoField: "contents"},
					},
				},
			},
		},
	}
}
