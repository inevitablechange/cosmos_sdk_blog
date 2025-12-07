package blog

import "fmt"

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
    return &GenesisState{
        Params: DefaultParams(),
        IndexedBlogPostList: []IndexedBlogPost{},
    }
}

// Validate performs basic genesis state validation.
func (gs *GenesisState) Validate() error {
    if err := gs.Params.Validate(); err != nil {
		return err
	}
    seen := make(map[string]bool)

    for _, indexed := range gs.IndexedBlogPostList {
        // 1) id(=index)가 비어 있으면 안 됨
        if len(indexed.Index) == 0 {
            return fmt.Errorf("index cannot be empty")
        }

        // 2) 중복 index 체크
        if seen[indexed.Index] {
            return fmt.Errorf("duplicate index: %s", indexed.Index)
        }
        seen[indexed.Index] = true

        // 3) BlogPost 안의 필드도 필요하면 검증 가능
        if len(indexed.BlogPost.Title) == 0 {
            return fmt.Errorf("title cannot be empty for index: %s", indexed.Index)
        }
        // contents, creator 등도 규칙을 추가하고 싶으면 여기에
    }

    return nil
}