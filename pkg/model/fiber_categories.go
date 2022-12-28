package model

// FiberCategory as defined in
// https://www.ravelry.com/api#FiberCategory_list_result
type FiberCategory struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Permalink string          `json:"permalink"`
	ShortName string          `json:"short_name"`
	Children  []FiberCategory `json:"children"`
}
