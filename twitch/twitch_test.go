package twitch

import (
	"errors"
	"net/http"
	"reflect"
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

func TestErrorToOutput(t *testing.T) {

	err := errors.New("Test error")

	output := errorToOutput(err)
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
		t.Errorf("GetRootWithError error.Status was not 400: %d", errorOutput.Error)
	}
	if errorOutput.Message != "Some error message" {
		t.Errorf("GetRootWithError error.Message was not \"Some error message\": %s", errorOutput.Message)
	}
}

func TestGetBlocks(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/users/1/blocks",
		httpmock.NewStringResponder(200, `{"blocks":[{"updated_at":"2013-02-07T01:04:43Z","user":{"updated_at":"2013-02-06T22:44:19Z","display_name":"test_user_troll","type":"user","bio":"I'm a troll.. Kappa","name":"test_user_troll","_id":13460644,"logo":"http://something.net/foo.png","created_at":"2010-06-30T08:26:49Z"},"_id":970887}]}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.GetBlocks(&BlocksInput{
		UserID: 1,
		Limit:  25,
		Offset: 0,
	})

	if errorOutput != nil {
		t.Errorf("GetBlocks errorOutput should have been nil: %+v", errorOutput)
	}

	if len(output.Blocks) != 1 {
		t.Errorf("GetBlocks the blocks list was not 1 in length: %d", len(output.Blocks))
	}

	block := output.Blocks[0]

	if block.ID != 970887 {
		t.Errorf("GetBlocks the block id was not 970887: %d", block.ID)
	}
	if block.User.ID != 13460644 {
		t.Errorf("GetBlocks the block user id was not 13460644: %d", block.User.ID)
	}
	if block.User.DisplayName != "test_user_troll" {
		t.Errorf("GetBlocks the block user disply name was not test_user_troll: %d", block.User.DisplayName)
	}
	if block.User.Type != "user" {
		t.Errorf("GetBlocks the block user type was not user: %d", block.User.Type)
	}
	if block.User.Bio != "I'm a troll.. Kappa" {
		t.Errorf("GetBlocks the block user bio was not \"I'm a troll.. Kappa\": %d", block.User.Bio)
	}
	if block.User.Logo != "http://something.net/foo.png" {
		t.Errorf("GetBlocks the block user logo was not http://something.net/foo.png: %d", block.User.Logo)
	}
}

func TestBlockUser(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.twitch.tv/kraken/users/1/blocks/2",
		httpmock.NewStringResponder(200, `{"updated_at":"2013-02-07T01:04:43Z","user":{"updated_at":"2013-01-18T22:33:55Z","logo":"http://something.net/foo.png","type":"user","bio":"I'm a troll.. Kappa","display_name":"test_user_troll","name":"test_user_troll","_id":22125774,"created_at":"2011-05-01T14:50:12Z"},"_id":287813}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.BlockUser(&BlockUserInput{
		UserID:       1,
		TargetUserID: 2,
	})

	if errorOutput != nil {
		t.Errorf("GetBlocks errorOutput should have been nil: %+v", errorOutput)
	}

	if output.ID != 287813 {
		t.Errorf("GetBlocks the block id was not 287813: %d", output.ID)
	}
	if output.User.ID != 22125774 {
		t.Errorf("GetBlocks the block user id was not 22125774: %d", output.User.ID)
	}
	if output.User.DisplayName != "test_user_troll" {
		t.Errorf("GetBlocks the block user disply name was not test_user_troll: %d", output.User.DisplayName)
	}
	if output.User.Type != "user" {
		t.Errorf("GetBlocks the block user type was not user: %d", output.User.Type)
	}
	if output.User.Bio != "I'm a troll.. Kappa" {
		t.Errorf("GetBlocks the block user bio was not \"I'm a troll.. Kappa\": %d", output.User.Bio)
	}
	if output.User.Logo != "http://something.net/foo.png" {
		t.Errorf("GetBlocks the block user logo was not http://something.net/foo.png: %d", output.User.Logo)
	}
}
