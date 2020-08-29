package twitch

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestNewClient(t *testing.T) {
	oauthConfig := &OAuthConfig{
		ClientID:     "client-id",
		ClientSecret: "client-secret",
		AccessToken:  "access-token",
	}

	httpClient := &http.Client{}

	client := NewClient(oauthConfig, httpClient)

	if client.apiURL != "https://api.twitch.tv/kraken/" {
		t.Errorf("client.apiURL was not correct: %s", client.apiURL)
	}
	if client.apiVersion != 5 {
		t.Errorf("client.apiVersion was not 5: %d", client.apiVersion)
	}
	if client.uploadVersion != 4 {
		t.Errorf("client.uploadVersion was not 5: %d", client.uploadVersion)
	}
	if client.httpClient != httpClient {
		t.Errorf("client.httpClient was not correct: %+v", client.httpClient)
	}
	if client.oauthConfig != oauthConfig {
		t.Errorf("client.oauthConfig was not correct: %+v", client.oauthConfig)
	}
}

func TestErrorToOutput(t *testing.T) {

	client := NewClient(&OAuthConfig{
		ClientID:    "client-id",
		AccessToken: "access-token",
	}, &http.Client{})

	err := errors.New("Test error")

	output := client.errorToOutput(err)
	if output.Message != "Test error" {
		t.Errorf("error output message was not \"Test error\": %s", output.Message)
	}
	if output.Error != "Twitchy error" {
		t.Errorf("error output error was not \"Test error\": %s", output.Error)
	}
	if output.Status != -1 {
		t.Errorf("error output status was not -1: %d", output.Status)
	}
}

func TestCreateAPIRequest(t *testing.T) {
	client := NewClient(&OAuthConfig{
		ClientID:    "client-id",
		AccessToken: "access-token",
	}, &http.Client{})

	req := client.createAPIRequest("GET", "foo/bar", nil)

	if req.URL.Scheme != "https" {
		t.Errorf("createAPIRequest scheme was not https: %s", req.URL.Scheme)
	}
	if req.URL.Host != "api.twitch.tv" {
		t.Errorf("createAPIRequest scheme was not api.twitch.tv: %s", req.URL.Host)
	}
	if req.URL.Path != "/kraken/foo/bar" {
		t.Errorf("createAPIRequest scheme was not /kraken/foo/bar: %s", req.URL.Path)
	}
	if req.Header.Get("Authorization") != "OAuth access-token" {
		t.Errorf("createAPIRequest Authorization header was not was not \"OAuth access-token\": %s", req.Header.Get("Authorization"))
	}
	if req.Header.Get("Client-ID") != "client-id" {
		t.Errorf("createAPIRequest Authorization header was not was not \"client-id\": %s", req.Header.Get("Client-ID"))
	}
	if req.Header.Get("Accept") != "application/vnd.twitchtv.v5+json" {
		t.Errorf("createAPIRequest Accept header was not was not \"application/vnd.twitchtv.v5+json\": %s", req.Header.Get("Accept"))
	}
}

func TestCreateAPIRequestPostRequest(t *testing.T) {
	client := NewClient(&OAuthConfig{
		ClientID:    "client-id",
		AccessToken: "access-token",
	}, &http.Client{})

	params := map[string]string{}
	params["foo"] = "bar"
	params["hello"] = "world"

	req := client.createAPIRequest("POST", "foo/bar", params)

	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	requestBody := buf.String()

	expectedRequestBody := "foo=bar&hello=world"

	if requestBody != expectedRequestBody {
		t.Errorf("createAPIRequestPostRequest body did not match %s : %s", expectedRequestBody, requestBody)
	}
}

func TestCreateAPIRequestGetRequest(t *testing.T) {
	client := NewClient(&OAuthConfig{
		ClientID:    "client-id",
		AccessToken: "access-token",
	}, &http.Client{})

	params := map[string]string{}
	params["foo"] = "bar"
	params["hello"] = "world"

	req := client.createAPIRequest("GET", "foo/bar", params)
	requestQuerystring := req.URL.RawQuery
	expectedRequestQuerystring := "foo=bar&hello=world"

	if requestQuerystring != expectedRequestQuerystring {
		t.Errorf("createAPIRequest querytring did not match %s : %s", expectedRequestQuerystring, requestQuerystring)
	}
}

func TestCreateUploadRequest(t *testing.T) {
	client := NewClient(&OAuthConfig{
		ClientID:    "client-id",
		AccessToken: "access-token",
	}, &http.Client{})

	req := client.createUploadRequest("GET", "foo/bar", "some/content-type", bytes.NewBufferString("foobar"))

	if req.URL.Scheme != "https" {
		t.Errorf("createUploadRequest scheme was not https: %s", req.URL.Scheme)
	}
	if req.URL.Host != "uploads.twitch.tv" {
		t.Errorf("createUploadRequest scheme was not api.twitch.tv: %s", req.URL.Host)
	}
	if req.URL.Path != "/foo/bar" {
		t.Errorf("createUploadRequest scheme was not /kraken/foo/bar: %s", req.URL.Path)
	}
	if req.Header.Get("Authorization") != "OAuth access-token" {
		t.Errorf("createUploadRequest Authorization header was not was not \"OAuth access-token\": %s", req.Header.Get("Authorization"))
	}
	if req.Header.Get("Client-ID") != "client-id" {
		t.Errorf("createUploadRequest Authorization header was not was not \"client-id\": %s", req.Header.Get("Client-ID"))
	}
	if req.Header.Get("Content-Type") != "some/content-type" {
		t.Errorf("createUploadRequest Content-Type header was not was not \"some/content-type\": %s", req.Header.Get("Content-Type"))
	}
	if req.Header.Get("Accept") != "application/vnd.twitchtv.v4+json" {
		t.Errorf("createUploadRequest Accept header was not was not \"application/vnd.twitchtv.v4+json\": %s", req.Header.Get("Accept"))
	}
}

func TestPerformRequestClientError(t *testing.T) {
	req := &http.Request{}
	client := NewClient(&OAuthConfig{}, &http.Client{})
	errorOutput := client.performRequest(req, new(RootOutput))

	if errorOutput == nil {
		t.Errorf("performRequestClientError errorOutput was nil")
	}
	if errorOutput.Status != -1 {
		t.Errorf("performRequestClientError error status was not -1: %d", errorOutput.Status)
	}
	if errorOutput.Error != "Twitchy error" {
		t.Errorf("performRequestClientError error status was \"Twitchy error\": %s", errorOutput.Error)
	}
	if errorOutput.Message != "Get \"\": http: nil Request.URL" {
		t.Errorf("performRequestClientError error message was \"http: nil Request.URL\": %s", errorOutput.Message)
	}
}

func TestPerformRequestJSONError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/",
		httpmock.NewStringResponder(200, `{"foo":{{]]"}}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	_, errorOutput := client.GetRoot()

	if errorOutput == nil {
		t.Errorf("TestPerformRequestJSONError errorOutput was nil")
	}
	if errorOutput.Status != -1 {
		t.Errorf("TestPerformRequestJSONError error status was not -1: %d", errorOutput.Status)
	}
	if errorOutput.Error != "Twitchy error" {
		t.Errorf("TestPerformRequestJSONError error status was \"Twitchy error\": %s", errorOutput.Error)
	}
	if errorOutput.Message != "invalid character '{' looking for beginning of object key string" {
		t.Errorf("TestPerformRequestJSONError error message was \"invalid character '{' looking for beginning of object key string\": %s", errorOutput.Message)
	}
}
