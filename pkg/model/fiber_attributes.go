package model

// FiberAttribute as defined in
// https://www.ravelry.com/api#FiberAttribute__result
type FiberAttribute struct {
	ID               int    `json:"id"`
	FiberAttrGroupID int    `json:"fiber_attribute_group_id"`
	Name             string `json:"name"`
	Permalink        string `json:"permalink"`
}
