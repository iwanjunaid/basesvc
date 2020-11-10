package gravatar

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strconv"

	"gopkg.in/resty.v1"
)

const BASEURL string = "https://www.gravatar.com"

type (
	Gravatar struct {
		email        string
		hash         string
		defaultURL   string
		defaultValue string
		size         int
		forceDefault bool
		rating       string
	}
	photo struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	}
	profiles struct {
		Entry []profile `json:"entry"`
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

func New(email string) *Gravatar {
	bEmail := []byte(email)
	hash := md5.Sum(bEmail)

	return &Gravatar{
		email:        email,
		hash:         fmt.Sprintf("%x", hash),
		defaultURL:   "",
		defaultValue: "",
		size:         0,
		forceDefault: false,
		rating:       "",
	}
}

// URL return profile url
func (g *Gravatar) URL() (string, error) {
	baseURL, err := url.Parse((BASEURL))
	if err != nil {
		return "", err
	}
	baseURL.Path += g.hash
	return baseURL.String(), nil
}

// JSONURL return profile url in json
func (g *Gravatar) JSONURL() (string, error) {
	baseURL, err := url.Parse((BASEURL))
	if err != nil {
		return "", err
	}
	baseURL.Path += fmt.Sprintf("%s.json", g.hash)

	return baseURL.String(), nil
}

// AvatarURL return url of avatar
func (g *Gravatar) AvatarURL() (avatar string, err error) {
	baseURL, err := url.Parse(BASEURL)
	if err != nil {
		return
	}

	// Add path segment avatar
	baseURL.Path += "avatar"

	// Add path segment hash email
	baseURL.Path += g.hash

	// Prepare Query params
	params := url.Values{}

	if g.forceDefault {
		params.Add("f", "y")
	}

	if g.defaultURL != "" {
		params.Add("d", g.defaultURL)
	} else if g.defaultValue != "" {
		params.Add("d", g.defaultValue)
	}

	if g.rating != "" {
		params.Add("r", g.rating)
	}

	if g.size > 0 {
		params.Add("s", strconv.Itoa(g.size))
	}

	// Add Query params to the URL
	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}

// GetProfile return Gravatar profile struct
func (g *Gravatar) GetProfile() (res *profiles, err error) {
	client := resty.New()

	// Because Gravatar API use redirect method,
	// We have to set resty redirect policy.
	// Assign Client Redirect Policy. Create one as per you need
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(3))

	url, err := g.JSONURL()
	if err != nil {
		return
	}

	res = &profiles{}
	resp, err := client.R().
		SetResult(res).
		Get(url)

	if err != nil {
		return
	}

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())

	return res, err
}
