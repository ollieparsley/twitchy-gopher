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

func TestGetChannelTeams(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/channels/1234/teams",
		httpmock.NewStringResponder(200, `{"teams":[{"_id":10,"background":null,"banner":"https://static-cdn.jtvnw.net/jtv_user_pictures/team-staff-banner_image-606ff5977f7dc36e-640x125.png","created_at":"2011-10-25T23:55:47Z","display_name":"Twitch Staff","info":"Twitch staff stream here. Drop in and say \"hi\" sometime :)","logo":"https://static-cdn.jtvnw.net/jtv_user_pictures/team-staff-team_logo_image-76418c0c93a9d48b-300x300.png","name":"staff","updated_at":"2014-10-16T00:44:11Z"}]}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.GetChannelTeams(&GetChannelTeamsInput{
		ChannelID: 1234,
	})

	if errorOutput != nil {
		t.Errorf("GetChannelTeams errorOutput should have been nil: %+v", errorOutput)
	}

	if len(output.Teams) != 1 {
		t.Errorf("GetChannelTeams the teams list was not 1 in length: %d", len(output.Teams))
	}

	team := output.Teams[0]

	if team.ID != 10 {
		t.Errorf("GetChannelTeams the first team id was not 10: %d", team.ID)
	}
	if team.Name != "staff" {
		t.Errorf("GetChannelTeams the first team name was not \"staff\": %s", team.Name)
	}
}

func TestGetChannelSubscribers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/channels/1234/subscriptions?direction=desc&limit=10&offset=10",
		httpmock.NewStringResponder(200, `{"_total":4,"subscriptions":[{"_id":"e5e2ddc37e74aa9636625e8d2cc2e54648a30418","created_at":"2016-04-06T04:44:31Z","sub_plan":"1000","sub_plan_name":"Channel Subscription (mr_woodchuck)","user":{"_id":89614178,"bio":"Twitch staff member who is a heimerdinger main on the road to diamond.","created_at":"2015-04-26T18:45:34Z","display_name":"Mr_Woodchuck","logo":"https://static-cdn.jtvnw.net/jtv_user_pictures/mr_woodchuck-profile_image-a8b10154f47942bc-300x300.jpeg","name":"mr_woodchuck","type":"staff","updated_at":"2017-04-06T00:14:13Z"}}]}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.GetChannelSubscribers(&GetChannelSubscribersInput{
		ChannelID: 1234,
		Limit:     10,
		Offset:    10,
		Direction: "desc",
	})

	if errorOutput != nil {
		t.Errorf("GetChannelSubscribers errorOutput should have been nil: %+v", errorOutput)
	}

	if len(output.Subscriptions) != 1 {
		t.Errorf("GetChannelSubscribers the subscribers list was not 1 in length: %d", len(output.Subscriptions))
	}

	subscription := output.Subscriptions[0]

	if subscription.ID != "e5e2ddc37e74aa9636625e8d2cc2e54648a30418" {
		t.Errorf("GetChannelSubscribers the first subscriber id was not \"e5e2ddc37e74aa9636625e8d2cc2e54648a30418\": %s", subscription.ID)
	}
	if subscription.User.ID != 89614178 {
		t.Errorf("GetChannelSubscribers the first user id was not 89614178: %d", subscription.User.ID)
	}
}

func TestCheckChannelSubscriptionByUser(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/channels/1234/subscriptions/123456",
		httpmock.NewStringResponder(200, `{"_id":"8bded3af51046d2d365279fe92a976b6a4ceb006","created_at":"2017-04-08T19:15:39Z","sub_plan":"3000","sub_plan_name":"Channel Subscription (mr_woodchuck) - $24.99 Sub","user":{"_id":13405587,"bio":"Software Engineer at Twitch and casual speedrunner. I play video games!","created_at":"2010-06-27T05:33:45Z","display_name":"TWW2","logo":"https://static-cdn.jtvnw.net/jtv_user_pictures/tww2-profile_image-6af5324eddee1468-300x300.png","name":"tww2","type":"staff","updated_at":"2017-04-06T03:31:55Z"}}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.CheckChannelSubscriptionByUser(&CheckChannelSubscriptionByUserInput{
		ChannelID: 1234,
		UserID:    123456,
	})

	if errorOutput != nil {
		t.Errorf("CheckChannelSubscriptionByUser errorOutput should have been nil: %+v", errorOutput)
	}

	if output.ID != "8bded3af51046d2d365279fe92a976b6a4ceb006" {
		t.Errorf("CheckChannelSubscriptionByUser the ID was not 8bded3af51046d2d365279fe92a976b6a4ceb006: %s", output.ID)
	}

	if output.SubPlan != "3000" {
		t.Errorf("CheckChannelSubscriptionByUser the language was not en: %s", output.SubPlan)
	}

	if output.User.Name != "tww2" {
		t.Errorf("CheckChannelSubscriptionByUser the user name was not tww2: %s", output.User.Name)
	}
}

func TestGetChannelVideos(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.twitch.tv/kraken/channels/1234/videos?broadcast_type=upload%2Chighlight&langage=&limit=10&offset=10&sort=time",
		httpmock.NewStringResponder(200, `{"_total":583,"videos":[{"_id":"v102381501","broadcast_id":23711574096,"broadcast_type":"highlight","channel":{"_id":"20694610","display_name":"Towelliee","name":"towelliee"},"created_at":"2016-11-20T23:46:06Z","description":"Last minutes of stream","description_html":"Last minutes of stream<br>","fps":{"chunked":59.9997939597903,"high":30.2491085172346,"low":30.249192959941,"medium":30.2491085172346,"mobile":30.249192959941},"game":"World of Warcraft","language":"en","length":201,"preview":{"large":"https://static-cdn.jtvnw.net/s3_vods/664fa5856b_towelliee_23711574096_550644271//thumb/thumb102381501-640x360.jpg","medium":"https://static-cdn.jtvnw.net/s3_vods/664fa5856b_towelliee_23711574096_550644271//thumb/thumb102381501-320x180.jpg","small":"https://static-cdn.jtvnw.net/s3_vods/664fa5856b_towelliee_23711574096_550644271//thumb/thumb102381501-80x45.jpg","template":"https://static-cdn.jtvnw.net/s3_vods/664fa5856b_towelliee_23711574096_550644271//thumb/thumb102381501-{width}x{height}.jpg"},"published_at":"2016-11-20T23:46:06Z","resolutions":{"chunked":"1920x1080","high":"1280x720","low":"640x360","medium":"852x480","mobile":"400x226"},"status":"recorded","tag_list":"","thumbnails":{"large":[{"type":"generated","url":"https://static-cdn.jtvnw.net/s3_vods/664fa5856b_towelliee_23711574096_550644271//thumb/thumb102381501-640x360.jpg"}],"medium":[{"type":"generated","url":"https://static-cdn.jtvnw.net/s3_vods/664fa5856b_towelliee_23711574096_550644271//thumb/thumb102381501-320x180.jpg"}],"small":[{"type":"generated","url":"https://static-cdn.jtvnw.net/s3_vods/664fa5856b_towelliee_23711574096_550644271//thumb/thumb102381501-80x45.jpg"}],"template":[{"type":"generated","url":"https://static-cdn.jtvnw.net/s3_vods/664fa5856b_towelliee_23711574096_550644271//thumb/thumb102381501-{width}x{height}.jpg"}]},"title":"Last minutes of stream","url":"https://www.twitch.tv/towelliee/v/102381501","viewable":"public","viewable_at":null,"views":1761}]}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.GetChannelVideos(&GetChannelVideosInput{
		ChannelID:      1234,
		Limit:          10,
		Offset:         10,
		BroadcastTypes: []string{"upload", "highlight"},
		Sort:           "time",
	})

	if errorOutput != nil {
		t.Errorf("GetChannelVideos errorOutput should have been nil: %+v", errorOutput)
	}

	if len(output.Videos) != 1 {
		t.Errorf("GetChannelVideos the follows list was not 1 in length: %d", len(output.Videos))
	}

	video := output.Videos[0]

	if video.ID != "v102381501" {
		t.Errorf("GetChannelVideos the first user id was not v102381501: %s", video.ID)
	}
	if video.FPS.Chunked != float64(59.9997939597903) {
		t.Errorf("GetChannelVideos the first video chunked FPS 59.9997939597903: %v", video.FPS.Chunked)
	}
	if video.Resolutions.Chunked != "1920x1080" {
		t.Errorf("GetChannelVideos the first video chunked resolution \"\": %v", video.Resolutions.Chunked)
	}
}

func TestStartChannelCommercial(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.twitch.tv/kraken/channel/1234/commercial",
		httpmock.NewStringResponder(200, `{"Length":30,"Message":"foo","RetryAfter":480}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.StartChannelCommercial(&StartChannelCommercialInput{
		ChannelID: 1234,
		Length:    30,
	})

	if errorOutput != nil {
		t.Errorf("StartChannelCommercial errorOutput should have been nil: %+v", errorOutput)
	}

	if output.Length != 30 {
		t.Errorf("StartChannelCommercial the length was not 30: %d", output.Length)
	}

	if output.Message != "foo" {
		t.Errorf("StartChannelCommercial the message was not \"foo\": %s", output.Message)
	}

	if output.RetryAfter != 480 {
		t.Errorf("StartChannelCommercial the retry after was not 480: %d", output.RetryAfter)
	}
}

func TestResetStreamKey(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("DELETE", "https://api.twitch.tv/kraken/channels/1234/stream_key",
		httpmock.NewStringResponder(200, `{"_id":"44322889","broadcaster_language":"en","created_at":"2013-06-03T19:12:02Z","display_name":"dallas","email":"dttester@twitch.tv","followers":42,"game":"Final Fantasy XV","language":"en","logo":"https://static-cdn.jtvnw.net/jtv_user_pictures/dallas-profile_image-1a2c906ee2c35f12-300x300.png","mature":true,"name":"dallas","partner":false,"profile_banner":null,"profile_banner_background_color":null,"status":"The Finalest of Fantasies","stream_key":"live_44322889_nCGwsCl38pt21oj4UJJZbFQ9nrVIU5","updated_at":"2016-12-14T18:32:07Z","url":"https://www.twitch.tv/dallas","video_banner":null,"views":315}`))

	client := NewClient(&OAuthConfig{}, &http.Client{})

	output, errorOutput := client.ResetStreamKey(&ResetStreamKeyInput{
		ChannelID: 1234,
	})

	if errorOutput != nil {
		t.Errorf("ResetStreamKey errorOutput should have been nil: %+v", errorOutput)
	}

	if output.ID != "44322889" {
		t.Errorf("ResetStreamKey the ID was not 44322889: %s", output.ID)
	}

	if output.Language != "en" {
		t.Errorf("ResetStreamKey the language was not en: %s", output.Language)
	}

	if output.Name != "dallas" {
		t.Errorf("ResetStreamKey the name was not dallas: %s", output.Name)
	}
}
