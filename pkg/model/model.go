package model

// ColorFamily model as defined in
// https://www.ravelry.com/api#/_color_families
type ColorFamily struct {
	ID            int    `json:"id"`
	SpectrumOrder int    `json:"spectrum_order"`
	Color         string `json:"color"`
	Name          string `json:"name"`
	Permalink     string `json:"permalink"`
}
