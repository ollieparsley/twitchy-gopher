package twitch

import (
	"bytes"
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

func TestCreateRequest(t *testing.T) {
	client := NewClient(&OAuthConfig{
		ClientID:    "client-id",
		AccessToken: "access-token",
	}, &http.Client{})

	req := client.createRequest("GET", "foo/bar", nil)

	if req.URL.Scheme != "https" {
		t.Errorf("createRequest scheme was not https: %s", req.URL.Scheme)
	}
	if req.URL.Host != "api.twitch.tv" {
		t.Errorf("createRequest scheme was not api.twitch.tv: %s", req.URL.Host)
	}
	if req.URL.Path != "/kraken/foo/bar" {
		t.Errorf("createRequest scheme was not /kraken/foo/bar: %s", req.URL.Path)
	}
	if req.Header.Get("Authorization") != "OAuth access-token" {
		t.Errorf("createRequest Authorization header was not was not \"OAuth access-token\": %s", req.Header.Get("Authorization"))
	}
	if req.Header.Get("Client-ID") != "client-id" {
		t.Errorf("createRequest Authorization header was not was not \"client-id\": %s", req.Header.Get("Client-ID"))
	}
	if req.Header.Get("Accept") != "application/vnd.twitchtv.v5+json" {
		t.Errorf("createRequest Accept header was not was not \"application/vnd.twitchtv.v5+json\": %s", req.Header.Get("Accept"))
	}
}

func TestCreateRequestPostRequest(t *testing.T) {
	client := NewClient(&OAuthConfig{
		ClientID:    "client-id",
		AccessToken: "access-token",
	}, &http.Client{})

	params := map[string]string{}
	params["foo"] = "bar"
	params["hello"] = "world"

	req := client.createRequest("POST", "foo/bar", params)

	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	requestBody := buf.String()

	expectedRequestBody := "foo=bar&hello=world"

	if requestBody != expectedRequestBody {
		t.Errorf("createRequestPostRequest body did not match %s : %s", expectedRequestBody, requestBody)
	}
}

func TestCreateRequestGetRequest(t *testing.T) {
	client := NewClient(&OAuthConfig{
		ClientID:    "client-id",
		AccessToken: "access-token",
	}, &http.Client{})

	params := map[string]string{}
	params["foo"] = "bar"
	params["hello"] = "world"

	req := client.createRequest("GET", "foo/bar", params)
	requestQuerystring := req.URL.RawQuery
	expectedRequestQuerystring := "foo=bar&hello=world"

	if requestQuerystring != expectedRequestQuerystring {
		t.Errorf("createRequestGetRequest querytring did not match %s : %s", expectedRequestQuerystring, requestQuerystring)
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
	if errorOutput.Message != "http: nil Request.URL" {
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

func TestListChannelFeedPosts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/feed/12345/posts",
		httpmock.NewStringResponder(200, `{"_total":8,"_cursor":"1454101643075611000","posts":[{"id":"20","created_at":"2016-01-29T21:07:23.075611Z","deleted":false,"emotes":[],"reactions":{"endorse":{"count":2,"user_ids":[]}},"body":"Test post","user":{"display_name":"bangbangalang","_id":104447238,"name":"bangbangalang","type":"user","bio":"i like turtles and cats","created_at":"2015-10-15T19:52:17Z","updated_at":"2016-01-29T21:06:42Z","logo":"http://something.net/foo.png"}}]}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.ListChannelFeedPosts(&ListChannelFeedPostsInput{
		ChannelID: 12345,
		Limit:     25,
		Cursor:    "0123456789",
	})

	if errorOutput != nil {
		t.Errorf("ListChannelFeedPosts errorOutput should have been nil: %+v", errorOutput)
	}

	if output.Total != 8 {
		t.Errorf("ListChannelFeedPosts the total was not 8: %d", output.Total)
	}

	if output.Cursor != "1454101643075611000" {
		t.Errorf("ListChannelFeedPosts the cursor was not \"1454101643075611000\": %s", output.Cursor)
	}

	if len(output.Posts) != 1 {
		t.Errorf("ListChannelFeedPosts the posts list was not 1 in length: %d", len(output.Posts))
	}

	post := output.Posts[0]

	if post.ID != "20" {
		t.Errorf("ListChannelFeedPosts the post id was not \"20\": %s", post.ID)
	}
	if post.Deleted != false {
		t.Errorf("ListChannelFeedPosts the post deleted was not false: %t", post.Deleted)
	}
	if post.Body != "Test post" {
		t.Errorf("ListChannelFeedPosts the post body was not \"Test post\": %s", post.Body)
	}

	if post.User.ID != 104447238 {
		t.Errorf("ListChannelFeedPosts the block user id was not 104447238: %d", post.User.ID)
	}
	if post.User.DisplayName != "bangbangalang" {
		t.Errorf("ListChannelFeedPosts the block user disply name was not bangbangalang: %s", post.User.DisplayName)
	}
	if post.User.Bio != "i like turtles and cats" {
		t.Errorf("ListChannelFeedPosts the block user bio was not \"i like turtles and cats\": %s", post.User.Bio)
	}
	if post.User.Logo != "http://something.net/foo.png" {
		t.Errorf("ListChannelFeedPosts the block user logo was not http://something.net/foo.png: %s", post.User.Logo)
	}
}

func TestGetChannelFeedPost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/feed/12345/posts/20",
		httpmock.NewStringResponder(200, `{"id":"20","created_at":"2016-01-29T21:07:23.075611Z","deleted":false,"emotes":[],"reactions":{"endorse":{"count":2,"user_ids":[]}},"body":"Test post","user":{"display_name":"bangbangalang","_id":104447238,"name":"bangbangalang","type":"user","bio":"i like turtles and cats","created_at":"2015-10-15T19:52:17Z","updated_at":"2016-01-29T21:06:42Z","logo":"http://something.net/foo.png"}}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.GetChannelFeedPost(&GetChannelFeedPostInput{
		ChannelID: 12345,
		PostID:    "20",
	})

	if errorOutput != nil {
		t.Errorf("GetChannelFeedPost errorOutput should have been nil: %+v", errorOutput)
	}

	if output.ID != "20" {
		t.Errorf("GetChannelFeedPost the post id was not \"20\": %s", output.ID)
	}
	if output.Deleted != false {
		t.Errorf("GetChannelFeedPost the post deleted was not false: %t", output.Deleted)
	}
	if output.Body != "Test post" {
		t.Errorf("GetChannelFeedPost the post body was not \"Test post\": %s", output.Body)
	}
	if output.User.ID != 104447238 {
		t.Errorf("GetChannelFeedPost the block user id was not 104447238: %d", output.User.ID)
	}
	if output.User.DisplayName != "bangbangalang" {
		t.Errorf("GetChannelFeedPost the block user disply name was not bangbangalang: %s", output.User.DisplayName)
	}
	if output.User.Bio != "i like turtles and cats" {
		t.Errorf("GetChannelFeedPost the block user bio was not \"i like turtles and cats\": %s", output.User.Bio)
	}
	if output.User.Logo != "http://something.net/foo.png" {
		t.Errorf("GetChannelFeedPost the block user logo was not http://something.net/foo.png: %s", output.User.Logo)
	}
}

func TestCreateChannelFeedPost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.twitch.tv/kraken/feed/12345/posts",
		httpmock.NewStringResponder(200, `{"tweet":"http://twitter.com/blah/status/23469487ruo4","post":{"id":"20","created_at":"2016-01-29T21:07:23.075611Z","deleted":false,"emotes":[],"reactions":{"endorse":{"count":2,"user_ids":[]}},"body":"Test post","user":{"display_name":"bangbangalang","_id":104447238,"name":"bangbangalang","type":"user","bio":"i like turtles and cats","created_at":"2015-10-15T19:52:17Z","updated_at":"2016-01-29T21:06:42Z","logo":"http://something.net/foo.png"}}}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.CreateChannelFeedPost(&CreateChannelFeedPostInput{
		ChannelID: 12345,
		Content:   "This is a foo bar",
		Share:     true,
	})

	if errorOutput != nil {
		t.Errorf("CreateChannelFeedPost errorOutput should have been nil: %+v", errorOutput)
	}

	post := output.Post

	if output.Tweet != "http://twitter.com/blah/status/23469487ruo4" {
		t.Errorf("CreateChannelFeedPost the tweet was not \"http://twitter.com/blah/status/23469487ruo4\": %s", output.Tweet)
	}
	if post.ID != "20" {
		t.Errorf("CreateChannelFeedPost the post id was not \"20\": %s", post.ID)
	}
	if post.Deleted != false {
		t.Errorf("CreateChannelFeedPost the post deleted was not false: %t", post.Deleted)
	}
	if post.Body != "Test post" {
		t.Errorf("CreateChannelFeedPost the post body was not \"Test post\": %s", post.Body)
	}
	if post.User.ID != 104447238 {
		t.Errorf("CreateChannelFeedPost the block user id was not 104447238: %d", post.User.ID)
	}
	if post.User.DisplayName != "bangbangalang" {
		t.Errorf("CreateChannelFeedPost the block user disply name was not bangbangalang: %s", post.User.DisplayName)
	}
	if post.User.Bio != "i like turtles and cats" {
		t.Errorf("CreateChannelFeedPost the block user bio was not \"i like turtles and cats\": %s", post.User.Bio)
	}
	if post.User.Logo != "http://something.net/foo.png" {
		t.Errorf("CreateChannelFeedPost the block user logo was not http://something.net/foo.png: %s", post.User.Logo)
	}
}

func TestCreateChannelFeedPostNoShare(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.twitch.tv/kraken/feed/12345/posts",
		httpmock.NewStringResponder(200, `{"tweet":"http://twitter.com/blah/status/23469487ruo4","post":{"id":"20","created_at":"2016-01-29T21:07:23.075611Z","deleted":false,"emotes":[],"reactions":{"endorse":{"count":2,"user_ids":[]}},"body":"Test post","user":{"display_name":"bangbangalang","_id":104447238,"name":"bangbangalang","type":"user","bio":"i like turtles and cats","created_at":"2015-10-15T19:52:17Z","updated_at":"2016-01-29T21:06:42Z","logo":"http://something.net/foo.png"}}}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	_, errorOutput := client.CreateChannelFeedPost(&CreateChannelFeedPostInput{
		ChannelID: 12345,
		Content:   "This is a foo bar",
		Share:     false,
	})

	if errorOutput != nil {
		t.Errorf("CreateChannelFeedPostNoShare errorOutput should have been nil: %+v", errorOutput)
	}
}

func TestDeleteChannelFeedPost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "https://api.twitch.tv/kraken/feed/12345/posts/20",
		httpmock.NewStringResponder(204, ``))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.DeleteChannelFeedPost(&DeleteChannelFeedPostInput{
		ChannelID: 12345,
		PostID:    "20",
	})

	if errorOutput != nil {
		t.Errorf("DeleteChannelFeedPost errorOutput should have been nil: %+v", errorOutput)
	}

	if output == nil {
		t.Errorf("DeleteChannelFeedPost output shouldn't have been nil")
	}
}

func TestCreateChannelFeedPostReaction(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.twitch.tv/kraken/feed/12345/posts/54321/reactions",
		httpmock.NewStringResponder(200, `{"id":"20","created_at":"2016-01-29T21:07:23.075611Z","emote_id":"25","user":{"display_name":"bangbangalang","_id":104447238,"name":"bangbangalang","type":"user","bio":"i like turtles and cats","created_at":"2015-10-15T19:52:17Z","updated_at":"2016-01-29T21:06:42Z","logo":"http://something.net/foo.png"}}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.CreateChannelFeedPostReaction(&CreateChannelFeedPostReactionInput{
		ChannelID: 12345,
		PostID:    "54321",
		EmoteID:   "25",
	})

	if errorOutput != nil {
		t.Errorf("CreateChannelFeedPostReaction errorOutput should have been nil: %+v", errorOutput)
	}

	if output.ID != "20" {
		t.Errorf("CreateChannelFeedPostReaction the reaction id was not \"20\": %s", output.ID)
	}
	if output.EmoteID != "25" {
		t.Errorf("CreateChannelFeedPostReaction the reaction emote ID was not 25: %s", output.EmoteID)
	}
	if output.User.ID != 104447238 {
		t.Errorf("CreateChannelFeedPostReaction the block user id was not 104447238: %d", output.User.ID)
	}
	if output.User.DisplayName != "bangbangalang" {
		t.Errorf("CreateChannelFeedPostReaction the block user disply name was not bangbangalang: %s", output.User.DisplayName)
	}
	if output.User.Bio != "i like turtles and cats" {
		t.Errorf("CreateChannelFeedPostReaction the block user bio was not \"i like turtles and cats\": %s", output.User.Bio)
	}
	if output.User.Logo != "http://something.net/foo.png" {
		t.Errorf("CreateChannelFeedPostReaction the block user logo was not http://something.net/foo.png: %s", output.User.Logo)
	}
}

func TestDeleteChannelFeedPostReaction(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "https://api.twitch.tv/kraken/feed/12345/posts/20/reactions",
		httpmock.NewStringResponder(204, ``))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.DeleteChannelFeedPostReaction(&DeleteChannelFeedPostReactionInput{
		ChannelID: 12345,
		PostID:    "20",
		EmoteID:   "25",
	})

	if errorOutput != nil {
		t.Errorf("DeleteChannelFeedPostReaction errorOutput should have been nil: %+v", errorOutput)
	}

	if output == nil {
		t.Errorf("DeleteChannelFeedPostReaction output shouldn't have been nil")
	}
}
