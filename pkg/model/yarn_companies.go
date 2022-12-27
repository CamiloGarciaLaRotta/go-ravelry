package model

// YarnCompany as defined in
// https://www.ravelry.com/api#YarnCompany_list_result
type YarnCompany struct {
	ID         int    `json:"id"`
	YarnsCount int    `json:"yarns_count"`
	LogoURL    string `json:"logo_url"`
	Name       string `json:"name"`
	Permalink  string `json:"permalink"`
	URL        string `json:"url"`
}

// YarnCompanySearchResponse represents the expected API response
// when querying https://www.ravelry.com/api#yarn_companies_search
type YarnCompanySearchResponse struct {
	Companies []YarnCompany `json:"yarn_companies"`
	Paginator Paginator     `json:"paginator"`
}
