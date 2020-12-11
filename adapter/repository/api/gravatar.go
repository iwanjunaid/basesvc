package api

import (
	"context"
	"crypto/md5"
	"fmt"
	"net/url"
	"strconv"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/internal/telemetry"
	"github.com/iwanjunaid/basesvc/usecase/gravatar/repository"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"

	"github.com/go-resty/resty/v2"
)

const BASEURL string = "https://www.gravatar.com"

type (
	GravatarRepositoryImpl struct {
		email        string
		hash         string
		defaultURL   string
		defaultValue string
		size         int
		forceDefault bool
		rating       string
		context      context.Context
	}
)

func NewGravatarRepository(ctx context.Context, email string) repository.GravatarRepository {
	bEmail := []byte(email)
	hash := md5.Sum(bEmail)

	return &GravatarRepositoryImpl{
		hash:    fmt.Sprintf("%x", hash),
		context: ctx,
	}
}

// URL return profile url
func (g *GravatarRepositoryImpl) URL() (string, error) {
	baseURL, err := url.Parse((BASEURL))
	if err != nil {
		return "", err
	}
	baseURL.Path += g.hash
	return baseURL.String(), nil
}

// JSONURL return profile url in json
func (g *GravatarRepositoryImpl) JSONURL() (string, error) {
	baseURL, err := url.Parse((BASEURL))
	if err != nil {
		return "", err
	}
	baseURL.Path += fmt.Sprintf("%s.json", g.hash)

	return baseURL.String(), nil
}

// AvatarURL return url of avatar
func (g *GravatarRepositoryImpl) AvatarURL() (avatar string, err error) {
	baseURL, err := url.Parse(BASEURL)
	if err != nil {
		return
	}

	// Add path segment avatar
	baseURL.Path += "avatar/"

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
func (g *GravatarRepositoryImpl) GetProfile() (res *model.GravatarProfiles, err error) {
	txn := telemetry.GetTelemetry(g.context)
	defer txn.End()

	client := resty.New()

	// Because Gravatar API use redirect method,
	// We have to set resty redirect policy.
	// Assign Client Redirect Policy. Create one as per you need
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(3))

	// Registering Request Middleware
	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		// Now you have access to Client and current Request object
		// manipulate it as per your need

		// Newrelic roundtripper
		c.SetTransport(newrelic.NewRoundTripper(c.GetClient().Transport))

		// Put transaction in the request's context:
		ctx := req.Context()
		ctx = newrelic.NewContext(ctx, txn)
		req.SetContext(ctx)

		return nil // if its success otherwise return error
	})

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
