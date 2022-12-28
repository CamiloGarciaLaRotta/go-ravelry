package model

import "errors"

var ErrNoUserID = errors.New("user ID can't be empty")

// User model as defined in
// https://www.ravelry.com/api#/_current_user
type User struct {
	ID            int           `json:"id"`
	Username      string        `json:"username"`
	AboutMe       string        `json:"about_me"`
	AboutMeHTML   string        `json:"about_me_html"`
	FaveColors    string        `json:"fave_colors"`
	FaveCurse     string        `json:"fave_curse"`
	FirstName     string        `json:"first_name"`
	Location      string        `json:"location"`
	LargePhotoURL string        `json:"large_photo_url"`
	PhotoURL      string        `json:"photo_url"`
	SmallPhotoURL string        `json:"small_photo_url"`
	TinyPhotoURL  string        `json:"tiny_photo_url"`
	UserSites     []UserSite    `json:"user_sites"`
	PatternAuthor PatternAuthor `json:"pattern_author"`
}

// UserSite model as defined in
// https://www.ravelry.com/api#UserSite_full_result
type UserSite struct {
	ID         int        `json:"id"`
	URL        string     `json:"url"`
	Username   string     `json:"username"`
	SocialSite SocialSite `json:"social_site"`
}

// SocialSite model as defined in
// https://www.ravelry.com/api#SocialSite__result
type SocialSite struct {
	ID         int    `json:"id"`
	Active     bool   `json:"active"`
	FaviconURL string `json:"favicon_url"`
	Name       string `json:"favinamecon_url"`
}

// PatternAuthor model as defined in
// https://www.ravelry.com/api#PatternAuthor_for_user_result
type PatternAuthor struct {
	ID             int    `json:"id"`
	FavoritesCount int    `json:"favorites_count"`
	PatternsCount  int    `json:"patterns_count"`
	Name           string `json:"name"`
	Permalink      string `json:"permalink"`
}
