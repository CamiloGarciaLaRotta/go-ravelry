package model

// YarnAttributeParent as defined in
// https://www.ravelry.com/api#AttributeGroup_list_result.
// Note that this is a recursive type.
type YarnAttributeParent struct {
	ID             int                   `json:"id"`
	Name           string                `json:"name"`
	Permalink      string                `json:"permalink"`
	YarnAttributes []YarnAttributeNode   `json:"yarn_attributes"`
	Children       []YarnAttributeParent `json:"children"`
}

// YarnAttributeNode as defined in
// https://www.ravelry.com/api#YarnAttributeGroup_result
type YarnAttributeNode struct {
	ID                   int    `json:"id"`
	SortOrder            int    `json:"sort_order"`
	YarnAttributeGroupID int    `json:"yarn_attribute_group_id"`
	Name                 string `json:"name"`
	Permalink            string `json:"permalink"`
	Description          string `json:"description"`
	CreatedAt            string `json:"created_at"`
}
