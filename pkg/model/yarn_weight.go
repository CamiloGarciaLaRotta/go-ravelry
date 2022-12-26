package model

// YarnWeight model as defined in
// https://www.ravelry.com/api#YarnWeight__result
type YarnWeight struct {
	ID           int    `json:"id"`
	CrochetGauge string `json:"crochet_gauge"`
	KnitGauge    string `json:"knit_gauge"`
	MaxGauge     string `json:"max_gauge"`
	MinGauge     string `json:"min_gauge"`
	Name         string `json:"name"`
	Ply          string `json:"ply"`
	Wpi          string `json:"wpi"`
}
