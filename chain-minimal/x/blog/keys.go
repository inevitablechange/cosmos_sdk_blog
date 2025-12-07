package blog

import "cosmossdk.io/collections"

const ModuleName = "blog"

var (
	ParamsKey      = collections.NewPrefix("Params")           // 나중에 쓸 수도 있음
	BlogPostsKey   = collections.NewPrefix("BlogPosts/value/") // 모듈에서 글들을 저장할 때 'BlogPosts/value/{id}' 이런 식의 key로 저장되게끔 해줌
	MaxIDLength    = 256
)

