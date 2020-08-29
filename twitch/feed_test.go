package twitch

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

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
