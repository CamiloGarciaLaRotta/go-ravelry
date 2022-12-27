package model

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyQuery    = errors.New("search query can't be empty")
	ErrNegativeLimit = errors.New("search limit can't be negative")
	ErrAboveLimitMax = fmt.Errorf("search limit is above max %d", SearchLimitMax)
)

// Search Type filters as defined in
// https://www.ravelry.com/api#/_search
const (
	SearchTypeUser          = "User"
	SearchTypePatternAuthor = "PatternAuthor"
	SearchTypePatternSource = "PatternSource"
	SearchTypePattern       = "Pattern"
	SearchTypeYarnCompany   = "YarnCompany"
	SearchTypeYarn          = "Yarn"
	SearchTypeGroup         = "Group"
	SearchTypeEvent         = "Event"
	SearchTypeProject       = "Project"
	SearchTypePage          = "Page"
	SearchTypeTopic         = "Topic"
	SearchTypeShop          = "Shop"
)

// Search limits as defined in
// https://www.ravelry.com/api#/_search
const (
	SearchLimitDefault = 50
	SearchLimitMax     = 500
)

// URL parameters used in different search endpoints.
const (
	SearchQueryParamQuery    = "query"
	SearchQueryParamLimit    = "limit"
	SearchQueryParamType     = "type"
	SearchQueryParamPage     = "page"
	SearchQueryParamPageSize = "page_size"
)

// SearchObject model as defined in
// https://www.ravelry.com/api#Object__result
type SearchObject struct {
	Title        string       `json:"title"`
	TypeName     string       `json:"type_name"`
	Caption      string       `json:"caption"`
	TinyImageURL string       `json:"tiny_image_url"`
	ImageURL     string       `json:"image_url"`
	Record       SearchRecord `json:"record"`
}

// SearchRecord model as defined in
// https://www.ravelry.com/api#Object__result
type SearchRecord struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Permalink string `json:"permalink"`
	URI       string `json:"uri"`
}

// SavedSearch model as defined in
// https://www.ravelry.com/api#SavedSearch_list_result
type SavedSearch struct {
	ID                    int               `json:"id"`
	Subscribed            bool              `json:"subscribed"`
	Title                 string            `json:"title"`
	SearchPath            string            `json:"search_path"`
	SearchType            string            `json:"search_type"`
	LastLoaded            string            `json:"last_loaded"`
	SubscriptionUpdatedAt string            `json:"subscription_updated_at"`
	CreatedAt             string            `json:"created_at"`
	UpdatedAt             string            `json:"updated_at"`
	SearchParams          map[string]string `json:"search_parameters"`
}
