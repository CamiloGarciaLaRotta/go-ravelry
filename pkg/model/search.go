package model

// title	Title of the object
// type_name	Type of object, suitable for displaying to end users
// caption	(optional) Caption/description related to object, suitable for displaying to end users
// tiny_image_url	(optional) Image suitable for displaying inline with text. Typically 24x24.
// image_url	(optional) Primary image associated with the object. Typically 500px on the longest side.
// record	Nested object with information about the matching record
// Key	Description
// type	Type of record, corresponds with API Result Types
// id	id of record
// permalink	current permalink for record
// uri	API URI for retrieving information about the record

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

const (
	SearchQueryParamQuery = "query"
	SearchQueryParamLimit = "limit"
	SearchQueryParamType  = "type"
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