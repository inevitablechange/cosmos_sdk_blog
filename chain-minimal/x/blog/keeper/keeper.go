package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"blockchainvalley.io/bvchain/x/blog"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	Schema    collections.Schema
	BlogPosts collections.Map[string, blog.BlogPost]
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authority string) Keeper {
    //블록체인 저장소 위에다가 여러 컬렉션 (Map, Item)을 올리기 위한 빌더
    sb := collections.NewSchemaBuilder(storeService)


    k := Keeper{
        cdc:          cdc,
        addressCodec: addressCodec,
        BlogPosts: collections.NewMap(
            sb, //방금 만든 스키마 빌더
            blog.BlogPostsKey, // prefix ("BlogPosts/value/")
            "blog_posts", // 이 컬렉션의 이름(디버깅, 메타 용),
            collections.StringKey, // key 타입이 string 이라는 정보
            codec.CollValue[blog.BlogPost](cdc), // value 타입이 blog.BlogPost 라는 정보
        ),
    }

    schema, err := sb.Build()
    if err != nil {
        panic(err)
    }

    k.Schema = schema

    return k
}

// CreatePost 글 생성 로직
func (k Keeper) CreatePost(
    ctx context.Context,
    msg *blog.MsgCreatePost,
) error {
    // ID 길이 검증
	if length := len([]byte(msg.Id)); length < 1 || blog.MaxIDLength < length {
		return fmt.Errorf("id too long or empty")
	}

    // 같은 id가 이미 있는지 확인
    if _, err := k.BlogPosts.Get(ctx, msg.Id); err == nil {
        return fmt.Errorf("post with id %s already exists", msg.Id)
    }

    //BlogPost 생성해서 SET
    post := blog.BlogPost{
        Id: msg.Id,
        Title: msg.Title,
        Contents: msg.Contents,
        Creator: msg.Creator,
    }

    return k.BlogPosts.Set(ctx, msg.Id, post)
}

// 글 조회 - Query 핸들러에서 사용
func (k Keeper) GetPost(
    ctx context.Context,
    id string,
) (blog.BlogPost, error) {
    return k.BlogPosts.Get(ctx, id)
}