package model

type (
	GravatarProfiles struct {
		Entry []profile `json:"entry"`
	}
	photo struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	}
	profile struct {
		Id           string   `json:"id"`
		Hash         string   `json:"hash"`
		RequestHash  string   `json:"requestHash"`
		ProfileUrl   string   `json:"profileUrl"`
		ThumbnailUrl string   `json:"thumbnailUrl"`
		Photos       []photo  `json:"photos"`
		Name         []string `json:"name"`
		DisplayName  string   `json:"displayName"`
		Urls         []string `json:"urls"`
	}
)
