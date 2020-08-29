package twitch

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetChannel(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/channel",
		httpmock.NewStringResponder(200, `{"mature":true,"status":"Struggle Bus 5: The Fight to Stay Alive","broadcaster_language":"en","display_name":"dallas","game":"Nioh","language":"en","_id":"1234","name":"dallas","created_at":"2013-06-03T19:12:02Z","updated_at":"2017-04-24T10:03:34Z","partner":false,"logo":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-profile_image-1a2c906ee2c35f12-300x300.png","video_banner":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-channel_offline_image-2e82c1df2a464df7-1920x1080.jpeg","profile_banner":null,"profile_banner_background_color":null,"url":"https://www.twitch.tv/dallas","views":2000,"followers":79,"broadcaster_type":"affiliate","stream_key":"live_44322889_a34ub37c8ajv98a0","email":"email@provider.com"}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.GetChannel()

	if errorOutput != nil {
		t.Errorf("GetChannel errorOutput should have been nil: %+v", errorOutput)
	}

	if output.ID != "1234" {
		t.Errorf("GetChannel the ID was not 1234: %s", output.ID)
	}

	if output.Language != "en" {
		t.Errorf("GetChannel the language was not en: %s", output.Language)
	}

	if output.Name != "dallas" {
		t.Errorf("GetChannel the name was not dallas: %s", output.Name)
	}
}

func TestGetChannelByID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/channels/1234",
		httpmock.NewStringResponder(200, `{"mature":true,"status":"Struggle Bus 5: The Fight to Stay Alive","broadcaster_language":"en","display_name":"dallas","game":"Nioh","language":"en","_id":"1234","name":"dallas","created_at":"2013-06-03T19:12:02Z","updated_at":"2017-04-24T10:03:34Z","partner":false,"logo":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-profile_image-1a2c906ee2c35f12-300x300.png","video_banner":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-channel_offline_image-2e82c1df2a464df7-1920x1080.jpeg","profile_banner":null,"profile_banner_background_color":null,"url":"https://www.twitch.tv/dallas","views":2000,"followers":79,"broadcaster_type":"affiliate","stream_key":"live_44322889_a34ub37c8ajv98a0","email":"email@provider.com"}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.GetChannelByID(&GetChannelByIDInput{
		ChannelID: 1234,
	})

	if errorOutput != nil {
		t.Errorf("GetChannelByID errorOutput should have been nil: %+v", errorOutput)
	}

	if output.ID != "1234" {
		t.Errorf("GetChannelByID the ID was not 1234: %s", output.ID)
	}

	if output.Language != "en" {
		t.Errorf("GetChannelByID the language was not en: %s", output.Language)
	}

	if output.Name != "dallas" {
		t.Errorf("GetChannelByID the name was not dallas: %s", output.Name)
	}
}

func TestUpdateChannel(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", "https://api.twitch.tv/kraken/channel/1234",
		httpmock.NewStringResponder(200, `{"mature":true,"status":"Struggle Bus 5: The Fight to Stay Alive","broadcaster_language":"en","display_name":"dallas","game":"Nioh","language":"en","_id":"1234","name":"dallas","created_at":"2013-06-03T19:12:02Z","updated_at":"2017-04-24T10:03:34Z","partner":false,"logo":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-profile_image-1a2c906ee2c35f12-300x300.png","video_banner":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-channel_offline_image-2e82c1df2a464df7-1920x1080.jpeg","profile_banner":null,"profile_banner_background_color":null,"url":"https://www.twitch.tv/dallas","views":2000,"followers":79,"broadcaster_type":"affiliate","stream_key":"live_44322889_a34ub37c8ajv98a0","email":"email@provider.com"}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.UpdateChannel(&UpdateChannelInput{
		ChannelID:          1234,
		Status:             "Wooo this is cool",
		Game:               "Minecraft",
		Delay:              60,
		ChannelFeedEnabled: true,
	})

	if errorOutput != nil {
		t.Errorf("GetChannelByID errorOutput should have been nil: %+v", errorOutput)
	}

	if output.ID != "1234" {
		t.Errorf("GetChannelByID the ID was not 1234: %s", output.ID)
	}

	if output.Language != "en" {
		t.Errorf("GetChannelByID the language was not en: %s", output.Language)
	}

	if output.Name != "dallas" {
		t.Errorf("GetChannelByID the name was not dallas: %s", output.Name)
	}
}

func TestGetChannelEditors(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/channels/1234/editors",
		httpmock.NewStringResponder(200, `{"users":[{"_id":129454141,"bio":null,"created_at":"2016-07-13T14:40:42Z","display_name":"dallasnchains","logo":null,"name":"dallasnchains","type":"user","updated_at":"2016-12-09T22:02:56Z"}]}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.GetChannelEditors(&GetChannelEditorsInput{
		ChannelID: 1234,
	})

	if errorOutput != nil {
		t.Errorf("GetChannelEditors errorOutput should have been nil: %+v", errorOutput)
	}

	if len(output.Users) != 1 {
		t.Errorf("GetChannelEditors the posts list was not 1 in length: %d", len(output.Users))
	}

	user := output.Users[0]

	if user.ID != 129454141 {
		t.Errorf("GetChannelEditors the first user id was not 129454141: %d", user.ID)
	}
	if user.Name != "dallasnchains" {
		t.Errorf("GetChannelEditors the first user name was not \"dallasnchains\": %s", user.Name)
	}
}

func TestGetChannelFollowers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/channels/1234/follows?cursor=123456789&direction=desc&limit=10&offset=10",
		httpmock.NewStringResponder(200, `{"_cursor":"1481675542963907000","_total":41,"follows":[{"created_at":"2016-12-14T00:32:22.963907Z","notifications":false,"user":{"_id":129454141,"bio":null,"created_at":"2016-07-13T14:40:42.398257Z","display_name":"dallasnchains","logo":null,"name":"dallasnchains","type":"user","updated_at":"2016-12-14T00:32:16.263122Z"}}]}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.GetChannelFollowers(&GetChannelFollowersInput{
		ChannelID: 1234,
		Limit:     10,
		Offset:    10,
		Cursor:    "123456789",
		Direction: "desc",
	})

	if errorOutput != nil {
		t.Errorf("GetChannelFollowers errorOutput should have been nil: %+v", errorOutput)
	}

	if len(output.Follows) != 1 {
		t.Errorf("GetChannelFollowers the follows list was not 1 in length: %d", len(output.Follows))
	}

	follow := output.Follows[0]

	if follow.User.ID != 129454141 {
		t.Errorf("GetChannelFollowers the first user id was not 129454141: %d", follow.User.ID)
	}
	if follow.User.Name != "dallasnchains" {
		t.Errorf("GetChannelFollowers the first user name was not \"dallasnchains\": %s", follow.User.Name)
	}
}
