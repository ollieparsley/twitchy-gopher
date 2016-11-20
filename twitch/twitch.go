package twitch

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
}

type Client struct {
	apiURL      string
	version     int
	httpClient  *http.Client
	oauthConfig *OAuthConfig
}

func NewClient(oauthConfig *OAuthConfig, httpClient *http.Client) *Client {
	apiURL := "https://api.twitch.tv/kracken/"

	return &Client{
		apiURL:      apiURL,
		version:     5,
		httpClient:  httpClient,
		oauthConfig: oauthConfig,
	}
}

func (c *Client) getSling() *sling.Sling {

	s := sling.New().Base(c.apiURL).Client(c.httpClient)

	// User-agent so twitch can track where this has come from
	s.Set("User-Agent", "Twitchy Gopher (https://github.com/ollieparsley/twitchy-gopher")

	// Specify the API version
	s.Set("Accept", fmt.Sprintf("application/vnd.twitchtv.v%d+json", c.version))

	// Add authentication
	s.Set("Authorization", "OAuth "+c.oauthConfig.AccessToken)

	return s
}
