package twitch

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	oauthConfig := &OAuthConfig{
		ClientID:     "client-id",
		ClientSecret: "client-secret",
		AccessToken:  "access-token",
	}

	httpClient := &http.Client{}

	client := NewClient(oauthConfig, httpClient)

	if client.apiURL != "https://api.twitch.tv/kracken/" {
		t.Errorf("client.apiURL was not correct: %s", client.apiURL)
	}
	if client.version != 5 {
		t.Errorf("client.version was not 5: %d", client.version)
	}
	if client.httpClient != httpClient {
		t.Errorf("client.httpClient was not correct: %+v", client.httpClient)
	}
	if client.oauthConfig != oauthConfig {
		t.Errorf("client.oauthConfig was not correct: %+v", client.oauthConfig)
	}
}

func TestGetSling(t *testing.T) {
	oauthConfig := &OAuthConfig{
		ClientID:     "client-id",
		ClientSecret: "client-secret",
		AccessToken:  "access-token",
	}

	httpClient := &http.Client{}

	client := NewClient(oauthConfig, httpClient)

	// get sling and get the raw http request
	request, err := client.getSling().Request()
	if err != nil {
		t.Errorf("error with sling http request: %s", err.Error())
	}
	if request.URL.Scheme != "https" {
		t.Errorf("request scheme was not https: %s", request.URL.Scheme)
	}
	if request.URL.Host != "api.twitch.tv" {
		t.Errorf("request scheme was not api.twitch.tv: %s", request.URL.Host)
	}
	if request.URL.Path != "/kracken/" {
		t.Errorf("request scheme was not /kracken/: %s", request.URL.Path)
	}
	if request.Header.Get("Authorization") != "OAuth access-token" {
		t.Errorf("request Authorization header was not was not \"OAuth access-token\": %s", request.Header.Get("Authorization"))
	}
	if request.Header.Get("Accept") != "application/vnd.twitchtv.v5+json" {
		t.Errorf("request Accept header was not was not \"application/vnd.twitchtv.v5+json\": %s", request.Header.Get("Accept"))
	}
}
