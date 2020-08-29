package twitch

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestListBlocks(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/users/1/blocks",
		httpmock.NewStringResponder(200, `{"blocks":[{"updated_at":"2013-02-07T01:04:43Z","user":{"updated_at":"2013-02-06T22:44:19Z","display_name":"test_user_troll","type":"user","bio":"I'm a troll.. Kappa","name":"test_user_troll","_id":13460644,"logo":"http://something.net/foo.png","created_at":"2010-06-30T08:26:49Z"},"_id":970887}]}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.ListBlocks(&ListBlocksInput{
		UserID: 1,
		Limit:  25,
		Offset: 50,
	})

	if errorOutput != nil {
		t.Errorf("ListBlocks errorOutput should have been nil: %+v", errorOutput)
	}

	if len(output.Blocks) != 1 {
		t.Errorf("ListBlocks the blocks list was not 1 in length: %d", len(output.Blocks))
	}

	block := output.Blocks[0]

	if block.ID != 970887 {
		t.Errorf("ListBlocks the block id was not 970887: %d", block.ID)
	}
	if block.User.ID != 13460644 {
		t.Errorf("ListBlocks the block user id was not 13460644: %d", block.User.ID)
	}
	if block.User.DisplayName != "test_user_troll" {
		t.Errorf("ListBlocks the block user disply name was not test_user_troll: %s", block.User.DisplayName)
	}
	if block.User.Type != "user" {
		t.Errorf("ListBlocks the block user type was not user: %s", block.User.Type)
	}
	if block.User.Bio != "I'm a troll.. Kappa" {
		t.Errorf("ListBlocks the block user bio was not \"I'm a troll.. Kappa\": %s", block.User.Bio)
	}
	if block.User.Logo != "http://something.net/foo.png" {
		t.Errorf("ListBlocks the block user logo was not http://something.net/foo.png: %s", block.User.Logo)
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
		t.Errorf("BlockUser errorOutput should have been nil: %+v", errorOutput)
	}

	if output.ID != 287813 {
		t.Errorf("BlockUser the block id was not 287813: %d", output.ID)
	}
	if output.User.ID != 22125774 {
		t.Errorf("BlockUser the block user id was not 22125774: %d", output.User.ID)
	}
	if output.User.DisplayName != "test_user_troll" {
		t.Errorf("BlockUser the block user disply name was not test_user_troll: %s", output.User.DisplayName)
	}
	if output.User.Type != "user" {
		t.Errorf("BlockUser the block user type was not user: %s", output.User.Type)
	}
	if output.User.Bio != "I'm a troll.. Kappa" {
		t.Errorf("BlockUser the block user bio was not \"I'm a troll.. Kappa\": %s", output.User.Bio)
	}
	if output.User.Logo != "http://something.net/foo.png" {
		t.Errorf("BlockUser the block user logo was not http://something.net/foo.png: %s", output.User.Logo)
	}
}

func TestUnblockUser(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "https://api.twitch.tv/kraken/users/1/blocks/2", httpmock.NewStringResponder(204, ``))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.UnblockUser(&UnblockUserInput{
		UserID:       1,
		TargetUserID: 2,
	})

	if errorOutput != nil {
		t.Errorf("UnblockUser errorOutput should have been nil: %+v", errorOutput)
	}
	if output == nil {
		t.Errorf("UnblockUser the output was nil")
	}
}
