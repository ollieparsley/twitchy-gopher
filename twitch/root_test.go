package twitch

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetRoot(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/",
		httpmock.NewStringResponder(200, `{"token":{"valid":true,"authorization":{"scopes":["channel_read"],"created_at":"2016-11-20T17:26:13Z","updated_at":"2016-11-20T17:26:13Z"},"user_name":"testing_user","client_id":"testing-client-id"}}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.GetRoot()

	if errorOutput != nil {
		t.Errorf("GetRoot errorOutput should have been null: %+v", errorOutput)
	}

	if output.Token.Valid != true {
		t.Errorf("GetRoot token.valid was not true")
	}
	if !reflect.DeepEqual(output.Token.Authorization.Scopes, []string{"channel_read"}) {
		t.Errorf("GetRoot token.authorization.scopes was not correct %+v", output.Token.Authorization.Scopes)
	}
	if output.Token.Username != "testing_user" {
		t.Errorf("GetRoot token.valid was not \"testing_user\": %s", output.Token.Username)
	}
	if output.Token.ClientID != "testing-client-id" {
		t.Errorf("GetRoot token.clientid was not \"testing-client-id\": %s", output.Token.ClientID)
	}
}

func TestGetRootWithError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/",
		httpmock.NewStringResponder(400, `{"error":"Bad Request","status":400,"message":"Some error message"}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	_, errorOutput := client.GetRoot()

	if errorOutput == nil {
		t.Errorf("GetRootWithError errorOutput shouldn't have been nil")
	}

	if errorOutput.Error != "Bad Request" {
		t.Errorf("GetRootWithError error.Error was not \"Bad Request\": %s", errorOutput.Error)
	}
	if errorOutput.Status != 400 {
		t.Errorf("GetRootWithError error.Status was not 400: %s", errorOutput.Error)
	}
	if errorOutput.Message != "Some error message" {
		t.Errorf("GetRootWithError error.Message was not \"Some error message\": %s", errorOutput.Message)
	}
}
