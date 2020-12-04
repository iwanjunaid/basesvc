package model

type (
	GravatarProfiles struct {
		Entry []Profile `json:"entry"`
	}
	Photo struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	}
	Profile struct {
		Id           string   `json:"id"`
		Hash         string   `json:"hash"`
		RequestHash  string   `json:"requestHash"`
		ProfileUrl   string   `json:"profileUrl"`
		ThumbnailUrl string   `json:"thumbnailUrl"`
		Photos       []Photo  `json:"photos"`
		Name         []string `json:"name"`
		DisplayName  string   `json:"displayName"`
		Urls         []string `json:"urls"`
	}
)
