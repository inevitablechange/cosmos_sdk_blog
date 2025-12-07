package blog

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterInterfaces registers the module's concrete types on the codec.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	// 1) Msg 타입을 sdk.Msg 구현체로 등록
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreatePost{},
	)

	// 2) Msg 서비스(_Msg_serviceDesc) 등록
	//    → 이게 없으면 지금 같은 panic이 난다
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
