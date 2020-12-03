package api

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strconv"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/repository"
	"gopkg.in/resty.v1"
)

const BASEURL string = "https://www.gravatar.com"

type (
	AuthorGravatarRepositoryImpl struct {
		email        string
		hash         string
		defaultURL   string
		defaultValue string
		size         int
		forceDefault bool
		rating       string
	}
)

func NewAuthorGravatar(email string) repository.AuthorGravatarRepository {
	bEmail := []byte(email)
	hash := md5.Sum(bEmail)

	return &AuthorGravatarRepositoryImpl{
		hash: fmt.Sprintf("%x", hash),
	}
}

// URL return profile url
func (g *AuthorGravatarRepositoryImpl) URL() (string, error) {
	baseURL, err := url.Parse((BASEURL))
	if err != nil {
		return "", err
	}
	baseURL.Path += g.hash
	return baseURL.String(), nil
}

// JSONURL return profile url in json
func (g *AuthorGravatarRepositoryImpl) JSONURL() (string, error) {
	baseURL, err := url.Parse((BASEURL))
	if err != nil {
		return "", err
	}
	baseURL.Path += fmt.Sprintf("%s.json", g.hash)

	return baseURL.String(), nil
}

// AvatarURL return url of avatar
func (g *AuthorGravatarRepositoryImpl) AvatarURL() (avatar string, err error) {
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
func (g *AuthorGravatarRepositoryImpl) GetProfile() (res *model.GravatarProfiles, err error) {
	client := resty.New()

	// Because Gravatar API use redirect method,
	// We have to set resty redirect policy.
	// Assign Client Redirect Policy. Create one as per you need
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(3))

	url, err := g.JSONURL()
	if err != nil {
		return
	}

	res = &model.GravatarProfiles{}
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
